package json

import (
	"encoding/json"
	"io"

	"github.com/growthers/mu-enshita/pkg/types/data"
)

// DataReadWriter データストアのJSONファイルを操作するための構造体
type DataReadWriter struct {
	f io.ReadWriteSeeker
}

func NewDataReadWriter(f io.ReadWriteSeeker) *DataReadWriter {
	return &DataReadWriter{f: f}
}

// write 対象に書き込み
func (w DataReadWriter) write(d []byte) error {
	defer func(f io.ReadWriteSeeker, offset int64, whence int) {
		_, err := f.Seek(offset, whence)
		if err != nil {
			return
		}
	}(w.f, 0, 0)

	// 上書きのために書き込み対象の先頭までシークしておく
	_, err := w.f.Seek(0, 0)
	if err != nil {
		return err
	}

	_, err = w.f.Write(d)
	if err != nil {
		return err
	}

	return nil
}

// read 対象を読み込んでマッピングして返す
func (w DataReadWriter) read() (data.StoreJSON, error) {
	defer func(f io.ReadWriteSeeker, offset int64, whence int) {
		_, err := f.Seek(offset, whence)
		if err != nil {
			return
		}
	}(w.f, 0, 0)
	// 対象をすべて読み込み
	d, err := io.ReadAll(w.f)
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
