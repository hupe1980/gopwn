package tubes

import (
	"bufio"
	"io"
)

type tube struct {
	stdin   io.WriteCloser
	stdout  io.ReadCloser
	stderr  io.ReadCloser
	newLine byte
}

func (t *tube) SendLine(input interface{}) (int, error) {
	b := Bytes(input)
	b = append(b, t.NewLine())
	return t.stdin.Write(b)
}

func (t *tube) RecvLine() ([]byte, error) {
	rd := bufio.NewReader(t.stdout)
	return rd.ReadBytes(t.NewLine())
}

func (t *tube) NewLine() byte {
	if t.newLine != byte(0) {
		return t.newLine
	}
	return '\n'

}

// Bytes takes type interface{} and converts it to []byte, if it can't convert
// to []byte it will panic
func Bytes(t interface{}) []byte {
	switch x := t.(type) {
	case string:
		return []byte(x)
	case []byte:
		return x
	case byte:
		return []byte{x}
	case rune:
		return []byte(string(x))
	default:
		panic("failed to convert t to type []byte")
	}
}
