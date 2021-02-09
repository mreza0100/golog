package golog

import (
	"fmt"
	"sync"

	"github.com/mreza0100/golog/colors"
	"github.com/mreza0100/golog/helpers"
	wr "github.com/mreza0100/golog/writer"
)

type dataT struct {
	logPath string
	addLog  []interface{}
	hooks   []func(lgr *Core) interface{}
	color   string
	mu      *sync.Mutex
}

type Core struct {
	IsDebugMode bool
	WR          wr.Writer
	Debug       *Core

	data *dataT
}

func (lgr Core) Copy() *Core {
	newData := &dataT{
		logPath: lgr.data.logPath,
		addLog:  lgr.data.addLog,
		hooks:   lgr.data.hooks,
		color:   colors.ColorWhite,
		mu:      &sync.Mutex{},
	}

	lgr.data = newData

	return &lgr
}

func (lgr *Core) With(add ...interface{}) *Core {
	newLogger := lgr.Copy()

	newLogger.data.addLog = helpers.Combine(lgr.data.addLog, add)

	return newLogger
}

func (lgr *Core) AddHook(fn ...func(lgr *Core) interface{}) *Core {
	lgr.data.mu.Lock()
	lgr.data.hooks = append(lgr.data.hooks, fn...)
	lgr.data.mu.Unlock()

	return lgr
}

func (lgr *Core) Log(msgs ...interface{}) *Core {
	lgr.data.mu.Lock()
	if !lgr.IsDebugMode {
		return lgr
	}
	msgs = lgr.getFullMsgs(msgs...)
	lgr.data.mu.Unlock()

	fmt.Print(lgr.data.color)
	for _, i := range msgs {
		fmt.Print(i)
	}

	wrErr := lgr.WR.Write(msgs...)
	if wrErr != nil {
		fmt.Println("internal error in `golog/writer`")
		fmt.Println("error: ", wrErr)
	}

	return lgr
}

func (lgr *Core) getFullMsgs(msgs ...interface{}) []interface{} {
	msgs = helpers.Combine(lgr.data.addLog, lgr.getHooksVals(), msgs)

	msgs = append(msgs, "\n")

	return msgs
}

func (lgr *Core) getHooksVals() []interface{} {
	result := make([]interface{}, len(lgr.data.hooks))

	for idx, hook := range lgr.data.hooks {
		result[idx] = hook(lgr)
	}

	return result
}

func (lgr *Core) setColor(color string) {
	lgr.data.mu.Lock()
	lgr.data.color = color
	lgr.data.mu.Unlock()
}

func (lgr *Core) colorLog(color string, msgs ...interface{}) *Core {
	if !lgr.IsDebugMode {
		fmt.Println(22)
		return lgr
	}

	lgr.setColor(color)
	defer lgr.setColor(colors.ColorWhite)

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
