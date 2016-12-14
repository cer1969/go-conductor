// Copyright Cristian Echeverría Rabí

package conductor

import (
	"fmt"
	"math"
	"testing"
)

func getConductor2() Conductor {
	return Conductor{CC_AAAC, "AAAC 740,8 MCM FLINT", 25.17, 0.00, 0.0, 0.0, 0.089360, 0, ""}
}

//----------------------------------------------------------------------------------------

func TestNewCurrentCalcDefaults2(t *testing.T) {
	cond := getConductor2()
	cc := NewCurrentCalc2(cond)

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

/*
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
*/
//----------------------------------------------------------------------------------------

func TestMethodResistance2(t *testing.T) {
	cc := NewCurrentCalc2(getConductor2())
	r25 := cc.R25
	alp := cc.Alpha

	cc.R25 = -0.001
	r := cc.Resistance(30.0)
	if !math.IsNaN(r) {
		t.Error("NaN expected, got ", r)
	}

	cc.R25 = r25
	cc.Alpha = 0
	r = cc.Resistance(30.0)
	if !math.IsNaN(r) {
		t.Error("NaN expected, got ", r)
	}
	cc.Alpha = 1
	r = cc.Resistance(30.0)
	if !math.IsNaN(r) {
		t.Error("NaN expected, got ", r)
	}

	cc.Alpha = alp
	r = cc.Resistance(-300.0)
	if r != 0 {
		t.Error("0.0 expected, got ", r)
	}

	r = cc.Resistance(25)
	if r != cc.R25 {
		t.Errorf("%v expected, got %v", cc.R25, r)
	}
}

func ExampleResistance2() {
	cc := NewCurrentCalc2(getConductor2())
	r1 := cc.Resistance(100)
	r2 := cc.Resistance(50)
	fmt.Printf("%.4f\n%.4f", r1, r2)
	// Output:
	// 0.1121
	// 0.0970
}

func ExampleCheckR25() {
	cc := NewCurrentCalc2(getConductor2())
	fmt.Printf("%v", cc.CheckR25())
	// Output:
	// true
}

func ExampleCheckAlpha() {
	cc := NewCurrentCalc2(getConductor2())
	cc.Alpha = 0.0
	fmt.Printf("%v", cc.CheckAlpha())
	// Output:
	// false
}

//----------------------------------------------------------------------------------------

/*
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
*/

//----------------------------------------------------------------------------------------

func BenchmarkCurrentCalc2(b *testing.B) {
	cc := NewCurrentCalc2(getConductor2())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cc.Current(25, 50)
	}
}
