package simulated

import (
	"time"

	"github.com/sirupsen/logrus"
)

// Sleep - Sleeps for a random duration.
func Sleep() {
	min := 100
	max := 300

	roll, err := generateRandomNumber(max - min + 1)
	if err != nil {
		sleepDuration := time.Millisecond * time.Duration(min)
		logrus.WithField("sleep", sleepDuration).
			WithError(err).
			Info("Sleeping...")
		time.Sleep(sleepDuration)

		return
	}

	sleepDuration := time.Millisecond * time.Duration(roll+min)
	logrus.WithField("sleep", sleepDuration).Info("Sleeping...")
	time.Sleep(sleepDuration)
}
