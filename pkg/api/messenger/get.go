package messenger

import (
	"encoding/json"
	"net/http"

	"github.com/george-e-shaw-iv/gophonologue/pkg/database"
	"github.com/george-e-shaw-iv/gophonologue/pkg/message"
)

func Get(res http.ResponseWriter, req *http.Request, dir string) {
	if req.Method != "GET" {
		http.Error(res, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	ds, err := database.Open(dir + database.DB_MAIN)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	defer ds.Close()

	messages := make(map[string]message.Message)

	err = ds.GetMessages(database.BUCKET_MESSAGES, messages)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(messages)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "application/json")
	res.Write(jsonData)
}
