package slopers

import (
	"math"

	"github.com/tailored-style/pattern-generator/geometry"
	"github.com/tailored-style/pattern-generator/pieces"
	"github.com/tobyjsullivan/catalogue/anchors"
)

type Torso struct {
	*TorsoMeasurements
}

type TorsoMeasurements struct {
	Height float64
	NeckCircumference float64
	ShoulderToShoulder float64
	ChestCircumference float64
	ShirtLength float64
	BellyButtonWaistCircumference float64
	NaturalWaistCircumference float64
	HipCircumference float64
	ShirtSleeveLength float64
	BicepCircumference float64
	WristCircumference float64
}

func (s *Torso) P0() *geometry.Point {
	return &geometry.Point{X: 0.0, Y: 0.0}
}

func (s *Torso) P1() *geometry.Point {
	return s.P0().SquareDown(s.TorsoMeasurements.ShirtLength)
}

func (s *Torso) P2() *geometry.Point {
	return s.P1().SquareRight(s.TorsoMeasurements.ChestCircumference / 2.0)
}

func (s *Torso) P3() *geometry.Point {
	return s.P2().SquareToHorizontalLine(s.P0().Y)
}

func (s *Torso) P4() *geometry.Point {
	return s.P0().SquareDown(s.TorsoMeasurements.ChestCircumference / 4.0)
}

func (s *Torso) P5() *geometry.Point {
	return s.P1().SquareUp(s.TorsoMeasurements.ShirtLength / 3.0)
}

func (s *Torso) P6() *geometry.Point {
	return s.P4().SquareToVerticalLine(s.P3().X)
}

func (s *Torso) P7() *geometry.Point {
	return s.P5().SquareToVerticalLine(s.P3().X)
}

func (s *Torso) P8() *geometry.Point {
	return s.P4().SquareRight((s.TorsoMeasurements.ChestCircumference / 6.0) + 2.0)
}

func (s *Torso) P9() *geometry.Point {
	return s.P8().SquareToHorizontalLine(s.P0().Y)
}

func (s *Torso) P10() *geometry.Point {
	return s.P6().SquareLeft(s.P4().DistanceTo(s.P8()) - 1.27)
}

func (s *Torso) P11() *geometry.Point {
	return s.P10().SquareToHorizontalLine(s.P0().Y)
}

func (s *Torso) P12() *geometry.Point {
	return s.P4().SquareRight(s.TorsoMeasurements.ChestCircumference/4.0 + 0.635)
}

func (s *Torso) P13() *geometry.Point {
	n := s.P12().SquareToHorizontalLine(s.P5().Y)

	ang := s.P6().AngleRelativeTo(s.P12()).Perpendicular()
	a := n.DistanceTo(s.P12())
	o := a * ang.Tan()

	return n.SquareRight(o)
}

func (s *Torso) P14() *geometry.Point {
	return s.P12().SquareToHorizontalLine(s.P1().Y)
}

func (s *Torso) P15() *geometry.Point {
	if s.TorsoMeasurements.BellyButtonWaistCircumference <= s.TorsoMeasurements.ChestCircumference {
		return s.P14()
	}

	return s.P14().SquareLeft((s.TorsoMeasurements.BellyButtonWaistCircumference - s.TorsoMeasurements.ChestCircumference) / 2.0)
}

func (s *Torso) P16() *geometry.Point {
	return s.P0().SquareRight(s.backNeckWidth())
}

func (s *Torso) P17() *geometry.Point {
	return s.P16().SquareUp(s.P0().DistanceTo(s.P16()) / 3.0)
}

func (s *Torso) P18() *geometry.Point {
	return s.P9().SquareDown(1.59)
}

func (s *Torso) P19() *geometry.Point {
	bb := s.P17()

	return bb.DrawAt(s.P18().AngleRelativeTo(bb),  s.shoulderLength())
}

