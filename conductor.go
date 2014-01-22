// Copyright Cristian Echeverría Rabí

// Calculations for high voltage conductors
package conductor

// Constants values
const (
	TA_MIN      = -90       // Minumum value for ambient temperature (World lowest -82.2°C Vostok Antartica 21/07/1983)
	TA_MAX      = 90        // Maximum value for ambient temperature (World highest 58.2°C Libia 13/09/1922)
	TC_MIN      = -90       // Minumum value for conductor temperature
	TC_MAX      = 2000      // Maximum value for conductor temperature (Copper melt at 1083 °C)
	ITER_MAX    = 20000     // Maximun iterations number
	TENSION_MAX = 50000     // Maximun condutor tension
	CF_CLASSIC  = "CLASSIC" // CLASSIC Formula
	CF_IEEE     = "IEEE"    // IEEE Formula
)

type Category struct {
	Name    string  // Name of conductor category
	Modelas float64 // Modulus of elasticity [kg/mm2]
	Coefexp float64 // Coefficient of Thermal Expansion [1/°C]
	Creep   float64 // Creep °C
	Alpha   float64 // Temperature coefficient of resistance [1/°C]
	Id      string  // Optional database id
}

type Conductor struct {
	Category         // Category
	Name     string  // Name of conductor
	Diameter float64 // Diameter [mm]
	Area     float64 // Cross section area [mm2]
	Weight   float64 // Weight per unit [kg/m]
	Strength float64 // Rated strength [kg]
	R25      float64 // Resistance at 25°C [Ohm/km]
	Hcap     float64 // Heat capacity [kcal/(ft*°C)]
	Id       string  // Optional database id
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
