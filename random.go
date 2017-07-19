package shuffle

import (
	"crypto/rand"
	"encoding/binary"
)

type cryptoSrc struct{}

func (cryptoSrc) Int63() int64 {
	var ret uint64
	binary.Read(rand.Reader, binary.BigEndian, &ret)
	return int64(ret &^ (1 << 63)) // unset high bit
}

func (cryptoSrc) Seed(_ int64) {}
