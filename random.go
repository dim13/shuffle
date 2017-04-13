package shuffle

import (
	"crypto/rand"
	"encoding/binary"
)

type cryptoSrc struct{}

func (_ cryptoSrc) Int63() int64 {
	var ret uint64
	binary.Read(rand.Reader, binary.BigEndian, &ret)
	ret &= ^(uint64(1) << 63) // unset high bit
	return int64(ret)
}

func (_ cryptoSrc) Seed(_ int64) {}
