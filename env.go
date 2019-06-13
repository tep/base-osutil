// Copyright © 2017 Timothy E. Peoples <eng@toolman.org>
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of the GNU General Public License
// as published by the Free Software Foundation; either version 2
// of the License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

// Package osutil provides convenience functions for OS interaction
package osutil // import "toolman.org/base/osutil"

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
