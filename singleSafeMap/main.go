package main

import (
	"fmt"
	"sync"
)

type safeMap struct {
	mux sync.RWMutex
	m   map[string]any
}

func newSafeMap() *safeMap {
	return &safeMap{
		m: make(map[string]any),
	}
}

func (sm *safeMap) Get(key string) any {
	sm.mux.RLock()
	defer sm.mux.RUnlock()

	value := sm.m[key]
	return value
}

func (sm *safeMap) Set(key string, value any) {
	sm.mux.Lock()
	defer sm.mux.Unlock()

	sm.m[key] = value
}

func (sm *safeMap) Delete(key string) {
	sm.mux.Lock()
	defer sm.mux.Unlock()

	delete(sm.m, key)
}

func main() {

	// CREATE A NEW INSTANCE OF SAFE MAP
	safeMap := newSafeMap()

	// SETTING VALUES TO THE SAFE MAP
	safeMap.Set("key1", "value1")
	safeMap.Set("key2", []int{1, 2, 3, 4, 5})

	// GETTING THE VALUE AND PRINTING THE VALUE
	result := safeMap.Get("key1")
	result2 := safeMap.Get("key2")
	fmt.Printf("Result of Key 1: %v", result)
	fmt.Printf("Result of Key 2: %+v", result2)

	// DELETE THE VALUE FROM THE MAP
	safeMap.Delete("key1")
}
