// Copyright Cristian Echeverría Rabí

package conductor

import (
	"fmt"
	"math"
	"testing"
)

func getConductor() Conductor {
	return Conductor{CC_AAAC, "AAAC 740,8 MCM FLINT", 25.17, 0.00, 0.0, 0.0, 0.089360, 0, ""}
}

//----------------------------------------------------------------------------------------

func TestConstructorDefaults(t *testing.T) {
	cond := getConductor()
	cc, _ := NewCurrentCalc(cond) // No se verifica error porque los parámetros son correctos

	if cc.Conductor != cond {
		t.Error("!=")
	}
	if cc.Altitude() != 300 {
		t.Error("!=")
	}
	if cc.AirVelocity() != 2 {
		t.Error("!=")
	}
	if cc.SunEffect() != 1 {
		t.Error("!=")
	}
	if cc.Emissivity() != 0.5 {
		t.Error("!=")
	}
	if cc.Formula() != CF_IEEE {
		t.Error("!=")
	}
	if cc.DeltaTemp() != 0.0001 {
		t.Error("!=")
	}
}

func TestConstructorR25(t *testing.T) {
	cond := getConductor()

	cond.R25 = 0.001
	cc, err := NewCurrentCalc(cond)
	if err != nil {
		t.Error(err)
	}
	if cc == nil {
		t.Error("CurrentCalc expected")
	}

	cond.R25 = 0.0
	cc, err = NewCurrentCalc(cond)
	if err == nil {
		t.Error("R25=0 error expected")
	}
	if cc != nil {
		t.Error("nil expected with R25=0")
	}

	cond.R25 = -0.001
	_, err = NewCurrentCalc(cond)
	if err == nil {
		t.Error("R25<0 error expected")
	}
	if cc != nil {
		t.Error("nil expected with R25<0")
	}
}

func TestConstructorDiameter(t *testing.T) {
	cond := getConductor()

	cond.Diameter = 0.001
	cc, err := NewCurrentCalc(cond)
	if err != nil {
		t.Error(err)
	}
	if cc == nil {
		t.Error("CurrentCalc expected")
	}

	cond.Diameter = 0.0
	cc, err = NewCurrentCalc(cond)
	if err == nil {
		t.Error("Diameter=0 error expected")
	}
	if cc != nil {
		t.Error("nil expected with Diameter=0")
	}

	cond.Diameter = -0.001
	cc, err = NewCurrentCalc(cond)
	if err == nil {
		t.Error("Diameter<0 error expected")
	}
	if cc != nil {
		t.Error("nil expected with Diameter<0")
	}
}

func TestConstructurAlpha(t *testing.T) {
	cond := getConductor()

	cond.Category.Alpha = 0.001
	cc, err := NewCurrentCalc(cond)
	if err != nil {
		t.Error(err)
	}
	if cc == nil {
		t.Error("CurrentCalc expected")
	}

	cond.Category.Alpha = 0.999
	cc, err = NewCurrentCalc(cond)
	if err != nil {
		t.Error(err)
	}
	if cc == nil {
		t.Error("CurrentCalc expected")
	}

	cond.Category.Alpha = 0.0
	cc, err = NewCurrentCalc(cond)
	if err == nil {
		t.Error("Alpha=0 error expected")
	}
	if cc != nil {
		t.Error("nil expected with Alpha=0")
	}

	cond.Category.Alpha = -0.001
	cc, err = NewCurrentCalc(cond)
	if err == nil {
		t.Error("Alpha<0 error expected")
	}
	if cc != nil {
		t.Error("nil expected with Alpha<0")
	}

	cond.Category.Alpha = 1.0
	cc, err = NewCurrentCalc(cond)
	if err == nil {
		t.Error("Alpha=1 error expected")
	}
	if cc != nil {
		t.Error("nil expected with Alpha=1")
	}

	cond.Category.Alpha = 1.001
	cc, err = NewCurrentCalc(cond)
	if err == nil {
		t.Error("Alpha>1 error expected")
	}
	if cc != nil {
		t.Error("nil expected with Alpha>1")
	}
}

