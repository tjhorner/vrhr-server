package main

import (
	"os"
	"time"
)

type State struct {
	CurrentHeartRate int       `json:"current_heart_rate"`
	CurrentAccuracy  int       `json:"current_accuracy"`
	ReportedByDevice string    `json:"reported_by_device"`
	DataReceivedAt   time.Time `json:"data_received_at"`
}

func NewState() State {
	return State{
		CurrentAccuracy: -1,
	}
}

type Context struct {
	State      State
	WriteToCsv bool
	CsvFile    *os.File
}
