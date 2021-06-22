package main

// https://semaphoreci.com/community/tutorials/how-to-deploy-a-go-web-application-with-docker

import (
	"./game"
	"./handlers"
	"log"
	"net/http"
)

func main() {
	gameState := game.NewGame(50, 50)
	port := ":3000"
	mux := handlers.NewMux(gameState)
	log.Printf("Starting the server on %s\n", port)
	log.Fatal(http.ListenAndServe(port, mux))
}
