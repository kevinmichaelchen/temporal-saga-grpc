// Package logging provides an FX module for logging.
package logging

import (
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

// Module - An FX module for logging.
var Module = fx.Module("logging",
	fx.Invoke(ConfigureLogger),
)

// ConfigureLogger - Configures the global logger.
func ConfigureLogger() {
	// Logs the event in colors if stdout is a tty, otherwise without colors.
	logrus.SetFormatter(
		&logrus.TextFormatter{},
	)
}
