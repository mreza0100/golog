package golog

import (
	"fmt"

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
	opts.Name = "name: " + opts.Name
	add := []interface{}{opts.Name}

	writer, err := wr.NewWriter(wr.NewOpts{
		LogPath:    opts.LogPath,
		PanicOnErr: opts.PanicOnErr,
	})
	if err != nil {
		panic(err)
	}

	err = writer.Write("mamad")
	fmt.Println(err)

	return &Core{
		logPath: opts.LogPath,
		name:    opts.Name,
		add:     add,
		wr:      writer,
	}
}
