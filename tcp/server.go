package tcp

import (
	"TcpKeyValueStore/coordinator"
	"TcpKeyValueStore/fomatting/requestFormatting"
	"TcpKeyValueStore/fomatting/responseFormatting"
	"TcpKeyValueStore/globals"
	"TcpKeyValueStore/logging"
	"TcpKeyValueStore/repository"
	"fmt"
	"net"
)

func Run(repo repository.Repo) {

	logging.InfoLogger.Println("opening listener")
	listener, err := net.Listen("tcp4", ":1234")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer func() {
		listener.Close()
		logging.InfoLogger.Println("listener closed")
	}()



	for {
		logging.InfoLogger.Println("Accepting")
		conn, err := listener.Accept()
		if err != nil{
			logging.ErrorLogger.Println(err)
			continue
		}

		go handle(conn, repo)
		logging.InfoLogger.Println("Handled")
	}
}

func handle(conn net.Conn, repo repository.Repo) {
	defer conn.Close()
	logging.InfoLogger.Println("Handling connection")
	for {
		method, err := requestFormatting.GetRequestMethod(conn)
		if err != nil {
			logging.ErrorLogger.Println(err)
			conn.Write([]byte("err"))
			return
		}
		if method == globals.RequestMethodBye {
			logging.InfoLogger.Println("Bye request received")
			return
		}
		requiredArgumentsCount, err := requestFormatting.GetRequiredArgumentsCount(method)
		if err != nil {
			logging.ErrorLogger.Println(err)
			conn.Write([]byte("err"))
			return
		}
		if arguments, err := requestFormatting.GetRequestArguments(conn, requiredArgumentsCount); err != nil {
			logging.ErrorLogger.Println(err)
			conn.Write([]byte("err"))
		} else {
			if len(arguments) < requiredArgumentsCount {
				logging.ErrorLogger.Println("too few arguments to complete request")
				conn.Write([]byte("err"))
			} else {
				responseString, err := coordinator.ExecuteAction(repo, method, arguments)
				conn.Write(responseFormatting.FormatResponse(method, responseString, err))
			}
		}
	}
}
