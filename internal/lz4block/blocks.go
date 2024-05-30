// Package lz4block provides LZ4 BlockSize types and pools of buffers.
package lz4block

const (
	Block64Kb uint32 = 1 << (16 + iota*2)
	Block256Kb
	Block1Mb
	Block4Mb
	Block8Mb = 2 * Block4Mb
)

func Index(b uint32) BlockSizeIndex {
	switch b {
	case Block64Kb:
		return 4
	case Block256Kb:
		return 5
	case Block1Mb:
		return 6
	case Block4Mb:
		return 7
	case Block8Mb: // only valid in legacy mode
		return 3
	}
	return 0
}

func IsValid(b uint32) bool {
	return Index(b) > 0
}

type BlockSizeIndex uint8

func (b BlockSizeIndex) IsValid() bool {
	switch b {
	case 4, 5, 6, 7:
		return true
	}
	return false
}

func (b BlockSizeIndex) Get() []byte {
	switch b {
	case 4:
		return make([]byte, Block64Kb)
	case 5:
		return make([]byte, Block256Kb)
	case 6:
		return make([]byte, Block1Mb)
	case 7:
		return make([]byte, Block4Mb)
	case 3:
		return make([]byte, Block8Mb)
	}

	return nil
}

func Put(buf []byte) {
	return
}

type CompressionLevel uint32

const Fast CompressionLevel = 0
