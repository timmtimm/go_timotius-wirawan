package main

import (
	"praktikum_section_20/config"
	"praktikum_section_20/middlewares"
	"praktikum_section_20/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	middlewares.LogMiddleware(e)
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
