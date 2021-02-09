package golog

import (
	"fmt"
	"sync"

	"github.com/mreza0100/golog/helpers"
	wr "github.com/mreza0100/golog/writer"
)

type Core struct {
	mu          *sync.Mutex
	color       string
	isDebugMode bool
	LogPath     string
	Add         []interface{}
	WR          wr.Writer
	Hooks       []func(lgr *Core) interface{}
	Debug       *Core
}

func (lgr Core) Copy() *Core {
	return &lgr
}

func (lgr *Core) With(add ...interface{}) *Core {
	newLogger := lgr.Copy()

	newLogger.Add = helpers.Combine(lgr.Add, add)
	newLogger.mu = &sync.Mutex{}

	return newLogger
}

func (lgr *Core) AddHook(fn ...func(lgr *Core) interface{}) *Core {
	lgr.mu.Lock()
	lgr.Hooks = append(lgr.Hooks, fn...)
	lgr.mu.Unlock()

	return lgr
}

func (lgr *Core) Log(msgs ...interface{}) *Core {
	lgr.mu.Lock()
	msgs = lgr.getFullMsgs(msgs...)
	lgr.mu.Unlock()

	fmt.Print(lgr.color)
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
	msgs = helpers.Combine(lgr.Add, lgr.getHooksVals(), msgs)
	msgs = append(msgs, "\n")

	return msgs
}

func (lgr *Core) getHooksVals() []interface{} {
	result := make([]interface{}, len(lgr.Hooks))

	for idx, hook := range lgr.Hooks {
		result[idx] = hook(lgr)
	}

	return result
}

func (lgr *Core) getColor() string {
	lgr.mu.Lock()
	defer lgr.mu.Unlock()
	return lgr.color
}

func (lgr *Core) setColor(color string) {
	lgr.mu.Lock()
	defer lgr.mu.Unlock()
	lgr.color = color
}

func (lgr *Core) colorLog(color string, msgs ...interface{}) *Core {
	lgr.setColor(color)
	defer lgr.setColor(helpers.ColorWhite)

	return lgr.Log(msgs...)
}

func (lgr *Core) RedLog(msgs ...interface{}) *Core {
	return lgr.colorLog(helpers.ColorRed, msgs...)
}
func (lgr *Core) YellowLog(msgs ...interface{}) *Core {
	return lgr.colorLog(helpers.ColorYellow, msgs...)
}
func (lgr *Core) GreenLog(msgs ...interface{}) *Core {
	return lgr.colorLog(helpers.ColorGreen, msgs...)
}
func (lgr *Core) BlueLog(msgs ...interface{}) *Core {
	return lgr.colorLog(helpers.ColorBlue, msgs...)
}
func (lgr *Core) PurpleLog(msgs ...interface{}) *Core {
	return lgr.colorLog(helpers.ColorPurple, msgs...)
}
func (lgr *Core) CyanLog(msgs ...interface{}) *Core {
	return lgr.colorLog(helpers.ColorCyan, msgs...)
}
