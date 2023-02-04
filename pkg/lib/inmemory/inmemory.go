package inmemory

import "io"

type DummyFile struct {
	readByte  int
	dataArray []byte
}

func (d *DummyFile) Read(p []byte) (n int, err error) {
	for i, _ := range p {
		// データ末尾に来たら脱出
		if i+d.readByte == len(d.dataArray) {
			return d.readByte + i, io.EOF
		}
		p[i] = d.dataArray[d.readByte+i]
	}
	return len(p), nil
}

func (d *DummyFile) Write(p []byte) (n int, err error) {
	for i := range p {
		d.dataArray = append(d.dataArray, p[i])
	}
	return len(p), nil
}

func (d *DummyFile) Seek(offset int64, whence int) (int64, error) {
	d.readByte = whence
	return int64(whence), nil
}
