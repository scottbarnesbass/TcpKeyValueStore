package main

import (
	"TcpKeyValueStore/logging"
	"TcpKeyValueStore/store"
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
	storeRepo := store.CreateNewStore()

	logging.InfoLogger.Println("Starting client")
	tcp.Run(storeRepo)
}

