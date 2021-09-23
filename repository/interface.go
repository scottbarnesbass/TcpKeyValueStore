package repository

import (
	"errors"
	"time"
)

var (
	ErrKeyNotFound = errors.New("key not found")
	ErrNilStore = errors.New("store has not been initialised correctly, nil store")
)

type Key string

type Value interface {
	UpdatedNow()
}

type Data struct {
	Content string `json:"value"`
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

type Repo interface {
	Put(key Key, value Value) error
	Get(key Key) (Value, error)
	Delete(key Key) error
}