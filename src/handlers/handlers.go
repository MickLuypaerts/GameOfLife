package handlers

import (
	"../game"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http" // https://golang.org/pkg/net/http/#pkg-constants
)

type gameHandlers struct {
	gameState *game.Game
}

func NewMux(gameState *game.Game) *http.ServeMux {
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

	if !checkMethod(r.Method, http.MethodPost, w) {
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if !handleReadAllError(err, w) {
		return
	}
	log.Println(string(reqBody))

	data := game.CellData{}
	if !handleUnmashalError(json.Unmarshal(reqBody, &data), w) {
		return
	}

	json.NewEncoder(w).Encode(data)
	if err = gs.gameState.Set(data.X, data.Y, data.State); err != nil {
		log.Printf(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (gs *gameHandlers) step(w http.ResponseWriter, r *http.Request) {
	if !checkMethod(r.Method, http.MethodGet, w) {
		return
	}
	gs.gameState.Tick()
	log.Printf("Changed cells: %+v\n", gs.gameState.ChangedCells)
	//w.WriteHeader(200)
	json.NewEncoder(w).Encode(gs.gameState.ChangedCells)
}

func (gs *gameHandlers) resetBoard(w http.ResponseWriter, r *http.Request) {
	if !checkMethod(r.Method, http.MethodGet, w) {
		return
	}
	log.Println("reseting board.")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	//w.Write([]byte("the board has been reset."))
	fmt.Fprintln(w, "the board has been reset.")
	gs.gameState.Board.Reset()
}

func (gs *gameHandlers) getBoardSize(w http.ResponseWriter, r *http.Request) {
	if !checkMethod(r.Method, http.MethodGet, w) {
		return
	}
	log.Println("Sending board size.")
	json.NewEncoder(w).Encode(gs.gameState.Board.BoardSize)
}

func (gs *gameHandlers) createNewBoard(w http.ResponseWriter, r *http.Request) {

	if !checkMethod(r.Method, http.MethodPost, w) {
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if !handleReadAllError(err, w) {
		return
	}
	log.Printf("Creating new board: %+v", string(reqBody))

	data := game.Option{}
	log.Println(string(reqBody))
	if !handleUnmashalError(json.Unmarshal(reqBody, &data), w) {
		return
	}
	gs.gameState = game.NewGame(data.Column, data.Row)
	//w.WriteHeader(200)
	json.NewEncoder(w).Encode(reqBody) // TODO CHANGE RESPONSE
}

// error funcs

func checkMethod(received string, expected string, w http.ResponseWriter) bool {
	if received != expected {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed) // can change StatusText to string
		return false
	} else {
		return true
	}
}

func handleReadAllError(err error, w http.ResponseWriter) bool {
	if err != nil {
		log.Printf("Body read error, %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return false
	} else {
		return true
	}
}

func handleUnmashalError(err error, w http.ResponseWriter) bool {
	if err != nil {
		log.Printf("Body parse error, %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return false
	} else {
		return true
	}
}
