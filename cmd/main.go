package main

import (
	"ImransProfoiloWebsite/db"
	"ImransProfoiloWebsite/renderers"
	"ImransProfoiloWebsite/routes"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	db, err := db.NewDatabase()
	if err != nil {
		log.Fatal("error opening the database: %w", err)
	}
	defer db.Close()
	log.Println("we are connected to the database")

	e := echo.New()

	e.Renderer = &renderers.HTMLTemplate{
		Dir: "templates",
		Ext: ".html",
	}

	routes.SetupRoutes(e)
	err = e.Start(":8080")

	if err != nil {
		log.Fatal("the server isnt running: %v", err)
	}

}
