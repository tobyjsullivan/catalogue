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

func (s *Sleeve) A() *geometry.Point {
	return &geometry.Point{X: 0.0, Y: 0.0}
}

func (s *Sleeve) B() *geometry.Point {
	return s.A().SquareDown((s.TorsoMeasurements.ShirtSleeveLength - s.torso().shoulderLength()) + 2.54)
}

func (s *Sleeve) frontArmholeLength() float64 {
	return s.torso().frontArmhole().Length()
}

func (s *Sleeve) backArmholeLength() float64 {
	torso := s.torso()
	return torso.backArmhole().Length() + torso.yokeArmhole().Length()
}

func (s *Sleeve) C() *geometry.Point {
	armholeLength := s.frontArmholeLength() + s.backArmholeLength()
	return s.A().SquareDown(armholeLength / 3.0 - 2.86)
}

func (s *Sleeve) D() *geometry.Point {
	return s.C().MidpointTo(s.B()).SquareUp(3.81)
}

func (s *Sleeve) E() *geometry.Point {
	h := s.frontArmholeLength() - 0.635
	a := s.A().DistanceTo(s.C())

	b := math.Sqrt((h * h) - (a * a))

	return s.C().SquareRight(b)
}

func (s *Sleeve) F() *geometry.Point {
	h := s.backArmholeLength() - 0.3175
	a := s.A().DistanceTo(s.C())

	b := math.Sqrt(h * h - a * a)

	return s.C().SquareLeft(b)
}

func (s *Sleeve) G() *geometry.Point {
	return s.E().SquareToHorizontalLine(s.B().Y)
}

func (s *Sleeve) H() *geometry.Point {
	return s.F().SquareToHorizontalLine(s.B().Y)
}

func (s *Sleeve) I() *geometry.Point {
	return s.E().SquareToHorizontalLine(s.D().Y)
}

func (s *Sleeve) J() *geometry.Point {
	return s.F().SquareToHorizontalLine(s.D().Y)
}

func (s *Sleeve) K() *geometry.Point {
	return s.A().MidpointTo(s.A().MidpointTo(s.E()))
}

func (s *Sleeve) L() *geometry.Point {
	l := s.A().MidpointTo(s.E())
	return l.DrawAt(l.AngleRelativeTo(s.A()), 1.27)
}

func (s *Sleeve) M() *geometry.Point {
	return s.A().MidpointTo(s.E()).MidpointTo(s.E())
}

func (s *Sleeve) N() *geometry.Point {
	return s.K().DrawAt(s.K().AngleRelativeTo(s.A()).Perpendicular(), 1.59)
}

func (s *Sleeve) O() *geometry.Point {
	return s.M().DrawAt(s.M().AngleRelativeTo(s.A()).Perpendicular().Opposite(), 1.27)
}

func (s *Sleeve) P() *geometry.Point {
	return s.A().MidpointTo(s.A().MidpointTo(s.F()))
}

func (s *Sleeve) Q() *geometry.Point {
	return s.A().MidpointTo(s.F())
}

func (s *Sleeve) R() *geometry.Point {
	return s.F().MidpointTo(s.A().MidpointTo(s.F()))
}

func (s *Sleeve) S() *geometry.Point {
	return s.P().DrawAt(s.A().AngleRelativeTo(s.P()).Perpendicular(), 1.905)
}

func (s *Sleeve) T() *geometry.Point {
	return s.Q().DrawAt(s.A().AngleRelativeTo(s.Q()).Perpendicular(), 0.9525)
}

func (s *Sleeve) U() *geometry.Point {
	return s.R().DrawAt(s.A().AngleRelativeTo(s.R()), 1.43)
}

func (s *Sleeve) V() *geometry.Point {
	p := s.R().MidpointTo(s.F())
	return p.DrawAt(s.F().AngleRelativeTo(s.R()).Perpendicular(), 0.635)
}

func (s *Sleeve) W() *geometry.Point {
	return s.G().SquareLeft(5.715)
}

func (s *Sleeve) X() *geometry.Point {
	return s.H().SquareRight(5.715)
}

func (s *Sleeve) sleeveCap() *geometry.Polyline {
	poly := &geometry.Polyline{}

	poly.AddLine(
		&geometry.PolyNCurve{
			Points: []*geometry.Point{
				s.F(),
				s.V(),
				s.U(),
				s.T(),
				s.S(),
				s.A(),
			},
			StartAngle: s.X().AngleRelativeTo(s.F()).Perpendicular(),
			EndAngle: &geometry.Angle{Rads: 0.0},
		},
		&geometry.PolyNCurve{
			Points: []*geometry.Point{
				s.A(),
				s.N(),
				s.L(),
				s.O(),
				s.E(),
			},
			StartAngle: &geometry.Angle{Rads: 0.0},
			EndAngle: s.W().AngleRelativeTo(s.E()).Perpendicular(),
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
			Start: s.E(),
			End: s.W(),
		},
		&geometry.StraightLine{
			Start: s.F(),
			End: s.X(),
		},
		&geometry.StraightLine{
			Start: s.W(),
			End: s.X(),
		},

	)

	a := make(map[string]*geometry.Point)
	a["A"] = s.A()
	a["B"] = s.B()
	a["C"] = s.C()
	a["D"] = s.D()
	a["E"] = s.E()
	a["F"] = s.F()
	a["G"] = s.G()
	a["H"] = s.H()
	a["I"] = s.I()
	a["J"] = s.J()
	a["K"] = s.K()
	a["L"] = s.L()
	a["M"] = s.M()
	a["N"] = s.N()
	a["O"] = s.O()
	a["P"] = s.P()
	a["Q"] = s.Q()
	a["R"] = s.R()
	a["S"] = s.S()
	a["T"] = s.T()
	a["U"] = s.U()
	a["V"] = s.V()
	a["W"] = s.W()
	a["X"] = s.X()
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
