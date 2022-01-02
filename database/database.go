package database

import (
	"database/sql"
	"h02/structs"
	"os"
)

func GetDbConnection() *sql.DB {
	dbConnection, _ := sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@/"+os.Getenv("DB_NAME"))

	return dbConnection
}

func WriteToDatabase(data structs.TrackerData, connection *sql.DB) {
	connection.Query("UPDATE trackers SET lat = ?, lng = ?  where imei = ?", data.Lat, data.Long, data.Imei)
}
