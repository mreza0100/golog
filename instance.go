package golog

import (
	wr "github.com/mreza0100/golog/writer"
)

type InitOprions struct {
	LogPath    string
	Name       string
	WithTime   bool
	PanicOnErr bool
	wr         wr.Writer
}

func New(opts InitOprions) *Core {
	add := []interface{}{opts.Name + ": "}

	writer := wr.New(wr.NewOpts{
		LogPath: opts.LogPath,
	})

	return &Core{
		LogPath:    opts.LogPath,
		Add:        add,
		Hooks:      []func(logger *Core) interface{}{},
		WR:         writer,
		panicOnErr: opts.PanicOnErr,
	}
}
