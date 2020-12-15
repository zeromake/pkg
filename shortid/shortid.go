package shortid

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"strconv"
	"sync"
	"unsafe"
)

const (
	idSize  = aes.BlockSize / 2 // 64 bits
	keySize = aes.BlockSize     // 128 bits
)

// ShortID 生成短ID
type ShortID struct {
	ctr    []byte
	index  int
	buffer []byte
	cipher cipher.Block
	sync.Mutex
}

// New 创建短ID实例
func New() (*ShortID, error) {
	buf := make([]byte, keySize+aes.BlockSize)
	_, err := io.ReadFull(rand.Reader, buf)
	if err != nil {
		// /dev/urandom had better work
		return nil, err
	}
	c, err := aes.NewCipher(buf[:keySize])
	if err != nil {
		// AES had better work
		return nil, err
	}
	n := aes.BlockSize
	ctr := buf[keySize:]
	b := make([]byte, aes.BlockSize)
	return &ShortID{
		ctr:    ctr,
		index:  n,
		buffer: b,
		cipher: c,
	}, nil
}

// Generate 生成短ID字符串
func (s *ShortID) Generate() string {
	s.Lock()
	if s.index == aes.BlockSize {
		s.cipher.Encrypt(s.buffer, s.ctr)
		// increment ctr
		for i := aes.BlockSize - 1; i >= 0; i-- {
			s.ctr[i]++
			if s.ctr[i] != 0 {
				break
			}
		}
		s.index = 0
	}
	// zero-copy b/c we're arch-neutral
	id := *(*uint64)(unsafe.Pointer(&s.buffer[s.index]))
	s.index += idSize
	s.Unlock()
	return strconv.FormatUint(id, 36)
}

var defaultShortID, _ = New()

// Generate 生成短ID字符串
func Generate() string {
	return defaultShortID.Generate()
}
