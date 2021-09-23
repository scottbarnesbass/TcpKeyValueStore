package main

import (
	"TcpKeyValueStore/logging"
	"TcpKeyValueStore/storeRepo"
	"TcpKeyValueStore/tcp"
	"log"
)

func main() {

	file, err := logging.SetupLogging("tcpKeyValueStore")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	logging.InfoLogger.Println("Opening repo")
	storeRepo.Open()
	defer storeRepo.Close()

	logging.InfoLogger.Println("Starting client")
	tcp.Run()
}

