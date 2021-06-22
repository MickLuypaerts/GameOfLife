package main

import (
	"gameoflife/game"
	"gameoflife/handlers"
	"log"
	"net/http"
)

func main() {
	gameState := game.NewGame(50, 50)
	port := ":8080"
	mux := handlers.NewMux(gameState)
	log.Printf("Starting the server on %s\n", port)
	log.Fatal(http.ListenAndServe(port, mux))
}
