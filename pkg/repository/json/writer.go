package json

import (
	"encoding/json"
	"github.com/growthers/mu-enshita/pkg/types/data"
	"io"
	"os"
)

// JSON Writer
type DataReadWriter struct {
	f *os.File
}

func NewDataReadWriter(f *os.File) *DataReadWriter {
	return &DataReadWriter{f: f}
}

func (w DataReadWriter) write(d []byte) error {
	_, err := w.f.Seek(0, 0)
	if err != nil {
		return err
	}

	_, err = w.f.Write(d)
	if err != nil {
		return err
	}

	_, err = w.f.Seek(0, 0)
	if err != nil {
		return err
	}

	return nil
}

func (w DataReadWriter) read() (data.StoreJSON, error) {
	d, err := io.ReadAll(w.f)
	if err != nil {
		return data.StoreJSON{}, err
	}

	_, err = w.f.Seek(0, 0)
	if err != nil {
		return data.StoreJSON{}, err
	}

	res := data.StoreJSON{}
	err = json.Unmarshal(d, &res)
	if err != nil {
		return data.StoreJSON{}, err
	}

	return res, nil
}
