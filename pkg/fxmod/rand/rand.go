package rand

import (
	"go.uber.org/fx"
	"math/rand"
	"time"
)

var Module = fx.Module("rand",
	fx.Invoke(SeedRand),
)

func SeedRand() {
	rand.Seed(time.Now().UnixNano())
}
