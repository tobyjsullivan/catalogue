package main

import (
	"fmt"

	"github.com/tailored-style/pattern-generator/rendering"
	"github.com/tobyjsullivan/catalogue/slopers"
	"github.com/tobyjsullivan/catalogue/styles"
)

type bodyLabsMeasurments struct {
	height                         float64
	neck                           float64
	chest                          float64
	bellyButtonWaist               float64
	naturalWaist                   float64
	glutealHip                     float64
	acrossBackShoulderNeck         float64
	alongFrontNeckBaseToGlutealHip float64
	midUpperArm                    float64
	shoulderToMidhand			   float64
	neckToGlutealHip               float64
	underBust					   float64
	maximumHip                     float64
	neckShoulderElbowWrist         float64
	wrist                          float64
}

func (m *bodyLabsMeasurments) toTorsoSloperMeasurements() *slopers.TorsoMeasurements {
	assertAmount(m.height)
	assertAmount(m.neck)
	assertAmount(m.acrossBackShoulderNeck)
	assertAmount(m.chest)
	assertAmount(m.neckToGlutealHip)
	assertAmount(m.bellyButtonWaist)
	assertAmount(m.naturalWaist)
	assertAmount(m.maximumHip)
	assertAmount(m.neckShoulderElbowWrist)
	assertAmount(m.midUpperArm)
	assertAmount(m.wrist)

	return &slopers.TorsoMeasurements{
		Height:                        m.height,
		NeckCircumference:             m.neck,
		ShoulderToShoulder:            m.acrossBackShoulderNeck,
		ChestCircumference:            m.chest,
		ShirtLength:                   m.neckToGlutealHip,
		BellyButtonWaistCircumference: m.bellyButtonWaist,
		NaturalWaistCircumference:     m.naturalWaist,
		HipCircumference:              m.maximumHip,
		ShirtSleeveLength:             m.neckShoulderElbowWrist,
		BicepCircumference:            m.midUpperArm,
		WristCircumference:            m.wrist,
	}
}

func assertAmount(v float64) {
	if v <= 0.0 {
		panic("Amount must be greater than zero")
	}
}

func main() {
	// 42" CHEST STANDARD
	//measurements := &bodyLabsMeasurments{
	//	height: 182.9, // 72"
	//	neck: 41.9, // 16 1/2"
	//	chest: 106.7, // 42"
	//	bellyButtonWaist: 91.4, // 36"
	//	naturalWaist: 91.4, // 36"
	//	acrossBackShoulderNeck: 46.4, // 18 1/4
	//	midUpperArm: 36.2, // 14 1/4
	//	neckToGlutealHip: 64.1, // 25 1/4
	//	maximumHip: 109.2, // 43"
	//	neckShoulderElbowWrist: 87.0, // 25 1/8" + (18 1/4" / 2) = 34 1/4
	//	wrist: 19.1, // 7 1/2"
	//}

	// PERSONAL MEASUREMENTS
	//m["height"] = 182.0
	//m["neck"] = 43.0
	//m["chest"] = 110.0
	//m["waist"] = 96.5
	//m["hip"] = 110.5
	//m["sleeve"] = 92.0
	//m["wrist"] = 17.0

	// PERSONAL - BODYxLABS ESTIMATE
	measurements := &bodyLabsMeasurments{
		height: 183.0,
		neck: 38.8,
		chest: 110.25,
		bellyButtonWaist: 104.8,
		naturalWaist: 98.89,
		glutealHip: 109.42,
		acrossBackShoulderNeck: 43.32,
		alongFrontNeckBaseToGlutealHip: 57.02,
		midUpperArm: 33.89,
		neckToGlutealHip: 64.28,
		shoulderToMidhand: 63.76,
		underBust: 101.3,
		maximumHip: 111.04,
		neckShoulderElbowWrist: 84.23,
		wrist: 18.33,
	}

	torsoMeasure := measurements.toTorsoSloperMeasurements()
	style := styles.NewSN11001Shirt(torsoMeasure)

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
	pieceRender := &rendering.PieceRender{
		Piece: &slopers.Torso{
			TorsoMeasurements: torsoMeasure,
		},
	}
	err = pieceRender.SavePDF("/Users/toby/sandbox/torso-out.pdf")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Generating PDF of Sleeve sloper")
	pieceRender = &rendering.PieceRender{
		Piece: &slopers.Sleeve{
			TorsoMeasurements: torsoMeasure,
		},
	}
	err = pieceRender.SavePDF("/Users/toby/sandbox/sleeve-out.pdf")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Done.")
}
