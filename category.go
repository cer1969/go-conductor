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

// Get Returns Category object from attributes values
func (ca *CategoryArgs) Get() (*category, error) {
	vc := checker.New("CategoryArgs Get")
	cat, err := NewCategory(ca.Name, ca.Modelas, ca.Coefexp, ca.Creep, ca.Alpha, ca.Id)
	if err != nil {
		vc.AppendSub(err)
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
func NewCategory(name string, modelas float64, coefexp float64, creep float64, alpha float64, id string) (*category, error) {
	vc := checker.New("NewCategory")
	vc.Ck("Modelas", modelas).Gt(0.0)
	vc.Ck("Coefexp", coefexp).Gt(0.0)
	vc.Ck("Creep", creep).Ge(0.0)
	vc.Ck("Alpha", alpha).Gt(0.0).Lt(1.0)
	err := vc.Error()
	if err != nil {
		return nil, err
	}
	return &category{name, modelas, coefexp, creep, alpha, id}, nil
}

//----------------------------------------------------------------------------------------

// Category Container for category characteristics. Groups similar conductors.
type category struct {
	name    string  // Name of conductor category
	modelas float64 // Modulus of elasticity [kg/mm2]
	coefexp float64 // Coefficient of Thermal Expansion [1/°C]
	creep   float64 // Creep °C
	alpha   float64 // Temperature coefficient of resistance [1/°C]
	id      string  // Optional database id
}

func (cat *category) Name() string {
	return cat.name
}
func (cat *category) Modelas() float64 {
	return cat.modelas
}
func (cat *category) Coefexp() float64 {
	return cat.coefexp
}
func (cat *category) Creep() float64 {
	return cat.creep
}
func (cat *category) Alpha() float64 {
	return cat.alpha
}
func (cat *category) Id() string {
	return cat.id
}

// CategoryArgs instances to use as constants
var (
	CC_CU     = CategoryArgs{"COPPER", 12000.0, 0.0000169, 0.0, 0.00374, "CU"}
	CC_AAAC   = CategoryArgs{"AAAC (AASC)", 6450.0, 0.0000230, 20.0, 0.00340, "AAAC"}
	CC_ACAR   = CategoryArgs{"ACAR", 6450.0, 0.0000250, 20.0, 0.00385, "ACAR"}
	CC_ACSR   = CategoryArgs{"ACSR", 8000.0, 0.0000191, 20.0, 0.00395, "ACSR"}
	CC_AAC    = CategoryArgs{"ALUMINUM", 5600.0, 0.0000230, 20.0, 0.00395, "AAC"}
	CC_CUWELD = CategoryArgs{"COPPERWELD", 16200.0, 0.0000130, 0.0, 0.00380, "CUWELD"}
	CC_AASC   = CC_AAAC
	CC_ALL    = CC_AAC
)
