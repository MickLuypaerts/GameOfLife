package main

import (
	"./game"
	"./handlers"
	"log"
	"net/http"
)

var gameState game.Game

func main() {
	port := ":3000"
	gameState.InitGame(50, 50)
	mux := handlers.Mux(&gameState)
	log.Printf("Starting the server on %s\n", port)
	log.Fatal(http.ListenAndServe(port, mux))
}
