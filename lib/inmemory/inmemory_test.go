package inmemory

import (
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

var T = DummyFile{
	dataArray: []byte("ABCDE"),
}

func TestDummyFile_Read(t *testing.T) {
	T.dataArray = []byte("ABCDE")
	T.readByte = 0

	// 読み込み元 == 読み込み先のとき
	b := make([]byte, 5)
	T.Read(b)
	assert.Equal(t, []byte("ABCDE"), b)

	// 読み込み元 > 読み込み先
	b1 := make([]byte, 10)
	T.Read(b1)
	assert.Equal(t, []byte{65, 66, 67, 68, 69, 0, 0, 0, 0, 0}, b1)

	// 読み込み元 < 読み込み先
	b2 := make([]byte, 1)
	T.Read(b2)
	assert.Equal(t, []byte("A"), b2)
}

func TestDummyFile_Write(t *testing.T) {
	T.dataArray = []byte{}
	T.readByte = 0
	b := []byte("ABCDE")
	T.Write(b)
	assert.Equal(t, []byte("ABCDE"), T.dataArray)
}

func TestDummyFile_Seek(t *testing.T) {
	T.readByte = 0
	T.dataArray = []byte("ABCDEFGHIJK")
	b := make([]byte, 5)
	T.Read(b)

	// 先頭に持っていく
	T.Seek(0, 0)
	assert.Equal(t, 0, T.readByte)
}

func Test(t *testing.T) {
	T.readByte = 0
	T.dataArray = []byte{}
	T.dataArray = []byte("ABCDEFGHIJK")

	e, _ := io.ReadAll(&T)
	assert.Equal(t, []byte("ABCDEFGHIJK"), e)
}
