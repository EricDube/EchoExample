package customMiddleware

import (
	"os"

	"github.com/google/logger"
	"github.com/labstack/echo"
)

type (
	Logging struct {
		Logger       *logger.Logger
	}
)

func NewLogger() *Logging {
	//Get current working directory to dump logs to
	path, err := os.Getwd()
	if err != nil {
		return nil
	}

	//Open a file to write logs to
	lf, err := os.OpenFile(path+"/log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		logger.Fatalf("Failed to open log file: %v", err)
	}
	//Return a new Logging struct with a Logger already Init-ed
	return &Logging{
		Logger:   logger.Init("LoggerExample", true, true, lf),
	}
}

// Process is the customMiddleware function.
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