func (s *Torso) P20() *geometry.Point {
	return s.P19().SquareToVerticalLine(s.P0().X)
}

func (s *Torso) P21() *geometry.Point {
	return s.P8().MidpointTo(s.P18())
}

func (s *Torso) P22() *geometry.Point {
	a := s.P4().SquareUp(s.P4().DistanceTo(s.P0()) / 4.0)
	return a.SquareToVerticalLine(s.P8().X).SquareRight(0.635)
}

func (s *Torso) P23() *geometry.Point {
	d := s.P12().DistanceTo(s.P8()) / 2.0
	return s.P8().DrawAt(&geometry.Angle{Rads: math.Pi / 4.0}, d)
}

func (s *Torso) P24() *geometry.Point {
	return s.P3().SquareDown(s.backNeckWidth() + 2.635)
}

func (s *Torso) P25() *geometry.Point {
	return s.P24().SquareLeft(s.backNeckWidth() + 0.6825)
}

func (s *Torso) P26() *geometry.Point {
	return s.P25().SquareToHorizontalLine(s.P3().Y)
}

func (s *Torso) P27() *geometry.Point {
	return s.P24().MidpointTo(s.P26())
}

func (s *Torso) P28() *geometry.Point {
	ang := s.P26().AngleRelativeTo(s.P24()).Perpendicular()

	return s.P27().DrawAt(ang, 2.22)
}

func (s *Torso) P29() *geometry.Point {
	return s.P11().SquareDown(3.81)
}

func (s *Torso) P30() *geometry.Point {
	return s.P26().DrawAt(s.P29().AngleRelativeTo(s.P26()), s.shoulderLength())
}

func (s *Torso) P31() *geometry.Point {
	return s.P10().SquareUp(s.P29().DistanceTo(s.P10()) / 3.0 + 1.59)
}

func (s *Torso) P32() *geometry.Point {
	dist := s.P8().DistanceTo(s.P23()) - 0.635

	return s.P10().DrawAt(&geometry.Angle{Rads: math.Pi * 3.0 / 4.0}, dist)
}

func (s *Torso) P33() *geometry.Point {
	ang := s.P18().AngleRelativeTo(s.P17())
	dist := s.P17().DistanceTo(s.P16()) * ang.Tan()

	return s.P16().SquareLeft(math.Abs(dist))
}

func (s *Torso) P34() *geometry.Point {
	ang := s.P19().AngleRelativeTo(s.P18())
	o := s.P19().DistanceTo(s.P18())

	h := o / ang.Sin()

	return s.P18().SquareDown(math.Abs(h))
}

func (s *Torso) P35() *geometry.Point {
	ang := s.P30().AngleRelativeTo(s.P29())
	o := s.P30().DistanceTo(s.P29())

	h := o / ang.Sin()

	return s.P29().SquareDown(math.Abs(h))
}

func (s *Torso) P36() *geometry.Point {
	return s.P0().SquareDown(7.6)
}

func (s *Torso) P37() *geometry.Point {
	return s.P36().SquareToVerticalLine(s.P9().X)
}

func (s *Torso) P38() *geometry.Point {
	ang := s.P26().AngleRelativeTo(s.P29())
	adj := s.P25().DistanceTo(s.P26())

	opp := adj * ang.Tan()

	return s.P25().SquareRight(math.Abs(opp))
}

func (s *Torso) P39() *geometry.Point {
	return s.P12().SquareToHorizontalLine(s.P5().Y)
}

func (s *Torso) backNeckWidth() float64 {
	return s.TorsoMeasurements.NeckCircumference / 6.0
}

func (s *Torso) shoulderLength() float64 {
	return (s.TorsoMeasurements.ShoulderToShoulder/2.0)-s.P16().DistanceTo(s.P0())
}

