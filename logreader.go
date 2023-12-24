package main

import (
	"bufio"
	"fmt"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
)

func tickTime() {
	for {
		tick <- true
		time.Sleep(100 * time.Millisecond)
	}
}

func readLog() {
	// home, err := os.UserHomeDir()

	go tickTime()
	readlog := false
	// go addWatcher(home + "\\AppData\\LocalLow\\Boundless Dynamics, LLC\\VTOLVR\\Player.log")go addWatcher(home + "\\AppData\\LocalLow\\Boundless Dynamics, LLC\\VTOLVR\\Player.log")
	for {
		<-tick
		logFile := readLogFile()
		scanLog := bufio.NewScanner(strings.NewReader(logFile))
		var logLinesTmp []string
		for scanLog.Scan() {
			logLinesTmp = append(logLinesTmp, scanLog.Text())
		}

		for y, x := range logLinesTmp {
			if y > (len(logLines) - 1) {
				if readlog {
					logHandler(x)
				}

			}
		}
		logLines = logLinesTmp
		readlog = true
	}
}

var logLines []string
var currentTrack int

func logHandler(newline string) bool {
	if strings.Contains(newline, "Playing song:") {
		splitline := strings.SplitAfter(newline, "clip length")
		if strings.Contains(splitline[0], "0") {
			track0(currentTrack)
			fmt.Println("Spike bearing 0, angels " + fmt.Sprint(currentTrack) + ".")
		} else if strings.Contains(splitline[0], "1") {
			track1(currentTrack)
			fmt.Println("Spike bearing 0, angels " + fmt.Sprint(currentTrack) + ".")
		} else if strings.Contains(splitline[0], "2") {
			track2(currentTrack)
			fmt.Println("Spike bearing 2, angels " + fmt.Sprint(currentTrack) + ".")
		}
		return true
	} else if strings.Contains(newline, "FlightLogger:") && strings.Contains(newline, "has spawned.") {
		currentTrack = 0
	}
	return false
}

func track2(track int) {
	switch track {
	case 0:
		Track2{}.TrackZero()
	case 1:
		Track2{}.TrackOne()
	case 2:
		Track2{}.TrackTwo()
	}
}

func track1(track int) {
	switch track {
	case 0:
		Track1{}.TrackZero()
	case 1:
		Track1{}.TrackOne()
	case 2:
		Track1{}.TrackTwo()
	}
}

func track0(track int) {
	switch track {
	case 0:
		Track0{}.TrackZero()
	case 1:
		Track0{}.TrackOne()
	case 2:
		Track0{}.TrackTwo()
	}
}

func (Track0) TrackZero() {
	robotgo.KeyTap(robotgo.AudioPlay)
	currentTrack = 0
}
func (Track0) TrackOne() {
	robotgo.KeyTap(robotgo.AudioNext)
	currentTrack = 0
}
func (Track0) TrackTwo() {
	robotgo.KeyTap(robotgo.AudioPrev)
	currentTrack = 0
}

type Track0 struct{}

func (Track1) TrackZero() {
	robotgo.KeyTap(robotgo.AudioPrev)
	currentTrack = 1
}

func (Track1) TrackOne() {
	robotgo.KeyTap(robotgo.AudioPlay)
	currentTrack = 1
}

func (Track1) TrackTwo() {
	robotgo.KeyTap(robotgo.AudioNext)
	currentTrack = 1
}

type Track1 struct{}

func (Track2) TrackZero() {
	robotgo.KeyTap(robotgo.AudioNext)
	currentTrack = 2
}

func (Track2) TrackOne() {
	robotgo.KeyTap(robotgo.AudioPrev)
	currentTrack = 2
}
func (Track2) TrackTwo() {
	robotgo.KeyTap(robotgo.AudioPlay)
}

type Track2 struct{}
