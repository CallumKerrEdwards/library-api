package main

import (
	"context"
	"os"

	"github.com/CallumKerrEdwards/library/server/internal/adapters/logrus"
	bookService "github.com/CallumKerrEdwards/library/server/internal/book/service"
	"github.com/CallumKerrEdwards/library/server/internal/db"
	transportHttp "github.com/CallumKerrEdwards/library/server/internal/transport/http"
	"github.com/CallumKerrEdwards/library/server/pkg/log"
)

// Run - sets up our application
func Run(logger log.Logger) error {
	logger.Infoln("Setting up Library server")

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

	bookService := bookService.NewService(database, logger)

	httpHandler := transportHttp.NewHandler(bookService, logger)
	if err := httpHandler.Serve(); err != nil {
		logger.WithError(err).Errorln("Server error")
		return err
	}

	// bookService.CreateBook(
	// 	context.Background(),
	// 	book.Book{ID: "2ef5a9ec-3464-44a6-97d6-a71625612049",
	// 		Title:  "The Way of Kings",
	// 		Author: "Brandon Sanderson",
	// 		Series: book.Series{
	// 			ID:       "cd1816d9-9e19-44f7-ac8d-d9525dd9f367",
	// 			Title:    "The Stormlight Archive",
	// 			Sequence: 1,
	// 		},
	// 	},
	// )

	// logger.Infoln(bookService.GetBook(context.Background(), "2ef5a9ec-3464-44a6-97d6-a71625612049"))

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
