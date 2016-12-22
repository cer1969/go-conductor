// Copyright Cristian Echeverría Rabí

package conductor

import (
	"bitbucket.org/tormundo/go.value/checker"
)

//----------------------------------------------------------------------------------------

// CategoryArgs Container for arguments for Category constructor
// Is a mutable version of a Category objects that allow to change attributes
type CategoryArgs struct {
	Name    string
	Modelas float64
	Coefexp float64
	Creep   float64
	Alpha   float64
	Id      string
}

func (ca *CategoryArgs) Get() (*Category, error) {
	vc := checker.New("CategoryArgs Get")
	cat, err1 := NewCategory(ca.Name, ca.Modelas, ca.Coefexp, ca.Creep, ca.Alpha, ca.Id)
	if err1 != nil {
		if ckerr, ok := vc.Error().(*checker.CheckError); ok {
			vc.AppendCheckError(ckerr)
		} else {
			vc.Append("Internal error")
		}
		return nil, vc.Error()
	}
	return cat, nil
}

//----------------------------------------------------------------------------------------

// NewCategory Returns Category object from attributes values
// name string     : Name of conductor category
// modelas float64 : Modulus of elasticity [kg/mm2] (required modelas > 0)
// coefexp float64 : Coefficient of Thermal Expansion [1/°C] (required coefexp > 0)
// creep float64   : Creep °C (required creep >= 0)
// alpha float64   : Temperature coefficient of resistance [1/°C] (required 0 < alpha < 1)
// id string       : Database id
func NewCategory(name string, modelas float64, coefexp float64, creep float64, alpha float64, id string) (*Category, error) {
	vc := checker.New("NewCategory")
	vc.Ck("Modelas", modelas).Gt(0.0)
	vc.Ck("Coefexp", coefexp).Gt(0.0)
	vc.Ck("Creep", creep).Ge(0.0)
	vc.Ck("Alpha", alpha).Gt(0.0).Lt(1.0)
	err := vc.Error()
	if err != nil {
		return nil, err
	}
	return &Category{name, modelas, coefexp, creep, alpha, id}, nil
}

// NewCategoryFromArgs Returns Category object from CategoryArgs object
func NewCategoryFromArgs(ca CategoryArgs) (*Category, error) {
	vc := checker.New("NewCategoryFromArgs")
	cat, err1 := NewCategory(ca.Name, ca.Modelas, ca.Coefexp, ca.Creep, ca.Alpha, ca.Id)
	if err1 != nil {
		if ckerr, ok := vc.Error().(*checker.CheckError); ok {
			vc.AppendCheckError(ckerr)
		} else {
			vc.Append("Internal error")
		}
		return nil, vc.Error()
	}
	return cat, nil
}

//----------------------------------------------------------------------------------------

// Category Container for category characteristics. Groups similar conductors.
type Category struct {
	Name    string  // Name of conductor category
	Modelas float64 // Modulus of elasticity [kg/mm2]
	Coefexp float64 // Coefficient of Thermal Expansion [1/°C]
	Creep   float64 // Creep °C
	Alpha   float64 // Temperature coefficient of resistance [1/°C]
	Id      string  // Optional database id
}

// Category instances to use as constants
var (
	CC_CU     = Category{"COPPER", 12000.0, 0.0000169, 0.0, 0.00374, "CU"}
	CC_AAAC   = Category{"AAAC (AASC)", 6450.0, 0.0000230, 20.0, 0.00340, "AAAC"}
	CC_ACAR   = Category{"ACAR", 6450.0, 0.0000250, 20.0, 0.00385, "ACAR"}
	CC_ACSR   = Category{"ACSR", 8000.0, 0.0000191, 20.0, 0.00395, "ACSR"}
	CC_AAC    = Category{"ALUMINUM", 5600.0, 0.0000230, 20.0, 0.00395, "AAC"}
	CC_CUWELD = Category{"COPPERWELD", 16200.0, 0.0000130, 0.0, 0.00380, "CUWELD"}
	CC_AASC   = CC_AAAC
	CC_ALL    = CC_AAC
)
