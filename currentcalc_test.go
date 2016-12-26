// Copyright Cristian Echeverría Rabí

package conductor

import (
	"fmt"
	"math"
	"testing"
)

func getConductorArgs() ConductorArgs {
	return ConductorArgs{"AAAC 740,8 MCM FLINT", CC_AAAC, 25.17, 0.00, 0.0, 0.0, 0.089360, 1e-10, ""}
}

func getConductor() *Conductor {
	args := getConductorArgs()
	return args.Get()
}

func getConductorFromCategoryArgs(args CategoryArgs) *Conductor {
	return NewConductor("AAAC 740,8 MCM FLINT", args.Get(), 25.17, 0.00, 0.0, 0.0, 0.089360, 1e-10, "")
}

//----------------------------------------------------------------------------------------

func TestConstructorDefaults(t *testing.T) {
	cond := getConductor()
	cc, _ := NewCurrentCalc(cond) // No se verifica error porque los parámetros son correctos

	if cc.Conductor() != cond {
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
	if cc.DeltaTemp() != 0.01 {
		t.Error("!=")
	}
}

func TestConstructorR25(t *testing.T) {
	args := getConductorArgs()

	args.R25 = 0.001
	cc, err := NewCurrentCalc(args.Get())
	if err != nil {
		t.Error(err)
	}
	if cc == nil {
		t.Error("CurrentCalc expected")
	}

	args.R25 = 0.0
	cc, err = NewCurrentCalc(args.Get())
	if err == nil {
		t.Error("R25=0 error expected")
	}
	if cc != nil {
		t.Error("nil expected with R25=0")
	}

	args.R25 = -0.001
	_, err = NewCurrentCalc(args.Get())
	if err == nil {
		t.Error("R25<0 error expected")
	}
	if cc != nil {
		t.Error("nil expected with R25<0")
	}
}

func TestConstructorDiameter(t *testing.T) {
	args := getConductorArgs()

	args.Diameter = 0.001
	cc, err := NewCurrentCalc(args.Get())
	if err != nil {
		t.Error(err)
	}
	if cc == nil {
		t.Error("CurrentCalc expected")
	}

	args.Diameter = 0.0
	cc, err = NewCurrentCalc(args.Get())
	if err == nil {
		t.Error("Diameter=0 error expected")
	}
	if cc != nil {
		t.Error("nil expected with Diameter=0")
	}

	args.Diameter = -0.001
	cc, err = NewCurrentCalc(args.Get())
	if err == nil {
		t.Error("Diameter<0 error expected")
	}
	if cc != nil {
		t.Error("nil expected with Diameter<0")
	}
}

func TestConstructurAlpha(t *testing.T) {
	cat_args := CategoryArgs{"ALUMINUM", 5600.0, 0.0000230, 20.0, 0.00395, "AAC"}

	cat_args.Alpha = 0.001
	cc, err := NewCurrentCalc(getConductorFromCategoryArgs(cat_args))
	if err != nil {
		t.Error(err)
	}
	if cc == nil {
		t.Error("CurrentCalc expected")
	}

	cat_args.Alpha = 0.999
	cc, err = NewCurrentCalc(getConductorFromCategoryArgs(cat_args))
	if err != nil {
		t.Error(err)
	}
	if cc == nil {
		t.Error("CurrentCalc expected")
	}

	cat_args.Alpha = 0.0
	cc, err = NewCurrentCalc(getConductorFromCategoryArgs(cat_args))
	if err == nil {
		t.Error("Alpha=0 error expected")
	}
	if cc != nil {
		t.Error("nil expected with Alpha=0")
	}

	cat_args.Alpha = -0.001
	cc, err = NewCurrentCalc(getConductorFromCategoryArgs(cat_args))
	if err == nil {
		t.Error("Alpha<0 error expected")
	}
	if cc != nil {
		t.Error("nil expected with Alpha<0")
	}

	cat_args.Alpha = 1.0
	cc, err = NewCurrentCalc(getConductorFromCategoryArgs(cat_args))
	if err == nil {
		t.Error("Alpha=1 error expected")
	}
	if cc != nil {
		t.Error("nil expected with Alpha=1")
	}

	cat_args.Alpha = 1.001
	cc, err = NewCurrentCalc(getConductorFromCategoryArgs(cat_args))
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
		t.Errorf("Altitud=0 error not expected: %v", err)
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
		t.Errorf("AirVelocity=0 error not expected: %v", err)
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
		t.Errorf("SunEffect=0 error not expected: %v", err)
	}
	err = cc.SetSunEffect(1.0)
	if err != nil {
		t.Errorf("SunEffect=1 error not expected: %v", err)
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
		t.Errorf("Emissivity=0 error not expected: %v", err)
	}
	err = cc.SetEmissivity(1.0)
	if err != nil {
		t.Errorf("Emissivity=1 error not expected: %v", err)
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
		t.Errorf("DeltaTemp>0 error not expected: %v", err)
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

////----------------------------------------------------------------------------------------

func TestMethodResistance(t *testing.T) {
	cc, _ := NewCurrentCalc(getConductor())

	_, err := cc.Resistance(TC_MIN)
	if err != nil {
		t.Errorf("tc >= TC_MIN error not expected: %v", err)
	}
	_, err = cc.Resistance(TC_MAX)
	if err != nil {
		t.Errorf("tc <= TC_MAX error not expected: %v", err)
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
		t.Errorf("ta >= TA_MIN error not expected: %v", err)
	}
	c, err := cc.Current(TA_MAX, 50)
	if err != nil {
		t.Errorf("ta <= TA_MAX error not expected: %v", err)
	}
	if c < 0 {
		t.Error("Current >=0 expected")
	}
	c, err = cc.Current(25, TC_MIN)
	if err != nil {
		t.Errorf("tc >= TC_MIN error not expected: %v", err)
	}
	if c < 0 {
		t.Error("Current >=0 expected")
	}
	c, err = cc.Current(25, TC_MAX)
	if err != nil {
		t.Errorf("tc <= TC_MAX error not expected: %v", err)
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
		t.Errorf("Current 517.7 expected got: %f", c)
	}
	c, err = cc.Current(30, 60)
	if math.Abs(c-585.4) > 1 {
		t.Errorf("Current 585.4 expected got: %f", c)
	}
}

func TestMethodTc(t *testing.T) {
	cc, _ := NewCurrentCalc(getConductor())

	// Verifica rangos de entrada para ta
	cur, _ := cc.Current(TA_MIN, TC_MAX)
	tc, err := cc.Tc(TA_MIN, cur)
	if err != nil {
		t.Errorf("Error not expected: %v", err)
	}
	tc, err = cc.Tc(TA_MIN-0.0001, cur)
	if err == nil {
		t.Error("ta < TA_MIN error expected")
	}
	cur, _ = cc.Current(TA_MAX, TC_MAX)
	tc, err = cc.Tc(TA_MAX, cur)
	if err != nil {
		t.Errorf("Error not expected: %v", err)
	}
	tc, err = cc.Tc(TA_MAX+0.0001, cur)
	if err == nil {
		t.Error("ta > TA_MAX error expected")
	}

	// Verifica rangos de entrada para ic
	tc, err = cc.Tc(30, -0.0001)
	if err == nil {
		t.Error("ic < 0 error expected")
	}
	cur, _ = cc.Current(30, TC_MAX)
	tc, err = cc.Tc(30, cur)
	if err != nil {
		t.Errorf("Error not expected: %v", err)
	}
	tc, err = cc.Tc(30, cur+0.001)
	if err == nil {
		t.Error("ic > icmax error expected")
	}

	// Verifica que los cálculos de Tc() sean coherentes Current()
	cur, _ = cc.Current(25, 50)
	tc, _ = cc.Tc(25, cur)
	if math.Abs(tc-50) > cc.deltaTemp {
		t.Errorf("Expected difference < %f [%f received]", cc.deltaTemp, math.Abs(tc-50))
	}
	cur, _ = cc.Current(35, 65)
	tc, _ = cc.Tc(35, cur)
	if math.Abs(tc-65) > cc.deltaTemp {
		t.Errorf("Expected difference < %f [%f received]", cc.deltaTemp, math.Abs(tc-50))
	}
}

func TestMethodTa(t *testing.T) {
	cc, _ := NewCurrentCalc(getConductor())

	// Verifica rangos de entrada para tc
	ta, err := cc.Ta(TC_MIN, 0)
	if err != nil {
		t.Errorf("Error not expected: %v", err)
	}
	ta, err = cc.Ta(TC_MIN-0.0001, 0)
	if err == nil {
		t.Error("tc < TC_MIN error expected")
	}
	cur, _ := cc.Current(TA_MIN, TC_MAX)
	ta, err = cc.Ta(TC_MAX, cur)
	if err != nil {
		t.Errorf("Error not expected: %v", err)
	}
	ta, err = cc.Ta(TC_MAX+0.0001, cur)
	if err == nil {
		t.Error("tc > TC_MAX error expected")
	}

	// Verifica rangos de entrada para ic
	cur, _ = cc.Current(TA_MAX, 100) // ic min
	ta, err = cc.Ta(100, cur)
	if err != nil {
		t.Errorf("Error not expected: %v", err)
	}
	ta, err = cc.Ta(100, cur-0.0001)
	if err == nil {
		t.Error("ic < Icmin error expected")
	}
	cur, _ = cc.Current(TA_MIN, 100) // ic max
	ta, err = cc.Ta(100, cur)
	if err != nil {
		t.Errorf("Error not expected: %v", err)
	}
	ta, err = cc.Ta(100, cur+0.0001)
	if err == nil {
		t.Error("ic > Icmax error expected")
	}

	// Verifica que los cálculos de Ta() sean coherentes Current()
	cur, _ = cc.Current(25, 50)
	ta, _ = cc.Ta(50, cur)
	if math.Abs(ta-25) > cc.deltaTemp {
		t.Errorf("Expected difference < %f [%f received]", cc.deltaTemp, math.Abs(ta-50))
	}
	cur, _ = cc.Current(35, 65)
	ta, _ = cc.Ta(65, cur)
	if math.Abs(ta-35) > cc.deltaTemp {
		t.Errorf("Expected difference < %f [%f received]", cc.deltaTemp, math.Abs(ta-50))
	}

}

////----------------------------------------------------------------------------------------

func ExampleResistance() {
	cc, _ := NewCurrentCalc(getConductor())
	r, _ := cc.Resistance(100)
	fmt.Printf("%.4f", r)
	// Output:
	// 0.1121
}

func ExampleCurrent() {
	cc, _ := NewCurrentCalc(getConductor())
	cur, _ := cc.Current(25, 50)
	fmt.Printf("%.2f", cur)
	// Output:
	// 517.68
}

func ExampleTc() {
	cc, _ := NewCurrentCalc(getConductor())
	tc, _ := cc.Tc(25, 100)
	fmt.Printf("%.2f", tc)
	// Output:
	// 33.87
}

func ExampleTa() {
	cc, _ := NewCurrentCalc(getConductor())
	ta, _ := cc.Ta(35, 100)
	fmt.Printf("%.2f", ta)
	// Output:
	// 26.14
}

////----------------------------------------------------------------------------------------

func BenchmarkResistance(b *testing.B) {
	cc, _ := NewCurrentCalc(getConductor())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cc.Resistance(50)
	}
}

func BenchmarkCurrent(b *testing.B) {
	cc, _ := NewCurrentCalc(getConductor())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cc.Current(25, 50)
	}
}

func BenchmarkTc(b *testing.B) {
	cc, _ := NewCurrentCalc(getConductor())
	//cc.SetDeltaTemp(0.0001)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cc.Tc(25, 100)
	}
}

func BenchmarkTa(b *testing.B) {
	cc, _ := NewCurrentCalc(getConductor())
	//cc.SetDeltaTemp(0.0001)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cc.Ta(35, 100)
	}
}
