package main

import (
	"praktikum_section_18/config"
	"praktikum_section_18/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}