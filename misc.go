package gopwn

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"os"
	"strings"
)

// Hex encodes the bytes hexadecimal.
func Hex(src []byte) []byte {
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	return dst
}

// HexToString encodes the bytes into a hexadecimal string.
func HexToString(src []byte) string {
	return hex.EncodeToString(src)
}

// UnHex decodes the hexadecimal bytes.
func UnHex(src []byte) ([]byte, error) {
	dst := make([]byte, hex.DecodedLen(len(src)))
	n, err := hex.Decode(dst, src)
	if err != nil {
		return nil, err
	}
	return dst[:n], nil
}

// UnHexString decodes the hexadecimal string into representative bytes.
func UnHexString(src string) ([]byte, error) {
	decoded, err := hex.DecodeString(src)
	if err != nil {
		return nil, err
	}
	return decoded, nil
}

func Base64E(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func Base64D(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

func ROT13(s string) string {
	return strings.Map(func(c rune) rune {
		if c >= 'a' && c <= 'm' || c >= 'A' && c <= 'M' {
			return c + 13
		} else if c >= 'n' && c <= 'z' || c >= 'N' && c <= 'Z' {
			return c - 13
		}
		return c
	}, s)
}

func OpenFile(path string) (*os.File, Bintype, error) {
	fh, err := os.Open(path)
	if err != nil {
		return nil, BINTYPE_UNKNOWN, err
	}

	var ident = make([]byte, 4)
	if _, err := fh.ReadAt(ident[0:], 0); err != nil {
		return nil, BINTYPE_UNKNOWN, err
	}

	var binType Bintype
	if bytes.HasPrefix(ident, []byte("\x7FELF")) {
		binType = BINTYPE_ELF
	} else if bytes.HasPrefix(ident, []byte("MZ")) {
		binType = BINTYPE_PE
	} else if bytes.HasPrefix(ident, P32L(0xfeedfacf)) {
		binType = BINTYPE_MACHO
	} else {
		binType = BINTYPE_UNKNOWN
	}

	return fh, binType, nil
}
