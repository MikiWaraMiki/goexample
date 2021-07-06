package main

import (
	"fmt"
	"sync"
)

type KeyValue struct {
	store map[string]string // map of key value
	mu    sync.RWMutex      // mutex
}

func NewKeyValue() *KeyValue {
	return &KeyValue{store: make(map[string]string)}
}

func (kv *KeyValue) Set(key, val string) {
	kv.mu.Lock()         // Lock
	defer kv.mu.Unlock() // メソッド抜ける際にUnLock

	kv.store[key] = val
}

func (kv *KeyValue) Get(key string) (string, bool) {
	kv.mu.RLock() // read lock
	defer kv.mu.RUnlock()

	val, ok := kv.store[key]

	return val, ok
}

func main() {
	kv := NewKeyValue()

	kv.Set("key", "value")
	value, ok := kv.Get("key")

	if ok {
		fmt.Println(value)
	}
}
