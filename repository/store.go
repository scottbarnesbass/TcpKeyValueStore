package store

import (
	"sync"
	"TcpKeyValueStore/repository"
)

type KvStore struct {
	store map[repository.Key]repository.Value
	mutex sync.RWMutex
}

func CreateNewStore() *KvStore {
	return &KvStore{ store: make(map[repository.Key]repository.Value) }
}

func (kvs *KvStore) Put(key repository.Key, value repository.Value) error {
	if kvs.store == nil {
		return repository.ErrNilStore
	}
	kvs.mutex.Lock()
	defer kvs.mutex.Unlock()
	kvs.store[key] = value
	return nil
}

func (kvs *KvStore) Get(key repository.Key) (repository.Value, error) {
	kvs.mutex.Lock()
	defer kvs.mutex.Unlock()
	return kvs.get(key)
}

func (kvs *KvStore) get(key repository.Key) (repository.Value, error){
	if data, ok := kvs.store[key]; !ok {
		return nil, repository.ErrKeyNotFound
	} else {
		return data, nil
	}
}

func (kvs *KvStore) Delete(key repository.Key) error {
	kvs.mutex.Lock()
	defer kvs.mutex.Unlock()
	if _, err := kvs.get(key); err != nil {
		return err
	} else {
		delete(kvs.store, key)
		return nil
	}
}
