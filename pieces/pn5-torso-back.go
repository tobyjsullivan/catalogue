package pieces

import (
	"fmt"
	"math"

	"github.com/tailored-style/pattern-generator/geometry"
	"github.com/tailored-style/pattern-generator/pieces"
)

type pn5TorsoBack struct {
	height             float64
	chestCircumference float64
	waistCircumference float64
	hipCircumference   float64
}

func NewPN5TorsoBack(height float64, chest float64, waist float64, hip float64) pieces.Piece {
	return &pn5TorsoBack{
		height:             height,
		chestCircumference: chest,
		waistCircumference: waist,
		hipCircumference:   hip,
	}
}

func (p *pn5TorsoBack) Details() *pieces.Details {
	return &pieces.Details{
		PieceNumber: 5,
		Description: "Torso Back",
	}
}

func (p *pn5TorsoBack) OnFold() bool {
	return true
}

func (p *pn5TorsoBack) a() *geometry.Point {
	return &geometry.Point{X: 0.0, Y: 0.0}
}

func (p *pn5TorsoBack) b() *geometry.Point {
	return p.a().SquareDown(p.chestCircumference/4.0 - 8.6)
}

func (p *pn5TorsoBack) c() *geometry.Point {
	return p.b().SquareRight(p.chestCircumference/4.0 + 1.6)
}

func (p *pn5TorsoBack) d() *geometry.Point {
	return p.a().SquareDown(p.height/4.0 - 11.8)
}

func (p *pn5TorsoBack) e() *geometry.Point {
	return p.d().SquareRight(p.waistCircumference/4.0 + 3.2)
}

func (p *pn5TorsoBack) f() *geometry.Point {
	return p.a().SquareDown(p.height/3.0 - 14.3)
}

func (p *pn5TorsoBack) g() *geometry.Point {
	return p.f().SquareRight(p.waistCircumference/4.0 + 4.3)
}

func (p *pn5TorsoBack) h() *geometry.Point {
	return p.a().SquareDown(p.height*(3.0/8.0) - 5.4)
}

func (p *pn5TorsoBack) i() *geometry.Point {
	return p.h().SquareRight(p.hipCircumference/4.0 + 0.6)
}

func (p *pn5TorsoBack) j() *geometry.Point {
	return p.i().SquareUp(7.3)
}

func (p *pn5TorsoBack) k() *geometry.Point {
	return p.h().SquareDown(5.4)
}

func (p *pn5TorsoBack) l() *geometry.Point {
	return p.k().SquareRight(7.6)
}

func (p *pn5TorsoBack) m() *geometry.Point {
	return p.b().SquareRight(p.chestCircumference/6.0 + 6.2)
}

func (p *pn5TorsoBack) n() *geometry.Point {
	return p.m().SquareToHorizontalLine(p.a().Y)
}

func (p *pn5TorsoBack) o() *geometry.Point {
	return p.n().SquareDown(1.1)
}

func (p *pn5TorsoBack) p() *geometry.Point {
	m := p.m()
	return m.SquareUp(m.DistanceTo(p.o())*(2.0/3.0) + 1.3)
}

func (p *pn5TorsoBack) q() *geometry.Point {
	return p.n().SquareLeft(8.4)
}

func (p *pn5TorsoBack) r() *geometry.Point {
	return p.p().SquareLeft(0.5)
}

func (p *pn5TorsoBack) s() *geometry.Point {
	return p.b().MidpointTo(p.m())
}

func (p *pn5TorsoBack) t() *geometry.Point {
	return p.s().SquareDown(p.height/8.0 - 2.5)
}

func (p *pn5TorsoBack) u() *geometry.Point {
	s := p.s()
	return s.SquareDown(s.DistanceTo(p.t())*2.0 - 3.8)
}

func (p *pn5TorsoBack) v() *geometry.Point {
	return p.t().SquareLeft(1.3)
}

func (p *pn5TorsoBack) w() *geometry.Point {
	return p.t().SquareRight(1.3)
}

func (p *pn5TorsoBack) centreBack() geometry.Line {
	return &geometry.StraightLine{
		Start: p.a(),
		End:   p.k(),
	}
}

func (p *pn5TorsoBack) yokeSeamStitch() geometry.Line {
	yokeSeamA := &geometry.StraightLine{
		Start: p.a(),
		End:   p.q(),
	}

	yokeSeamB := &geometry.EllipseCurve{
		Start:         p.q(),
		End:           p.o(),
		StartingAngle: &geometry.Angle{Rads: math.Pi * (3.0 / 2.0)},
		ArcAngle:      &geometry.Angle{Rads: math.Pi / 8.0},
	}

	line := &geometry.Polyline{}
	line.AddLine(
		yokeSeamA,
		yokeSeamB,
	)

	return line
}

