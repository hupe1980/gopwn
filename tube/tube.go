package tube

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
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

// RecvUntil receives data until the specified sequence of bytes is detected.
func (t *tube) RecvUntil(needle []byte, drop bool) ([]byte, error) {
	return t.RecvUntilWithContext(context.Background(), needle, drop)
}

// RecvUntilWithContext receives data until the specified sequence of bytes is detected or the context is done.
func (t *tube) RecvUntilWithContext(ctx context.Context, needle []byte, drop bool) ([]byte, error) {
	data := make([]byte, len(needle))
	b := bufio.NewReader(t.stdout)

	_, err := io.ReadFull(b, data)
	if err != nil {
		return nil, err
	}

	idx := 0
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		if bytes.Equal(data[idx:idx+len(needle)], needle) {
			if drop {
				return data[0 : len(data)-len(needle)], nil
			}
			return data, nil
		}

		byt, err := b.ReadByte()
		if err != nil {
			return nil, err
		}

		data = append(data, byt)
		idx++
	}
}

// RecvLine receives data until a newline delimiter is detected.
func (t *tube) RecvLine() ([]byte, error) {
	return t.RecvUntil([]byte{t.NewLine()}, true)
}

// Interactive allows the user to interact with the tube manually.
func (t *tube) Interactive() error {
	go func() {
		for {
			data, err := t.RecvN(1)
			if err != nil {
				break
			}
			fmt.Printf("%c", data[0])
		}
	}()

	for {
		var line string
		fmt.Scanln(&line)
		if line == "quit()" {
			fmt.Println("Exiting interactive mode...")
			return nil
		}

		_, err := t.SendLine([]byte(line))
		if err != nil {
			return err
		}
	}
}

func (t *tube) NewLine() byte {
	if t.newLine != byte(0) {
		return t.newLine
	}
	return '\n'

}

// Bytes takes type interface{} and converts it to []byte, if it can't convert
// to []byte it will panic.
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
		panic("Failed to convert t to type []byte")
	}
}
