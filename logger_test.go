package golog_test

import (
	"fmt"
	"testing"

	logger "github.com/mreza0100/golog"
)

var log = logger.New(logger.InitOprions{
	WithTime: true,
	LogPath:  "./log/out.log",
	Name:     "test",
})

func TestCopy(t *testing.T) {
	logCopy := log.Copy()
	if logCopy == log {
		fmt.Println("Instance was the same")
		panic("'logCopy == log' was 'true'")
	}

	logCopy.LogPath = "./log2/all.log"
	if logCopy.LogPath == log.LogPath {
		fmt.Println("Instance was the same")
		panic("'logCopy.LogPath == log.LogPath' was 'true'")
	}

}

func TestLogger(t *testing.T) {

	log = log.With("hamishe", "mamad")
	log.Log("second mamad")

}

func TestStructLogger(t *testing.T) {
	log.Log(struct {
		a int
		b string
	}{
		a: 2143124,
		b: "oawodfihaiorhwf",
	})
}
