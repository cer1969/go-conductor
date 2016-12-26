// Copyright Cristian Echeverría Rabí

package conductor

import (
	//"fmt"
	"testing"
)

//----------------------------------------------------------------------------------------

func TestNewConductor(t *testing.T) {
	//cu300, err := NewConductor("CU 300 MCM", cx.CC_CU, 15.95, 152.00, 1.378, 6123.0, 0.12270, 1e-10, "")

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

//func TestNewCategoryModelas(t *testing.T) {
//	args := CategoryArgs{"COPPER", 12000.0, 0.0000169, 0.0, 0.00374, "CU"}

//	args.Modelas = 0.01
//	cat, err := args.Get()
//	if cat == nil {
//		t.Error("category expected")
//	}
//	if err != nil {
//		t.Error("Error not expected")
//	}

//	args.Modelas = 0
//	cat, err = args.Get()
//	if cat != nil {
//		t.Error("nil expected")
//	}
//	if err == nil {
//		t.Error("Error expected")
//	}

//	args.Modelas = -0.01
//	cat, err = args.Get()
//	if cat != nil {
//		t.Error("nil expected")
//	}
//	if err == nil {
//		t.Error("Error expected")
//	}
//}

//func TestNewCategoryCoefexp(t *testing.T) {
//	args := CategoryArgs{"COPPER", 12000.0, 0.0000169, 0.0, 0.00374, "CU"}

//	args.Coefexp = 0.01
//	cat, err := args.Get()
//	if cat == nil {
//		t.Error("category expected")
//	}
//	if err != nil {
//		t.Error("Error not expected")
//	}

//	args.Coefexp = 0
//	cat, err = args.Get()
//	if cat != nil {
//		t.Error("nil expected")
//	}
//	if err == nil {
//		t.Error("Error expected")
//	}

//	args.Coefexp = -0.01
//	cat, err = args.Get()
//	if cat != nil {
//		t.Error("nil expected")
//	}
//	if err == nil {
//		t.Error("Error expected")
//	}
//}

//func TestNewCategoryCreep(t *testing.T) {
//	args := CategoryArgs{"COPPER", 12000.0, 0.0000169, 0.0, 0.00374, "CU"}

//	args.Creep = 0.01
//	cat, err := args.Get()
//	if cat == nil {
//		t.Error("category expected")
//	}
//	if err != nil {
//		t.Error("Error not expected")
//	}

//	args.Creep = 0
//	cat, err = args.Get()
//	if cat == nil {
//		t.Error("category expected")
//	}
//	if err != nil {
//		t.Error("Error not expected")
//	}

//	args.Creep = -0.01
//	cat, err = args.Get()
//	if cat != nil {
//		t.Error("nil expected")
//	}
//	if err == nil {
//		t.Error("Error expected")
//	}
//}

//func TestNewCategoryAlpha(t *testing.T) {
//	args := CategoryArgs{"COPPER", 12000.0, 0.0000169, 0.0, 0.00374, "CU"}

//	args.Alpha = 0.01
//	cat, err := args.Get()
//	if cat == nil {
//		t.Error("category expected")
//	}
//	if err != nil {
//		t.Error("Error not expected")
//	}

//	args.Alpha = 0
//	cat, err = args.Get()
//	if cat != nil {
//		t.Error("nil expected")
//	}
//	if err == nil {
//		t.Error("Error expected")
//	}

//	args.Alpha = -0.01
//	cat, err = args.Get()
//	if cat != nil {
//		t.Error("nil expected")
//	}
//	if err == nil {
//		t.Error("Error expected")
//	}

//	args.Alpha = 0.99
//	cat, err = args.Get()
//	if cat == nil {
//		t.Error("category expected")
//	}
//	if err != nil {
//		t.Error("Error not expected")
//	}

//	args.Alpha = 1.0
//	cat, err = args.Get()
//	if cat != nil {
//		t.Error("nil expected")
//	}
//	if err == nil {
//		t.Error("Error expected")
//	}

//	args.Alpha = 1.01
//	cat, err = args.Get()
//	if cat != nil {
//		t.Error("nil expected")
//	}
//	if err == nil {
//		t.Error("Error expected")
//	}
//}

////----------------------------------------------------------------------------------------

//func ExampleCategory() {
//	_, err := NewCategory("COPPER", 0.0, 0.0000169, 0.0, 0.00374, "CU")
//	fmt.Print(err)
//	// Output:
//	// NewCategory COPPER: Modelas (0) required value > 0,
//}
