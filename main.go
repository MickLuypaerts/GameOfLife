package main

import (
	"./game"
	"./handlers"
	"log"
	"net/http"
)

func main() {
	gameState := game.Game{}
	port := ":3000"
	gameState.InitGame(50, 50)
	mux := handlers.Mux(&gameState)
	log.Printf("Starting the server on %s\n", port)
	log.Fatal(http.ListenAndServe(port, mux))
}
