package main

// https://semaphoreci.com/community/tutorials/how-to-deploy-a-go-web-application-with-docker

import (
	"./game"
	"./handlers"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

func main() {
	gameState := game.NewGame(50, 50)
	port := ":3000"
	mux := handlers.Mux(gameState)
	log.Printf("Starting the server on %s\n", port)
	// err := openbrowser("http://localhost:3000")
	//if err != nil {
	//	log.Fatal(err)
	//}
	log.Fatal(http.ListenAndServe(port, mux))
}

func openbrowser(url string) error {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		return err
	} else {
		return nil
	}
}
