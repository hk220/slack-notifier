package reader

import (
	"io"
	"os"
)

type StdinReader struct{}

func NewStdinReader() *StdinReader {
	return &StdinReader{}
}

func (sr *StdinReader) Read() (string, error) {
	buf, err := io.ReadAll(os.Stdin)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}
