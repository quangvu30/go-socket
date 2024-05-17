package logger

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	log *logrus.Logger
}

var Log *Logger

func NewLogger() (*Logger, error) {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{})
	log.SetLevel(logrus.DebugLevel)
	// Get the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	file, err := os.OpenFile(path.Join(currentDir, fmt.Sprintf("/logs/%s.log", time.Now().Format("02-01-2006"))), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
	Log = &Logger{
		log: log,
	}
	return Log, nil
}

func (l *Logger) Info(msg ...interface{}) {
	fmt.Printf("[INFO][%s] %v\n", time.Now().UTC().Format("2006-01-02 15:04:05"), msg)
	l.log.Info(msg...)
}

func (l *Logger) Error(msg ...interface{}) {
	fmt.Printf("[ERROR][%s] %v\n", time.Now().UTC().Format("2006-01-02 15:04:05"), msg)
	l.log.Error(msg...)
}

func (l *Logger) Warn(msg ...interface{}) {
	fmt.Printf("[WARN][%s] %v\n", time.Now().UTC().Format("2006-01-02 15:04:05"), msg)
	l.log.Warn(msg...)
}
