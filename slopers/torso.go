package slopers

import (
	"math"

	"github.com/tailored-style/pattern-generator/geometry"
	"github.com/tailored-style/pattern-generator/pieces"
	catalogue_pieces "github.com/tobyjsullivan/catalogue/pieces"
	"fmt"
)

type Torso struct {
	Height                        float64
	NeckCircumference             float64
	ShoulderToShoulder            float64
	ChestCircumference            float64
	ShirtLength                   float64
	BellyButtonWaistCircumference float64
	NaturalWaistCircumference     float64
	HipCircumference              float64
	ShirtSleeveLength             float64
	BicepCircumference            float64
	WristCircumference            float64
}

func (s *Torso) A() *geometry.Point {
	return &geometry.Point{X: 0.0, Y: 0.0}
}

func (s *Torso) B() *geometry.Point {
	return s.A().SquareDown(s.ShirtLength)
}

func (s *Torso) C() *geometry.Point {
	return s.B().SquareRight(s.ChestCircumference / 2.0)
}

func (s *Torso) D() *geometry.Point {
	return s.C().SquareToHorizontalLine(s.A().Y)
}

func (s *Torso) E() *geometry.Point {
	return s.A().SquareDown(s.ChestCircumference / 4.0)
}

func (s *Torso) F() *geometry.Point {
	return s.B().SquareUp(s.ShirtLength / 3.0)
}

func (s *Torso) G() *geometry.Point {
	return s.E().SquareToVerticalLine(s.D().X)
}

func (s *Torso) H() *geometry.Point {
	return s.F().SquareToVerticalLine(s.D().X)
}

func (s *Torso) I() *geometry.Point {
	return s.E().SquareRight((s.ChestCircumference / 6.0) + 2.0)
}

func (s *Torso) J() *geometry.Point {
	return s.I().SquareToHorizontalLine(s.A().Y)
}

func (s *Torso) K() *geometry.Point {
	return s.G().SquareLeft(s.E().DistanceTo(s.I()) - 1.27)
}

func (s *Torso) L() *geometry.Point {
	return s.K().SquareToHorizontalLine(s.A().Y)
}

func (s *Torso) M() *geometry.Point {
	return s.E().SquareRight(s.ChestCircumference/4.0 + 0.635)
}

func (s *Torso) N() *geometry.Point {
	n := s.M().SquareToHorizontalLine(s.F().Y)

	ang := s.P().AngleRelativeTo(s.M()).Perpendicular()
	a := n.DistanceTo(s.M())
	o := a * ang.Tan()

	return n.SquareRight(o)
}

func (s *Torso) O() *geometry.Point {
	return s.M().SquareToHorizontalLine(s.B().Y)
}

func (s *Torso) P() *geometry.Point {
	if s.BellyButtonWaistCircumference <= s.ChestCircumference {
		return s.O()
	}

	return s.O().SquareLeft((s.BellyButtonWaistCircumference - s.ChestCircumference) / 2.0)
}

func (s *Torso) backNeckWidth() float64 {
	return s.NeckCircumference / 6.0
}

func (s *Torso) BA() *geometry.Point {
	return s.A().SquareRight(s.backNeckWidth())
}

func (s *Torso) BB() *geometry.Point {
	return s.BA().SquareUp(s.A().DistanceTo(s.BA()) / 3.0)
}

func (s *Torso) BC() *geometry.Point {
	return s.J().SquareDown(1.59)
}

func (s *Torso) BD() *geometry.Point {
	bb := s.BB()
	bc := s.BC()
	shoulderLength :=  (s.ShoulderToShoulder/2.0)-s.BA().DistanceTo(s.A())

	return bb.DrawAt(bc.AngleRelativeTo(bb), shoulderLength)
}

func (s *Torso) BE() *geometry.Point {
	return s.BD().SquareToVerticalLine(s.A().X)
}

func (s *Torso) BF() *geometry.Point {
	return s.I().MidpointTo(s.BC())
}

func (s *Torso) BG() *geometry.Point {
	a := s.E().SquareUp(s.E().DistanceTo(s.A()) / 4.0)
	return a.SquareToVerticalLine(s.I().X).SquareRight(0.635)
}

func (s *Torso) BH() *geometry.Point {
	d := s.M().DistanceTo(s.I()) / 2.0
	return s.I().DrawAt(&geometry.Angle{Rads: math.Pi / 4.0}, d)
}

func (s *Torso) FI() *geometry.Point {
	return s.D().SquareDown(s.backNeckWidth() + 0.635)
}

func (s *Torso) FJ() *geometry.Point {
	return s.FI().SquareLeft(s.backNeckWidth() - 0.3175)
}

func (s *Torso) FK() *geometry.Point {
	return s.FJ().SquareToHorizontalLine(s.D().Y)
}

func (s *Torso) FL() *geometry.Point {
	return s.FI().MidpointTo(s.FK())
}

func (s *Torso) FM() *geometry.Point {
	ang := s.FK().AngleRelativeTo(s.FI()).Perpendicular()

	return s.FL().DrawAt(ang, 2.22)
}

func (s *Torso) FN() *geometry.Point {
	return s.L().SquareDown(3.81)
}

