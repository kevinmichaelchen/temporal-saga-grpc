package rand

import (
	"math/rand"
	"time"

	"go.uber.org/fx"
)

var Module = fx.Module("rand",
	fx.Invoke(SeedRand),
)

func SeedRand() {
	rand.Seed(time.Now().UnixNano())
}
