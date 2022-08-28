package main

import (
	"context"
	"os"

	"github.com/CallumKerrEdwards/loggerrific"

	"github.com/CallumKerrEdwards/library-api/internal/adapters/logrus"
	bookService "github.com/CallumKerrEdwards/library-api/internal/book/service"
	"github.com/CallumKerrEdwards/library-api/internal/db"
	transportHttp "github.com/CallumKerrEdwards/library-api/internal/transport/http"
)

// Run - sets the application.
func Run(logger loggerrific.Logger) error {
	logger.Infoln("Setting up Library API")

	database, err := db.NewDatabase(context.Background(), logger)
	if err != nil {
		logger.Errorln("failed to connect to the database")
		return err
	}

	err = database.Ping(context.Background())
	if err != nil {
		logger.Errorln("failed to ping the databas e")
		return err
	}

	logger.Infoln("Successfully connected to database")

	mainBookService := bookService.NewService(database, logger)

	httpHandler := transportHttp.NewHandler(mainBookService, logger)
	if err := httpHandler.Serve(); err != nil {
		logger.WithError(err).Errorln("Server error")
		return err
	}

	return nil
}

func main() {
	logger := logrus.NewLogger()
	logger.SetLevelDebug()

	if err := Run(logger); err != nil {
		logger.WithError(err).Errorln("Error starting up our Library server")
		os.Exit(1)
	}
}
