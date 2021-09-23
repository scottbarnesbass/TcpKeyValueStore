package storeRepo_test

import (
	"TcpKeyValueStore/storeRepo"
	"fmt"
	"testing"
)

func TestOpenAndCloseKvStore(t *testing.T) {
	t.Run("TestOpeningNewRepo", func(t *testing.T) {
		if err := storeRepo.Open(); err != nil {
			t.Error("should not error", err)
		}
	})
	t.Run("TestClosingOpenRepo", func(t *testing.T) {
		if err := storeRepo.Close(); err != nil {
			t.Error("should not error", err)
		}
	})
}

func TestCloseStoreThatIsNotOpen(t *testing.T) {
	if err := storeRepo.Close(); err == nil {
		t.Error("should have thrown an error")
	}
}

func TestOpenAStoreThatIsAlreadyOpen(t *testing.T) {
	_ = storeRepo.Open()
	defer storeRepo.Close()
	if err := storeRepo.Open(); err == nil {
		t.Error("should have thrown an error")
	}
}

func TestPutGetAndDeleteValue(t *testing.T) {

	data := storeRepo.NewData("testValue")
	t.Run("PutBeforeStoreOpenShouldThrowError", func(t *testing.T) {
		if err := storeRepo.Put("1234", &data); err == nil {
			t.Error("should have thrown an error")
		}
	})

	storeRepo.Open()
	defer storeRepo.Close()

	t.Run("PutToOpenStoreShouldNotThrowError", func(t *testing.T) {
		if err := storeRepo.Put("1234", &data); err != nil {
			t.Error("should not have thrown an error", err)
		}
	})
	t.Run("GetValue", func(t *testing.T) {
		if val, err := storeRepo.Get("1234"); err != nil {
			t.Error("should not have thrown an error", err)
		} else {
			data := val.(*storeRepo.Data)
			if data.Content != "testValue" {
				t.Error("incorrect content returned", data.Content)

			}
		}
	})
	t.Run("DeleteValue", func(t *testing.T) {
		if err := storeRepo.Delete("1234"); err != nil {
			t.Error("should not have thrown an error", err)
		}
		if _, err := storeRepo.Get("1234"); err != storeRepo.ErrKeyNotFound {
			t.Error("should have thrown an ErrKeyNotFound error")
		}
		if err := storeRepo.Delete("1234"); err != storeRepo.ErrKeyNotFound {
			t.Error("should have thrown an ErrKeyNotFound error")
		}
	})
}

func TestGetAll(t *testing.T) {
	_ = storeRepo.Open()
	defer storeRepo.Close()

	if res, _ := storeRepo.GetAll(); len(res) > 0 {
		t.Error("should return empty map")
	}

	for i := 0; i < 4; i++ {
		data := storeRepo.NewData(fmt.Sprint("testValue", i))
		t.Run("PutBeforeStoreOpenShouldThrowError", func(t *testing.T) {
			key := storeRepo.Key(fmt.Sprint(1234 + i))
			if err := storeRepo.Put(key, &data); err != nil {
				t.Error("should not have thrown an error", err)
			}
		})
	}
	if res, _ := storeRepo.GetAll(); len(res) != 4  {
		fmt.Println(res)

		t.Error("should have returned a map containing 3 entries. Returned ", len(res))
	}
}
