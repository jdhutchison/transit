package transit

import (
	"github.com/sirupsen/logrus"
	"os"
)

func initLog(logfile string) {
  log := logrus.New()
  file, _ := os.OpenFile(logfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
  log.Out = file
}

func loadConfiguration() Configuration {
	return Configuration{logfile: "/var/log/transit.log"}
}

func main() {
  // Load configuration
  config := loadConfiguration()
  initLog(config.logfile)
}

type Configuration struct {
  logfile string
}


type TransitType string
const (
	TRAM TransitType = "TRAM"
	TRAIN TransitType = "TRAIN"
	BUS TransitType = "Bus"
	
)

type Direction  string
const (
	CITY Direction = "CITY"
	OUT Direction = "OUT"
)

type Stop struct {
	id string
	label string
}

type Way struct {
	id string
	label string
	theType TransitType
	direction Direction
	routes []uint8
	line string
	stops []Stop
}
