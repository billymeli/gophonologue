package application

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/george-e-shaw-iv/gophonologue/pkg/contenttype"
)

type Application struct {
	Directory    string
	DocumentRoot string
	Port         int
}

var server http.Server

func New(dir string, root string, port int) *Application {
	app := &Application{
		Directory:    dir,
		DocumentRoot: root,
		Port:         port,
	}

	server = http.Server{
		Addr:         "localhost:" + strconv.Itoa(app.Port),
		Handler:      app,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	return app
}

func (app *Application) Start() error {
	return server.ListenAndServe()
}

func (app *Application) Stop(graceful bool) error {
	if graceful {
		context, cancel := context.WithTimeout(context.Background(), 5)
		defer cancel()

		err := server.Shutdown(context)
		if err == nil {
			return nil
		}

		fmt.Printf("Graceful shutdown failed attempting forced: %v\n", err)
	}

	if err := server.Close(); err != nil {
		return err
	}

	return nil
}

func (app *Application) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path[1:]
	if len(path) == 0 {
		path = (app.Directory + app.DocumentRoot + "index.html")
	} else {
		path = (app.Directory + app.DocumentRoot + path)
	}

	f, err := os.Open(path)

	if err != nil {
		http.Error(res, err.Error(), http.StatusNotFound)
		return
	}

	ct, err := contenttype.GetContentType(path)

	if err != nil {
		http.Error(res, err.Error(), http.StatusUnsupportedMediaType)
		return
	}

	res.Header().Add("Content-Type", ct)
	_, err = io.Copy(res, f)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}
