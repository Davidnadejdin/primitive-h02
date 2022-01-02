package server

import (
	"h02/structs"
	"net"
	"strconv"
	"strings"
)

func StartServer(address string, handle func(data structs.TrackerData)) {
	ln, _ := net.Listen("tcp", address)

	for {
		conn, _ := ln.Accept()
		buffer := make([]byte, 96)

		n, _ := conn.Read(buffer)

		data := parse(string(buffer[:n]))

		go handle(data)

		conn.Close()
	}
}

func parse(rawData string) structs.TrackerData {
	data := strings.Split(strings.TrimSpace(rawData), ",")

	trackerData := structs.TrackerData{
		Imei: data[1],
		Lat:  fixLat(data[5]),
		Long: fixLong(data[7]),
	}

	return trackerData
}

func fixLat(lat string) string {
	minutes, _ := strconv.ParseFloat(lat[2:9], 64)
	degrees, _ := strconv.ParseFloat(lat[0:2], 64)

	return strconv.FormatFloat(degrees+(minutes/60), 'f', 6, 64)
}

func fixLong(lat string) string {
	minutes, _ := strconv.ParseFloat(lat[3:10], 64)
	degrees, _ := strconv.ParseFloat(lat[0:3], 64)

	return strconv.FormatFloat(degrees+(minutes/60), 'f', 6, 64)
}
