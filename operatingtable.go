// Copyright Cristian Echeverría Rabí

package conductor

import (
	"math"
)

//----------------------------------------------------------------------------------------

// NewOperatingItem Returns OperatingItem object
// currentCalc *CurrentCalc: *CurrentCalc instance
// tempMaxOp   float64     : Maximux operating temperature for currentcalc.conductor [°C]
// nsc         int         : Number of subconductor per fase
func NewOperatingItem(currentCalc *CurrentCalc, tempMaxOp float64,
	nsc int) (*OperatingItem, error) {
	if tempMaxOp < TC_MIN {
		return nil, &ValueError{"NewOperatingItem: tempMaxOp < TC_MIN"}
	}
	if tempMaxOp > TC_MAX {
		return nil, &ValueError{"NewOperatingItem: tempMaxOp > TC_MAX"}
	}
	if nsc < 1 {
		return nil, &ValueError{"NewOperatingItem: nsc < 1"}
	}
	return &OperatingItem{currentCalc, tempMaxOp, nsc}, nil
}

////----------------------------------------------------------------------------------------

// OperatingItem Object to calculate current and temperatures for CurrentCalc and operating
// 				 condictions.
type OperatingItem struct {
	currentCalc *CurrentCalc // *CurrentCalc instance
	tempMaxOp   float64      // Maximux operating temperature for currentcalc.conductor [°C]
	nsc         int          // Number of subconductor per fase
}

func (oi *OperatingItem) Current(ta float64) (float64, error) {
	cur, err := oi.currentCalc.Current(ta, oi.tempMaxOp)
	if err != nil {
		return math.NaN(), &ValueError{"OperatingItem.Current: " + err.Error()}
	}
	return (cur * float64(oi.nsc)), nil
}

func (oi *OperatingItem) CurrentCalc() *CurrentCalc {
	return oi.currentCalc
}

func (oi *OperatingItem) TempMaxOp() float64 {
	return oi.tempMaxOp
}

func (oi *OperatingItem) SetTempMaxOp(t float64) error {
	if t < TC_MIN {
		return &ValueError{"OperatingItem.SetTempMaxOp: tempMaxOp < TC_MIN"}
	}
	if t > TC_MAX {
		return &ValueError{"OperatingItem.SetTempMaxOp: tempMaxOp > TC_MAX"}
	}
	oi.tempMaxOp = t
	return nil
}

func (oi *OperatingItem) Nsc() int {
	return oi.nsc
}

func (oi *OperatingItem) SetNsc(n int) error {
	if n < 1 {
		return &ValueError{"OperatingItem.SetNsc: nsc < 1"}
	}
	oi.nsc = n
	return nil
}
