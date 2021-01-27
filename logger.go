package logger

import (
	"fmt"

	wr "github.com/mreza0100/logger/writer"
)

type Core struct {
	logPath string
	wr      wr.Writer
	name    string
	add     []interface{}
	hooks   []func(logger *Core) interface{}
}

func (logger *Core) Copy() *Core {
	newLogger := *logger

	return &newLogger
}

func (logger *Core) With(add ...interface{}) *Core {
	newLogger := logger.Copy()

	newLogger.add = combine(add, logger.add)

	return newLogger
}

func (logger *Core) AddHook(fn ...func(logger *Core) interface{}) *Core {
	logger.hooks = append(logger.hooks, fn...)

	return logger
}

func (logger *Core) getHooksVals() []interface{} {
	result := make([]interface{}, len(logger.hooks))

	for idx, hook := range logger.hooks {
		result[idx] = hook(logger)
	}

	return result
}

func (logger *Core) Log(msgs ...interface{}) *Core {
	msgs = combine(logger.add, logger.getHooksVals(), msgs)

	fmt.Println(msgs...)

	logger.wr.Write(msgs...)

	return logger
}
