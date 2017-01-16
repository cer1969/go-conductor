// Copyright Cristian Echeverría Rabí

package conductor

//----------------------------------------------------------------------------------------

// ConductorMaker Container for arguments for Conductor constructor
// Is a mutable version of a Conductor object that allow to change attributes
type ConductorMaker struct {
	Name     string
	Category *Category
	Diameter float64
	Area     float64
	Weight   float64
	Strength float64
	R25      float64
	Hcap     float64
	Id       string
}

// Get Returns *conductor object from attributes values
func (ca *ConductorMaker) Get() *Conductor {
	return &Conductor{ca.Name, ca.Category, ca.Diameter, ca.Area, ca.Weight, ca.Strength, ca.R25,
		ca.Hcap, ca.Id}
}

//----------------------------------------------------------------------------------------

// NewConductor Returns *Conductor object from arguments values
// name     string    : Name of conductor
// category *category : *category instance
// diameter float64   : Diameter [mm]
// area     float64   : Cross section area [mm2]
// weight   float64   : Weight per unit [kg/m]
// strength float64   : Rated strength [kg]
// r25      float64   : Resistance at 25°C [Ohm/km]
// hcap     float64   : Heat capacity [kcal/(ft*°C)]
// id       string    : Database id
func NewConductor(name string, category *Category, diameter float64, area float64, weight float64,
	strength float64, r25 float64, hcap float64, id string) *Conductor {
	return &Conductor{name, category, diameter, area, weight, strength, r25, hcap, id}
}

//----------------------------------------------------------------------------------------

// Conductor Container for conductor characteristics
type Conductor struct {
	name     string    // Name of conductor
	category *Category // *category instance
	diameter float64   // Diameter [mm]
	area     float64   // Cross section area [mm2]
	weight   float64   // Weight per unit [kg/m]
	strength float64   // Rated strength [kg]
	r25      float64   // Resistance at 25°C [Ohm/km]
	hcap     float64   // Heat capacity [kcal/(ft*°C)]
	id       string    // Optional database id
}

func (c *Conductor) Name() string {
	return c.name
}
func (c *Conductor) Category() *Category {
	return c.category
}
func (c *Conductor) Diameter() float64 {
	return c.diameter
}
func (c *Conductor) Area() float64 {
	return c.area
}
func (c *Conductor) Weight() float64 {
	return c.weight
}
func (c *Conductor) Strength() float64 {
	return c.strength
}
func (c *Conductor) R25() float64 {
	return c.r25
}
func (c *Conductor) Hcap() float64 {
	return c.hcap
}
func (c *Conductor) Id() string {
	return c.id
}
