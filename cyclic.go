package gopwn

import (
	"bytes"
	"context"
)

const (
	alphabet = "abcdefgbhijklmnopqstuvwxyz"
)

type AdvancedCyclicParams struct {
	Length       int
	Alphabet     string
	SubseqLength int
}

type AdvancedCyclicFindParams struct {
	Subseq       []byte
	Alphabet     string
	SubseqLength int
}

func AdvancedCylic(params AdvancedCyclicParams) string {
	l := params.Length
	a := alphabet
	n := 4

	if params.Alphabet != "" {
		a = params.Alphabet
	}

	if params.SubseqLength != 0 {
		n = params.SubseqLength
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	chn := make(chan byte)
	go func() {
		defer close(chn)
		deBruijn(ctx, chn, a, n)
	}()

	var buf bytes.Buffer
	for i := range chn {
		buf.WriteByte(a[i])
		if buf.Len() == l {
			cancel()
			break
		}
	}
	return buf.String()
}

func AdvancedCylicFind(params AdvancedCyclicFindParams) int {
	s := params.Subseq
	a := alphabet
	n := 4

	if params.Alphabet != "" {
		a = params.Alphabet
	}

	if params.SubseqLength != 0 {
		n = params.SubseqLength
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	chn := make(chan byte)
	go func() {
		defer close(chn)
		deBruijn(ctx, chn, a, n)
	}()

	var seq []byte
	pos := 0
	for i := range chn {
		seq = append(seq, a[i])
		if len(seq) > len(s) {
			seq = seq[1:]
			pos += 1
			if bytes.Compare(seq, s) == 0 {
				cancel()
				return pos
			}
		}
	}
	return -1
}

func Cyclic(length int) string {
	return AdvancedCylic(AdvancedCyclicParams{
		Alphabet:     alphabet,
		Length:       length,
		SubseqLength: 4,
	})
}

func CyclicFind(subseq []byte) int {
	return AdvancedCylicFind(AdvancedCyclicFindParams{
		Subseq:       subseq,
		Alphabet:     alphabet,
		SubseqLength: 4,
	})
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
