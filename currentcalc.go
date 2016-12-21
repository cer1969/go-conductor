// Copyright Cristian Echeverría Rabí

package conductor

import (
	"math"

	"bitbucket.org/tormundo/go.value/checker"
)

//----------------------------------------------------------------------------------------

// NewCurrentCalc Returns CurrentCalc object
// c Conductor: Conductor instance
//
// Conductor fields R25, Diameter and Alpha are copied to CurrentCal.
// Subsequent changes in the Conductor will not be reflected in CurrentCalc.
func NewCurrentCalc(c Conductor) (*CurrentCalc, error) {
	vc := checker.New("NewCurrentCalc")
	vc.Ck("R25", c.R25).Gt(0.0)
	vc.Ck("Diameter", c.Diameter).Gt(0.0)
	vc.Ck("Alpha", c.Alpha).Gt(0.0).Lt(1.0)

	err := vc.Error()
	if err != nil {
		return nil, err
	}

	return &CurrentCalc{c.R25, c.Diameter, c.Alpha, 300.0, 2.0, 1.0, 0.5, CF_IEEE, 0.01}, nil
}

//----------------------------------------------------------------------------------------

// CurrentCalc Object to calculate conductor current and temperatures.
type CurrentCalc struct {
	r25         float64 // Conductor.R25
	diameter    float64 // Conductor.Diameter
	alpha       float64 // Conductor Alpha
	altitude    float64 // Altitude [m] = 300.0
	airVelocity float64 // Velocity of air stream [ft/seg] =   2.0
	sunEffect   float64 // Sun effect factor (0 to 1) = 1.0
	emissivity  float64 // Emissivity (0 to 1) = 0.5
	formula     string  // Define formula for current calculation = CF_IEEE
	deltaTemp   float64 // Temperature difference to determine equality [°C] = 0.0001
}

func (cc *CurrentCalc) Resistance(tc float64) (float64, error) {
	vc := checker.New("CurrentCalc Resistance")
	vc.Ck("tc", tc).Ge(TC_MIN).Le(TC_MAX)

	err := vc.Error()
	if err != nil {
		return math.NaN(), err
	}

	return cc.r25 * (1 + cc.alpha*(tc-25.0)), nil
}

func (cc *CurrentCalc) Current(ta float64, tc float64) (q float64, err error) {
	q = math.NaN()

	vc := checker.New("CurrentCalc Current")
	vc.Ck("ta", ta).Ge(TA_MIN).Le(TA_MAX)
	vc.Ck("tc", tc).Ge(TC_MIN).Le(TC_MAX)
	err = vc.Error()
	if err != nil {
		return
	}

	if ta >= tc {
		q = 0.0
		return
	}

	D := cc.diameter / 25.4                                                 // Diámetro en pulgadas
	Pb := math.Pow(10, (1.880813592 - cc.altitude/18336.0))                 // Presión barométrica en cmHg
	V := cc.airVelocity * 3600                                              // Vel. viento en pies/hora
	res, _ := cc.Resistance(tc)                                             // No necesito verificar error porque el valor de tc ya se verificó
	Rc := res * 0.0003048                                                   // Resistencia en ohm/pies
	Tm := 0.5 * (tc + ta)                                                   // Temperatura media
	Rf := 0.2901577 * Pb / (273 + Tm)                                       // Densidad rel.aire [lb/ft^3]
	Uf := 0.04165 + 0.000111*Tm                                             // Viscosidad abs. aire [lb/(ft x hora)]
	Kf := 0.00739 + 0.0000227*Tm                                            // Coef. conductividad term. aire [Watt/(ft x °C)]
	Qc := 0.283 * math.Sqrt(Rf) * math.Pow(D, 0.75) * math.Pow(tc-ta, 1.25) // watt/ft

	if V != 0 {
		factor := D * Rf * V / Uf
		Qc1 := 0.1695 * Kf * (tc - ta) * math.Pow(factor, 0.6)
		Qc2 := Kf * (tc - ta) * (1.01 + 0.371*math.Pow(factor, 0.52))
		if cc.formula == CF_IEEE {
			// IEEE criteria
			Qc = math.Max(Qc, Qc1)
			Qc = math.Max(Qc, Qc2)
		} else {
			// CLASSIC criteria
			if factor < 12000 {
				Qc = Qc2
			} else {
				Qc = Qc1
			}
		}
	}
	LK := math.Pow((tc+273)/100, 4)
	MK := math.Pow((ta+273)/100, 4)
	Qr := 0.138 * D * cc.emissivity * (LK - MK)
	Qs := 3.87 * D * cc.sunEffect

	if (Qc + Qr) < Qs {
		q = 0.0
	} else {
		q = math.Sqrt((Qc + Qr - Qs) / Rc)
	}
	return
}

