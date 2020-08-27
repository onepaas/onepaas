package ulid

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid"
)

func Generate() string {
	seed := time.Now().UnixNano()
	entropy := rand.New(rand.NewSource(seed))
	monotonicEntropy := ulid.Monotonic(entropy, 0)

	return ulid.MustNew(ulid.Timestamp(time.Now()), monotonicEntropy).String()
}
