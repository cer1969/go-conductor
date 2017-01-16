// Copyright Cristian Echeverría Rabí

package conductor

import (
	"fmt"
	"testing"
	//. "bitbucket.org/tormundo/go.conductor"
)

func getCurrentCalc() *CurrentCalc {
	cmk := ConductorMaker{"AAAC 740,8 MCM FLINT", CC_AAAC, 25.17, 0.00, 0.0, 0.0, 0.089360, 1e-10, ""}
	cc, _ := NewCurrentCalc(cmk.Get())
	return cc
}

//----------------------------------------------------------------------------------------

func Test_OperatingItem_ConstructorDefaults(t *testing.T) {
	cc := getCurrentCalc()
	opi, _ := NewOperatingItem(cc, 50, 2)

	if opi.CurrentCalc() != cc {
		t.Error("!=")
	}
	if opi.TempMaxOp() != 50 {
		t.Error("!=")
	}
	if opi.Nsc() != 2 {
		t.Error("!=")
	}
}

func Test_OperatingItem_ConstructorCurrentCalc(t *testing.T) {
	opi, err := NewOperatingItem(nil, 50, 2)
	if err == nil {
		t.Error("CurrentCalc=nil error expected")
	}
	if opi != nil {
		t.Error("nil expected with CurrentCalc=nil")
	}
}

func Test_OperatingItem_ConstructorTempMaxOp(t *testing.T) {
	opi, err := NewOperatingItem(getCurrentCalc(), TC_MIN, 2)
	if err != nil {
		t.Error(err)
	}
	if opi == nil {
		t.Error("OperatingItem expected")
	}
	opi, err = NewOperatingItem(getCurrentCalc(), TC_MAX, 2)
	if err != nil {
		t.Error(err)
	}
	if opi == nil {
		t.Error("OperatingItem expected")
	}
	opi, err = NewOperatingItem(getCurrentCalc(), TC_MIN-0.01, 2)
	if err == nil {
		t.Error("TempMaxOp < TC_MIN error expected")
	}
	if opi != nil {
		t.Error("nil expected with TempMaxOp < TC_MIN")
	}
	opi, err = NewOperatingItem(getCurrentCalc(), TC_MAX+0.01, 2)
	if err == nil {
		t.Error("TempMaxOp > TC_MAX error expected")
	}
	if opi != nil {
		t.Error("nil expected with TempMaxOp > TC_MAX")
	}
}

func Test_OperatingItem_ConstructorNsc(t *testing.T) {
	opi, err := NewOperatingItem(getCurrentCalc(), 50, 1)
	if err != nil {
		t.Error(err)
	}
	if opi == nil {
		t.Error("OperatingItem expected")
	}
	opi, err = NewOperatingItem(getCurrentCalc(), 50, 0)
	if err == nil {
		t.Error("Nsc < 1 error expected")
	}
	if opi != nil {
		t.Error("nil expected with Nsc < 1")
	}
}

//----------------------------------------------------------------------------------------

func Test_OperatingTable_ConstructorItems(t *testing.T) {
	ot, err := NewOperatingTable(nil, "")
	if err == nil {
		t.Error("items = nil error expected")
	}
	if ot != nil {
		t.Error("nil expected with items = nil")
	}
	ot, err = NewOperatingTable([]*OperatingItem{}, "")
	if err == nil {
		t.Error("len(items) < 1 error expected")
	}
	if ot != nil {
		t.Error("nil expected with len(items) < 1")
	}
	ot, err = NewOperatingTable([]*OperatingItem{nil, nil}, "")
	if err == nil {
		t.Error("item = nil error expected")
	}
	if ot != nil {
		t.Error("nil expected with item = nil")
	}
}

func Test_OperatingTable_Append(t *testing.T) {
	opi1, _ := NewOperatingItem(getCurrentCalc(), 50, 1)
	ot, _ := NewOperatingTable([]*OperatingItem{opi1}, "")

	opi2, _ := NewOperatingItem(getCurrentCalc(), 60, 1)
	err := ot.Append(opi2)
	if err != nil {
		t.Error(err)
	}
	if ot.Len() != 2 {
		t.Error("Len = 2 expected")
	}

	err = ot.Append(nil)
	if err == nil {
		t.Error("items = nil error expected")
	}
	if ot.Len() != 2 {
		t.Error("Len = 2 expected")
	}
}

//----------------------------------------------------------------------------------------

func Test_OperatingItem_Current(t *testing.T) {
	cc := getCurrentCalc()
	opi, _ := NewOperatingItem(cc, 50, 1)
	x1, _ := cc.Current(30, 50)
	x2, _ := opi.Current(30)
	if x1 != x2 {
		t.Error("!=")
	}
}

//----------------------------------------------------------------------------------------

func Example_OperatingItem_Current() {
	opi, _ := NewOperatingItem(getCurrentCalc(), 50, 1)

	curr, err := opi.Current(TA_MIN - 0.01)
	fmt.Printf("%.4f - %v", curr, err)
	// Output:
	// NaN - OperatingItem.Current: ta < TA_MIN
}

func Example_OperatingTable_Current() {
	opi1, _ := NewOperatingItem(getCurrentCalc(), 50, 1)
	opi2, _ := NewOperatingItem(getCurrentCalc(), 60, 1)
	ot, _ := NewOperatingTable([]*OperatingItem{opi1, opi2}, "")

	curr, err := ot.Current(30)
	fmt.Printf("%.4f - %v", curr, err)
	// Output:
	// 435.5280 - <nil>
}

func Example_OperatingTable_Append() {
	opi1, _ := NewOperatingItem(getCurrentCalc(), 50, 1)
	ot, _ := NewOperatingTable([]*OperatingItem{opi1}, "")

	opi2, _ := NewOperatingItem(getCurrentCalc(), 60, 1)
	_ = ot.Append(opi2)

	curr, err := ot.Current(30)
	fmt.Printf("%.4f - %v - %d", curr, err, ot.Len())
	// Output:
	// 435.5280 - <nil> - 2
}
