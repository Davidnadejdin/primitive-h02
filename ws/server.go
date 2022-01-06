package ws

import (
	"encoding/json"
	_ "encoding/json"
	"flag"
	"github.com/gorilla/websocket"
	"h02/structs"
	"log"
	"net/http"
)

var addr = flag.String("addr", "0.0.0.0:1338", "http service address")
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Server struct {
	clients map[*websocket.Conn]bool
}

func StartServer() *Server {
	server := Server{
		make(map[*websocket.Conn]bool),
	}

	http.HandleFunc("/", server.echo)

	go func() {
		flag.Parse()

		err := http.ListenAndServe(*addr, nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	return &server
}

func (server *Server) SendMessage(data *structs.TrackerData) {
	for conn := range server.clients {
		jsoned, _ := json.Marshal(*data)

		err := conn.WriteMessage(1, jsoned)

		if err != nil {
			log.Println(err)
		}
	}
}

func (server *Server) echo(w http.ResponseWriter, r *http.Request) {
	connection, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
	}

	server.clients[connection] = true

	for {
		mt, _, err := connection.ReadMessage()

		if err != nil || mt == websocket.CloseMessage {
			break
		}
	}

	delete(server.clients, connection)
	err = connection.Close()

	if err != nil {
		log.Println(err)
	}
}
