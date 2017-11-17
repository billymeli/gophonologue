package main

import (
	"fmt"

	"github.com/george-e-shaw-iv/gophonologue/pkg/application"
)

func main() {
	app := application.New("application/", "document_root/", 3000)

	fmt.Println("CTRL+C to terminate application")
	fmt.Println("Attempting to start listening at localhost:3000")

	err := app.Start()
	if err != nil {
		fmt.Errorf("Fatal error running application: %v", err.Error())
	}
}
