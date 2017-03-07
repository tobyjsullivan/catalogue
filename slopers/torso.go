package slopers

import (
	"math"

	"github.com/tailored-style/pattern-generator/geometry"
	"github.com/tailored-style/pattern-generator/pieces"
	catalogue_pieces "github.com/tobyjsullivan/catalogue/pieces"
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
	return s.E().SquareRight(s.ShoulderToShoulder / 2.0)
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

func (s *Torso) O() *geometry.Point {
	return s.M().SquareToHorizontalLine(s.B().Y)
}

func (s *Torso) P() *geometry.Point {
	if s.BellyButtonWaistCircumference <= s.ChestCircumference {
		return s.O()
	}

	return s.O().SquareLeft((s.BellyButtonWaistCircumference - s.ChestCircumference) / 2.0)
}

func (s *Torso) BA() *geometry.Point {
	return s.A().SquareRight(s.NeckCircumference / 6.0)
}

func (s *Torso) BB() *geometry.Point {
	return s.BA().SquareUp(s.A().DistanceTo(s.BA()) / 3.0)
}

func (s *Torso) BC() *geometry.Point {
	return s.J().SquareDown(1.59)
}

func (s *Torso) BD() *geometry.Point {
	bb := s.BB()
	return bb.DrawAt(s.BC().AngleRelativeTo(bb), (s.ShoulderToShoulder/2.0)-s.BA().DistanceTo(s.A()))
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

func (s *Torso) Stitch() *geometry.Block {
	return &geometry.Block{}
}

func (s *Torso) InnerCut() *geometry.Block {
	return &geometry.Block{}
}

func (s *Torso) Ink() *geometry.Block {
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
		&geometry.EllipseCurve{
			Start:         s.A(),
			End:           s.BB(),
			StartingAngle: &geometry.Angle{Rads: math.Pi * 3.0 / 2.0},
			ArcAngle:      s.BD().AngleRelativeTo(s.BB()),
		},
		&geometry.EllipseCurve{
			Start:         s.M(),
			End:           s.BF(),
			StartingAngle: &geometry.Angle{Rads: math.Pi * 3.0 / 2.0},
			ArcAngle:      &geometry.Angle{Rads: math.Pi / 2.0},
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
	anchors["P"] = s.P()
	anchors["BA"] = s.BA()
	anchors["BB"] = s.BB()
	anchors["BC"] = s.BC()
	anchors["BD"] = s.BD()
	anchors["BE"] = s.BE()
	anchors["BF"] = s.BF()
	anchors["BG"] = s.BG()
	anchors["BH"] = s.BH()
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
