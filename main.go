package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	"h02/database"
	"h02/server"
	"h02/structs"
	"h02/ws"
	"log"
	"os"
)

var dbConnection = database.GetDbConnection()
var updatesChannel = make(chan *structs.TrackerData)

func main() {
	fmt.Println("Hello Akmal")

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	go ws.StartServer(updatesChannel)

	server.StartServer(":"+os.Getenv("SERVER_PORT"), func(data *structs.TrackerData) {
		go database.WriteToDatabase(data, dbConnection)

		updatesChannel <- data
	})
}
