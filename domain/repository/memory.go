package domain

import (
	"fmt"
	"log"
)

type IMemoryRepository interface {
	Get(key string) (string, error)
	Set(key, value string) error
}

type MemoryDB struct {
	DB map[string]string
}

func NewInMemoryRepository() MemoryDB {
	return MemoryDB{
		DB: make(map[string]string),
	}
}

func (m MemoryDB) Get(key string) (string, error) {
	log.Printf("m: %+v", m)
	value, ok := m.DB[key]
	if !ok {
		return "", fmt.Errorf("key %s not found", key)
	}
	return value, nil
}

func (m MemoryDB) Set(key, value string) error {
	m.DB[key] = value
	return nil
}
