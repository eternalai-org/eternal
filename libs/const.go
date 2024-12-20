package libs

import (
	"crypto/sha256"
	"encoding/hex"
	"math"
	"math/big"
	"time"
)

const TimeToWating time.Duration = 10
const (
	MODE_MINER     = "miner"
	MODE_VALIDATOR = "validator"
)

func CreateSeed(params string, requestID string) uint64 {
	seed := hex.EncodeToString([]byte(params + requestID))

	h := sha256.New()

	h.Write([]byte(seed))

	bs := h.Sum(nil)

	seedHex := hex.EncodeToString(bs)

	i := new(big.Int)
	i.SetString(seedHex, 16)

	// check if the seed is too large for uint64

	if i.BitLen() > 64 {
		i = i.Mod(i, new(big.Int).SetUint64(math.MaxUint64))
	}

	return i.Uint64()
}
