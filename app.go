package main

import (
	"fmt"
	"github.com/tailored-style/pattern-generator/patternfile"
	"github.com/tobyjsullivan/catalogue/styles"
	"github.com/tailored-style/pattern-generator/marker"
)

func main() {
	m := make(map[string]float64)

	// 42" CHEST STANDARD
	//m["height"] = 182.9 // 72"
	//m["neck"] = 41.9 // 16 1/2"
	//m["chest"] = 106.7 // 42"
	//m["waist"] = 91.4 // 36"
	//m["hip"] = 109.2 // 43"
	//m["sleeve"] = 87.0 // 25 1/8" + (18 1/4" / 2) = 34 1/4
	//m["wrist"] = 19.1 // 7 1/2"

	// PERSONAL MEASUREMENTS
	//m["height"] = 182.0
	//m["neck"] = 43.0
	//m["chest"] = 110.0
	//m["waist"] = 96.5
	//m["hip"] = 110.5
	//m["sleeve"] = 92.0
	//m["wrist"] = 17.0

	// PERSONAL - BODYxLABS ESTIMATE
	m["height"] = 183.0 // height (input)
	m["neck"] = 38.8 // neck (input)
	m["chest"] = 110.25 // chest
	m["waist"] = 104.8 // belly_button_waist
	m["hip"] = 111.04 // maximum_hip
	m["sleeve"] = 84.23 // neck_shoulder_elbow_wrist
	m["wrist"] = 18.33 // wrist

	style := styles.NewSN11001Shirt(
		m["height"],
		m["neck"],
		m["chest"],
		m["waist"],
		m["hip"],
		m["sleeve"],
		m["wrist"],
	)

	fmt.Println("Generating DXF...")
	pf := &patternfile.PatternFile{
		Style: style,
	}
	err := pf.SaveDXF("/Users/toby/sandbox/v3-out.dxf")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Generating PDF...")
	pf = &patternfile.PatternFile{
		Style: style,
	}
	err = pf.SavePDF("/Users/toby/sandbox/v3-out.pdf")
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
