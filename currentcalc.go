// Copyright Cristian Echeverría Rabí

package conductor

import (
	"math"
)

//----------------------------------------------------------------------------------------

// NewCurrentCalc Returns CurrentCalc object
// c *Conductor: *Conductor instance
func NewCurrentCalc(c *Conductor) (*CurrentCalc, error) {
	if c.diameter <= 0 {
		return nil, &ValueError{"NewCurrentCalc: Conductor.Diameter <= 0"}
	}
	if c.r25 <= 0 {
		return nil, &ValueError{"NewCurrentCalc: Conductor.R25 <= 0"}
	}
	if c.category.alpha <= 0 {
		return nil, &ValueError{"NewCurrentCalc: Conductor.Category.Alpha <= 0"}
	}
	if c.category.alpha >= 1 {
		return nil, &ValueError{"NewCurrentCalc: Conductor.Category.Alpha >=1"}
	}
	return &CurrentCalc{c, 300.0, 2.0, 1.0, 0.5, CF_IEEE, 0.01}, nil
}

//----------------------------------------------------------------------------------------

// CurrentCalc Object to calculate conductor current and temperatures.
type CurrentCalc struct {
	conductor   *Conductor // *Conductor instance
	altitude    float64    // Altitude [m] = 300.0
	airVelocity float64    // Velocity of air stream [ft/seg] =   2.0
	sunEffect   float64    // Sun effect factor (0 to 1) = 1.0
	emissivity  float64    // Emissivity (0 to 1) = 0.5
	formula     string     // Define formula for current calculation = CF_IEEE
	deltaTemp   float64    // Temperature difference to determine equality [°C] = 0.0001
}

func (cc *CurrentCalc) Resistance(tc float64) (float64, error) {
	if tc < TC_MIN {
		return math.NaN(), &ValueError{"CurrentCalc.Resistance: tc < TC_MIN"}
	}
	if tc > TC_MAX {
		return math.NaN(), &ValueError{"CurrentCalc.Resistance: tc > TC_MAX"}
	}
	return cc.conductor.r25 * (1 + cc.conductor.category.alpha*(tc-25.0)), nil
}

func (cc *CurrentCalc) Current(ta float64, tc float64) (float64, error) {
	if ta < TA_MIN {
		return math.NaN(), &ValueError{"CurrentCalc.Current: ta < TA_MIN"}
	}
	if ta > TA_MAX {
		return math.NaN(), &ValueError{"CurrentCalc.Current: ta > TA_MAX"}
	}
	if tc < TC_MIN {
		return math.NaN(), &ValueError{"CurrentCalc.Current: tc < TC_MIN"}
	}
	if tc > TC_MAX {
		return math.NaN(), &ValueError{"CurrentCalc.Current: tc > TC_MAX"}
	}
	if ta >= tc {
		return 0, nil
	}

	D := cc.conductor.diameter / 25.4                                       // Diámetro en pulgadas
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
		return 0, nil
	} else {
		return math.Sqrt((Qc + Qr - Qs) / Rc), nil
	}
}

func (cc *CurrentCalc) Tc(ta float64, ic float64) (float64, error) {
	if ta < TA_MIN {
		return math.NaN(), &ValueError{"CurrentCalc.Tc: ta < TA_MIN"}
	}
	if ta > TA_MAX {
		return math.NaN(), &ValueError{"CurrentCalc.Tc: ta > TA_MAX"}
	}

	icmax, _ := cc.Current(ta, TC_MAX) // No debe haber error: ta está verificado
	if ic < 0 {
		return math.NaN(), &ValueError{"CurrentCalc.Tc: ic < 0"}
	}
	if ic > icmax {
		return math.NaN(), &ValueError{"CurrentCalc.Tc: ic > Imax (TC_MAX)"}
	}

	var tcmin, tcmax, tcmed, icmed float64
	var err error
	tcmin = ta
	tcmax = TC_MAX
	cuenta := 0

	for (tcmax - tcmin) > cc.deltaTemp {
		tcmed = 0.5 * (tcmin + tcmax)
		icmed, err = cc.Current(ta, tcmed)
		if err != nil {
			return math.NaN(), &ValueError{"CurrentCalc.Tc: " + err.Error()}
		}
		if icmed > ic {
			tcmax = tcmed
		} else {
			tcmin = tcmed
		}
		cuenta += 1
		//		if cuenta > ITER_MAX {
		//			return math.NaN(), &ValueError{"CurrentCalc.Tc: ITERA MAX exeeded"}
		//		}
	}
	return tcmed, nil
}

