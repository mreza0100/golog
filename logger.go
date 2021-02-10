package golog

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/mreza0100/golog/colors"
	"github.com/mreza0100/golog/helpers"
	wr "github.com/mreza0100/golog/writer"
)

type dataT struct {
	LogPath string
	AddLog  []interface{}
	Hooks   []func(lgr *Core) interface{}
	mu      *sync.Mutex
}

type Core struct {
	IsDebugMode bool
	WR          wr.Writer
	Debug       *Core

	Data *dataT
}

func (lgr *Core) Copy() *Core {
	newLogger := *lgr
	newData := *lgr.Data

	newLogger.Data = &newData

	return &newLogger
}

func (lgr *Core) Fatal(msgs ...interface{}) {
	lgr.YellowLog("Fatal error ", strings.Repeat("-", 25))
	lgr.RedLog(msgs...)
	lgr.YellowLog("Fatal error ", strings.Repeat("-", 25))
	os.Exit(1)
}

func (lgr *Core) With(add ...interface{}) *Core {
	newLogger := lgr.Copy()

	newLogger.Data.AddLog = helpers.Combine(lgr.Data.AddLog, add)

	return newLogger
}

func (lgr *Core) AddHook(fn ...func(lgr *Core) interface{}) *Core {
	lgr.Data.mu.Lock()
	lgr.Data.Hooks = append(lgr.Data.Hooks, fn...)
	lgr.Data.mu.Unlock()

	return lgr
}

func (lgr *Core) Log(msgs ...interface{}) *Core {
	{
		lgr.Data.mu.Lock()
		if !lgr.IsDebugMode {
			return lgr
		}
		msgs = lgr.getFullMsgs(msgs...)
		lgr.Data.mu.Unlock()
	}

	fmt.Print(msgs...)

	wrErr := lgr.WR.Write(msgs...)
	if wrErr != nil {
		fmt.Println("internal error in `golog/writer`")
		fmt.Println("error: ", wrErr)
	}

	return lgr
}

func (lgr *Core) getFullMsgs(msgs ...interface{}) []interface{} {
	msgs = helpers.Combine(lgr.Data.AddLog, lgr.getHooksVals(), msgs)
	msgs = append(msgs, "\n")

	return msgs
}

func (lgr *Core) getHooksVals() []interface{} {
	result := make([]interface{}, len(lgr.Data.Hooks))

	for idx, hook := range lgr.Data.Hooks {
		result[idx] = hook(lgr)
	}

	return result
}

func (lgr *Core) setColor(color string) {
	fmt.Print(color)
}

func (lgr *Core) unSet() {
	fmt.Printf("%s[%dm", "\x1b", 0)
}

func (lgr *Core) colorLog(color string, msgs ...interface{}) *Core {
	if !lgr.IsDebugMode {
		return lgr
	}

	lgr.setColor(color)
	defer lgr.unSet()

	return lgr.Log(msgs...)
}

func (lgr *Core) RedLog(msgs ...interface{}) *Core {
	return lgr.colorLog(colors.ColorRed, msgs...)
}
func (lgr *Core) YellowLog(msgs ...interface{}) *Core {
	return lgr.colorLog(colors.ColorYellow, msgs...)
}
func (lgr *Core) GreenLog(msgs ...interface{}) *Core {
	return lgr.colorLog(colors.ColorGreen, msgs...)
}
func (lgr *Core) BlueLog(msgs ...interface{}) *Core {
	return lgr.colorLog(colors.ColorBlue, msgs...)
}
func (lgr *Core) PurpleLog(msgs ...interface{}) *Core {
	return lgr.colorLog(colors.ColorPurple, msgs...)
}
func (lgr *Core) CyanLog(msgs ...interface{}) *Core {
	return lgr.colorLog(colors.ColorCyan, msgs...)
}
