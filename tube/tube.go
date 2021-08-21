package tube

import (
	"bufio"
	"bytes"
	"io"
)

type tube struct {
	stdin   io.WriteCloser
	stdout  io.ReadCloser
	stderr  io.ReadCloser
	newLine byte
}

// SendLine sends data with a trailing newline character
func (t *tube) SendLine(input interface{}) (int, error) {
	b := Bytes(input)
	b = append(b, t.NewLine())
	return t.stdin.Write(b)
}

// RecvN receives a specified number of bytes
func (t *tube) RecvN(n int) ([]byte, error) {
	b := make([]byte, n)
	rn, err := t.stdout.Read(b)
	if err != nil {
		return nil, err
	}
	return b[:rn], nil
}

func (t *tube) RecvLine() ([]byte, error) {
	rd := bufio.NewReader(t.stdout)
	b, err := rd.ReadBytes(t.NewLine())
	if err != nil {
		return nil, err
	}
	b = bytes.TrimSuffix(b, []byte{t.NewLine()})
	return b, nil
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
