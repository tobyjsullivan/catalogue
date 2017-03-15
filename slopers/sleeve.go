package slopers

import (
	"github.com/tailored-style/pattern-generator/geometry"
	"github.com/tailored-style/pattern-generator/pieces"
	"math"
	"github.com/tobyjsullivan/catalogue/anchors"
)

type Sleeve struct {
	*TorsoMeasurements
}

func (s *Sleeve) torso() *Torso {
	return &Torso{
		TorsoMeasurements: s.TorsoMeasurements,
	}
}

func (s *Sleeve) P0() *geometry.Point {
	return &geometry.Point{X: 0.0, Y: 0.0}
}

func (s *Sleeve) P1() *geometry.Point {
	return s.P0().SquareDown((s.TorsoMeasurements.ShirtSleeveLength - s.torso().shoulderLength()) + 2.54)
}

func (s *Sleeve) frontArmholeLength() float64 {
	return s.torso().frontArmhole().Length()
}

func (s *Sleeve) backArmholeLength() float64 {
	torso := s.torso()
	return torso.backArmhole().Length() + torso.yokeArmhole().Length()
}

func (s *Sleeve) P2() *geometry.Point {
	armholeLength := s.frontArmholeLength() + s.backArmholeLength()
	return s.P0().SquareDown(armholeLength / 3.0 - 2.86)
}

func (s *Sleeve) P3() *geometry.Point {
	return s.P2().MidpointTo(s.P1()).SquareUp(3.81)
}

func (s *Sleeve) P4() *geometry.Point {
	h := s.frontArmholeLength() - 0.635
	a := s.P0().DistanceTo(s.P2())

	b := math.Sqrt((h * h) - (a * a))

	return s.P2().SquareRight(b)
}

func (s *Sleeve) P5() *geometry.Point {
	h := s.backArmholeLength() - 0.3175
	a := s.P0().DistanceTo(s.P2())

	b := math.Sqrt(h * h - a * a)

	return s.P2().SquareLeft(b)
}

func (s *Sleeve) P6() *geometry.Point {
	return s.P4().SquareToHorizontalLine(s.P1().Y)
}

func (s *Sleeve) P7() *geometry.Point {
	return s.P5().SquareToHorizontalLine(s.P1().Y)
}

func (s *Sleeve) P8() *geometry.Point {
	return s.P4().SquareToHorizontalLine(s.P3().Y)
}

func (s *Sleeve) P9() *geometry.Point {
	return s.P5().SquareToHorizontalLine(s.P3().Y)
}

func (s *Sleeve) P10() *geometry.Point {
	return s.P0().MidpointTo(s.P0().MidpointTo(s.P4()))
}

func (s *Sleeve) P11() *geometry.Point {
	l := s.P0().MidpointTo(s.P4())
	return l.DrawAt(l.AngleRelativeTo(s.P0()), 1.27)
}

func (s *Sleeve) P12() *geometry.Point {
	return s.P0().MidpointTo(s.P4()).MidpointTo(s.P4())
}

func (s *Sleeve) P13() *geometry.Point {
	return s.P10().DrawAt(s.P10().AngleRelativeTo(s.P0()).Perpendicular(), 1.59)
}

func (s *Sleeve) P14() *geometry.Point {
	return s.P12().DrawAt(s.P12().AngleRelativeTo(s.P0()).Perpendicular().Opposite(), 1.27)
}

func (s *Sleeve) P15() *geometry.Point {
	return s.P0().MidpointTo(s.P0().MidpointTo(s.P5()))
}

func (s *Sleeve) P16() *geometry.Point {
	return s.P0().MidpointTo(s.P5())
}

func (s *Sleeve) P17() *geometry.Point {
	return s.P5().MidpointTo(s.P0().MidpointTo(s.P5()))
}

func (s *Sleeve) P18() *geometry.Point {
	return s.P15().DrawAt(s.P0().AngleRelativeTo(s.P15()).Perpendicular(), 1.905)
}

func (s *Sleeve) P19() *geometry.Point {
	return s.P16().DrawAt(s.P0().AngleRelativeTo(s.P16()).Perpendicular(), 0.9525)
}

func (s *Sleeve) P20() *geometry.Point {
	return s.P17().DrawAt(s.P0().AngleRelativeTo(s.P17()), 1.43)
}

func (s *Sleeve) P21() *geometry.Point {
	p := s.P17().MidpointTo(s.P5())
	return p.DrawAt(s.P5().AngleRelativeTo(s.P17()).Perpendicular(), 0.635)
}

func (s *Sleeve) P22() *geometry.Point {
	return s.P6().SquareLeft(5.715)
}

func (s *Sleeve) P23() *geometry.Point {
	return s.P7().SquareRight(5.715)
}

func (s *Sleeve) sleeveCap() *geometry.Polyline {
	poly := &geometry.Polyline{}

	poly.AddLine(
		&geometry.PolyNCurve{
			Points: []*geometry.Point{
				s.P5(),
				s.P21(),
				s.P20(),
				s.P19(),
				s.P18(),
				s.P0(),
			},
			StartAngle: s.P23().AngleRelativeTo(s.P5()).Perpendicular(),
			EndAngle: &geometry.Angle{Rads: 0.0},
		},
		&geometry.PolyNCurve{
			Points: []*geometry.Point{
				s.P0(),
				s.P13(),
				s.P11(),
				s.P14(),
				s.P4(),
			},
			StartAngle: &geometry.Angle{Rads: 0.0},
			EndAngle: s.P22().AngleRelativeTo(s.P4()).Perpendicular(),
		},
	)

	return poly
}

func (s *Sleeve) Stitch() *geometry.Block {
	return &geometry.Block{}
}

func (p *Sleeve) OuterCut() *geometry.Polyline {
	return &geometry.Polyline{}
}

func (s *Sleeve) InnerCut() *geometry.Block {
	return &geometry.Block{}
}

func (s *Sleeve) Ink() *geometry.Block {
	return &geometry.Block{}
}

func (s *Sleeve) Reference() *geometry.Block {
	layer := &geometry.Block{}

	layer.AddLine(
		s.sleeveCap(),
		&geometry.StraightLine{
			Start: s.P4(),
			End: s.P22(),
		},
		&geometry.StraightLine{
			Start: s.P5(),
			End: s.P23(),
		},
		&geometry.StraightLine{
			Start: s.P22(),
			End: s.P23(),
		},

	)

	a := make(map[string]*geometry.Point)
	a["0"] = s.P0()
	a["1"] = s.P1()
	a["2"] = s.P2()
	a["3"] = s.P3()
	a["4"] = s.P4()
	a["5"] = s.P5()
	a["6"] = s.P6()
	a["7"] = s.P7()
	a["8"] = s.P8()
	a["9"] = s.P9()
	a["10"] = s.P10()
	a["11"] = s.P11()
	a["12"] = s.P12()
	a["13"] = s.P13()
	a["14"] = s.P14()
	a["15"] = s.P15()
	a["16"] = s.P16()
	a["17"] = s.P17()
	a["18"] = s.P18()
	a["19"] = s.P19()
	a["20"] = s.P20()
	a["21"] = s.P21()
	a["22"] = s.P22()
	a["23"] = s.P23()
	anchors.AddAnchors(layer, a)

	return layer
}

func (s *Sleeve) CutCount() int {
	return 1
}

func (s *Sleeve) OnFold() bool {
	return false
}

func (s *Sleeve) Mirrored() bool {
	return true
}

func (s *Sleeve) Details() *pieces.Details {
	return &pieces.Details{}
}