//----------------------------------------------------------------------------------------

func TestPropAltitude(t *testing.T) {
	cc, _ := NewCurrentCalc(getConductor())

	err := cc.SetAltitude(150.0)
	if err != nil {
		t.Error(err)
	}
	if cc.altitude != cc.Altitude() {
		t.Error("Error en SetAltitude")
	}
	if cc.Altitude() != 150.0 {
		t.Error("Error en SetAltitude")
	}
	err = cc.SetAltitude(0.0)
	if err != nil {
		t.Error("Altitud=0 error not expected")
	}
	err = cc.SetAltitude(-0.001)
	if err == nil {
		t.Error("Altitud<0 error expected")
	}
}

func TestPropAirVelocity(t *testing.T) {
	cc, _ := NewCurrentCalc(getConductor())

	err := cc.SetAirVelocity(2.0)
	if err != nil {
		t.Error(err)
	}
	if cc.airVelocity != cc.AirVelocity() {
		t.Error("Error en SetAirVelocity")
	}
	if cc.AirVelocity() != 2.0 {
		t.Error("Error en SetAirVelocity")
	}
	err = cc.SetAirVelocity(0.0)
	if err != nil {
		t.Error("AirVelocity=0 error not expected")
	}
	err = cc.SetAirVelocity(-0.001)
	if err == nil {
		t.Error("AirVelocity<0 error expected")
	}
}

func TestPropSunEffect(t *testing.T) {
	cc, _ := NewCurrentCalc(getConductor())

	err := cc.SetSunEffect(0.5)
	if err != nil {
		t.Error(err)
	}
	if cc.sunEffect != cc.SunEffect() {
		t.Error("Error en SetSunEffect")
	}
	if cc.SunEffect() != 0.5 {
		t.Error("Error en SetSunEffect")
	}
	err = cc.SetSunEffect(0.0)
	if err != nil {
		t.Error("SunEffect=0 error not expected")
	}
	err = cc.SetSunEffect(1.0)
	if err != nil {
		t.Error("SunEffect=1 error not expected")
	}
	err = cc.SetSunEffect(-0.001)
	if err == nil {
		t.Error("SunEffect<0 error expected")
	}
	err = cc.SetSunEffect(1.001)
	if err == nil {
		t.Error("SunEffect>1 error expected")
	}
}

func TestPropEmissivity(t *testing.T) {
	cc, _ := NewCurrentCalc(getConductor())

	err := cc.SetEmissivity(0.7)
	if err != nil {
		t.Error(err)
	}
	if cc.emissivity != cc.Emissivity() {
		t.Error("Error en SetEmissivity")
	}
	if cc.Emissivity() != 0.7 {
		t.Error("Error en SetEmissivity")
	}
	err = cc.SetEmissivity(0.0)
	if err != nil {
		t.Error("Emissivity=0 error not expected")
	}
	err = cc.SetEmissivity(1.0)
	if err != nil {
		t.Error("Emissivity=1 error not expected")
	}
	err = cc.SetEmissivity(-0.001)
	if err == nil {
		t.Error("Emissivity<0 error expected")
	}
	err = cc.SetEmissivity(1.001)
	if err == nil {
		t.Error("Emissivity>1 error expected")
	}
}

func TestPropFormula(t *testing.T) {
	cc, _ := NewCurrentCalc(getConductor())

	cc.SetFormula(CF_IEEE)
	if cc.formula != cc.Formula() {
		t.Error("Error en SetFormula")
	}
	if cc.Formula() != CF_IEEE {
		t.Error("Error en SetFormula")
	}
	cc.SetFormula("")
	if cc.formula != cc.Formula() {
		t.Error("Error en SetFormula")
	}
	if cc.Formula() != CF_IEEE {
		t.Error("Error en SetFormula")
	}
	cc.SetFormula(CF_CLASSIC)
	if cc.formula != cc.Formula() {
		t.Error("Error en SetFormula")
	}
	if cc.Formula() != CF_CLASSIC {
		t.Error("Error en SetFormula")
	}
}

