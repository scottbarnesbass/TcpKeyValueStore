package coordinator

import (
	"TcpKeyValueStore/globals"
	"TcpKeyValueStore/logging"
	"TcpKeyValueStore/repository"
	"errors"
)

func ExecuteAction(repo repository.Repo, method string, arguments []string, ) (string, error) {
	switch method {
	case globals.RequestMethodPut:
		if err := repo.Put(repository.Key(arguments[0]), repository.NewData(arguments[1])); err != nil {
			logging.ErrorLogger.Println(logging.PrepLogArgs(
				"Failed to execute method :", method, "arguments :", arguments, err))
			return "err", err
		} else {
			logging.InfoLogger.Println(logging.PrepLogArgs(
				"Put value executed for method :", method, "arguments :", arguments))
			return "ack", nil
		}
	case globals.RequestMethodGet:
		if result, err := repo.Get(repository.Key(arguments[0])); err == repository.ErrKeyNotFound {
			logging.ErrorLogger.Println(logging.PrepLogArgs(
				"Failed to execute method :", method, "arguments :", arguments, err))
			return "nil", err
		} else if err != nil {
			logging.ErrorLogger.Println(logging.PrepLogArgs(
				"Failed to execute method :", method, "arguments :", arguments, err))
			return "err", err
		} else {
			data := result.(repository.Data)
			response := data.Content
			logging.InfoLogger.Println(logging.PrepLogArgs(
				"Get by key executed for method :", method, ", arguments :", arguments, ", response :", response))
			return response, nil
		}
	case globals.RequestMethodDelete:
		if err := repo.Delete(repository.Key(arguments[0]));
			err == repository.ErrKeyNotFound || err == nil {
			logging.InfoLogger.Println(logging.PrepLogArgs(
				"Deleting by key executed without error for method :", method, "arguments :", arguments))
			return "ack", nil
		} else if err != nil {
			logging.ErrorLogger.Println(logging.PrepLogArgs(
				"Failed to execute method :", method, "arguments :", arguments, err))
			return "err", err
		}
	}
	logging.ErrorLogger.Println(logging.PrepLogArgs(
		"Error executing method :", method, "arguments :", arguments))
	return "err", errors.New("unresolved action method")
}
