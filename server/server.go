package server

import (
	"h02/structs"
	"net"
	"strconv"
	"strings"
)

func StartServer(address string, handle func(data *structs.TrackerData)) {
	ln, _ := net.Listen("tcp", address)

	for {
		conn, err := ln.Accept()

		if err != nil {
			panic(err)
		}

		go func() {
			buffer := make([]byte, 96)

			n, err := conn.Read(buffer)

			if err != nil {
				return
			}

			err = conn.Close()

			if err != nil {
				return
			}

			rawData := string(buffer[:n])

			if len(rawData) > 96 {
				data := parse(rawData)

				handle(data)
			}
		}()
	}
}

func parse(rawData string) *structs.TrackerData {
	data := strings.Split(strings.TrimSpace(rawData), ",")

	trackerData := structs.TrackerData{
		Imei: data[1],
		Lat:  fixLat(data[5]),
		Long: fixLong(data[7]),
	}

	return &trackerData
}

func fixLat(lat string) string {
	minutes, err := strconv.ParseFloat(lat[2:9], 64)

	if err != nil {
		panic(err)
	}

	degrees, err := strconv.ParseFloat(lat[0:2], 64)

	if err != nil {
		panic(err)
	}

	return strconv.FormatFloat(degrees+(minutes/60), 'f', 6, 64)
}

func fixLong(lat string) string {
	minutes, err := strconv.ParseFloat(lat[3:10], 64)

	if err != nil {
		panic(err)
	}

	degrees, err := strconv.ParseFloat(lat[0:3], 64)

	if err != nil {
		panic(err)
	}

	return strconv.FormatFloat(degrees+(minutes/60), 'f', 6, 64)
}
