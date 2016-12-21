// Copyright Cristian Echeverría Rabí

package conductor

//----------------------------------------------------------------------------------------

// Conductor Container for conductor characteristics
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
