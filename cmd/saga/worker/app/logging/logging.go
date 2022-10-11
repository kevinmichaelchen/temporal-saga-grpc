package logging

import (
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

var Module = fx.Module("logging",
	fx.Invoke(ConfigureLogger),
)

func ConfigureLogger() {
	// Logs the event in colors if stdout is a tty, otherwise without colors.
	logrus.SetFormatter(&logrus.TextFormatter{})
}
