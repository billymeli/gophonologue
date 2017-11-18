package messenger

import (
	"encoding/json"
	"html"
	"net/http"
	"strconv"
	"time"

	"github.com/george-e-shaw-iv/gophonologue/pkg/database"
	"github.com/george-e-shaw-iv/gophonologue/pkg/message"
)

func Post(res http.ResponseWriter, req *http.Request, dir string) {
	if req.Method != "POST" {
		http.Error(res, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	var requestData struct {
		Username string `json:"username"`
		Message  string `json:"message"`
	}

	err := json.NewDecoder(req.Body).Decode(&requestData)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	var responseData message.Message
	responseData.Username = html.EscapeString(requestData.Username)
	responseData.Message = html.EscapeString(requestData.Message)

	timestamp := strconv.Itoa(int(time.Now().Unix()))
	response := make(map[string]message.Message, 1)
	response[timestamp] = responseData

	ds, err := database.Open(dir + database.DB_MAIN)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	defer ds.Close()

	err = ds.Put(database.BUCKET_MESSAGES, []byte(timestamp), responseData)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "application/json")
	res.Write(jsonData)
}