func (cc *CurrentCalc) Ta(tc float64, ic float64) (float64, error) {
	if tc < TC_MIN {
		return math.NaN(), &ValueError{"CurrentCalc.Ta: tc < TC_MIN"}
	}
	if tc > TC_MAX {
		return math.NaN(), &ValueError{"CurrentCalc.Ta: tc > TC_MAX"}
	}

	imin, _ := cc.Current(TA_MAX, tc) // No debe haber error: tc está verificado
	imax, _ := cc.Current(TA_MIN, tc) // No debe haber error: tc está verificado
	if ic < imin {
		return math.NaN(), &ValueError{"CurrentCalc.Ta: ic < Imin (TA_MAX)"}
	}
	if ic > imax {
		return math.NaN(), &ValueError{"CurrentCalc.Ta: ic > Imax (TA_MIN)"}
	}

	var tamin, tamax, tamed, icmed float64
	var err error
	tamin = TA_MIN
	tamax = math.Min(TA_MAX, tc)

	if tamin > tamax {
		return 0, nil
	}

	cuenta := 0

	for (tamax - tamin) > cc.deltaTemp {
		tamed = 0.5 * (tamin + tamax)
		icmed, err = cc.Current(tamed, tc)
		if err != nil {
			return math.NaN(), &ValueError{"CurrentCalc.Ta: " + err.Error()}
		}
		if icmed > ic {
			tamin = tamed
		} else {
			tamax = tamed
		}
		cuenta += 1
		//		if cuenta > ITER_MAX {
		//			return math.NaN(), &ValueError{"CurrentCalc.Ta: ITERA MAX exeeded"}
		//		}
	}
	return tamed, nil
}

func (cc *CurrentCalc) Conductor() *Conductor {
	return cc.conductor
}

func (cc *CurrentCalc) Altitude() float64 {
	return cc.altitude
}

func (cc *CurrentCalc) SetAltitude(h float64) error {
	if h < 0 {
		return &ValueError{"CurrentCalc.SetAltitude: h < 0"}
	}
	cc.altitude = h
	return nil
}

func (cc *CurrentCalc) AirVelocity() float64 {
	return cc.airVelocity
}

func (cc *CurrentCalc) SetAirVelocity(v float64) error {
	if v < 0 {
		return &ValueError{"CurrentCalc.SetAirVelocity: v < 0"}
	}
	cc.airVelocity = v
	return nil
}

func (cc *CurrentCalc) SunEffect() float64 {
	return cc.sunEffect
}

func (cc *CurrentCalc) SetSunEffect(se float64) error {
	if se < 0 {
		return &ValueError{"CurrentCalc.SetSunEffect: se < 0"}
	}
	if se > 1 {
		return &ValueError{"CurrentCalc.SetSunEffect: se > 1"}
	}
	cc.sunEffect = se
	return nil
}

func (cc *CurrentCalc) Emissivity() float64 {
	return cc.emissivity
}

func (cc *CurrentCalc) SetEmissivity(e float64) error {
	if e < 0 {
		return &ValueError{"CurrentCalc.SetEmissivity: e < 0"}
	}
	if e > 1 {
		return &ValueError{"CurrentCalc.SetEmissivity: e > 1"}
	}
	cc.emissivity = e
	return nil
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
	if t <= 0 {
		return &ValueError{"CurrentCalc.SetDeltaTemp: t <= 0"}
	}
	cc.deltaTemp = t
	return nil
}
