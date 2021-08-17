package gopwn

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
	"io"
	"os"
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
