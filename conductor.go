// Copyright Cristian Echeverría Rabí

package conductor

import (
	"fmt"

	"bitbucket.org/tormundo/go.value/checker"
)

//----------------------------------------------------------------------------------------

// ConductorArgs Container for arguments for Conductor constructor
// Is a mutable version of a Conductor object that allow to change attributes
type ConductorArgs struct {
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
func (ca *ConductorArgs) Get() (*Conductor, error) {
	vc := checker.New(fmt.Sprintf("ConductorArgs Get %s", ca.Name))
	cond, errsub := NewConductor(ca.Name, ca.Category, ca.Diameter, ca.Area, ca.Weight,
		ca.Strength, ca.R25, ca.Hcap, ca.Id)
	vc.AppendSub(errsub)
	if err := vc.Error(); err != nil {
		return nil, err
	}
	return cond, nil
}

//----------------------------------------------------------------------------------------

// NewConductor Returns *conductor object from arguments values
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
	strength float64, r25 float64, hcap float64, id string) (*Conductor, error) {

	vc := checker.New(fmt.Sprintf("NewConductor %s", name))
	vc.Ck("Diameter", diameter).Gt(0.0)
	vc.Ck("Area", area).Gt(0.0)
	vc.Ck("Weight", weight).Gt(0.0)
	vc.Ck("Strength", strength).Gt(0.0)
	vc.Ck("R25", r25).Gt(0.0)
	vc.Ck("Hcap", hcap).Gt(0.0)
	if err := vc.Error(); err != nil {
		return nil, err
	}
	return &Conductor{*category, name, diameter, area, weight, strength, r25, hcap, id}, nil
}

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

// Conductor Container for conductor characteristics
type conductor struct {
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

func (c *conductor) Name() string {
	return c.name
}
func (c *conductor) Category() *Category {
	return c.category
}
func (c *conductor) Diameter() float64 {
	return c.diameter
}
func (c *conductor) Area() float64 {
	return c.area
}
func (c *conductor) Weight() float64 {
	return c.weight
}
func (c *conductor) Strength() float64 {
	return c.strength
}
func (c *conductor) R25() float64 {
	return c.r25
}
func (c *conductor) Hcap() float64 {
	return c.hcap
}
func (c *conductor) Id() string {
	return c.id
}
