package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv"
	"h02/database"
	"h02/server"
	"h02/structs"
	"h02/ws"
	"os"
)

var dbConnection = database.GetDbConnection()
var updatesChannel = make(chan *structs.TrackerData, 10)

func main() {
	fmt.Println("Hello Akmal")

	go ws.StartServer(updatesChannel)

	server.StartServer(":"+os.Getenv("SERVER_PORT"), func(data *structs.TrackerData) {
		updatesChannel <- data

		go database.WriteToDatabase(data, dbConnection)
	})
}
