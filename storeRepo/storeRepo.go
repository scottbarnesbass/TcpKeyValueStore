package storeRepo

import (
	"errors"
	"sync"
	"time"
)

var (
	kvStore = KvStore{ store: nil, open: false }
	ErrKeyNotFound      = errors.New("key not found")
	ErrStoreAlreadyOpen = errors.New("store already open")
	ErrStoreNotOpen     = errors.New("store not open")
)

type Key string

type Value interface {
	UpdatedNow()
}

type Data struct {
	Content     string `json:"value"`
	Timestamp time.Time `json:"timestamp"`
	CreatedBy string
}

func (d Data) UpdatedNow() {
	d.Timestamp = time.Now()
}

func NewData(
	newValue string) Data {
	return Data {
		Content: newValue,
		Timestamp: time.Now(),
	}
}

type KvStore struct {
	store map[Key]Value
	open bool
	mutex *sync.RWMutex
}

func (kvs *KvStore) isOpen() bool {
	return kvStore.open
}

func (kvs *KvStore) isClosed() bool {
	return !kvStore.open
}

func (kvs *KvStore) openStore() {
	kvStore.store = make(map[Key]Value)
	kvs.open = true
	kvs.mutex = &sync.RWMutex{}
}

func (kvs *KvStore) closeStore() {
	kvStore.store = nil
	kvs.open = false
	kvs.mutex = nil
}

// API METHODS
func Open() error {
	if kvStore.isOpen() {
		return ErrStoreAlreadyOpen
	} else {
		kvStore.openStore()
		return nil
	}
}

func Close() error {
	if kvStore.isOpen() {
		kvStore.closeStore()
		return nil
	} else {
		return ErrStoreNotOpen
	}
}

func Put(key Key, value Value) error {
	if kvStore.open {
		kvStore.mutex.Lock()
		defer kvStore.mutex.Unlock()
		kvStore.store[key] = value
		return nil
	} else {
		return ErrStoreNotOpen
	}
}

func Get(key Key) (Value, error) {
	if kvStore.isClosed() {
		return nil, ErrStoreNotOpen
	}
	kvStore.mutex.Lock()
	defer kvStore.mutex.Unlock()

	return get(key)
}

func get(key Key) (Value, error){
	if data, ok := kvStore.store[key]; !ok {
		return nil, ErrKeyNotFound
	} else {
		return data, nil
	}
}

func Delete(key Key) error {
	if kvStore.isClosed() {
		return ErrStoreNotOpen
	}
	kvStore.mutex.Lock()
	defer kvStore.mutex.Unlock()
	if _, err := get(key); err != nil {
		return err
	} else {
		delete(kvStore.store, Key(key))
		return nil
	}
}

func GetAll() (map[Key]Value, error) {
	if kvStore.isClosed() {
		return map[Key]Value{}, ErrStoreNotOpen
	}

	kvStore.mutex.Lock()
	defer kvStore.mutex.Unlock()

	return kvStore.store, nil
}