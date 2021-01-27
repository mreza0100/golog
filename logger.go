package golog

import (
	"fmt"
	"log"

	wr "github.com/mreza0100/golog/writer"
)

type Core struct {
	LogPath    string
	WR         wr.Writer
	panicOnErr bool
	Add        []interface{}
	Hooks      []func(logger *Core) interface{}
}

func (logger *Core) Copy() *Core {
	newLogger := *logger

	return &newLogger
}

func (logger *Core) With(add ...interface{}) *Core {
	newLogger := logger.Copy()

	newLogger.Add = combine(logger.Add, add)

	return newLogger
}

func (logger *Core) AddHook(fn ...func(logger *Core) interface{}) *Core {
	logger.Hooks = append(logger.Hooks, fn...)

	return logger
}

func (logger *Core) callHooksVals() []interface{} {
	result := make([]interface{}, len(logger.Hooks))

	for idx, hook := range logger.Hooks {
		result[idx] = hook(logger)
	}

	return result
}

func (logger *Core) Log(msgs ...interface{}) *Core {
	msgs = combine(logger.Add, logger.callHooksVals(), msgs)

	fmt.Print(msgs...)

	wrErr := logger.WR.Write(msgs...)
	if wrErr != nil {
		if logger.panicOnErr {
			panic(wrErr)
		}
		log.Fatal(wrErr)
	}

	return logger
}
