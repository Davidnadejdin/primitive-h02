package database

import (
	"database/sql"
	"h02/structs"
	"log"
	"os"
)

func GetDbConnection() *sql.DB {
	dbConnection, err := sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@/"+os.Getenv("DB_NAME"))

	if err != nil {
		log.Fatalln(err)
	}

	return dbConnection
}

func WriteToDatabase(data *structs.TrackerData, connection *sql.DB) {
	rows, err := connection.Query("UPDATE trackers SET lat = ?, lng = ?  where imei = ?", data.Lat, data.Long, data.Imei)

	if err != nil {
		log.Println(err)

		return
	}

	err = rows.Close()

	if err != nil {
		log.Println(err)
		return
	}
}
