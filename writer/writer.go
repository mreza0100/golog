package writer

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Writer struct {
	logPath    string
	panicOnErr bool
}

type NewOpts struct {
	LogPath    string
	PanicOnErr bool
}

func NewWriter(opt NewOpts) (Writer, error) {
	createDir(opt.LogPath)

	return Writer{
		logPath: opt.LogPath,
	}, nil
}

func (w *Writer) writeHandler(data []byte) error {
	f, _ := os.OpenFile(w.logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	defer f.Close()

	data = append(data, []byte("\n")...)
	_, err := f.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func (w *Writer) Write(msgs ...interface{}) (err error) {
	finallMsg := []byte("")
	for _, msg := range msgs {
		switch s := msg.(type) {
		case string:
			finallMsg = append(finallMsg, []byte(s)...)
		case int:
			finallMsg = append(finallMsg, []byte(strconv.Itoa(s))...)
		case int64:
			finallMsg = append(finallMsg, []byte(strconv.Itoa(int(s)))...)
		case uint64:
			finallMsg = append(finallMsg, []byte(strconv.Itoa(int(s)))...)
		default:
			finallMsg, err = json.Marshal(parseVal(s))
			if err != nil {
				fmt.Println("From writer: Cant Marshal value")
				return err
			}
		}
	}
	err = w.writeHandler(finallMsg)
	if err != nil && w.panicOnErr {
		return err
	}

	return nil
}
