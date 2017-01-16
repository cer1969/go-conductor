// Copyright Cristian Echeverría Rabí

package conductor

import (
	"math"
)

//----------------------------------------------------------------------------------------

// NewOperatingItem Returns *OperatingItem object
// currentCalc *CurrentCalc: *CurrentCalc instance
// tempMaxOp   float64     : Maximux operating temperature for currentcalc.conductor [°C]
// nsc         int         : Number of subconductor per fase
func NewOperatingItem(currentCalc *CurrentCalc, tempMaxOp float64,
	nsc int) (*OperatingItem, error) {
	if currentCalc == nil {
		return nil, &ValueError{"NewCurrentCalc: currentCalc == nil"}
	}
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

//----------------------------------------------------------------------------------------

// OperatingItem Object to calculate current and temperatures for CurrentCalc and operating
// 				 condictions.
type OperatingItem struct {
	currentCalc *CurrentCalc // *CurrentCalc instance
	tempMaxOp   float64      // Maximux operating temperature for currentcalc.conductor [°C]
	nsc         int          // Number of subconductor per fase
}

func (oi *OperatingItem) Current(ta float64) (float64, error) {
	if ta < TA_MIN {
		return math.NaN(), &ValueError{"OperatingItem.Current: ta < TA_MIN"}
	}
	if ta > TA_MAX {
		return math.NaN(), &ValueError{"OperatingItem.Current: ta > TA_MAX"}
	}
	return oi.currentCalc.Current(ta, oi.tempMaxOp)
}

func (oi *OperatingItem) CurrentCalc() *CurrentCalc {
	return oi.currentCalc
}

func (oi *OperatingItem) TempMaxOp() float64 {
	return oi.tempMaxOp
}

func (oi *OperatingItem) Nsc() int {
	return oi.nsc
}

//----------------------------------------------------------------------------------------

// NewOperatingTable Returns *OperatingTable object
//	items []*OperatingItem : Slice of *OperatingItem
//	id    string           : Database id
func NewOperatingTable(items []*OperatingItem, id string) (*OperatingTable, error) {
	if items == nil {
		return nil, &ValueError{"NewOperatingTable: items == nil"}
	}
	if len(items) < 1 {
		return nil, &ValueError{"NewOperatingTable: len(items) < 1"}
	}
	for _, i := range items {
		if i == nil {
			return nil, &ValueError{"NewOperatingTable: item == nil"}
		}
	}
	return &OperatingTable{items, id}, nil
}

//----------------------------------------------------------------------------------------
// TODO : DEFINIR FORMA DE ACCEDER A LOS ITEMS QUE NO PERMITA MODIFICARLOS

// OperatingTable Groups of OperatingItem objects to perform calculus over the group
type OperatingTable struct {
	items []*OperatingItem // Slice of *OperatingItem
	id    string           // Database id
}

func (ot *OperatingTable) Current(ta float64) (float64, error) {
	if ta < TA_MIN {
		return math.NaN(), &ValueError{"OperatingTable.Current: ta < TA_MIN"}
	}
	if ta > TA_MAX {
		return math.NaN(), &ValueError{"OperatingTable.Current: ta > TA_MAX"}
	}
	cur, _ := ot.items[0].Current(ta)
	for _, x := range ot.items[1:] {
		if xcur, _ := x.Current(ta); xcur < cur {
			cur = xcur
		}
	}
	return cur, nil
}

func (ot *OperatingTable) Append(item *OperatingItem) error {
	if item == nil {
		return &ValueError{"OperatingTable.Append: items == nil"}
	}
	ot.items = append(ot.items, item)
	return nil
}

func (ot *OperatingTable) Len() int {
	return len(ot.items)
}
