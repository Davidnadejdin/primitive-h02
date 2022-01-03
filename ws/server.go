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
var upgrader = websocket.Upgrader{}

type Server struct {
	updatesChannel chan *structs.TrackerData
}

func StartServer(updatesChannel chan *structs.TrackerData) {
	server := Server{
		updatesChannel: updatesChannel,
	}

	flag.Parse()

	http.HandleFunc("/", server.echo)

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) echo(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	connection, _ := upgrader.Upgrade(w, r, nil)

	for {
		jsoned, _ := json.Marshal(<-server.updatesChannel)

		err := connection.WriteMessage(1, jsoned)

		if err != nil {
			break
		}
	}

	err := connection.Close()
	if err != nil {
		log.Println(err)
	}
}
