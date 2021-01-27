package writer_test

import (
	"testing"

	wr "github.com/mreza0100/golog/writer"
)

func TestWriterType(t *testing.T) {
	w := wr.New(wr.NewOpts{
		LogPath: "./logs/out.log",
	})
	var err error

	type s struct {
		a int
		b uint
	}

	err = w.Write(s{
		a: -214,
		b: 22431,
	})
	if err != nil {
		panic(err)
	}

	err = w.Write(&s{
		a: -214,
		b: 22431,
	})
	if err != nil {
		panic(err)
	}

	err = w.Write(map[string]int{
		"a": 2,
		"b": 2,
	})

	if err != nil {
		panic(err)
	}

}
