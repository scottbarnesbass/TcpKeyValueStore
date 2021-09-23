package repository

import (
	"sync"
)

type KvStore struct {
	store map[Key]Value
	mutex sync.RWMutex
}

func CreateNewStore() *KvStore {
	return &KvStore{ store: make(map[Key]Value) }
}

func (kvs *KvStore) Put(key Key, value Value) error {
	if kvs.store == nil {
		return ErrNilStore
	}
	kvs.mutex.Lock()
	defer kvs.mutex.Unlock()
	kvs.store[key] = value
	return nil
}

func (kvs *KvStore) Get(key Key) (Value, error) {
	kvs.mutex.Lock()
	defer kvs.mutex.Unlock()
	return kvs.get(key)
}

func (kvs *KvStore) get(key Key) (Value, error){
	if data, ok := kvs.store[key]; !ok {
		return nil, ErrKeyNotFound
	} else {
		return data, nil
	}
}

func (kvs *KvStore) Delete(key Key) error {
	kvs.mutex.Lock()
	defer kvs.mutex.Unlock()
	if _, err := kvs.get(key); err != nil {
		return err
	} else {
		delete(kvs.store, key)
		return nil
	}
}
