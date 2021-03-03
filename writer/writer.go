package writer

import (
	"os"

	"github.com/mreza0100/golog/helpers"
)

type Writer struct {
	logPath string
}

type NewOpts struct {
	LogPath string
}

func New(opt NewOpts) Writer {
	helpers.CreateDir(opt.LogPath)

	return Writer{
		logPath: opt.LogPath,
	}
}

func (w *Writer) writeHandler(data []byte) error {
	f, _ := os.OpenFile(w.logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	defer f.Close()

	_, err := f.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func (w *Writer) Write(msgs ...interface{}) error {
	if w.logPath == "" {
		return nil
	}
	return w.writeHandler(helpers.ToByte(msgs...))
}

func (w *Writer) RemoveFile() error {
	if w.logPath == "" {
		return nil
	}

	return os.Remove(w.logPath)
}