func (cc *CurrentCalc) Tc(ta float64, ic float64) (tc float64, err error) {
	tc = math.NaN()

	vc := checker.New("CurrentCalc Tc")
	vc.Ck("ta", ta).Ge(TA_MIN).Le(TA_MAX)
	err = vc.Error()
	if err != nil {
		return
	}

	icmax, _ := cc.Current(ta, TC_MAX) // No debe haber error: ta está verificado

	vc.Ck("ic", ic).Ge(0).Le(icmax) // Asegura valor de ta <= tc <= TC_MAX
	err = vc.Error()
	if err != nil {
		return
	}

	var tcmin, tcmax, tcmed, icmed float64
	var err2 error
	tcmin = ta
	tcmax = TC_MAX
	cuenta := 0

	for (tcmax - tcmin) > cc.deltaTemp {
		tcmed = 0.5 * (tcmin + tcmax)
		icmed, err2 = cc.Current(ta, tcmed)
		if err2 != nil {
			vc.Append(err2.Error())
			return
		}
		if icmed > ic {
			tcmax = tcmed
		} else {
			tcmin = tcmed
		}
		cuenta += 1
		if cuenta > ITER_MAX {
			vc.Append("CurrentCalc Tc ITERA MAX exeeded")
			return
		}
	}
	tc = tcmed
	return
}

func (cc *CurrentCalc) Ta(tc float64, ic float64) (ta float64, err error) {
	ta = math.NaN()

	vc := checker.New("CurrentCalc Ta")
	vc.Ck("tc", tc).Ge(TC_MIN).Le(TC_MAX)
	err = vc.Error()
	if err != nil {
		return
	}

	imin, _ := cc.Current(TA_MAX, tc) // No debe haber error: tc está verificado
	imax, _ := cc.Current(TA_MIN, tc) // No debe haber error: tc está verificado
	vc.Ck("ic", ic).Ge(imin).Le(imax) // Asegura valor de TA_MIN <= ta <= TA_MAX
	err = vc.Error()
	if err != nil {
		return
	}

	var tamin, tamax, tamed, icmed float64
	var err2 error
	tamin = TA_MIN
	tamax = math.Min(TA_MAX, tc)

	if tamin > tamax {
		ta = 0.0
		return
	}

	cuenta := 0

	for (tamax - tamin) > cc.deltaTemp {
		tamed = 0.5 * (tamin + tamax)
		icmed, err2 = cc.Current(tamed, tc)
		if err2 != nil {
			vc.Append(err2.Error())
			return
		}
		if icmed > ic {
			tamin = tamed
		} else {
			tamax = tamed
		}
		cuenta += 1
		if cuenta > ITER_MAX {
			vc.Append("CurrentCalc Ta ITERA MAX exeeded")
			return
		}
	}
	ta = tamed
	return
}

func (cc *CurrentCalc) Altitude() float64 {
	return cc.altitude
}

func (cc *CurrentCalc) SetAltitude(h float64) error {
	vc := checker.New("CurrentCalc SetAltitude")
	vc.Ck("h", h).Ge(0)

	cc.altitude = h

	return vc.Error()
}

func (cc *CurrentCalc) AirVelocity() float64 {
	return cc.airVelocity
}

func (cc *CurrentCalc) SetAirVelocity(v float64) error {
	vc := checker.New("CurrentCalc SetAirVelocity")
	vc.Ck("v", v).Ge(0)

	cc.airVelocity = v

	return vc.Error()
}

func (cc *CurrentCalc) SunEffect() float64 {
	return cc.sunEffect
}

func (cc *CurrentCalc) SetSunEffect(se float64) error {
	vc := checker.New("CurrentCalc SetSunEffect")
	vc.Ck("se", se).Ge(0).Le(1)

	cc.sunEffect = se

	return vc.Error()
}

func (cc *CurrentCalc) Emissivity() float64 {
	return cc.emissivity
}

func (cc *CurrentCalc) SetEmissivity(e float64) error {
	vc := checker.New("CurrentCalc SetEmissivity")
	vc.Ck("e", e).Ge(0).Le(1)

	cc.emissivity = e

	return vc.Error()
}

func (cc *CurrentCalc) Formula() string {
	return cc.formula
}

func (cc *CurrentCalc) SetFormula(f string) {
	if f != CF_CLASSIC {
		f = CF_IEEE
	}
	cc.formula = f
}

func (cc *CurrentCalc) DeltaTemp() float64 {
	return cc.deltaTemp
}

func (cc *CurrentCalc) SetDeltaTemp(t float64) error {
	vc := checker.New("CurrentCalc SetDeltaTemp")
	vc.Ck("t", t).Gt(0)

	cc.deltaTemp = t

	return vc.Error()
}
