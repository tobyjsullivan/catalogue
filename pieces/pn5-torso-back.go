package pieces

import (
	"fmt"
	"math"

	"github.com/tailored-style/pattern-generator/geometry"
	"github.com/tailored-style/pattern-generator/pieces"
	"github.com/tobyjsullivan/catalogue/anchors"
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

func (p *pn5TorsoBack) CutCount() int {
	return 1
}

func (p *pn5TorsoBack) OnFold() bool {
	return true
}

func (p *pn5TorsoBack) Mirrored() bool {
	return false
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

func (p *pn5TorsoBack) bellyWaistGirth() float64 {
	return p.waistCircumference/4.0 + 4.3
}

func (p *pn5TorsoBack) g() *geometry.Point {
	return p.f().SquareRight(p.bellyWaistGirth())
}

func (p *pn5TorsoBack) h() *geometry.Point {
	return p.a().SquareDown(p.height*(3.0/8.0) - 5.4)
}

func (p *pn5TorsoBack) i() *geometry.Point {
	hip := p.hipCircumference/4.0 + 0.6
	bellyWaist := p.bellyWaistGirth()
	if hip < bellyWaist {
		hip = bellyWaist
	}

	return p.h().SquareRight(hip)
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

	armscyeB := &geometry.QuadraticBezierCurve{
		P0: p.r(),
		P1: p.m(),
		P2: p.c(),
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

func (p *pn5TorsoBack) OuterCut() *geometry.Polyline {
	return &geometry.Polyline{}
}

func (p *pn5TorsoBack) InnerCut() *geometry.Block {
	layer := &geometry.Block{}

	armholeStitch := p.armholeStitch()

	seamAllowance := pieces.SeamAllowance(false,
		pieces.AddSeamAllowance(p.yokeSeamStitch(), false),
		pieces.AddSeamAllowance(armholeStitch, false),
		pieces.AddSeamAllowance(p.sideSeamStitch(), false),
		&geometry.ReverseLine{InnerLine: pieces.AddSeamAllowance(p.hemLineStitch(), true)},
	)

	layer.AddLine(
		seamAllowance,
		pieces.Notch(armholeStitch, armholeStitch.Length() - 7.6, false),
		pieces.Notch(armholeStitch, armholeStitch.Length() - 8.6, false),
	)

	return layer
}

func (p *pn5TorsoBack) Stitch() *geometry.Block {
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

func (p *pn5TorsoBack) Ink() *geometry.Block {
	return &geometry.Block{}
}

func (p *pn5TorsoBack) Reference() *geometry.Block {
	layer := &geometry.Block{}

	layer.AddLine(
		p.centreBack(),
	)

	if DEBUG {
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

		// Draw all points
		a := make(map[string]*geometry.Point)
		a["A"] = p.a()
		a["B"] = p.b()
		a["C"] = p.c()
		a["D"] = p.d()
		a["E"] = p.e()
		a["F"] = p.f()
		a["G"] = p.g()
		a["H"] = p.h()
		a["I"] = p.i()
		a["J"] = p.j()
		a["K"] = p.k()
		a["L"] = p.l()
		a["M"] = p.m()
		a["N"] = p.n()
		a["O"] = p.o()
		a["P"] = p.p()
		a["Q"] = p.q()
		a["R"] = p.r()
		a["S"] = p.s()
		a["T"] = p.t()
		a["U"] = p.u()
		a["V"] = p.v()
		a["W"] = p.w()
		anchors.AddAnchors(layer, a)
	}

	return layer
}

func (p *pn5TorsoBack) String() string {
	return fmt.Sprintf("[PN: %d]", p.Details().PieceNumber)
}
