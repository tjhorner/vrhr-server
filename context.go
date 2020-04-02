package main

import (
	"os"
	"time"
)

type State struct {
	CurrentHeartRate int       `json:"current_heart_rate"`
	CurrentAccuracy  int       `json:"current_accuracy"`
	DataReceivedAt   time.Time `json:"data_received_at"`
}

type Context struct {
	State      State
	WriteToCsv bool
	CsvFile    *os.File
}
