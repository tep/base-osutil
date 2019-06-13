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

package osutil

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

type testcase struct {
	name string
	def  string
	vars []string
	want string
}

func varString(v []string) string {
	var vl []string
	for _, vs := range v {
		vl = append(vl, fmt.Sprintf("%q", vs))
	}

	return strings.Join(vl, ", ")
}

func newTest(n, w, d string, v ...string) *testcase {
	return &testcase{n, d, v, w}
}

func (tc *testcase) test(t *testing.T) {
	if got := FindEnvDefault(tc.def, tc.vars...); got != tc.want {
		t.Errorf("FindEnvDefault(%q, %s) := %q; Wanted: %q", tc.def, varString(tc.vars), got, tc.want)
	}
}

func TestFindEnvDefault(t *testing.T) {
	var emtstr string
	os.Clearenv()
	os.Setenv("VAR1", "val1")
	os.Setenv("VAR2", "val2")
	os.Setenv("EMPT", emtstr)

	tests := []*testcase{
		newTest("one-var-unset", "dflt", "dflt", "VAR0"),
		newTest("two-var-unset", "dflt", "dflt", "VAR0", "VAR3"),
		newTest("first-var-set", "val1", "dflt", "VAR1", "VAR2"),
		newTest("second-varset", "val2", "dflt", "VAR0", "VAR2", "VAR1"),
		newTest("first-v-empty", emtstr, "dflt", "EMPT", "VAR1"),
		newTest("unset-2-empty", emtstr, "dflt", "VAR0", "EMPT", "VAR1"),
	}

	var failed bool
	for _, tc := range tests {
		if bad := t.Run(tc.name, tc.test); bad {
			failed = true
		}
	}

	if failed {
		t.Logf("ENVIRONMENT: %v", os.Environ())
	}
}
