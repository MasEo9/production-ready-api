package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/MasEo9/go-rest-api-course/internal/transport/http"
)

// App - struct contains pointers to db connections
type App struct{}

func (app *App) Run() error {
	fmt.Println("Setting Up Our APP")

	handler = transportHTTP.NewHander()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to setup Server")
		return err
	}

	return nil
}

func main() {
	fmt.Println("Go REST API Course")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our REST API")
		fmt.Print(err)
	}
}
