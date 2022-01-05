package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

var htmlFile = "index.html"
var game Game

func wshandle(w http.ResponseWriter, r *http.Request) {
	upgrader := &websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade failed")
		return
	}
	defer conn.Close()

	game.init(5, "A", "B")
	log.Println("game init")

	for game.winner() == 3 {
		mt, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("readfailed")
			break
		}

		cor1, cor2 := handleInput(msg)
		moveAvailable := game.move(cor1, cor2)

		game.print(log.Printf)

		if moveAvailable {
			err = conn.WriteMessage(mt, []byte{'1'})
			if err != nil {
				log.Println("write failed:", err)
				break
			}
		} else {
			err = conn.WriteMessage(mt, []byte{'0'})
			if err != nil {
				log.Println("write failed:", err)
				break
			}
		}
	}
}

func handleInput(msg []byte) (string, string) {
	input := string(msg) // <cor1 cor2>
	inputArr := strings.Split(input, " ")

	return inputArr[0], inputArr[1]
}

func main() {
	// go broadcaster()
	http.HandleFunc("/chess", wshandle)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, htmlFile)
	})

	log.Println("server start at :8899")
	log.Fatal(http.ListenAndServe(":8899", nil))
}
