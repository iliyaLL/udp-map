package kvstorage

import (
	"strings"
	"sync"
	"time"
)

// expiry time.Time can be used to set an expiration time for the key-value pair.
// Example: SET key value PX 10000
// The PX option allows you to specify the expiration time in milliseconds.
type storageItem struct {
	value  string
	expiry time.Time
}

var (
	storage = make(map[string]storageItem)
	mu      sync.RWMutex
)

// SET Foo Bar
// SET foo bar px 10000
func Set(command string) string {
	parts := strings.Fields(command)
	if len(parts) < 3 {
		return "(error) ERR wrong number of arguments for 'SET' command"
	}

	mu.Lock()
	defer mu.Unlock()

	if len(parts) >= 5 && strings.ToUpper(parts[len(parts)-2]) == "PX" {
		// Parse the duration in milliseconds
		expirationTime, err := time.ParseDuration(parts[len(parts)-1] + "ms")
		if err != nil {
			return "(error) ERR invalid PX duration"
		}

		key, value := parts[1], strings.Join(parts[2:(len(parts)-2)], " ")
		expiryTime := time.Now().Add(expirationTime)
		storage[key] = storageItem{
			value:  value,
			expiry: expiryTime,
		}
	} else {
		key, value := parts[1], strings.Join(parts[2:], " ")
		storage[key] = storageItem{
			value:  value,
			expiry: time.Time{},
		}
	}

	return "OK"
}

// SET Foo Bar
// GET Foo
// Bar
func Get(command string) string {
	parts := strings.Fields(command)
	if len(parts) != 2 {
		return "(error) ERR wrong number of arguments for 'GET' command"
	}

	mu.RLock()
	defer mu.RUnlock()
	key := parts[1]
	if item, ok := storage[key]; ok {
		if !item.expiry.IsZero() && time.Now().After(item.expiry) {
			delete(storage, key)
			return "(nil)"
		}
		return item.value
	}

	return "(nil)"
}
