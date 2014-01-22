// Copyright Cristian Echeverría Rabí

package conductor

import (
	"errors"
	"fmt"
)

//----------------------------------------------------------------------------------------

type Tester struct {
	s string
	n int
}

func (ts *Tester) addError(ref string, sim string, val float64, limit float64) {
	ts.n += 1
	ts.s += fmt.Sprintf("\n%s: Requiered value %s %f [%f received]", ref, sim, limit, val)
}

func (ts *Tester) Lt(ref string, val float64, limit float64) {
	if val >= limit {
		ts.addError(ref, "<", val, limit)
	}
}

func (ts *Tester) Le(ref string, val float64, limit float64) {
	if val > limit {
		ts.addError(ref, "<=", val, limit)
	}
}

func (ts *Tester) Gt(ref string, val float64, limit float64) {
	if val <= limit {
		ts.addError(ref, ">", val, limit)
	}
}

func (ts *Tester) Ge(ref string, val float64, limit float64) {
	if val < limit {
		ts.addError(ref, ">=", val, limit)
	}
}

func (ts *Tester) GetError() error {
	if ts.n == 0 {
		return nil
	}
	return errors.New(fmt.Sprintf("ERROR [%d]%s", ts.n, ts.s))
}
