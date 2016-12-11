// Copyright Cristian Echeverría Rabí

package conductor

import (
	"math"

	"bitbucket.org/tormundo/go.utils/values"
)

//----------------------------------------------------------------------------------------

func NewCurrentCalc(cond Conductor) (cc CurrentCalc, err error) {
	vc := values.Checker("NewCurrentCalc cond")
	vc.Val("R25", cond.R25).Gt(0.0)
	vc.Val("Diameter", cond.Diameter).Gt(0.0)
	vc.Val("Alpha", cond.Alpha).Gt(0.0).Lt(1.0)

	err = vc.Error()
	cc = CurrentCalc{cond, 300.0, 2.0, 1.0, 0.5, CF_IEEE, 0.0001}

	return
}

//----------------------------------------------------------------------------------------

type CurrentCalc struct {
	Conductor           // Conductor instance
	altitude    float64 // Altitude [m] = 300.0
	airVelocity float64 // Velocity of air stream [ft/seg] =   2.0
	sunEffect   float64 // Sun effect factor (0 to 1) = 1.0
	emissivity  float64 // Emissivity (0 to 1) = 0.5
	formula     string  // Define formula for current calculation = CF_IEEE
	deltaTemp   float64 // Temperature difference to determine equality [°C] = 0.0001
}

func (cc *CurrentCalc) Resistance(tc float64) (r float64, err error) {
	vc := values.Checker("CurrentCalc Resistance")
	vc.Val("tc", tc).Ge(TC_MIN).Le(TC_MAX)

	r = math.NaN()
	err = vc.Error()

	if err != nil {
		return
	}

	r = cc.R25 * (1 + cc.Alpha*(tc-25.0))

	return
}

func (cc *CurrentCalc) Current(ta float64, tc float64) (q float64, err error) {
	vc := values.Checker("CurrentCalc Current")
	vc.Val("ta", ta).Ge(TA_MIN).Le(TA_MAX)
	vc.Val("tc", tc).Ge(TC_MIN).Le(TC_MAX)

	q = math.NaN()
	err = vc.Error()

	if err != nil {
		return
	}

	if ta >= tc {
		q = 0.0
		return
	}

	D := cc.Diameter / 25.4                                                 // Diámetro en pulgadas
	Pb := math.Pow(10, (1.880813592 - cc.Altitude()/18336.0))               // Presión barométrica en cmHg
	V := cc.AirVelocity() * 3600                                            // Vel. viento en pies/hora
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
		if cc.Formula() == CF_IEEE {
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
	Qr := 0.138 * D * cc.Emissivity() * (LK - MK)
	Qs := 3.87 * D * cc.SunEffect()

	if (Qc + Qr) < Qs {
		q = 0.0
	} else {
		q = math.Sqrt((Qc + Qr - Qs) / Rc)
	}
	return
}

func (cc *CurrentCalc) Altitude() float64 {
	return cc.altitude
}

func (cc *CurrentCalc) SetAltitude(h float64) error {
	vc := values.Checker("CurrentCalc SetAltitude")
	vc.Val("h", h).Ge(0)

	cc.altitude = h

	return vc.Error()
}

func (cc *CurrentCalc) AirVelocity() float64 {
	return cc.airVelocity
}

func (cc *CurrentCalc) SetAirVelocity(v float64) error {
	vc := values.Checker("CurrentCalc SetAirVelocity")
	vc.Val("v", v).Ge(0)

	cc.airVelocity = v

	return vc.Error()
}

func (cc *CurrentCalc) SunEffect() float64 {
	return cc.sunEffect
}

func (cc *CurrentCalc) SetSunEffect(se float64) error {
	vc := values.Checker("CurrentCalc SetSunEffect")
	vc.Val("se", se).Ge(0).Le(1)

	cc.sunEffect = se

	return vc.Error()
}

func (cc *CurrentCalc) Emissivity() float64 {
	return cc.emissivity
}

func (cc *CurrentCalc) SetEmissivity(e float64) error {
	vc := values.Checker("CurrentCalc SetEmissivity")
	vc.Val("e", e).Ge(0).Le(1)

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
	vc := values.Checker("CurrentCalc SetDeltaTemp")
	vc.Val("t", t).Gt(0)

	cc.deltaTemp = t

	return vc.Error()
}
