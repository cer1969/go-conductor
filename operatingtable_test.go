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
	// NaN - OperatingItem.Current: CurrentCalc.Current: ta < TA_MIN
}
