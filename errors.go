// Copyright Cristian Echeverría Rabí

package conductor

//----------------------------------------------------------------------------------------

type ValueError struct {
	msg string
}

func (e *ValueError) Error() string {
	return e.msg
}
