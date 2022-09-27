package main

import (
	"day-13-orm/routes"
)

func main() {

	e := routes.New()

	e.Logger.Fatal(e.Start(":8000"))
}
