// Copyright Cristian Echeverría Rabí

package conductor

//----------------------------------------------------------------------------------------

// CategoryArgs Container for arguments for Category constructor
// Is a mutable version of a Category object that allow to change attributes
type CategoryArgs struct {
	Name    string
	Modelas float64
	Coefexp float64
	Creep   float64
	Alpha   float64
	Id      string
}

// Get Returns *category object from attributes values
func (ca *CategoryArgs) Get() *Category {
	return &Category{ca.Name, ca.Modelas, ca.Coefexp, ca.Creep, ca.Alpha, ca.Id}
}

//----------------------------------------------------------------------------------------

// NewCategory Returns *category object from arguments values
// name    string  : Name of conductor category
// modelas float64 : Modulus of elasticity [kg/mm2] (required modelas > 0)
// coefexp float64 : Coefficient of Thermal Expansion [1/°C] (required coefexp > 0)
// creep   float64 : Creep °C (required creep >= 0)
// alpha   float64 : Temperature coefficient of resistance [1/°C] (required 0 < alpha < 1)
// id      string  : Database id
func NewCategory(name string, modelas float64, coefexp float64, creep float64, alpha float64,
	id string) *Category {
	return &Category{name, modelas, coefexp, creep, alpha, id}
}

//----------------------------------------------------------------------------------------

// Category Container for category characteristics. Groups similar conductors.
type Category struct {
	name    string  // Name of conductor category
	modelas float64 // Modulus of elasticity [kg/mm2]
	coefexp float64 // Coefficient of Thermal Expansion [1/°C]
	creep   float64 // Creep °C
	alpha   float64 // Temperature coefficient of resistance [1/°C]
	id      string  // Database id
}

func (cat *Category) Name() string {
	return cat.name
}
func (cat *Category) Modelas() float64 {
	return cat.modelas
}
func (cat *Category) Coefexp() float64 {
	return cat.coefexp
}
func (cat *Category) Creep() float64 {
	return cat.creep
}
func (cat *Category) Alpha() float64 {
	return cat.alpha
}
func (cat *Category) Id() string {
	return cat.id
}

//----------------------------------------------------------------------------------------
// *Category instances to use as constants
//	CC_CU     = &Category{"COPPER", 12000.0, 0.0000169, 0.0, 0.00374, "CU"}
//	CC_AAAC   = &Category{"AAAC (AASC)", 6450.0, 0.0000230, 20.0, 0.00340, "AAAC"}
//	CC_ACAR   = &Category{"ACAR", 6450.0, 0.0000250, 20.0, 0.00385, "ACAR"}
//	CC_ACSR   = &Category{"ACSR", 8000.0, 0.0000191, 20.0, 0.00395, "ACSR"}
//	CC_AAC    = &Category{"ALUMINUM", 5600.0, 0.0000230, 20.0, 0.00395, "AAC"}
//	CC_CUWELD = &Category{"COPPERWELD", 16200.0, 0.0000130, 0.0, 0.00380, "CUWELD"}
//	CC_AASC   = CC_AAAC
//	CC_ALL    = CC_AAC

var (
	CC_CU     = &Category{"COPPER", 12000.0, 0.0000169, 0.0, 0.00374, "CU"}
	CC_AAAC   = &Category{"AAAC (AASC)", 6450.0, 0.0000230, 20.0, 0.00340, "AAAC"}
	CC_ACAR   = &Category{"ACAR", 6450.0, 0.0000250, 20.0, 0.00385, "ACAR"}
	CC_ACSR   = &Category{"ACSR", 8000.0, 0.0000191, 20.0, 0.00395, "ACSR"}
	CC_AAC    = &Category{"ALUMINUM", 5600.0, 0.0000230, 20.0, 0.00395, "AAC"}
	CC_CUWELD = &Category{"COPPERWELD", 16200.0, 0.0000130, 0.0, 0.00380, "CUWELD"}
	CC_AASC   = CC_AAAC
	CC_ALL    = CC_AAC
)
