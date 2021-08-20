package gopwn

import (
	"encoding/hex"
)

//Hex encodes the bytes hexadecimal.
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
