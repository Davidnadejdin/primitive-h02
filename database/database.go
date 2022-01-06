package database

import (
	"database/sql"
	"github.com/joho/godotenv"
	"h02/structs"
	"log"
	"os"
)

type Database struct {
	connection *sql.DB
}

func InitDb() *Database {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	connection, err := sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@/"+os.Getenv("DB_NAME"))

	if err != nil {
		log.Fatalln(err)
	}

	db := Database{
		connection: connection,
	}

	return &db
}

func (database *Database) Write(data *structs.TrackerData) {
	rows, err := database.connection.Query("UPDATE trackers SET lat = ?, lng = ?  where imei = ?", data.Lat, data.Long, data.Imei)

	if err != nil {
		log.Println(err)

		return
	}

	err = rows.Close()

	if err != nil {
		log.Println(err)
	}
}
