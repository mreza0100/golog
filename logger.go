package golog

import (
	"fmt"
	"os"
	"strings"
	"sync"

	stuff "github.com/mreza0100/golog/fun_stuff"
	hp "github.com/mreza0100/golog/helpers"
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
	lgr.InfoLog("Fatal error ", strings.Repeat("-", 25))
	lgr.RedLog(msgs...)
	lgr.InfoLog("Fatal error ", strings.Repeat("-", 25))
	os.Exit(1)
}

func (lgr *Core) With(add ...interface{}) *Core {
	newLogger := lgr.Copy()

	newLogger.Data.AddLog = hp.Combine(lgr.Data.AddLog, add)

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
	msgs = hp.Combine(lgr.Data.AddLog, lgr.getHooksVals(), msgs)
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

func (lgr *Core) DebugPKG(where string, active bool) (func(msgs ...interface{}) func(error) error, func(...interface{}) interface{}) {
	lgr = lgr.Debug.With("\nDebugPKG " + where + ":")
	if !active {
		return func(msgs ...interface{}) func(error) error {
				return func(err error) error {
					if err != nil {
						lgr.RedLog("Error in deactive mode")
						lgr.RedLog("msgs:")
						lgr.RedLog(msgs...)
						lgr.RedLog(err)
					}
					return err
				}
			}, func(result ...interface{}) interface{} {
				return result
			}
	}

	lgr.InfoLog("----start----")
	return func(msgs ...interface{}) func(error) error {
			return func(err error) error {
				if err == nil {
					lgr.BlueLog(hp.Unshift(msgs, "Passed")...)
					return nil
				}
				lgr.BugHunter(append(msgs, "\nError: {{\n\n\t", err, "\n\n}} \n")...)
				return err
			}
		}, func(result ...interface{}) interface{} {
			lgr.SuccessLog("Done! no error :)")
			lgr.SuccessLog("result: ", result)
			return result
		}
}

func (lgr *Core) RedLog(msgs ...interface{}) *Core {
	return lgr.colorLog(stuff.ColorRed, msgs...)
}
func (lgr *Core) InfoLog(msgs ...interface{}) *Core {
	return lgr.colorLog(stuff.ColorYellow, msgs...)
}
func (lgr *Core) GreenLog(msgs ...interface{}) *Core {
	return lgr.colorLog(stuff.ColorGreen, msgs...)
}
func (lgr *Core) BlueLog(msgs ...interface{}) *Core {
	return lgr.colorLog(stuff.ColorBlue, msgs...)
}
func (lgr *Core) PurpleLog(msgs ...interface{}) *Core {
	return lgr.colorLog(stuff.ColorPurple, msgs...)
}
func (lgr *Core) CyanLog(msgs ...interface{}) *Core {
	return lgr.colorLog(stuff.ColorCyan, msgs...)
}

func (lgr *Core) SuccessLog(msgs ...interface{}) *Core {
	return lgr.GreenLog(hp.Unshift(msgs, stuff.Check)...)
}
func (lgr *Core) ErrorLog(msgs ...interface{}) *Core {
	return lgr.RedLog(hp.Unshift(msgs, stuff.Cross)...)
}
func (lgr *Core) BugHunter(msgs ...interface{}) *Core {
	return lgr.RedLog(hp.Unshift(msgs, stuff.Bug)...)
}
func (lgr *Core) Here(msgs ...interface{}) *Core {
	return lgr.InfoLog(hp.Unshift(msgs, "here!")...)
}
