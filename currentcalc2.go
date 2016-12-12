// Copyright Cristian Echeverría Rabí

package conductor

import (
	"math"

	"bitbucket.org/tormundo/go.utils/values"
)

//----------------------------------------------------------------------------------------

func NewCurrentCalc2(cond Conductor) *CurrentCalc2 {
	//vc := values.Checker("NewCurrentCalc2")
	//vc.Val("R25", cond.R25).Gt(0.0)
	//vc.Val("Diameter", cond.Diameter).Gt(0.0)
	//vc.Val("Alpha", cond.Alpha).Gt(0.0).Lt(1.0)
	//err := vc.Error()
	//if err != nil {
	//	return nil, err
	//}
	return &CurrentCalc2{cond, 300.0, 2.0, 1.0, 0.5, CF_IEEE, 0.0001}
}

//----------------------------------------------------------------------------------------

type CurrentCalc2 struct {
	Conductor           // Conductor instance
	altitude    float64 // Altitude [m] = 300.0
	airVelocity float64 // Velocity of air stream [ft/seg] =   2.0
	sunEffect   float64 // Sun effect factor (0 to 1) = 1.0
	emissivity  float64 // Emissivity (0 to 1) = 0.5
	formula     string  // Define formula for current calculation = CF_IEEE
	deltaTemp   float64 // Temperature difference to determine equality [°C] = 0.0001
}

func (cc *CurrentCalc2) Resistance(tc float64) float64 {
	vc := values.SimpleChecker()
	vc.Val(cc.R25).Gt(0)
	vc.Val(cc.Alpha).Gt(0.0).Lt(1.0)

	if !vc.Ok() {
		return math.NaN()
	}

	r := cc.R25 * (1 + cc.Alpha*(tc-25.0))

	if r < 0 {
		return 0.0
	}

	return r
}

func (cc *CurrentCalc2) Current(ta float64, tc float64) float64 {
	if ta >= tc {
		return 0.0
	}
	vc := values.SimpleChecker()
	//vc.Val(ta).Ge(TA_MIN).Le(TA_MAX)
	//vc.Val(tc).Ge(TC_MIN).Le(TC_MAX)
	vc.Val(cc.Diameter).Gt(0.0)
	vc.Val(cc.altitude).Ge(0.0)
	vc.Val(cc.airVelocity).Ge(0.0)
	vc.Val(cc.emissivity).Ge(0.0).Le(1.0)
	vc.Val(cc.sunEffect).Ge(0.0).Le(1.0)
	if !vc.Ok() {
		return math.NaN()
	}

	D := cc.Diameter / 25.4                                                 // Diámetro en pulgadas
	Pb := math.Pow(10, (1.880813592 - cc.altitude/18336.0))                 // Presión barométrica en cmHg
	V := cc.airVelocity * 3600                                              // Vel. viento en pies/hora
	res := cc.Resistance(tc)                                                // No necesito verificar error porque el valor de tc ya se verificó
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
		return 0.0
	}
	return math.Sqrt((Qc + Qr - Qs) / Rc)
}

func (cc *CurrentCalc2) Tc(ta float64, ic float64) (tc float64, err error) {
	vc := values.Checker("CurrentCalc Tc")
	vc.Val("ta", ta).Ge(TA_MIN).Le(TA_MAX)

	tc = math.NaN()
	err = vc.Error()
	if err != nil {
		return
	}

	icmax := cc.Current(ta, TC_MAX)
	//if err != nil {
	//	return
	//}

	vc.Val("ic", ic).Ge(0).Le(icmax) // Asegura valor de ta <= tc <= TC_MAX
	err = vc.Error()
	if err != nil {
		return
	}

	var tcmin, tcmax, tcmed, icmed float64
	tcmin = ta
	tcmax = TC_MAX
	cuenta := 0

	for (tcmax - tcmin) > cc.deltaTemp {
		tcmed = 0.5 * (tcmin + tcmax)
		icmed = cc.Current(ta, tcmed)
		//if err != nil {
		//	return
		//}
		if icmed > ic {
			tcmax = tcmed
		} else {
			tcmin = tcmed
		}
		cuenta += 1
		if cuenta > ITER_MAX {
			err = values.NewValError("CurrentCalc Tc ITERA MAX", 1)
			return
		}
	}
	tc = tcmed
	return

}

//----------------------------------------------------------------------------------------

func (cc *CurrentCalc2) CheckR25() bool {
	return values.Check(cc.R25).Gt(0).Ok()
}

func (cc *CurrentCalc2) CheckAlpha() bool {
	return values.Check(cc.Alpha).Gt(0.0).Lt(1.0).Ok()
}

//----------------------------------------------------------------------------------------

func (cc *CurrentCalc2) Altitude() float64 {
	return cc.altitude
}

func (cc *CurrentCalc2) SetAltitude(h float64) error {
	vc := values.Checker("CurrentCalc SetAltitude")
	vc.Val("h", h).Ge(0)

	cc.altitude = h

	return vc.Error()
}

func (cc *CurrentCalc2) AirVelocity() float64 {
	return cc.airVelocity
}

func (cc *CurrentCalc2) SetAirVelocity(v float64) error {
	vc := values.Checker("CurrentCalc SetAirVelocity")
	vc.Val("v", v).Ge(0)

	cc.airVelocity = v

	return vc.Error()
}

func (cc *CurrentCalc2) SunEffect() float64 {
	return cc.sunEffect
}

func (cc *CurrentCalc2) SetSunEffect(se float64) error {
	vc := values.Checker("CurrentCalc SetSunEffect")
	vc.Val("se", se).Ge(0).Le(1)

	cc.sunEffect = se

	return vc.Error()
}

func (cc *CurrentCalc2) Emissivity() float64 {
	return cc.emissivity
}

func (cc *CurrentCalc2) SetEmissivity(e float64) error {
	vc := values.Checker("CurrentCalc SetEmissivity")
	vc.Val("e", e).Ge(0).Le(1)

	cc.emissivity = e

	return vc.Error()
}

func (cc *CurrentCalc2) Formula() string {
	return cc.formula
}

func (cc *CurrentCalc2) SetFormula(f string) {
	if f != CF_CLASSIC {
		f = CF_IEEE
	}
	cc.formula = f
}

func (cc *CurrentCalc2) DeltaTemp() float64 {
	return cc.deltaTemp
}

func (cc *CurrentCalc2) SetDeltaTemp(t float64) error {
	vc := values.Checker("CurrentCalc SetDeltaTemp")
	vc.Val("t", t).Gt(0)

	cc.deltaTemp = t

	return vc.Error()
}
