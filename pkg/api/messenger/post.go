package messenger

import (
	"encoding/json"
	"html"
	"net/http"
	"time"
)

func Post(res http.ResponseWriter, req *http.Request) {
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

	var responseData struct {
		Username         string `json:"username"`
		SanitizedMessage string `json:"sanitized_message"`
		Timestamp        int    `json:"timestamp"`
	}

	responseData.Username = html.EscapeString(requestData.Username)
	responseData.SanitizedMessage = html.EscapeString(requestData.Message)
	responseData.Timestamp = int(time.Now().Unix())

	/* TODO: Latest 50 messages stored in database logic here */

	jsonData, err := json.Marshal(responseData)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "application/json")
	res.Write(jsonData)
}
