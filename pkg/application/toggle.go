package application

import (
	"context"
	"fmt"
)

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
