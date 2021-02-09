package golog_test

import (
	"fmt"
	"testing"

	logger "github.com/mreza0100/golog"
)

func TestCopy(t *testing.T) {
	log := logger.New(logger.InitOprions{
		WithTime: false,
		LogPath:  "./log/all.log",
		Name:     "test",
	})
	defer func() {}()
	logCopy := log.Copy()
	if logCopy == log {
		fmt.Println("Instance was the same")
		panic("'logCopy == log' was 'true'")
	}
	logCopy.With("some stupid text: ")

	log.LogPath = "./log2/all.log"
	if logCopy.LogPath == log.LogPath {
		fmt.Println("Instance was the same")
		panic("'logCopy.LogPath == log.LogPath' was 'true'")
	}

}

func TestLogger(t *testing.T) {
	log := logger.New(logger.InitOprions{
		WithTime: true,
		LogPath:  "./logs/out.log",
		Name:     "test",
	})
	extendedLogger := log.With("hamishe mamad")
	extendedLogger.Log("second mamad")

	extendedLogger.Log("third mamad")
}
