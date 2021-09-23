package coordinator

import (
	"TcpKeyValueStore/globals"
	"TcpKeyValueStore/logging"
	"TcpKeyValueStore/storeRepo"
	"errors"
)

func ExecuteAction(method string, arguments []string) (string, error) {
	switch method {
	case globals.RequestMethodPut:
		if err := storeRepo.Put(storeRepo.Key(arguments[0]), storeRepo.NewData(arguments[1])); err != nil {
			logging.ErrorLogger.Println(logging.PrepLogArgs("Failed to execute method :", method, "arguments :", arguments))
			return "err", err
		} else {
			logging.InfoLogger.Println(logging.PrepLogArgs("Put value executed for method :", method, "arguments :", arguments))
			return "ack", nil
		}
	case globals.RequestMethodGet:
		if result, err := storeRepo.Get(storeRepo.Key(arguments[0])); err == storeRepo.ErrKeyNotFound {
			logging.ErrorLogger.Println(logging.PrepLogArgs("Failed to execute method :", method, "arguments :", arguments, err))
			return "nil", err
		} else if err != nil {
			logging.ErrorLogger.Println(logging.PrepLogArgs("Failed to execute method :", method, "arguments :", arguments))
			return "err", err
		} else {
			data := result.(storeRepo.Data)
			response := data.Content
			logging.InfoLogger.Println(logging.PrepLogArgs("Get by key executed for method :", method, ", arguments :", arguments, ", response :", response))
			return response, nil
		}
	case globals.RequestMethodDelete:
		if err := storeRepo.Delete(storeRepo.Key(arguments[0]));
			err == storeRepo.ErrKeyNotFound || err == nil {
			logging.InfoLogger.Println(logging.PrepLogArgs("Deleting by key executed without error for method :", method, "arguments :", arguments))
			return "ack", nil
		} else if err != nil {
			logging.ErrorLogger.Println(logging.PrepLogArgs("Failed to execute method :", method, "arguments :", arguments))
			return "err", err
		}
	}
	logging.ErrorLogger.Println(logging.PrepLogArgs("Error executing method :", method, "arguments :", arguments))
	return "err", errors.New("unresolved action method")
}
