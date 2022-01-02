package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	"h02/database"
	"h02/server"
	"h02/structs"
	"os"
)

func main() {
	godotenv.Load()

	dbConnection := database.GetDbConnection()

	server.StartServer(":"+os.Getenv("SERVER_PORT"), func(data structs.TrackerData) {
		go database.WriteToDatabase(data, dbConnection)
	})
}
