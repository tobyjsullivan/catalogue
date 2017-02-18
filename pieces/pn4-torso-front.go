package pieces

import (
	"github.com/tailored-style/pattern-generator/geometry"
	"math"
	"github.com/tailored-style/pattern-generator/pieces"
)

type PN4TorsoFront struct {
	*pieces.Measurements
}

func (p *PN4TorsoFront) Details() *pieces.Details {
	return &pieces.Details{
		PieceNumber: "4",
		Description: "Torso Front",
	}
}

func (p *PN4TorsoFront) OnFold() bool {
	return false
}

func (p *PN4TorsoFront) a() *geometry.Point {
	return &geometry.Point{X: 0.0, Y: 0.0}
}

func (p *PN4TorsoFront) b() *geometry.Point {
	return p.a().SquareDown(p.ChestCircumference / 4.0)
}

func (p *PN4TorsoFront) c() *geometry.Point {
	return p.b().SquareLeft(p.ChestCircumference/4.0 + 1.4)
}

func (p *PN4TorsoFront) d() *geometry.Point {
	return p.a().SquareDown(p.Height/4.0 - 3.2)
}

func (p *PN4TorsoFront) e() *geometry.Point {
	return p.d().SquareLeft(p.WaistCircumference/4.0 + 3.2)
}

func (p *PN4TorsoFront) f() *geometry.Point {
	return p.a().SquareDown(p.Height/3.0 - 5.7)
}

func (p *PN4TorsoFront) g() *geometry.Point {
	return p.f().SquareLeft(p.HipCircumference/4.0 - 0.6)
}

func (p *PN4TorsoFront) h() *geometry.Point {
	return p.a().SquareDown(p.Height*(3.0/8.0) + 3.2)
}

func (p *PN4TorsoFront) i() *geometry.Point {
	return p.h().SquareLeft(p.HipCircumference/4.0 + 0.6)
}

func (p *PN4TorsoFront) j() *geometry.Point {
	return p.i().SquareUp(7.0)
}

func (p *PN4TorsoFront) k() *geometry.Point {
	return p.h().SquareDown(4.4)
}

func (p *PN4TorsoFront) l() *geometry.Point {
	return p.a().SquareDown(p.NeckCircumference/8.0 + 0.5)
}

func (p *PN4TorsoFront) m() *geometry.Point {
	return p.l().SquareLeft(p.NeckCircumference/8.0 + 2.2)
}

func (p *PN4TorsoFront) n() *geometry.Point {
	return p.m().SquareToHorizontalLine(p.a().Y)
}

func (p *PN4TorsoFront) o() *geometry.Point {
	return p.b().SquareLeft(p.ChestCircumference/6.0 + 4.1)
}

func (p *PN4TorsoFront) p() *geometry.Point {
	return p.o().SquareToHorizontalLine(p.a().Y)
}

func (p *PN4TorsoFront) q() *geometry.Point {
	return p.p().SquareDown(5.3)
}

func (p *PN4TorsoFront) r() *geometry.Point {
	n := p.n()
	q := p.q()
	return (&geometry.StraightLine{Start: n, End: q}).Resize(n.DistanceTo(q) + 2.3).End
}

func (p *PN4TorsoFront) s() *geometry.Point {
	o := p.o()
	return  o.SquareUp(o.DistanceTo(p.q()) / 2.0)
}

func (p *PN4TorsoFront) necklineStitch() geometry.Line {
	return &geometry.EllipseCurve{
		Start:         p.l(),
		End:           p.n(),
		StartingAngle: &geometry.Angle{Rads: math.Pi / 2.0},
		ArcAngle:      &geometry.Angle{Rads: math.Pi / 3.0},
	}
}

func (p *PN4TorsoFront) shoulderStitch() geometry.Line {
	return &geometry.StraightLine{
		Start: p.n(),
		End:   p.r(),
	}
}

func (p *PN4TorsoFront) armholeStitch() geometry.Line {
	top := &geometry.EllipseCurve{
		Start:         p.s(),
		End:           p.r(),
		StartingAngle: &geometry.Angle{Rads: 0.0},
		ArcAngle:      &geometry.Angle{Rads: math.Pi / 8.0},
	}
	bottom := &geometry.EllipseCurve{
		Start:         p.s(),
		End:           p.c(),
		StartingAngle: &geometry.Angle{Rads: math.Pi},
		ArcAngle:      &geometry.Angle{Rads: math.Pi * (2.0 / 5.0)},
	}

	armhole := &geometry.Polyline{}
	armhole.AddLine(
		&geometry.ReverseLine{InnerLine: top},
		bottom,
	)

	return armhole
}

func (p *PN4TorsoFront) sideSeamAStitch() geometry.Line {
	return &geometry.EllipseCurve{
		Start:         p.e(),
		End:           p.c(),
		StartingAngle: &geometry.Angle{Rads: 0.0},
		ArcAngle:      &geometry.Angle{Rads: math.Pi / 4.0},
	}
}

func (p *PN4TorsoFront) sideSeamBStitch() geometry.Line {
	return &geometry.ThreePointCurve{
		Start: p.j(),
		Middle: p.g(),
		End: p.e(),
		Rotation: &geometry.Angle{Rads: math.Pi / 2.0},
	}
}

func (p *PN4TorsoFront) sideSeamCStitch() geometry.Line {
	return &geometry.StraightLine{
		Start: p.j(),
		End:   p.i(),
	}
}


func (p *PN4TorsoFront) hemlineStitch() geometry.Line {
	return &geometry.SCurve{
		Start:         p.k(),
		End:           p.i(),
		StartingAngle: &geometry.Angle{Rads: math.Pi / 2.0},
		FinishAngle:   &geometry.Angle{Rads: math.Pi / 2.0},
		MaxAngle:      &geometry.Angle{Rads: math.Pi / 8.0},
	}
}

func (p *PN4TorsoFront) centreFront() geometry.Line {
	return &geometry.StraightLine{Start: p.l(), End: p.k()}
}

func (p *PN4TorsoFront) CutLayer() *geometry.Block {
	layer := &geometry.Block{}

	armholeCut := pieces.AddSeamAllowance(p.armholeStitch(), true)

	layer.AddLine(
		p.centreFront(),
		pieces.AddSeamAllowance(p.necklineStitch(), true),
		pieces.AddSeamAllowance(p.shoulderStitch(), true),
		armholeCut,
		pieces.Notch(armholeCut, 7.6),
		pieces.Notch(armholeCut, armholeCut.Length() - 7.6),
		pieces.Notch(armholeCut, armholeCut.Length() - 8.9),
		pieces.AddSeamAllowance(p.sideSeamAStitch(), false),
		pieces.AddSeamAllowance(p.sideSeamBStitch(), false),
		pieces.AddSeamAllowance(p.sideSeamCStitch(), true),
		pieces.AddSeamAllowance(p.hemlineStitch(), false),
	)

	return layer
}

func (p *PN4TorsoFront) StitchLayer() *geometry.Block {
	layer := &geometry.Block{}

	layer.AddLine(
		p.necklineStitch(),
		p.shoulderStitch(),
		p.armholeStitch(),
		p.sideSeamAStitch(),
		p.sideSeamBStitch(),
		p.sideSeamCStitch(),
		p.hemlineStitch(),
	)

	return layer
}

func (p *PN4TorsoFront) NotationLayer() *geometry.Block {
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
	addAnchors(layer, anchors)

	return layer
}
