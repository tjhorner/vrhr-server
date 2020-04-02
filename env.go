package main

import "os"

// GetEnv gets an environment variable with a fallback
func GetEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}

var envListenAddr = GetEnv("VRHR_LISTEN_ADDRESS", ":8000")
