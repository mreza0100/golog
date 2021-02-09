package golog_test

import (
	"fmt"
	"testing"

	"github.com/mreza0100/golog"
)

var lgr = golog.New(golog.InitOprions{
	WithTime:     true,
	LogPath:      "./log/out.log",
	Name:         "test",
	DebugMode:    false,
	ClearLogFile: true,
})

func TestCopy(t *testing.T) {
	logCopy := lgr.Copy()
	if logCopy == lgr {
		fmt.Println("Instance was the same")
		panic("'logCopy == log' was 'true'")
	}
}

func TestLogger(t *testing.T) {
	newLgr := lgr.With("hamishe", "mamad")
	newLgr.Log("second mamad")

}

func TestStructLogger(t *testing.T) {
	lgr.Log(struct {
		a int
		b string
	}{
		a: 2143124,
		b: "oawodfihaiorhwf",
	})
}

func TestDebugMode(t *testing.T) {
	lgr.Debug.GreenLog("mamad is here")
}
