package ws

import (
	"encoding/json"
	_ "encoding/json"
	"flag"
	"github.com/gorilla/websocket"
	"h02/structs"
	"log"
	"net/http"
	"sync"
)

var addr = flag.String("addr", "0.0.0.0:1338", "http service address")
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type clients struct {
	mu sync.Mutex
	v  map[*websocket.Conn]bool
}

type Server struct {
	clients clients
}

func StartServer() *Server {
	server := Server{
		clients{v: make(map[*websocket.Conn]bool)},
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
	server.clients.mu.Lock()
	defer server.clients.mu.Unlock()

	for conn := range server.clients.v {
		jsoned, _ := json.Marshal(*data)

		err := conn.WriteMessage(1, jsoned)

		if err != nil {
			log.Println(err)
		}
	}
}

func (server *Server) echo(w http.ResponseWriter, r *http.Request) {
	connection, err := upgrader.Upgrade(w, r, nil)

	defer func(connection *websocket.Conn) {
		err := connection.Close()

		if err != nil {
			log.Println(err)
		}
	}(connection)

	if err != nil {
		log.Println(err)
	}

	server.clients.mu.Lock()
	server.clients.v[connection] = true
	server.clients.mu.Unlock()

	defer func() {
		server.clients.mu.Lock()
		delete(server.clients.v, connection)
		server.clients.mu.Unlock()
	}()

	for {
		mt, _, err := connection.ReadMessage()

		if err != nil || mt == websocket.CloseMessage {
			break
		}
	}
}
