package main

import (
	"day-13-orm/configs"
	"day-13-orm/routes"
)

func main() {
	configs.Init()

	e := routes.New()

	e.Logger.Fatal(e.Start(":8000"))
}
