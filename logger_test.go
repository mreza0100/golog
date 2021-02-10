package golog_test

import (
	"testing"

	"github.com/mreza0100/golog"
)

var lgr = golog.New(golog.InitOpns{
	WithTime:     true,
	LogPath:      "./log/out.log",
	Name:         "test",
	DebugMode:    true,
	ClearLogFile: true,
})

func TestCopy(t *testing.T) {
	lgrCopy := lgr.Copy()
	lgrCopy.IsDebugMode = false
	lgrCopy.Data.LogPath = "./fakePath/mamad.log"

	t.Run("1", func(t *testing.T) {
		if lgrCopy == lgr {
			t.Error("Instance was the same\n")
			t.Error("'lgrCopy == log' was 'true'")
			t.FailNow()
		}
	})

	t.Run("2", func(t *testing.T) {
		if lgr.IsDebugMode == false {
			t.Error("Instance was the same\n")
			t.Error("'lgrCopy.IsDebugMode == log.IsDebugMode' was 'true'")
			t.FailNow()
		}
	})

	t.Run("Deeper", func(t *testing.T) {
		if lgrCopy.Data.LogPath == lgr.Data.LogPath || lgrCopy.Data == lgr.Data {
			t.Error("Instance was the same\n")
			t.Error("'lgrCopy.IsDebugMode == log.IsDebugMode' was 'true'")
			t.FailNow()
		}
	})
}

func TestLogger(t *testing.T) {
	newLgr := lgr.With("hamishe ", " mamad")
	newLgr.Log("second mamad")

}

func TestColor() {
	lgr.GreenLog("mamad is here and must be green")
	lgr.Log("this is from the other mamad and should be in default color")
}

func TestDebugMode(t *testing.T) {
	lgr.Debug.IsDebugMode = true

	lgr.Debug.Log("you should be able to see this")

	lgr.Debug.IsDebugMode = false
	lgr.Debug.Log("___you should not be able to see this___")
}
