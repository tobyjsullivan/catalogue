package pieces

import (
	"fmt"
	"math"

	"github.com/tailored-style/pattern-generator/geometry"
	"github.com/tailored-style/pattern-generator/pieces"
)

const PN4_BUTTON_WIDTH = 1.5

type pn4TorsoFront struct {
	height             float64
	neckCircumference  float64
	chestCircumference float64
	waistCircumference float64
	hipCircumference   float64
}

func NewPN4TorsoFront(height float64, neck float64, chest float64, waist float64, hip float64) pieces.Piece {
	return &pn4TorsoFront{
		height:             height,
		neckCircumference:  neck,
		chestCircumference: chest,
		waistCircumference: waist,
		hipCircumference:   hip,
	}
}

func (p *pn4TorsoFront) Details() *pieces.Details {
	return &pieces.Details{
		PieceNumber: 4,
		Description: "Torso Front",
	}
}

func (p *pn4TorsoFront) OnFold() bool {
	return false
}

func (p *pn4TorsoFront) a() *geometry.Point {
	return &geometry.Point{X: 0.0, Y: 0.0}
}

func (p *pn4TorsoFront) b() *geometry.Point {
	return p.a().SquareDown(p.chestCircumference / 4.0)
}

func (p *pn4TorsoFront) c() *geometry.Point {
	return p.b().SquareLeft(p.chestCircumference/4.0 + 1.6)
}

func (p *pn4TorsoFront) d() *geometry.Point {
	return p.a().SquareDown(p.height/4.0 - 3.2)
}

func (p *pn4TorsoFront) e() *geometry.Point {
	return p.d().SquareLeft(p.waistCircumference/4.0 + 3.2)
}

func (p *pn4TorsoFront) f() *geometry.Point {
	return p.a().SquareDown(p.height/3.0 - 5.7)
}

func (p *pn4TorsoFront) bellyWaistGirth() float64 {
	return p.waistCircumference/4.0 + 4.3
}

func (p *pn4TorsoFront) g() *geometry.Point {
	return p.f().SquareLeft(p.bellyWaistGirth())
}

func (p *pn4TorsoFront) h() *geometry.Point {
	return p.a().SquareDown(p.height*(3.0/8.0) + 3.2)
}

func (p *pn4TorsoFront) i() *geometry.Point {
	hipLine := p.hipCircumference/4.0 + 0.6
	bellyWaist := p.bellyWaistGirth()
	if hipLine < bellyWaist {
		hipLine = bellyWaist
	}

	return p.h().SquareLeft(hipLine)
}

func (p *pn4TorsoFront) j() *geometry.Point {
	return p.i().SquareUp(7.0)
}

func (p *pn4TorsoFront) k() *geometry.Point {
	return p.h().SquareDown(4.4)
}

func (p *pn4TorsoFront) l() *geometry.Point {
	return p.a().SquareDown(p.neckCircumference/8.0 + 0.5)
}

func (p *pn4TorsoFront) m() *geometry.Point {
	return p.l().SquareLeft(p.neckCircumference/8.0 + 2.2)
}

func (p *pn4TorsoFront) n() *geometry.Point {
	return p.m().SquareToHorizontalLine(p.a().Y)
}

func (p *pn4TorsoFront) o() *geometry.Point {
	return p.b().SquareLeft(p.chestCircumference/6.0 + 4.1)
}

func (p *pn4TorsoFront) p() *geometry.Point {
	return p.o().SquareToHorizontalLine(p.a().Y)
}

func (p *pn4TorsoFront) q() *geometry.Point {
	return p.p().SquareDown(5.3)
}

func (p *pn4TorsoFront) r() *geometry.Point {
	n := p.n()
	q := p.q()
	return (&geometry.StraightLine{Start: n, End: q}).Resize(n.DistanceTo(q) + 2.3).End
}

func (p *pn4TorsoFront) s() *geometry.Point {
	o := p.o()
	return o.SquareUp(o.DistanceTo(p.q()) / 2.0)
}

func (p *pn4TorsoFront) t() *geometry.Point {
	return p.l().SquareRight((PN4_BUTTON_WIDTH / 2.0) + 1.3)
}

func (p *pn4TorsoFront) u() *geometry.Point {
	l := p.l()
	return l.SquareLeft(l.DistanceTo(p.t())).SquareUpToLine(p.necklineStitch())
}

func (p *pn4TorsoFront) v() *geometry.Point {
	return p.u().MirrorHorizontally(p.t().X)
}

func (p *pn4TorsoFront) w() *geometry.Point {
	return p.t().MirrorHorizontally(p.v().X)
}

func (p *pn4TorsoFront) x() *geometry.Point {
	return p.t().SquareToHorizontalLine(p.k().Y)
}

func (p *pn4TorsoFront) y() *geometry.Point {
	return p.w().SquareToHorizontalLine(p.k().Y)
}

func (p *pn4TorsoFront) z() *geometry.Point {
	return p.v().SquareToHorizontalLine(p.k().Y)
}

func (p *pn4TorsoFront) frontNeckLine() geometry.Line {
	return &geometry.EllipseCurve{
		Start:         p.l(),
		End:           p.n(),
		StartingAngle: &geometry.Angle{Rads: math.Pi / 2.0},
		ArcAngle:      &geometry.Angle{Rads: math.Pi / 3.0},
	}
}

