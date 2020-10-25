package main

import (
	"github.com/picatz/roku"
	"log"
	"os"
)

func main() {
	devices, err := roku.Find(roku.DefaultWaitTime)
	if err != nil {
		log.Fatal(err)
	}

	if len(devices) < 1 {
		log.Fatalln("no device found")
	}

	devices[0].LaunchApp("tvinput.hdmi"+os.Args[1], nil)
}
