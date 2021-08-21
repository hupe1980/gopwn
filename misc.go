package gopwn

import (
	"encoding/base64"
	"encoding/hex"
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