func (p *pn4TorsoFront) necklineStitch() geometry.Line {
	neckline := &geometry.Polyline{}

	neckline.AddLine(
		&geometry.StraightLine{
			Start: p.t(),
			End:   p.l(),
		},
		p.frontNeckLine(),
	)

	return neckline
}

func (p *pn4TorsoFront) buttonStandTopA() geometry.Line {
	return geometry.SliceLineVertically(geometry.MirrorLineHorizontally(p.necklineStitch(), p.t().X), p.v().X)
}

func (p *pn4TorsoFront) buttonStandTopB() geometry.Line {
	return geometry.MirrorLineHorizontally(p.buttonStandTopA(), p.v().X)
}

func (p *pn4TorsoFront) buttonStandFoldA() geometry.Line {
	return &geometry.StraightLine{
		Start: p.t(),
		End:   p.x(),
	}
}

func (p *pn4TorsoFront) buttonStandFoldB() geometry.Line {
	return &geometry.StraightLine{
		Start: p.v(),
		End:   p.z(),
	}
}

func (p *pn4TorsoFront) buttonStandBottom() geometry.Line {
	return &geometry.StraightLine{
		Start: p.x(),
		End:   p.y(),
	}
}

func (p *pn4TorsoFront) buttonStandFront() geometry.Line {
	return &geometry.StraightLine{
		Start: p.w(),
		End:   p.y(),
	}
}

func (p *pn4TorsoFront) shoulderStitch() geometry.Line {
	return &geometry.StraightLine{
		Start: p.n(),
		End:   p.r(),
	}
}

func (p *pn4TorsoFront) armholeStitch() geometry.Line {
	top := &geometry.EllipseCurve{
		Start:         p.s(),
		End:           p.r(),
		StartingAngle: &geometry.Angle{Rads: 0.0},
		ArcAngle:      &geometry.Angle{Rads: math.Pi / 8.0},
	}

	bottom := &geometry.QuadraticBezierCurve{
		P0: p.s(),
		P1: p.o(),
		P2: p.c(),
	}

	armhole := &geometry.Polyline{}
	armhole.AddLine(
		&geometry.ReverseLine{InnerLine: top},
		bottom,
	)

	return armhole
}

func (p *pn4TorsoFront) sideSeamStitch() geometry.Line {
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

func (p *pn4TorsoFront) hemlineStitch() geometry.Line {
	line := &geometry.Polyline{}

	line.AddLine(
		&geometry.StraightLine{
			Start: p.x(),
			End:   p.k(),
		},
		&geometry.SCurve{
			Start:         p.k(),
			End:           p.i(),
			StartingAngle: &geometry.Angle{Rads: math.Pi / 2.0},
			FinishAngle:   &geometry.Angle{Rads: math.Pi / 2.0},
			MaxAngle:      &geometry.Angle{Rads: math.Pi / 8.0},
		},
	)

	return line
}

func (p *pn4TorsoFront) centreFront() geometry.Line {
	return &geometry.StraightLine{Start: p.l(), End: p.k()}
}

func (p *pn4TorsoFront) CutLayer() *geometry.Block {
	layer := &geometry.Block{}

	armholeCut := pieces.AddSeamAllowance(p.armholeStitch(), true)

	seamAllowance := pieces.SeamAllowance(
		&geometry.ReverseLine{InnerLine: pieces.AddSeamAllowance(p.buttonStandBottom(), true)},
		pieces.AddSeamAllowance(p.hemlineStitch(), false),
		&geometry.ReverseLine{InnerLine: pieces.AddSeamAllowance(p.sideSeamStitch(), true)},
		&geometry.ReverseLine{InnerLine: armholeCut},
		&geometry.ReverseLine{InnerLine: pieces.AddSeamAllowance(p.shoulderStitch(), true)},
		&geometry.ReverseLine{InnerLine: pieces.AddSeamAllowance(p.necklineStitch(), true)},
		pieces.AddSeamAllowance(p.buttonStandTopA(), false),
		&geometry.ReverseLine{InnerLine: pieces.AddSeamAllowance(p.buttonStandTopB(), true)},
	)

	front := p.buttonStandFront()

	layer.AddLine(
		front,
		seamAllowance,
		geometry.Connect(front, seamAllowance),
		geometry.Connect(seamAllowance, front),
		pieces.Notch(armholeCut, 7.6),
		pieces.Notch(armholeCut, armholeCut.Length()-7.6),
		pieces.Notch(armholeCut, armholeCut.Length()-8.9),
	)

	return layer
}

func (p *pn4TorsoFront) StitchLayer() *geometry.Block {
	layer := &geometry.Block{}

	layer.AddLine(
		p.necklineStitch(),
		p.buttonStandTopA(),
		p.buttonStandTopB(),
		p.buttonStandBottom(),
		p.shoulderStitch(),
		p.armholeStitch(),
		p.sideSeamStitch(),
		p.hemlineStitch(),
	)

	return layer
}

func (p *pn4TorsoFront) NotationLayer() *geometry.Block {
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
		p.centreFront(),
		p.buttonStandFoldA(),
		p.buttonStandFoldB(),
	)

	// Draw all points (DEBUG)
	if DEBUG {
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
		anchors["X"] = p.x()
		anchors["Y"] = p.y()
		anchors["Z"] = p.z()
		AddAnchors(layer, anchors)
	}

	return layer
}

func (p *pn4TorsoFront) String() string {
	return fmt.Sprintf("[PN: %d]", p.Details().PieceNumber)
}