func TestPropDeltaTemp(t *testing.T) {
	cc, _ := NewCurrentCalc(getConductor())

	err := cc.SetDeltaTemp(0.001)
	if err != nil {
		t.Error(err)
	}
	if cc.deltaTemp != cc.DeltaTemp() {
		t.Error("Error en SetDeltaTemp")
	}
	if cc.DeltaTemp() != 0.001 {
		t.Error("Error en SetDeltaTemp")
	}
	err = cc.SetDeltaTemp(0.0001)
	if err != nil {
		t.Error("DeltaTemp>0 error not expected")
	}
	err = cc.SetDeltaTemp(-0.0001)
	if err == nil {
		t.Error("DeltaTemp<0 error expected")
	}
	err = cc.SetDeltaTemp(0.0)
	if err == nil {
		t.Error("DeltaTemp=0 error expected")
	}
}

//----------------------------------------------------------------------------------------

func TestMethodResistance(t *testing.T) {
	cc, _ := NewCurrentCalc(getConductor())

	_, err := cc.Resistance(TC_MIN)
	if err != nil {
		t.Error("tc >= TC_MIN error not expected")
	}
	_, err = cc.Resistance(TC_MAX)
	if err != nil {
		t.Error("tc <= TC_MAX error not expected")
	}
	_, err = cc.Resistance(TC_MIN - 0.001)
	if err == nil {
		t.Error("tc < TC_MIN error expected")
	}
	_, err = cc.Resistance(TC_MAX + 0.001)
	if err == nil {
		t.Error("tc > TC_MAX error expected")
	}
}

func TestMethodCurrent(t *testing.T) {
	cc, _ := NewCurrentCalc(getConductor())

	_, err := cc.Current(TA_MIN, 50)
	if err != nil {
		t.Error("ta >= TA_MIN error not expected")
	}
	c, err := cc.Current(TA_MAX, 50)
	if err != nil {
		t.Error("ta <= TA_MAX error not expected")
	}
	if c < 0 {
		t.Error("Current >=0 expected")
	}
	c, err = cc.Current(25, TC_MIN)
	if err != nil {
		t.Error("tc >= TC_MIN error not expected")
	}
	if c < 0 {
		t.Error("Current >=0 expected")
	}
	c, err = cc.Current(25, TC_MAX)
	if err != nil {
		t.Error("tc <= TC_MAX error not expected ")
	}

	_, err = cc.Current(TA_MIN-0.001, 50)
	if err == nil {
		t.Error("ta < TA_MIN error expected")
	}
	_, err = cc.Current(TA_MAX+0.001, 50)
	if err == nil {
		t.Error("ta > TA_MAX error expected")
	}
	_, err = cc.Current(25, TC_MIN-0.001)
	if err == nil {
		t.Error("tc < TC_MIN error expected")
	}
	_, err = cc.Current(25, TC_MAX+0.001)
	if err == nil {
		t.Error("tc > TC_MAX error expected")
	}

	c, err = cc.Current(25, 25)
	if c != 0.0 {
		t.Error("Current=0 expected")
	}
	c, err = cc.Current(26, 25)
	if c != 0.0 {
		t.Error("Current=0 expected")
	}

	cc.SetFormula(CF_CLASSIC)
	c, err = cc.Current(25, 50)
	if math.Abs(c-517.7) > 1 {
		t.Error("Current 517.7 expected got ", c)
	}
	c, err = cc.Current(30, 60)
	if math.Abs(c-585.4) > 1 {
		t.Error("Current 585.4 expected got ", c)
	}

}

//----------------------------------------------------------------------------------------

func ExampleResistance() {
	cc, _ := NewCurrentCalc(getConductor())
	r, _ := cc.Resistance(100)
	fmt.Printf("%.4f", r)
	// Output:
	// 0.1121
}

//----------------------------------------------------------------------------------------

func BenchmarkCurrentCalc(b *testing.B) {
	cc, _ := NewCurrentCalc(getConductor())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cc.Current(25, 50)
	}
}
