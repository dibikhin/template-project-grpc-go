package main

import (
	"fmt"
	"log"

	"TODO/pkg/app"
	"TODO/pkg/client"
)

func main() {
	log.Println("app: starting...")

	teardown := func() {}
	onExit := func() {
		teardown()
		log.Println("app: stopped")
		fmt.Println("\nBye!")
	}
	go app.WaitForExit(onExit)
	log.Println("app: started")

	fmt.Print("\nWelcome to TODO\n\n")

	cfg := app.LoadConfig()
	cl, teardown := client.Connect(cfg)
	s := client.NewTODOService(cl, app.DefaultRead)

	client.RunLoop(s)

	onExit()
}
