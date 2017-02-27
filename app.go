package main

import (
	"fmt"

	"github.com/tailored-style/pattern-generator/rendering"
	"github.com/tobyjsullivan/catalogue/slopers"
	"github.com/tobyjsullivan/catalogue/styles"
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
	m["height"] = 183.0             // height (input)
	m["neck"] = 38.8                // neck (input)
	m["chest"] = 110.25             // chest
	m["belly_button_waist"] = 104.8 // belly_button_waist
	m["natural_waist"] = 98.89
	m["gluteal_hip"] = 109.42
	m["across_back_shoulder_neck"] = 43.32
	m["along_front_neck_base_to_gluteal_hip"] = 57.02
	m["mid_upper_arm"] = 33.89
	m["neck_to_gluteal_hip"] = 64.28
	m["shoulder_to_midhand"] = 63.76
	m["under_bust"] = 101.3
	m["maximum_hip"] = 111.04              // maximum_hip
	m["neck_shoulder_elbow_wrist"] = 84.23 // neck_shoulder_elbow_wrist
	m["wrist"] = 18.33                     // wrist

	style := styles.NewSN11001Shirt(
		m["height"],
		m["neck"],
		m["chest"],
		m["belly_button_waist"],
		m["maximum_hip"],
		m["neck_shoulder_elbow_wrist"],
		m["wrist"],
	)

	fmt.Println("Generating DXF...")
	pf := &rendering.Pattern{
		Style: style,
	}
	err := pf.SaveDXF("/Users/toby/sandbox/v3-out.dxf")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Generating PDF...")
	pf = &rendering.Pattern{
		Style: style,
	}
	err = pf.SavePDF("/Users/toby/sandbox/v3-out.pdf")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Generating PDF Marker...")
	marker := &rendering.Marker{
		Style: style,
	}
	err = marker.SavePDF("/Users/toby/sandbox/v3-marker-test.pdf")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Generating PDF of torso sloper")
	pieceRender := &rendering.Piece{
		Piece: &slopers.Torso{
			Height:                        m["height"],
			NeckCircumference:             m["neck"],
			ShoulderToShoulder:            m["across_back_shoulder_neck"],
			ChestCircumference:            m["chest"],
			ShirtLength:                   m["neck_to_gluteal_hip"],
			BellyButtonWaistCircumference: m["belly_button_waist"],
			NaturalWaistCircumference:     m["natural_waist"],
			HipCircumference:              m["maximum_hip"],
			ShirtSleeveLength:             m["shoulder_to_midhand"],
			BicepCircumference:            m["mid_upper_arm"],
			WristCircumference:            m["wrist"],
		},
	}
	err = pieceRender.SavePDF("/Users/toby/sandbox/torso-out.pdf")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Done.")
}
