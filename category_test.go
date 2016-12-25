// Copyright Cristian Echeverría Rabí

package conductor

import (
	"fmt"
	//"math"
	"testing"
)

//----------------------------------------------------------------------------------------

func TestNewCategory(t *testing.T) {
	cat, err := NewCategory("COPPER", 12000.0, 0.0000169, 0.0, 0.00374, "CU")
	if cat == nil {
		t.Error("category expected")
	}
	if err != nil {
		t.Error("Error not expected")
	}
	if cat.Name() != "COPPER" {
		t.Error("!=")
	}
	if cat.Modelas() != 12000 {
		t.Error("!=")
	}
	if cat.Coefexp() != 0.0000169 {
		t.Error("!=")
	}
	if cat.Creep() != 0.0 {
		t.Error("!=")
	}
	if cat.Alpha() != 0.00374 {
		t.Error("!=")
	}
	if cat.Id() != "CU" {
		t.Error("!=")
	}
}

func TestNewCategoryModelas(t *testing.T) {
	args := CategoryArgs{"COPPER", 12000.0, 0.0000169, 0.0, 0.00374, "CU"}

	args.Modelas = 0.01
	cat, err := args.Get()
	if cat == nil {
		t.Error("category expected")
	}
	if err != nil {
		t.Error("Error not expected")
	}

	args.Modelas = 0
	cat, err = args.Get()
	if cat != nil {
		t.Error("nil expected")
	}
	if err == nil {
		t.Error("Error expected")
	}

	args.Modelas = -0.01
	cat, err = args.Get()
	if cat != nil {
		t.Error("nil expected")
	}
	if err == nil {
		t.Error("Error expected")
	}
}

func TestNewCategoryCoefexp(t *testing.T) {
	args := CategoryArgs{"COPPER", 12000.0, 0.0000169, 0.0, 0.00374, "CU"}

	args.Coefexp = 0.01
	cat, err := args.Get()
	if cat == nil {
		t.Error("category expected")
	}
	if err != nil {
		t.Error("Error not expected")
	}

	args.Coefexp = 0
	cat, err = args.Get()
	if cat != nil {
		t.Error("nil expected")
	}
	if err == nil {
		t.Error("Error expected")
	}

	args.Coefexp = -0.01
	cat, err = args.Get()
	if cat != nil {
		t.Error("nil expected")
	}
	if err == nil {
		t.Error("Error expected")
	}
}

func TestNewCategoryCreep(t *testing.T) {
	args := CategoryArgs{"COPPER", 12000.0, 0.0000169, 0.0, 0.00374, "CU"}

	args.Creep = 0.01
	cat, err := args.Get()
	if cat == nil {
		t.Error("category expected")
	}
	if err != nil {
		t.Error("Error not expected")
	}

	args.Creep = 0
	cat, err = args.Get()
	if cat == nil {
		t.Error("category expected")
	}
	if err != nil {
		t.Error("Error not expected")
	}

	args.Creep = -0.01
	cat, err = args.Get()
	if cat != nil {
		t.Error("nil expected")
	}
	if err == nil {
		t.Error("Error expected")
	}
}

func TestNewCategoryAlpha(t *testing.T) {
	args := CategoryArgs{"COPPER", 12000.0, 0.0000169, 0.0, 0.00374, "CU"}

	args.Alpha = 0.01
	cat, err := args.Get()
	if cat == nil {
		t.Error("category expected")
	}
	if err != nil {
		t.Error("Error not expected")
	}

	args.Alpha = 0
	cat, err = args.Get()
	if cat != nil {
		t.Error("nil expected")
	}
	if err == nil {
		t.Error("Error expected")
	}

	args.Alpha = -0.01
	cat, err = args.Get()
	if cat != nil {
		t.Error("nil expected")
	}
	if err == nil {
		t.Error("Error expected")
	}

	args.Alpha = 0.99
	cat, err = args.Get()
	if cat == nil {
		t.Error("category expected")
	}
	if err != nil {
		t.Error("Error not expected")
	}

	args.Alpha = 1.0
	cat, err = args.Get()
	if cat != nil {
		t.Error("nil expected")
	}
	if err == nil {
		t.Error("Error expected")
	}

	args.Alpha = 1.01
	cat, err = args.Get()
	if cat != nil {
		t.Error("nil expected")
	}
	if err == nil {
		t.Error("Error expected")
	}
}

func TestCategoryArgsContants(t *testing.T) {
	_, err := CC_CU.Get()
	if err != nil {
		t.Errorf("CC_CU Error: %v", err)
	}
	_, err = CC_AAAC.Get()
	if err != nil {
		t.Errorf("CC_AAAC Error: %v", err)
	}
	_, err = CC_ACAR.Get()
	if err != nil {
		t.Errorf("CC_ACAR Error: %v", err)
	}
	_, err = CC_ACSR.Get()
	if err != nil {
		t.Errorf("CC_ACSR Error: %v", err)
	}
	_, err = CC_AAC.Get()
	if err != nil {
		t.Errorf("CC_AAC Error: %v", err)
	}
	_, err = CC_CUWELD.Get()
	if err != nil {
		t.Errorf("CC_CUWELD Error: %v", err)
	}
	_, err = CC_AASC.Get()
	if err != nil {
		t.Errorf("CC_AASC Error: %v", err)
	}
	_, err = CC_ALL.Get()
	if err != nil {
		t.Errorf("CC_ALL Error: %v", err)
	}
}

//----------------------------------------------------------------------------------------

func ExampleCategory() {
	CC_CU.Modelas = 0
	_, err := CC_CU.Get()
	fmt.Printf("%q", err)
	// Output:
	// "CategoryArgs Get: \n  NewCategory: Modelas (0) required value > 0, "
}
