package handlers

import (
	"../game"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type gameHandlers struct {
	gameState *game.Game
}

func Mux(gameState *game.Game) *http.ServeMux {
	mux := http.NewServeMux()
	var gameHandlers gameHandlers
	gameHandlers.gameState = gameState
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/", fs)
	mux.HandleFunc("/set", gameHandlers.setCell)
	mux.HandleFunc("/step", gameHandlers.step)
	mux.HandleFunc("/resetboard", gameHandlers.resetBoard)
	mux.HandleFunc("/getboardsize", gameHandlers.getBoardSize)
	mux.HandleFunc("/createnewboard", gameHandlers.createNewBoard)
	return mux
}

func (gs *gameHandlers) setCell(w http.ResponseWriter, r *http.Request) {
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
	gs.gameState.Set(data.X, data.Y, data.State)
}

func (gs *gameHandlers) step(w http.ResponseWriter, r *http.Request) {
	gs.gameState.Tick()
	log.Printf("Changed cells: %+v\n", gs.gameState.ChangedCells)
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(gs.gameState.ChangedCells)
}

func (gs *gameHandlers) resetBoard(w http.ResponseWriter, r *http.Request) {
	log.Println("reseting board.")
	w.WriteHeader(200)
	w.Write([]byte(`{"msg": "the board has been reset."}`))
	gs.gameState.Board.Reset()
}

func (gs *gameHandlers) getBoardSize(w http.ResponseWriter, r *http.Request) {
	log.Println("Sending board size.")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(gs.gameState.Board.BoardSize)
}

func (gs *gameHandlers) createNewBoard(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)
	log.Printf("Creating new board: %+v", string(reqBody))

	data := game.Option{}
	log.Println(string(reqBody))
	if err := json.Unmarshal(reqBody, &data); err != nil {
		log.Printf("Body parse error, %v", err)
		w.WriteHeader(400)
		return
	}
	gs.gameState = game.NewGame(data.Column, data.Row)
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(reqBody) // TODO CHANGE RESPONSE
}
