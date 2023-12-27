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
			fmt.Println("Spike bearing 0, angels " + fmt.Sprint(currentTrack) + ".")
			track0(currentTrack)
		} else if strings.Contains(splitline[0], "1") {
			fmt.Println("Spike bearing 1, angels " + fmt.Sprint(currentTrack) + ".")
			track1(currentTrack)
		} else if strings.Contains(splitline[0], "2") {
			fmt.Println("Spike bearing 2, angels " + fmt.Sprint(currentTrack) + ".")
			track2(currentTrack)
		}
		return true
	} else if strings.Contains(newline, "FlightLogger:") && strings.Contains(newline, "has spawned.") {
		currentTrack = 0
		fmt.Println("Splash 1, bearing 0")
	}
	return false
}

func track2(track int) {

	switch track {
	case 0:
		Track2{}.RW()

	case 1:
		Track2{}.FF()

	default:
		Track2{}.Play()
	}

}

func track1(track int) {

	switch track {
	case 0:
		Track1{}.FF()
	case 2:
		Track1{}.RW()

	default:
		Track1{}.Play()
	}

}

func track0(track int) {

	switch track {
	case 1:
		Track0{}.RW()

	case 2:
		Track0{}.FF()

	default:
		Track0{}.Play()
	}

}

func (Track0) Play() {
	robotgo.KeyTap(robotgo.AudioPlay)
	currentTrack = 0
}
func (Track0) RW() {
	robotgo.KeyTap(robotgo.AudioPrev)
	currentTrack = 0
}
func (Track0) FF() {
	robotgo.KeyTap(robotgo.AudioNext)
	currentTrack = 0
}

type Track0 struct{}

func (Track1) Play() {
	robotgo.KeyTap(robotgo.AudioPlay)
	currentTrack = 1
}

func (Track1) FF() {
	robotgo.KeyTap(robotgo.AudioNext)
	currentTrack = 1
}

func (Track1) RW() {
	robotgo.KeyTap(robotgo.AudioPrev)
	currentTrack = 1
}

type Track1 struct{}

func (Track2) Play() {
	robotgo.KeyTap(robotgo.AudioPlay)
	currentTrack = 2
}

func (Track2) FF() {
	robotgo.KeyTap(robotgo.AudioNext)
	currentTrack = 2
}

func (Track2) RW() {
	robotgo.KeyTap(robotgo.AudioPrev)
	currentTrack = 2
}

type Track2 struct{}
