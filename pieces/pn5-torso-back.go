package pieces

import (
	"github.com/tailored-style/pattern-generator/geometry"
	"math"
	"github.com/tailored-style/pattern-generator/pieces"
)

type PN5TorsoBack struct {
	*pieces.Measurements
}

func (p *PN5TorsoBack) Details() *pieces.Details {
	return &pieces.Details{
		PieceNumber: "5",
		Description: "Torso Back",
	}
}

func (p *PN5TorsoBack) OnFold() bool {
	return true
}

func (p *PN5TorsoBack) a() *geometry.Point {
	return &geometry.Point{X: 0.0, Y: 0.0}
}

func (p *PN5TorsoBack) b() *geometry.Point {
	return p.a().SquareDown(p.ChestCircumference/4.0 - 8.6)
}

func (p *PN5TorsoBack) c() *geometry.Point {
	return p.b().SquareRight(p.ChestCircumference/4.0 + 1.6)
}

func (p *PN5TorsoBack) d() *geometry.Point {
	return p.a().SquareDown(p.Height/4.0 - 11.4)
}

func (p *PN5TorsoBack) e() *geometry.Point {
	return p.d().SquareRight(p.WaistCircumference/4.0 + 3.2)
}

func (p *PN5TorsoBack) f() *geometry.Point {
	return p.a().SquareDown(p.Height*(7.0/24.0) - 6.4)
}

func (p *PN5TorsoBack) g() *geometry.Point {
	return p.f().SquareRight(p.WaistCircumference/4.0 + 4.3)
}

func (p *PN5TorsoBack) h() *geometry.Point {
	return p.a().SquareDown(p.Height*(3.0/8.0) - 4.8)
}

func (p *PN5TorsoBack) i() *geometry.Point {
	return p.h().SquareRight(p.HipCircumference/4.0 + 0.6)
}

func (p *PN5TorsoBack) j() *geometry.Point {
	return p.i().SquareUp(7.3)
}

func (p *PN5TorsoBack) k() *geometry.Point {
	return p.h().SquareDown(5.4)
}

func (p *PN5TorsoBack) l() *geometry.Point {
	return p.k().SquareRight(7.6)
}

func (p *PN5TorsoBack) m() *geometry.Point {
	return p.b().SquareRight(p.ChestCircumference/6.0 + 6.2)
}

func (p *PN5TorsoBack) n() *geometry.Point {
	return p.m().SquareToHorizontalLine(p.a().Y)
}

func (p *PN5TorsoBack) o() *geometry.Point {
	return p.n().SquareDown(1.1)
}

func (p *PN5TorsoBack) p() *geometry.Point {
	m := p.m()
	return m.SquareUp(m.DistanceTo(p.o())*(2.0/3.0) + 1.3)
}

func (p *PN5TorsoBack) q() *geometry.Point {
	return p.n().SquareLeft(8.4)
}

func (p *PN5TorsoBack) r() *geometry.Point {
	return p.p().SquareLeft(0.5)
}

func (p *PN5TorsoBack) s() *geometry.Point {
	return p.b().MidpointTo(p.m())
}

func (p *PN5TorsoBack) t() *geometry.Point {
	return p.s().SquareDown(p.Height/8.0 - 2.5)
}

func (p *PN5TorsoBack) u() *geometry.Point {
	s := p.s()
	return s.SquareDown(s.DistanceTo(p.t())*2.0 - 3.8)
}

func (p *PN5TorsoBack) v() *geometry.Point {
	return p.t().SquareLeft(1.3)
}

func (p *PN5TorsoBack) w() *geometry.Point {
	return p.t().SquareRight(1.3)
}

func (p *PN5TorsoBack) centreBack() geometry.Line {
	return &geometry.StraightLine{
		Start: p.a(),
		End:   p.k(),
	}
}

func (p *PN5TorsoBack) yokeSeamStitch() geometry.Line {
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

func (p *PN5TorsoBack) armholeStitch() geometry.Line {
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

func (p *PN5TorsoBack) sideSeamStitch() geometry.Line {
	return &geometry.ThreePointCurve{
		Start: p.i(),
		Middle: p.g(),
		End: p.c(),
		Rotation: &geometry.Angle{Rads: math.Pi / 2.0},
	}
}

func (p *PN5TorsoBack) hemLineStitch() geometry.Line {
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

func (p *PN5TorsoBack) dartStitch() geometry.Line {
	dart := &geometry.Polyline{}

	dart.AddLine(
		&geometry.StraightLine{
			Start: p.s(),
			End: p.v(),
		},
		&geometry.StraightLine{
			Start: p.v(),
			End: p.u(),
		},
		&geometry.StraightLine{
			Start: p.u(),
			End: p.w(),
		},
		&geometry.StraightLine{
			Start: p.w(),
			End: p.s(),
		},
	)

	return dart
}

func (p *PN5TorsoBack) CutLayer() *geometry.Block {
	layer := &geometry.Block{}

	armscyeCut := pieces.AddSeamAllowance(p.armholeStitch(), false)

	layer.AddLine(
		p.centreBack(),
		pieces.AddSeamAllowance(p.yokeSeamStitch(), false),
		armscyeCut,
		pieces.Notch(armscyeCut, 7.6),
		pieces.Notch(armscyeCut, armscyeCut.Length() - 7.6),
		pieces.AddSeamAllowance(p.sideSeamStitch(), true),
		pieces.AddSeamAllowance(p.hemLineStitch(), true),
	)

	return layer
}

func (p *PN5TorsoBack) StitchLayer() *geometry.Block {
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

func (p *PN5TorsoBack) NotationLayer() *geometry.Block {
	layer := &geometry.Block{}

	chestLine := &geometry.StraightLine{
		Start: p.b(),
		End: p.c(),
	}

	naturalWaistLine := &geometry.StraightLine{
		Start: p.d(),
		End: p.e(),
	}

	bellyButtonWaistLine := &geometry.StraightLine{
		Start: p.f(),
		End: p.g(),
	}

	hipLine := &geometry.StraightLine{
		Start: p.h(),
		End: p.i(),
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
	addAnchors(layer, anchors)

	return layer
}
