package gopwn

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"os"
	"strings"
)

// MD5Sum calculates the md5 sum of a byte array
func MD5Sum(b []byte) string {
	return hashSum(b, md5.New())
}

// MD5File calculates the md5 sum of a file
func MD5File(path string) string {
	return hashFile(path, sha1.New())
}

// Sha1Sum calculates the md5 sum of a byte array
func Sha1Sum(b []byte) string {
	return hashSum(b, md5.New())
}

// Sha1Sum calculates the md5 sum of a file
func Sha1File(path string) string {
	return hashFile(path, sha1.New())
}

// Sha224Sum calculates the md5 sum of a byte array
func Sha224Sum(b []byte) string {
	return hashSum(b, sha256.New224())
}

// Sha224Sum calculates the md5 sum of a file
func Sha224File(path string) string {
	return hashFile(path, sha256.New224())
}

// Sha256Sum calculates the md5 sum of a byte array
func Sha256Sum(b []byte) string {
	return hashSum(b, sha256.New())
}

// Sha256Sum calculates the md5 sum of a file
func Sha256File(path string) string {
	return hashFile(path, sha256.New())
}

// Sha384Sum calculates the md5 sum of a byte array
func Sha384Sum(b []byte) string {
	return hashSum(b, sha512.New384())
}

// Sha384Sum calculates the md5 sum of a file
func Sha384File(path string) string {
	return hashFile(path, sha512.New384())
}

// Sha512Sum calculates the md5 sum of a byte array
func Sha512Sum(b []byte) string {
	return hashSum(b, sha512.New())
}

// Sha512Sum calculates the md5 sum of a file
func Sha512File(path string) string {
	return hashFile(path, sha512.New())
}

func hashSum(b []byte, h hash.Hash) string {
	h.Write(b)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func hashFile(path string, h hash.Hash) string {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err := io.Copy(h, f); err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

// Hex encodes the bytes hexadecimal.
func Hex(src []byte) []byte {
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	return dst
}

// HexString encodes the bytes into a hexadecimal string.
func HexString(src []byte) string {
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
