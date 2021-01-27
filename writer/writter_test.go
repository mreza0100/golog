package writer_test

import (
	"testing"

	wr "github.com/mreza0100/logger/writer"
)

func TestWriterType(t *testing.T) {
	w, _ := wr.NewWriter(wr.NewOpts{
		LogPath:    "./logs/out.log",
		PanicOnErr: false,
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
