package main

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

func initLog(logfile string) *logrus.Logger {
	log := logrus.New()
	file, err := os.OpenFile(logfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		fmt.Println("While Opening log file for writing: " + err.Error())
	}
	log.Out = file
	return log
}

func loadConfiguration(file string) Configuration {
	var config Configuration
	// Read entire file content, giving us little control but
	// making it very simple. No need to close the file.
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("Unable to read configuration file:\n" + err.Error())
	}

	err2 := json.Unmarshal(content, &config)
	if err2 != nil {
		fmt.Println("Error deserialising configuration JSON:")
		fmt.Println(err2)
	}
	return config
}

func main() {
	// Load configuration
	config := loadConfiguration("/data/code/transit/test-config.json")
	log := initLog(config.Logfile)
	log.Info("Starting Transit...")
	fmt.Println("Exiting.")
}

type Configuration struct {
	Logfile string
	Paths   []Path
}

type TransitType string

const (
	TRAM  TransitType = "TRAM"
	TRAIN TransitType = "TRAIN"
	BUS   TransitType = "BUS"
)

type Direction string

const (
	CITY Direction = "CITY"
	OUT  Direction = "OUT"
)

type Stop struct {
	Id    string // `json:"id"`
	Label string // `json:"label"`
}

type Path struct {
	Id        string      // `json:"id"`
	Label     string      // `json:"label"`
	Type      TransitType `json:"Type"`
	Direction Direction   // `json:"direction"`
	Routes    []uint8     // `json:"routes"`
	Line      string      // `json:"line"`
	Stops     []Stop      //`json:"stops"`
}
