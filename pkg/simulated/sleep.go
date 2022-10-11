package simulated

import (
	"github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

func Sleep() {
	min := 100
	max := 300
	sleepDuration := time.Millisecond * time.Duration(rand.Intn(max-min+1)+min)
	logrus.WithField("sleep", sleepDuration).Info("Sleeping...")
	time.Sleep(sleepDuration)
}
