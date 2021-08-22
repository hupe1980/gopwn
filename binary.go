package gopwn

import (
	"bytes"
	"io/ioutil"
	"os"
)

type BinaryReader interface {
	Read(p []byte) (n int, err error)
	ReadAt(b []byte, off int64) (n int, err error)
	Seek(offset int64, whence int) (int64, error)
}

type fileBytes struct {
	raw          []byte
	addrToOffset func(addr uint64) (uint64, error)
}

// Write copies data to the raw data at the specified virtual address
func (f *fileBytes) Write(data []byte, addr uint64) error {
	offset, err := f.addrToOffset(addr)
	if err != nil {
		return err
	}
	copy(f.raw[offset:offset+uint64(len(data))], data)
	return nil
}

// Read reads up to n bytes from the raw data at the specified virtual address
func (f *fileBytes) Read(addr uint64, n int) ([]byte, error) {
	offset, err := f.addrToOffset(addr)
	if err != nil {
		return nil, err
	}
	buf := make([]byte, n)
	r := bytes.NewReader(f.raw)
	if _, err := r.ReadAt(buf, int64(offset)); err != nil {
		return nil, err
	}
	return buf, nil
}

// Save saves the raw bytes to a specified file path
func (f *fileBytes) Save(filePath string, fileMode os.FileMode) error {
	return ioutil.WriteFile(filePath, f.raw, fileMode)
}

type Cave struct {
	SectionName string
	Begin       int
	End         int
	Size        int
	Addr        int
	Infos       string
}

func searchCaves(name string, body []byte, offset, addr uint64, infos string, caveSize int) []Cave {
	caveBytes := []byte("\x00")
	var caves []Cave
	caveCount := 0
	for currentOffset := 0; currentOffset < len(body); currentOffset++ {
		currentByte := body[currentOffset]
		if bytes.Contains([]byte{currentByte}, caveBytes) {
			caveCount++
		} else {
			if caveCount >= caveSize {
				caves = append(caves, Cave{
					SectionName: name,
					Size:        caveCount,
					Addr:        int(addr) + currentOffset - caveCount,
					Begin:       int(offset) + currentOffset - caveCount,
					End:         int(offset) + currentOffset,
					Infos:       infos,
				})
			}
			caveCount = 0
		}
	}
	return caves
}
