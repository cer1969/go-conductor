// Copyright Cristian Echeverría Rabí

package conductor

import (
	"testing"
)

func TestDefaults(t *testing.T) {
    cu300 := Conductor{CC_CU, "CU 300 MCM", 15.95, 152.00, 1.378, 6123.0, 0.12270, 0, ""}
    cc, _ := NewCurrentCalc(cu300)
    if cc.Conductor != cu300 {t.Error("!=")}
    if cc.Altitude() != 300 {t.Error("!=")}
    if cc.AirVelocity() != 2 {t.Error("!=")}
    if cc.SunEffect() != 1 {t.Error("!=")}
    if cc.Emissivity() != 0.5 {t.Error("!=")}
    if cc.Formula() != CF_IEEE {t.Error("!=")}
    if cc.DeltaTemp() != 0.0001 {t.Error("!=")}
}

func TestErrorR25(t *testing.T) {
    cu300 := Conductor{CC_CU, "CU 300 MCM", 15.95, 152.00, 1.378, 6123.0, 0.12270, 0, ""}
    
    cu300.R25 = 0.001
    _, err := NewCurrentCalc(cu300)
    if err != nil {t.Error(err)}
    
    cu300.R25 = 0.0
    _, err = NewCurrentCalc(cu300)
    if err == nil {t.Error("R25=0 error expected")}
    
    cu300.R25 = -0.001
    _, err = NewCurrentCalc(cu300)
    if err == nil {t.Error("R25<0 error expected")}
}

func TestErrorDiameter(t *testing.T) {
    cu300 := Conductor{CC_CU, "CU 300 MCM", 15.95, 152.00, 1.378, 6123.0, 0.12270, 0, ""}
    
    cu300.Diameter = 0.001
    _, err := NewCurrentCalc(cu300)
    if err != nil {t.Error(err)}
    
    cu300.Diameter = 0.0
    _, err = NewCurrentCalc(cu300)
    if err == nil {t.Error("Diameter=0 error expected")}
    
    cu300.Diameter = -0.001
    _, err = NewCurrentCalc(cu300)
    if err == nil {t.Error("Diameter<0 error expected")}
}

func TestErrorAlpha(t *testing.T) {
    cu300 := Conductor{CC_CU, "CU 300 MCM", 15.95, 152.00, 1.378, 6123.0, 0.12270, 0, ""}
    
    cu300.Category.Alpha = 0.001
    _, err := NewCurrentCalc(cu300)
    if err != nil {t.Error(err)}
    
    cu300.Category.Alpha = 0.999
    _, err = NewCurrentCalc(cu300)
    if err != nil {t.Error(err)}
    
    cu300.Category.Alpha = 0.0
    _, err = NewCurrentCalc(cu300)
    if err == nil {t.Error("Alpha=0 error expected")}
    
    cu300.Category.Alpha = -0.001
    _, err = NewCurrentCalc(cu300)
    if err == nil {t.Error("Alpha<0 error expected")}
    
    cu300.Category.Alpha = 1.001
    _, err = NewCurrentCalc(cu300)
    if err == nil {t.Error("Alpha>1 error expected")}
}


func BenchmarkCurrentCalc(b *testing.B) {
	cu300 := Conductor{CC_CU, "CU 300 MCM", 15.95, 152.00, 1.378, 6123.0, 0.12270, 0, ""}
    cc, _ := NewCurrentCalc(cu300)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cc.Current(25, 50)
		//current, _ := cc.Current(25, 50)
		//fmt.Println(current)
	}
}