package main

import (
	"log"
	"os"
)

var tick = make(chan bool)

func readLogFile() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Panic(err)
	}
	file, err := os.ReadFile(home + "\\AppData\\LocalLow\\Boundless Dynamics, LLC\\VTOLVR\\Player.log")
	if err != nil {
		log.Panic(err)
	}

	return string(file)
}
