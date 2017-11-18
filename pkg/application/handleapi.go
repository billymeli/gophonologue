package application

import (
	"net/http"
	"strings"

	"github.com/george-e-shaw-iv/gophonologue/pkg/api/messenger"
)

func (app *Application) handleAPI(res http.ResponseWriter, req *http.Request) bool {
	path := req.URL.Path[1:]
	path = (app.Directory + app.DocumentRoot + path)

	splitUrl := strings.SplitN(path, "api", 2)
	suspectApi := strings.ToLower(splitUrl[len(splitUrl)-1])

	switch suspectApi {
	case "/messenger/post":
		messenger.Post(res, req, app.Directory)
		return true
	case "/messenger/get":
		messenger.Get(res, req, app.Directory)
		return true
	default:
		return false
	}
}
