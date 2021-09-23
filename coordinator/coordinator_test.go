package coordinator

import (
	"TcpKeyValueStore/globals"
	"TcpKeyValueStore/logging"
	"TcpKeyValueStore/repository"
	"log"
	"testing"
)

func TestExecuteAction(t *testing.T) {
	store := repository.CreateNewStore()

	file, err := logging.SetupLogging("coordinator_test", logging.LogToConsole)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	key := "some"
	value := "thing"
	t.Run("put key/value", func(t *testing.T) {
		method := globals.RequestMethodPut
		arguments := []string{ key, value }

		if result, err := ExecuteAction(store, method, arguments); err != nil {
			t.Error("should not have thrown error", err)
		} else if result != "ack" {
			t.Error("incorrect result returned, expected ack but got", result)
		}
	})
	t.Run("get by key", func(t *testing.T) {
		method := globals.RequestMethodGet
		arguments := []string{ key }
		expectedResult := "thing"

		if result, err := ExecuteAction(store, method, arguments); err != nil {
			t.Error("should not have thrown error", err)
		}else if result != expectedResult {
			t.Error("incorrect result returned,", expectedResult, "but got", result)
		}
	})
	t.Run("delete by key", func(t *testing.T) {
		method := globals.RequestMethodDelete
		arguments := []string{ key }
		expectedResult := "ack"

		if result, err := ExecuteAction(store, method, arguments); err != nil {
			t.Error("should not have thrown error", err)
		} else if result != expectedResult {
			t.Error("incorrect result returned,", expectedResult, "but got", result)
		}
	})
	t.Run("delete by key not in repo should return ack (no-opp)", func(t *testing.T) {
		method := globals.RequestMethodDelete
		arguments := []string{ key }
		expectedResult := "ack"

		if result, err := ExecuteAction(store, method, arguments); err != nil {
			t.Error("should not have thrown error", err)
		}else if result != expectedResult {
			t.Error("incorrect result returned,", expectedResult, "but got", result)
		}
	})
	t.Run("get by key not in repo should return nil", func(t *testing.T) {
		method := globals.RequestMethodGet
		arguments := []string{ key }
		expectedResult := "nil"

		if result, err := ExecuteAction(store, method, arguments); err == nil {
			t.Error("should have thrown error")
		} else if err != repository.ErrKeyNotFound {
			t.Error(`should have thrown "ErrKeyNotFound" error, but got`, err)
		} else if result != expectedResult {
			t.Error("incorrect result returned,", expectedResult, "but got", result)
		}
	})
}


