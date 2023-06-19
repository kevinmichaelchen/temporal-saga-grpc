package simulated

import (
	"math/rand"
	"time"

	"github.com/sirupsen/logrus"
)

// Sleep - Sleeps for a random duration.
func Sleep() {
	min := 100
	max := 300
	sleepDuration := time.Millisecond * time.Duration(rand.Intn(max-min+1)+min)
	logrus.WithField("sleep", sleepDuration).Info("Sleeping...")
	time.Sleep(sleepDuration)
}
