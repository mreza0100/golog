package golog

import (
	"sync"

	"github.com/mreza0100/golog/helpers"
	wr "github.com/mreza0100/golog/writer"
)

type InitOprions struct {
	LogPath      string
	Name         string
	WithTime     bool
	Debug        bool
	DebugMode    bool
	ClearLogFile bool
	wr           wr.Writer
}

type hookT []func(*Core) interface{}

func New(opts InitOprions) *Core {
	add := []interface{}{"[[ ", opts.Name, " ]]"}

	writer := wr.New(wr.NewOpts{
		LogPath: opts.LogPath,
	})
	if opts.ClearLogFile {
		writer.RemoveFile()
	}

	hooks := make(hookT, 0)

	if opts.WithTime {
		hooks = append(hooks, timeHook)
	}

	lgr := &Core{
		LogPath:     opts.LogPath,
		Add:         add,
		Hooks:       hooks,
		WR:          writer,
		isDebugMode: false,
		mu:          &sync.Mutex{},
		color:       helpers.ColorWhite,
	}

	lgr.Debug = lgr
	lgr.Debug.isDebugMode = opts.Debug

	return lgr
}
