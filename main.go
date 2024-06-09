package main

import "sync"

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

func main() {

	// CREATE A NEW INSTANCE OF SAFE MAP
	safeMap := newSafeMap()


	// SETTING VALUES TO THE SAFE MAP
	safeMap.Set("key1", "value1")
}
