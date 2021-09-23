package repository_test

import (
	"TcpKeyValueStore/repository"
	"testing"
)


func TestPutGetAndDeleteValue(t *testing.T) {

	store := repository.CreateNewStore()
	data := repository.NewData("testValue")

	t.Run("PutToStoreShouldNotThrowError", func(t *testing.T) {
		if err := store.Put("1234", &data); err != nil {
			t.Error("should not have thrown an error", err)
		}
	})
	t.Run("GetValue", func(t *testing.T) {
		if val, err := store.Get("1234"); err != nil {
			t.Error("should not have thrown an error", err)
		} else {
			data := val.(*repository.Data)
			if data.Content != "testValue" {
				t.Error("incorrect content returned", data.Content)

			}
		}
	})
	t.Run("DeleteValue", func(t *testing.T) {
		if err := store.Delete("1234"); err != nil {
			t.Error("should not have thrown an error", err)
		}
		if _, err := store.Get("1234"); err != repository.ErrKeyNotFound {
			t.Error("should have thrown an ErrKeyNotFound error")
		}
		if err := store.Delete("1234"); err != repository.ErrKeyNotFound {
			t.Error("should have thrown an ErrKeyNotFound error")
		}
	})
}

