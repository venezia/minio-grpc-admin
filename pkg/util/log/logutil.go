package log

import (
	"github.com/juju/loggo"
)

// GetModuleLogger get a logger given a module name and level
func GetModuleLogger(module string) loggo.Logger {
	log := loggo.GetLogger(module)
	return log
}