func (s *Torso) frontArmhole() *geometry.Polyline {
	poly := &geometry.Polyline{}

	poly.AddLine(
		&geometry.QuadraticBezierCurve{
			P0: s.P30(),
			P1: s.P35(),
			P2: s.P31(),
		},
		&geometry.QuadraticBezierCurve{
			P0: s.P31(),
			P1: s.P10(),
			P2: s.P12(),
		},
	)

	return poly
}

func (s *Torso) yokeArmhole() *geometry.Polyline {
	poly := &geometry.Polyline{}

	poly.AddLine(
		&geometry.QuadraticBezierCurve{
			P0: s.P19(),
			P1: s.P34(),
			P2: s.P37(),
		},
	)

	return poly
}

func (s *Torso) backArmhole() *geometry.Polyline {
	poly := &geometry.Polyline{}

	poly.AddLine(
		&geometry.StraightLine{
			Start: s.P37(),
			End: s.P21(),
		},
		&geometry.QuadraticBezierCurve{
			P0: s.P21(),
			P1: s.P8(),
			P2: s.P12(),
		},
	)

	return poly
}

func (s *Torso) Stitch() *geometry.Block {
	return &geometry.Block{}
}

func (p *Torso) OuterCut() *geometry.Polyline {
	return &geometry.Polyline{}
}

func (s *Torso) InnerCut() *geometry.Block {
	return &geometry.Block{}
}

func (s *Torso) Ink() *geometry.Block {
	return &geometry.Block{}
}

func (s *Torso) Reference() *geometry.Block {
	layer := &geometry.Block{}

	neckLineFront := &geometry.QuadraticBezierCurve{
		P0: s.P24(),
		P1: s.P38(),
		P2: s.P26(),
	}
	neckLineBack := &geometry.QuadraticBezierCurve{
		P0: s.P0(),
		P1: s.P33(),
		P2: s.P17(),
	}

	layer.AddLine(
		&geometry.StraightLine{
			Start: s.P0(),
			End:   s.P1(),
		},
		&geometry.StraightLine{
			Start: s.P1(),
			End:   s.P2(),
		},
		&geometry.StraightLine{
			Start: s.P2(),
			End:   s.P3(),
		},
		&geometry.StraightLine{
			Start: s.P3(),
			End:   s.P0(),
		},
		&geometry.StraightLine{
			Start: s.P4(),
			End:   s.P6(),
		},
		&geometry.StraightLine{
			Start: s.P5(),
			End:   s.P7(),
		},
		&geometry.StraightLine{
			Start: s.P8(),
			End:   s.P9(),
		},
		&geometry.StraightLine{
			Start: s.P10(),
			End:   s.P11(),
		},
		&geometry.StraightLine{
			Start: s.P12(),
			End:   s.P15(),
		},
		&geometry.StraightLine{
			Start: s.P17(),
			End:   s.P19(),
		},
		neckLineBack,
		s.yokeArmhole(),
		s.backArmhole(),
		neckLineFront,
		&geometry.StraightLine{
			Start: s.P26(),
			End:   s.P30(),
		},
		s.frontArmhole(),
		&geometry.StraightLine{
			Start: s.P36(),
			End:   s.P37(),
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
	a["24"] = s.P24()
	a["25"] = s.P25()
	a["26"] = s.P26()
	a["27"] = s.P27()
	a["28"] = s.P28()
	a["29"] = s.P29()
	a["30"] = s.P30()
	a["31"] = s.P31()
	a["32"] = s.P32()
	a["33"] = s.P33()
	a["34"] = s.P34()
	a["35"] = s.P35()
	a["36"] = s.P36()
	a["37"] = s.P37()
	a["38"] = s.P38()
	anchors.AddAnchors(layer, a)

	return layer
}

func (s *Torso) CutCount() int {
	return 1
}

func (s *Torso) OnFold() bool {
	return true
}

func (s *Torso) Mirrored() bool {
	return false
}

func (s *Torso) Details() *pieces.Details {
	return &pieces.Details{}
}
