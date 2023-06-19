// Package rand configures the global random seed.
package rand

import (
	"math/rand"
	"time"

	"go.uber.org/fx"
)

// Module - An FX module to configure the global random seed.
var Module = fx.Module("rand",
	fx.Invoke(SeedRand),
)

// SeedRand - Configures the global random seed.
func SeedRand() {
	rand.Seed(time.Now().UnixNano())
}
