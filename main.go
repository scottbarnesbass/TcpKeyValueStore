package main

import (
	"TcpKeyValueStore/logging"
	"TcpKeyValueStore/repository"
	"TcpKeyValueStore/tcp"
	"log"
)

func main() {

	file, err := logging.SetupLogging("tcpKeyValueStore", logging.LogToFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	logging.InfoLogger.Println("Opening repo")
	storeRepo := repository.CreateNewStore()

	logging.InfoLogger.Println("Starting client")
	tcp.Run(storeRepo)
}

