package store_test

import (
	"TcpKeyValueStore/storeRepo"
	"fmt"
	"testing"
)

func TestOpenAndCloseKvStore(t *testing.T) {
	t.Run("TestOpeningNewRepo", func(t *testing.T) {
		if err := store.Open(); err != nil {
			t.Error("should not error", err)
		}
	})
	t.Run("TestClosingOpenRepo", func(t *testing.T) {
		if err := store.Close(); err != nil {
			t.Error("should not error", err)
		}
	})
}

func TestCloseStoreThatIsNotOpen(t *testing.T) {
	if err := store.Close(); err == nil {
		t.Error("should have thrown an error")
	}
}

func TestOpenAStoreThatIsAlreadyOpen(t *testing.T) {
	_ = store.Open()
	defer store.Close()
	if err := store.Open(); err == nil {
		t.Error("should have thrown an error")
	}
}

func TestPutGetAndDeleteValue(t *testing.T) {

	data := store.NewData("testValue")
	t.Run("PutBeforeStoreOpenShouldThrowError", func(t *testing.T) {
		if err := store.Put("1234", &data); err == nil {
			t.Error("should have thrown an error")
		}
	})

	store.Open()
	defer store.Close()

	t.Run("PutToOpenStoreShouldNotThrowError", func(t *testing.T) {
		if err := store.Put("1234", &data); err != nil {
			t.Error("should not have thrown an error", err)
		}
	})
	t.Run("GetValue", func(t *testing.T) {
		if val, err := store.Get("1234"); err != nil {
			t.Error("should not have thrown an error", err)
		} else {
			data := val.(*store.Data)
			if data.Content != "testValue" {
				t.Error("incorrect content returned", data.Content)

			}
		}
	})
	t.Run("DeleteValue", func(t *testing.T) {
		if err := store.Delete("1234"); err != nil {
			t.Error("should not have thrown an error", err)
		}
		if _, err := store.Get("1234"); err != store.ErrKeyNotFound {
			t.Error("should have thrown an ErrKeyNotFound error")
		}
		if err := store.Delete("1234"); err != store.ErrKeyNotFound {
			t.Error("should have thrown an ErrKeyNotFound error")
		}
	})
}

func TestGetAll(t *testing.T) {
	_ = store.Open()
	defer store.Close()

	if res, _ := store.GetAll(); len(res) > 0 {
		t.Error("should return empty map")
	}

	for i := 0; i < 4; i++ {
		data := store.NewData(fmt.Sprint("testValue", i))
		t.Run("PutBeforeStoreOpenShouldThrowError", func(t *testing.T) {
			key := store.Key(fmt.Sprint(1234 + i))
			if err := store.Put(key, &data); err != nil {
				t.Error("should not have thrown an error", err)
			}
		})
	}
	if res, _ := store.GetAll(); len(res) != 4  {
		fmt.Println(res)

		t.Error("should have returned a map containing 3 entries. Returned ", len(res))
	}
}