func (p *pn5TorsoBack) armholeStitch() geometry.Line {
	armscyeA := &geometry.EllipseCurve{
		Start:         p.r(),
		End:           p.o(),
		StartingAngle: &geometry.Angle{Rads: 0.0},
		ArcAngle:      &geometry.Angle{Rads: math.Pi / 8.0},
	}

	armscyeB := &geometry.EllipseCurve{
		Start:         p.r(),
		End:           p.c(),
		StartingAngle: &geometry.Angle{Rads: math.Pi},
		ArcAngle:      &geometry.Angle{Rads: math.Pi / 2.0},
	}

	line := &geometry.Polyline{}
	line.AddLine(
		&geometry.ReverseLine{InnerLine: armscyeA},
		armscyeB,
	)

	return line
}

func (p *pn5TorsoBack) sideSeamStitch() geometry.Line {
	return &geometry.PolyNCurve{
		Points: []*geometry.Point{
			p.c(),
			p.e(),
			p.g(),
			p.i(),
		},
		StartAngle: &geometry.Angle{Rads: -math.Pi / 2.0},
		EndAngle:   &geometry.Angle{Rads: -math.Pi / 2.0},
		Vertical:   true,
	}
}

func (p *pn5TorsoBack) hemLineStitch() geometry.Line {
	hemLineA := &geometry.StraightLine{
		Start: p.k(),
		End:   p.l(),
	}

	hemLineB := &geometry.SCurve{
		Start:         p.l(),
		End:           p.i(),
		StartingAngle: &geometry.Angle{Rads: math.Pi * (3.0 / 2.0)},
		FinishAngle:   &geometry.Angle{Rads: math.Pi * (3.0 / 2.0)},
		MaxAngle:      &geometry.Angle{Rads: math.Pi / 4.0},
	}

	line := &geometry.Polyline{}
	line.AddLine(
		hemLineA,
		hemLineB,
	)

	return line
}

func (p *pn5TorsoBack) dartStitch() geometry.Line {
	dart := &geometry.Polyline{}

	dart.AddLine(
		&geometry.StraightLine{
			Start: p.s(),
			End:   p.v(),
		},
		&geometry.StraightLine{
			Start: p.v(),
			End:   p.u(),
		},
		&geometry.StraightLine{
			Start: p.u(),
			End:   p.w(),
		},
		&geometry.StraightLine{
			Start: p.w(),
			End:   p.s(),
		},
	)

	return dart
}

func (p *pn5TorsoBack) CutLayer() *geometry.Block {
	layer := &geometry.Block{}

	armscyeCut := pieces.AddSeamAllowance(p.armholeStitch(), false)

	layer.AddLine(
		p.centreBack(),
		pieces.AddSeamAllowance(p.yokeSeamStitch(), false),
		armscyeCut,
		pieces.Notch(armscyeCut, 7.6),
		pieces.Notch(armscyeCut, armscyeCut.Length()-7.6),
		pieces.AddSeamAllowance(p.sideSeamStitch(), false),
		pieces.AddSeamAllowance(p.hemLineStitch(), true),
	)

	return layer
}

func (p *pn5TorsoBack) StitchLayer() *geometry.Block {
	layer := &geometry.Block{}

	layer.AddLine(
		p.yokeSeamStitch(),
		p.armholeStitch(),
		p.sideSeamStitch(),
		p.hemLineStitch(),
		p.dartStitch(),
	)

	return layer
}

func (p *pn5TorsoBack) NotationLayer() *geometry.Block {
	layer := &geometry.Block{}

	chestLine := &geometry.StraightLine{
		Start: p.b(),
		End:   p.c(),
	}

	naturalWaistLine := &geometry.StraightLine{
		Start: p.d(),
		End:   p.e(),
	}

	bellyButtonWaistLine := &geometry.StraightLine{
		Start: p.f(),
		End:   p.g(),
	}

	hipLine := &geometry.StraightLine{
		Start: p.h(),
		End:   p.i(),
	}

	layer.AddLine(
		chestLine,
		naturalWaistLine,
		bellyButtonWaistLine,
		hipLine,
	)

	// Draw all points (DEBUG)
	anchors := make(map[string]*geometry.Point)
	anchors["A"] = p.a()
	anchors["B"] = p.b()
	anchors["C"] = p.c()
	anchors["D"] = p.d()
	anchors["E"] = p.e()
	anchors["F"] = p.f()
	anchors["G"] = p.g()
	anchors["H"] = p.h()
	anchors["I"] = p.i()
	anchors["J"] = p.j()
	anchors["K"] = p.k()
	anchors["L"] = p.l()
	anchors["M"] = p.m()
	anchors["N"] = p.n()
	anchors["O"] = p.o()
	anchors["P"] = p.p()
	anchors["Q"] = p.q()
	anchors["R"] = p.r()
	anchors["S"] = p.s()
	anchors["T"] = p.t()
	anchors["U"] = p.u()
	anchors["V"] = p.v()
	anchors["W"] = p.w()
	AddAnchors(layer, anchors)

	return layer
}

func (p *pn5TorsoBack) String() string {
	return fmt.Sprintf("[PN: %d]", p.Details().PieceNumber)
}
