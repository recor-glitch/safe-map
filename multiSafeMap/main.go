package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	mux sync.RWMutex
	m   map[string]any
}

type multiSafeMap []*SafeMap

func NewMultiSafeMap(cap int) *multiSafeMap {
	myMultiSafeMap := make(multiSafeMap, cap)
	for i := range myMultiSafeMap {
		myMultiSafeMap[i] = &SafeMap{
			m: make(map[string]any),
		}
	}
	return &myMultiSafeMap
}

func (m multiSafeMap) Set(index int, key string, value any) {
	myMap := m[index]
	myMap.mux.Lock()
	defer myMap.mux.Unlock()
	myMap.m[key] = value
}

func main() {

	// Initialize the multi safe map
	multiSafeMap := NewMultiSafeMap(5)
	fmt.Printf("My safe map: %v\n", multiSafeMap)

	// SET VALUE TO THE MULTI SAFE MAP IN PARTICULAR INDEX

}
