package main

import (
	"./game"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

var gameState game.Game

func main() {
	gameState.InitGame(50, 50)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/set", setCell)
	http.HandleFunc("/step", step)
	http.HandleFunc("/resetboard", resetBoard)
	http.HandleFunc("/getboardsize", getBoardSize)
	gameState.Board.Reset()
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func getBoardSize(w http.ResponseWriter, r *http.Request) {
	log.Println("Sending board size")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(gameState.Board.BoardSize)
}

func resetBoard(w http.ResponseWriter, r *http.Request) {
	log.Println("reseting board.")
	w.WriteHeader(200)
	w.Write([]byte(`{"msg": "the board has been reset."}`))
	gameState.Board.Reset()
}

func setCell(w http.ResponseWriter, r *http.Request) {
	data := game.CellData{}
	reqBody, _ := ioutil.ReadAll(r.Body)
	log.Println(string(reqBody))

	if err := json.Unmarshal(reqBody, &data); err != nil {
		log.Printf("Body parse error, %v", err)
		w.WriteHeader(400)
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(data)
	gameState.Set(data.X, data.Y, data.State)
}

func step(w http.ResponseWriter, r *http.Request) {
	gameState.Tick()
	log.Printf("Changed cells: %+v\n", gameState.ChangedCells)
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(gameState.ChangedCells)
}
