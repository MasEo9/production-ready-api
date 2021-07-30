package main

import "fmt"

// App - struct contains pointers to db connections
type App struct{}

func (app *App) run() error {
	fmt.Println("Setting Up Our APP")
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
