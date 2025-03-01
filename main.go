package main

import (
	"log"
	"vm-maker/bootstrap"
)

func main() {
	app := bootstrap.NewApplication()
	log.Fatal(app.Listen(":3000"))
}
