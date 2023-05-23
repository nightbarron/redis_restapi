package helpers

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func InitLogger() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}
