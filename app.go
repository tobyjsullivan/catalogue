package main

import (
	"fmt"
	"github.com/tailored-style/pattern-generator/patternfile"
	"github.com/tobyjsullivan/catalogue/styles"
	"github.com/tailored-style/pattern-generator/pieces"
	"github.com/tailored-style/pattern-generator/marker"
)

func main() {
	// 42" CHEST STANDARD
	//measurements := &pieces.Measurements{
	//	ChestCircumference: 106.7, // 42"
	//	WaistCircumference: 91.4, // 36"
	//	HipCircumference: 109.2, // 43"
	//	NeckCircumference: 41.9, // 16 1/2"
	//	Height: 182.9, // 72"
	//  SleeveLength: 87.0, // 25 1/8" + (18 1/4" / 2) = 34 1/4
	//}

	// PERSONAL MEASUREMENTS
	//measurements := &pieces.Measurements{
	//	ChestCircumference: 110.0,
	//	WaistCircumference: 96.5,
	//	HipCircumference: 110.5,
	//	NeckCircumference: 43.0,
	//	Height: 182.0,
	//  SleeveLength: 92.0,
	//}

	// PERSONAL - BODYxLABS ESTIMATE
	measurements := &pieces.Measurements{
		ChestCircumference: 110.25, // chest
		WaistCircumference: 104.8, // belly_button_waist
		HipCircumference: 111.04, // maximum_hip
		NeckCircumference: 38.8, // neck (input)
		Height: 183.0, // height (input)
		SleeveLength: 84.23, // neck_shoulder_elbow_wrist
	}

	style := &styles.SN11001Shirt{
		Measurements: measurements,
	}

	fmt.Println("Generating DXF...")
	pf := patternfile.NewPatternFile()
	err := pf.DrawPattern(style)
	if err != nil {
		panic(err.Error())
	}

	err = pf.SaveAs("/Users/toby/sandbox/v3-out.dxf")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Generating PDF Marker...")
	marker := &marker.Marker{
		Style: style,
	}
	err = marker.SavePDF("/Users/toby/sandbox/v3-marker-test.pdf")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Done.")
}
