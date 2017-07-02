// Package osutil provides convenience functions for OS interaction
package osutil  // import "toolman.org/base/osutil"

import (
	"os"
)

// FindEnvDefault examines the environment variables specified by vars and
// returns the value of the first variable found to be set, even if the set
// value is the empty string. If none of the given vars are set in the current
// environment, defaultValue is returned instead.
func FindEnvDefault(defaultValue string, vars ...string) string {
	for _, v := range vars {
		if val, ok := os.LookupEnv(v); ok {
			return val
		}
	}

	return defaultValue
}