func (s *Torso) FO() *geometry.Point {
	return s.FK().DrawAt(s.FN().AngleRelativeTo(s.FK()), s.BD().DistanceTo(s.BB()))
}

func (s *Torso) FP() *geometry.Point {
	return s.K().SquareUp(s.FN().DistanceTo(s.K()) / 3.0 + 1.59)
}

func (s *Torso) FQ() *geometry.Point {
	dist := s.I().DistanceTo(s.BH()) - 0.635

	return s.K().DrawAt(&geometry.Angle{Rads: math.Pi * 3.0 / 4.0}, dist)
}

func (s *Torso) ZA() *geometry.Point {
	ang := s.BC().AngleRelativeTo(s.BB())
	dist := s.BB().DistanceTo(s.BA()) * ang.Tan()

	return s.BA().SquareLeft(math.Abs(dist))
}

func (s *Torso) ZB() *geometry.Point {
	ang := s.BD().AngleRelativeTo(s.BC())
	o := s.BD().DistanceTo(s.BC())

	h := o / ang.Sin()

	return s.BC().SquareDown(math.Abs(h))
}

func (s *Torso) ZC() *geometry.Point {
	ang := s.FO().AngleRelativeTo(s.FN())
	o := s.FO().DistanceTo(s.FN())

	h := o / ang.Sin()

	return s.FN().SquareDown(math.Abs(h))
}

func (s *Torso) ZD() *geometry.Point {
	return s.A().SquareDown(7.6)
}

func (s *Torso) ZE() *geometry.Point {
	return s.ZD().SquareToVerticalLine(s.J().X)
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

	layer.AddLine(
		&geometry.StraightLine{
			Start: s.A(),
			End:   s.B(),
		},
		&geometry.StraightLine{
			Start: s.B(),
			End:   s.C(),
		},
		&geometry.StraightLine{
			Start: s.C(),
			End:   s.D(),
		},
		&geometry.StraightLine{
			Start: s.D(),
			End:   s.A(),
		},
		&geometry.StraightLine{
			Start: s.E(),
			End:   s.G(),
		},
		&geometry.StraightLine{
			Start: s.F(),
			End:   s.H(),
		},
		&geometry.StraightLine{
			Start: s.I(),
			End:   s.J(),
		},
		&geometry.StraightLine{
			Start: s.K(),
			End:   s.L(),
		},
		&geometry.StraightLine{
			Start: s.M(),
			End:   s.P(),
		},
		&geometry.StraightLine{
			Start: s.BB(),
			End:   s.BD(),
		},
		&geometry.QuadraticBezierCurve{
			P0: s.A(),
			P1: s.ZA(),
			P2: s.BB(),
		},
		&geometry.QuadraticBezierCurve{
			P0: s.BD(),
			P1: s.ZB(),
			P2: s.ZE(),
		},
		&geometry.QuadraticBezierCurve{
			P0: s.M(),
			P1: s.I(),
			P2: s.BF(),
		},
		&geometry.QuadraticBezierCurve{
			P0: s.FI(),
			P1: s.FJ(),
			P2: s.FK(),
		},
		&geometry.StraightLine{
			Start: s.FK(),
			End:   s.FO(),
		},
		&geometry.QuadraticBezierCurve{
			P0: s.FO(),
			P1: s.ZC(),
			P2: s.FP(),
		},
		&geometry.QuadraticBezierCurve{
			P0: s.FP(),
			P1: s.K(),
			P2: s.M(),
		},
		&geometry.StraightLine{
			Start: s.ZD(),
			End:   s.ZE(),
		},
	)

	anchors := make(map[string]*geometry.Point)
	anchors["A"] = s.A()
	anchors["B"] = s.B()
	anchors["C"] = s.C()
	anchors["D"] = s.D()
	anchors["E"] = s.E()
	anchors["F"] = s.F()
	anchors["G"] = s.G()
	anchors["H"] = s.H()
	anchors["I"] = s.I()
	anchors["J"] = s.J()
	anchors["K"] = s.K()
	anchors["L"] = s.L()
	anchors["M"] = s.M()
	anchors["N"] = s.N()
	anchors["O"] = s.O()
	anchors["P"] = s.P()
	anchors["BA"] = s.BA()
	anchors["BB"] = s.BB()
	anchors["BC"] = s.BC()
	anchors["BD"] = s.BD()
	anchors["BE"] = s.BE()
	anchors["BF"] = s.BF()
	anchors["BG"] = s.BG()
	anchors["BH"] = s.BH()
	anchors["FI"] = s.FI()
	anchors["FJ"] = s.FJ()
	anchors["FK"] = s.FK()
	anchors["FL"] = s.FL()
	anchors["FM"] = s.FM()
	anchors["FN"] = s.FN()
	anchors["FO"] = s.FO()
	anchors["FP"] = s.FP()
	anchors["FQ"] = s.FQ()
	anchors["ZA"] = s.ZA()
	anchors["ZB"] = s.ZB()
	anchors["ZC"] = s.ZC()
	anchors["ZD"] = s.ZD()
	anchors["ZE"] = s.ZE()
	catalogue_pieces.AddAnchors(layer, anchors)

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
