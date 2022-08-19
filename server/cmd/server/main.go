package main

import (
	"os"

	"github.com/CallumKerrEdwards/library/server/internal/adapters/logrus"
	"github.com/CallumKerrEdwards/library/server/pkg/log"
)

// Run - sets up our application
func Run(logger log.Logger) error {
	logger.Infoln("Setting up Library server")
	return nil
}

func main() {
	logger := logrus.NewLogger()

	if err := Run(logger); err != nil {
		logger.WithError(err).Errorln("Error starting up our Library server")
		os.Exit(1)
	}
}
