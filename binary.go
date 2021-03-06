package gopwn

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/ianlancetaylor/demangle"
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
	SectionName   string
	SectionOffset uint64
	SectionSize   uint64
	Begin         int
	End           int
	Size          int
	Addr          int
	Infos         string
}

func (c *Cave) Dump() {
	fmt.Println("\n[+] CAVE DETECTED!")
	fmt.Printf("[!] Section Name: %s\n", c.SectionName)
	fmt.Printf("[!] Section Offset: %#x\n", c.SectionOffset)
	fmt.Printf("[!] Section Size: %#x (%d bytes)\n", c.SectionSize, int(c.SectionSize))
	fmt.Printf("[!] Section Flags: %s\n", c.Infos)
	fmt.Printf("[!] Virtual Address: %#x\n", c.Addr)
	fmt.Printf("[!] Cave Begin: %#x\n", c.Begin)
	fmt.Printf("[!] Cave End: %#x\n", c.End)
	fmt.Printf("[!] Cave Size: %#x (%d bytes)\n", c.Size, c.Size)
}

func searchCaves(name string, data []byte, offset, addr, size uint64, infos string, caveSize int) []Cave {
	caveBytes := []byte("\x00")
	var caves []Cave
	caveCount := 0
	for currentOffset := 0; currentOffset < len(data); currentOffset++ {
		currentByte := data[currentOffset]
		if bytes.Contains([]byte{currentByte}, caveBytes) {
			caveCount++
		} else {
			if caveCount >= caveSize {
				caves = append(caves, Cave{
					SectionName:   name,
					SectionOffset: offset,
					SectionSize:   size,
					Size:          caveCount,
					Addr:          int(addr) + currentOffset - caveCount,
					Begin:         int(offset) + currentOffset - caveCount,
					End:           int(offset) + currentOffset,
					Infos:         infos,
				})
			}
			caveCount = 0
		}
	}
	return caves
}

type dataReader interface {
	Data() ([]byte, error)
}

type StringsOptions struct {
	Min      int
	Max      int
	Regex    func(min, max int) *regexp.Regexp
	Sections []string
	Demangle bool
}

func parseStrings(sections []dataReader, optFns ...func(o *StringsOptions)) []string {
	options := StringsOptions{
		Min:      4,
		Max:      100,
		Demangle: false,
		Regex: func(min, max int) *regexp.Regexp {
			return regexp.MustCompile(fmt.Sprintf("([\x20-\x7E]{%d}[\x20-\x7E]*)", min))
		},
	}
	for _, fn := range optFns {
		fn(&options)
	}

	validString := options.Regex(options.Min, options.Max)

	var strs []string
	for _, s := range sections {
		b, err := s.Data()
		if err != nil {
			continue
		}
		var slice [][]byte
		if slice = bytes.Split(b, []byte("\x00")); slice == nil {
			return nil
		}
		for _, b := range slice {
			if len(b) == 0 || len(b) > options.Max {
				continue
			}
			str := string(b)
			if validString.MatchString(str) {
				if options.Demangle {
					str = demangle.Filter(str)
				}
				strs = append(strs, str)
			}
		}
	}
	return strs
}
