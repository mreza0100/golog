package golog

import (
	"sync"

	"github.com/mreza0100/golog/colors"
	wr "github.com/mreza0100/golog/writer"
)

type InitOprions struct {
	LogPath      string
	Name         string
	WithTime     bool
	DebugMode    bool
	ClearLogFile bool
	wr           wr.Writer
}

type hookT []func(*Core) interface{}

func New(opts InitOprions) *Core {
	var (
		addLog      = make([]interface{}, 0, 1)
		hooks       = make(hookT, 0, 1)
		isDebugMode = opts.DebugMode

		writer wr.Writer
	)

	{
		writer = wr.New(wr.NewOpts{
			LogPath: opts.LogPath,
		})
		if opts.ClearLogFile {
			writer.RemoveFile()
		}
	}

	addLog = append(addLog, "[[ "+opts.Name+" ]]")

	if opts.WithTime {
		hooks = append(hooks, timeHook)
	}

	data := &dataT{
		logPath: opts.LogPath,
		addLog:  addLog,
		hooks:   hooks,
		color:   colors.ColorWhite,
		mu:      &sync.Mutex{},
	}

	lgr := &Core{
		IsDebugMode: true, // always working
		data:        data,

		WR: writer,
	}
	lgr.Debug = &Core{
		IsDebugMode: isDebugMode,
		WR:          writer,
		data:        data,
	}

	return lgr
}
