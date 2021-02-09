package golog

import (
	"sync"

	wr "github.com/mreza0100/golog/writer"
)

type InitOprions struct {
	LogPath   string
	Name      string
	WithTime  bool
	Debug     bool
	DebugMode bool
	wr        wr.Writer
}

type hookT []func(*Core) interface{}

func New(opts InitOprions) *Core {
	add := []interface{}{"[[ " + opts.Name + " ]] => "}

	writer := wr.New(wr.NewOpts{
		LogPath: opts.LogPath,
	})

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
		color:       "",
	}

	lgr.Debug = lgr
	lgr.Debug.isDebugMode = opts.Debug

	return lgr
}
