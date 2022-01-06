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

func main() {
	fmt.Println("Hello Akmal")

	wsServer := ws.StartServer()
	db := database.InitDb()

	server.StartServer(":"+os.Getenv("SERVER_PORT"), func(data *structs.TrackerData) {
		go wsServer.SendMessage(data)
		go db.Write(data)
	})
}
