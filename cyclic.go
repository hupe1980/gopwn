package gopwn

import (
	"bytes"
	"context"
)

const (
	alphabet         = "abcdefgbhijklmnopqstuvwxyz"
	distSubseqLength = 4
)

type CyclicOptions struct {
	Alphabet         string
	DistSubseqLength int
}

func Cyclic(length int, optFns ...func(o *CyclicOptions)) []byte {
	options := CyclicOptions{
		Alphabet:         alphabet,
		DistSubseqLength: distSubseqLength,
	}
	for _, fn := range optFns {
		fn(&options)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	chn := make(chan byte)
	go func() {
		defer close(chn)
		deBruijn(ctx, chn, options.Alphabet, options.DistSubseqLength)
	}()

	var buf bytes.Buffer
	for i := range chn {
		buf.WriteByte(options.Alphabet[i])
		if buf.Len() == length {
			cancel()
			break
		}
	}
	return buf.Bytes()
}

func CyclicFind(subseq []byte, optFns ...func(o *CyclicOptions)) int {
	options := CyclicOptions{
		Alphabet:         alphabet,
		DistSubseqLength: distSubseqLength,
	}
	for _, fn := range optFns {
		fn(&options)
	}

	if len(subseq) > options.DistSubseqLength {
		subseq = subseq[:options.DistSubseqLength]
	}

	for _, c := range subseq {
		if !bytes.Contains([]byte(options.Alphabet), []byte{c}) {
			return -1
		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	chn := make(chan byte)
	go func() {
		defer close(chn)
		deBruijn(ctx, chn, options.Alphabet, options.DistSubseqLength)
	}()

	var seq []byte
	pos := 0
	for i := range chn {
		seq = append(seq, options.Alphabet[i])
		if len(seq) > len(subseq) {
			seq = seq[1:]
			pos += 1
			if bytes.Equal(seq, subseq) {
				cancel()
				return pos
			}
		}
	}
	return -1
}

// https://en.wikipedia.org/wiki/De_Bruijn_sequence
func deBruijn(ctx context.Context, chn chan byte, alphabet string, n int) {
	k := len(alphabet)
	a := make([]byte, k*n)
	var db func(int, int) // recursive closure
	db = func(t, p int) {
		if t > n {
			if n%p == 0 {
				for _, b := range a[1 : p+1] {
					select {
					case <-ctx.Done():
						return
					default:
						chn <- b
					}
				}
			}
		} else {
			select {
			case <-ctx.Done():
				return
			default:
				a[t] = a[t-p]
				db(t+1, p)
				for j := int(a[t-p] + 1); j < k; j++ {
					a[t] = byte(j)
					db(t+1, t)
				}
			}
		}
	}
	db(1, 1)
}
