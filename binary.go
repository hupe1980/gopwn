package gopwn

import "bytes"

type BinaryReader interface {
	Read(p []byte) (n int, err error)
	ReadAt(b []byte, off int64) (n int, err error)
	Seek(offset int64, whence int) (int64, error)
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
