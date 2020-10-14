package logger

import (
	"os"

	"github.com/google/logger"
	"github.com/labstack/echo"
)

//Declare a global logger so anyone can log using this one
var Logger *Logging

type (
	Logging struct {
		Logger       *logger.Logger
	}
)

// NewLogger
// @Summary Creates a new Logger
// @Description Creates a logger to be used to save logs
// @ID new-logger
func NewLogger() error {
	//Get current working directory to dump logs to
	path, err := os.Getwd()
	if err != nil {
		return err
	}

	//Open a file to write logs to
	lf, err := os.OpenFile(path+"/log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		logger.Fatalf("Failed to open log file: %v", err)
	}
	//Return a new Logging struct with a Logger already Init-ed
	Logger = &Logging{
		Logger:   logger.Init("LoggerExample", true, true, lf),
	}

	// Return nil to make sure the caller doesnt think theres an error
	return nil
}

// Process
// @Summary Handles an echo middleware to add logs
// @Description Processes Echo requests using the middleware and saves the logs using the created logger
// @ID logger-process
func (t *Logging) Process(next echo.HandlerFunc) echo.HandlerFunc {
	//Return a function of type echo Context to be used by the middleware
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			c.Error(err)
		}
		//Use logging framework to log to console and file
		t.Logger.Info("I'm about to do something!")
		return nil
	}
}
