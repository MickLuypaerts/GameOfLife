package main

// https://semaphoreci.com/community/tutorials/how-to-deploy-a-go-web-application-with-docker

import (
	"gameoflife/game"     // "github.com/MickLuypaerts/GameOfLife/tree/master/src/game"     // "./game"
	"gameoflife/handlers" // "github.com/MickLuypaerts/GameOfLife/tree/master/src/handlers" // "./handlers"
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
