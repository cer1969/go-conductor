// Copyright Cristian Echeverría Rabí

// Calculations for high voltage conductors
package conductor

//----------------------------------------------------------------------------------------

// Define constats for conductor library

// Formula to use in CurrentCalc for current calculations
// CF_CLASSIC = "CLASSIC"    Identifies CLASSIC formula
// CF_IEEE    = "IEEE"       Identifies IEEE formula
//
// Ambient temperature in °C
// TA_MIN = -90    Minimum value for ambient temperature
//                 World lowest -82.2°C Vostok Antartica 21/07/1983
// TA_MAX =  90    Maximum value for ambient temperature
//                 World highest 58.2°C Libia 13/09/1922
//
// Conductor temperature [°C]
// TC_MIN =  -90    Minimum value for conductor temperature
// TC_MAX = 2000    Maximum value for conductor temperature = 2000°C
//                 Copper melt at 1083 °C
//
// Iterations
// ITER_MAX = 20000    Maximum iterations number = 20000
//
// Conductor tension [kg]
// TENSION_MAX = 50000    Maximum conductor tension
//
const (
	TA_MIN      = -90.0
	TA_MAX      = 90.0
	TC_MIN      = -90.0
	TC_MAX      = 2000.0
	ITER_MAX    = 20000
	TENSION_MAX = 50000.0
	CF_CLASSIC  = "CLASSIC"
	CF_IEEE     = "IEEE"
)
