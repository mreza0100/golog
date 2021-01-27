package logger_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/mreza0100/logger"
)

func TestCopy(t *testing.T) {
	log := logger.New(logger.InitOprions{
		WithTime: false,
		LogPath:  "./log",
		Name:     "test",
	})
	logCopy := log.Copy()
	if logCopy == log {
		fmt.Println("Instance was the same")
		panic("'logCopy == log' was 'true'")
	}

}

func TestLogger(t *testing.T) {
	log := logger.New(logger.InitOprions{
		WithTime: false,
		LogPath:  "./logs/out.log",
		Name:     "test",
	})
	extendedLogger := log.With("hamishe mamad")
	extendedLogger.Log("second mamad")

	extendedLogger.AddHook(func(logger *logger.Core) interface{} {
		return "time: " + time.Now().String()
	})

	extendedLogger.Log("third mamad")
}
