package websocketServices

import "sync"

var resettingRegistry = struct {
	mu sync.Mutex
	m  map[int]bool
}{m: make(map[int]bool)}

func MarkUserResetting(userID int) {
	resettingRegistry.mu.Lock()
	defer resettingRegistry.mu.Unlock()
	resettingRegistry.m[userID] = true
}

func UnmarkUserResetting(userID int) {
	resettingRegistry.mu.Lock()
	defer resettingRegistry.mu.Unlock()
	delete(resettingRegistry.m, userID)
}

func IsUserResetting(userID int) bool {
	resettingRegistry.mu.Lock()
	defer resettingRegistry.mu.Unlock()
	return resettingRegistry.m[userID]
}