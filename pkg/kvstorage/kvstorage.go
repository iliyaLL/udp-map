package kvstorage

import (
	"strings"
	"sync"
	"time"
)

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

	if len(parts) > 3 && strings.ToUpper(parts[len(parts)-2]) == "PX" {
		if len(parts) < 5 {
			return "(error) ERR wrong number of arguments for 'SET' command with PX"
		}

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

func Get(command string) string {
	return ""
}
