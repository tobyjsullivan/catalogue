package pieces

import (
	"fmt"
	"math"

	"github.com/tailored-style/pattern-generator/geometry"
	"github.com/tailored-style/pattern-generator/pieces"
	"github.com/tobyjsullivan/catalogue/slopers"
	"github.com/tobyjsullivan/catalogue/anchors"
)

const PN4_BUTTON_WIDTH = 1.5

type pn4TorsoFront struct {
	*slopers.TorsoMeasurements
}

func (p *pn4TorsoFront) Details() *pieces.Details {
	return &pieces.Details{
		PieceNumber: 4,
		Description: "Torso Front",
	}
}

func (p *pn4TorsoFront) CutCount() int {
	return 1
}

func (p *pn4TorsoFront) OnFold() bool {
	return false
}

func (p *pn4TorsoFront) Mirrored() bool {
	return true
}

func (p *pn4TorsoFront) torsoSloper() *slopers.Torso {
	return &slopers.Torso{
		TorsoMeasurements: p.TorsoMeasurements,
	}
}

func (p *pn4TorsoFront) a() *geometry.Point {
	return p.torsoSloper().P3()
}

func (p *pn4TorsoFront) b() *geometry.Point {
	return p.torsoSloper().P6()
}

func (p *pn4TorsoFront) c() *geometry.Point {
	return p.torsoSloper().P12().SquareLeft(1.0)
}

func (p *pn4TorsoFront) f() *geometry.Point {
	return p.torsoSloper().P7()
}

func (p *pn4TorsoFront) g() *geometry.Point {
	return p.torsoSloper().P39().SquareLeft(1.0)
}

func (p *pn4TorsoFront) h() *geometry.Point {
	return p.torsoSloper().P2()
}

func (p *pn4TorsoFront) i() *geometry.Point {
	return p.torsoSloper().P15().SquareLeft(1.0)
}

func (p *pn4TorsoFront) k() *geometry.Point {
	return p.h().SquareDown(4.4)
}

func (p *pn4TorsoFront) l() *geometry.Point {
	return p.torsoSloper().P24()
}

func (p *pn4TorsoFront) m() *geometry.Point {
	return p.torsoSloper().P25()
}

func (p *pn4TorsoFront) n() *geometry.Point {
	return p.torsoSloper().P26()
}

func (p *pn4TorsoFront) o() *geometry.Point {
	return p.torsoSloper().P10()
}

func (p *pn4TorsoFront) r() *geometry.Point {
	return p.torsoSloper().P30()
}

func (p *pn4TorsoFront) s() *geometry.Point {
	return p.torsoSloper().P31()
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

func (p *pn4TorsoFront) buttons() []*geometry.Point {
	return []*geometry.Point{
		p.l().SquareDown(4.3),
		p.l().SquareDown(14.3),
		p.l().SquareDown(24.3),
		p.l().SquareDown(34.3),
		p.l().SquareDown(44.3),
		p.l().SquareDown(54.3),
	}
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

func (p *pn4TorsoFront) OuterCut() *geometry.Polyline {
	seamAllowance := pieces.SeamAllowance(false,
		&geometry.ReverseLine{InnerLine: pieces.AddSeamAllowance(p.buttonStandBottom(), true)},
		pieces.AddSeamAllowance(p.hemlineStitch(), false),
		&geometry.ReverseLine{InnerLine: pieces.AddSeamAllowance(p.sideSeamStitch(), true)},
		&geometry.ReverseLine{InnerLine: pieces.AddSeamAllowance(p.armholeStitch(), true)},
		&geometry.ReverseLine{InnerLine: pieces.AddSeamAllowance(p.shoulderStitch(), true)},
		&geometry.ReverseLine{InnerLine: pieces.AddSeamAllowance(p.necklineStitch(), true)},
		pieces.AddSeamAllowance(p.buttonStandTopA(), false),
		&geometry.ReverseLine{InnerLine: pieces.AddSeamAllowance(p.buttonStandTopB(), true)},
	)

	front := p.buttonStandFront()

	poly := &geometry.Polyline{}
	poly.AddLine(
		seamAllowance,
		front,
		geometry.Connect(front, seamAllowance),
		geometry.Connect(seamAllowance, front),
	)


	return poly
}

func (p *pn4TorsoFront) InnerCut() *geometry.Block {
	layer := &geometry.Block{}

	armholeStitch := p.armholeStitch()
	layer.AddLine(
		pieces.Notch(armholeStitch, 7.6, true),
		pieces.Notch(armholeStitch, armholeStitch.Length() - 7.6, true),
	)

	return layer
}

func (p *pn4TorsoFront) Stitch() *geometry.Block {
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

func (p *pn4TorsoFront) Ink() *geometry.Block {
	return &geometry.Block{}
}

func (p *pn4TorsoFront) Reference() *geometry.Block {
	layer := &geometry.Block{}

	if DEBUG {
		chestLine := &geometry.StraightLine{
			Start: p.b(),
			End:   p.c(),
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
			bellyButtonWaistLine,
			hipLine,
			p.centreFront(),
			p.buttonStandFoldA(),
			p.buttonStandFoldB(),
		)

		// Draw all points (DEBUG)
		a := make(map[string]*geometry.Point)
		a["A"] = p.a()
		a["B"] = p.b()
		a["C"] = p.c()
		a["F"] = p.f()
		a["G"] = p.g()
		a["H"] = p.h()
		a["I"] = p.i()
		a["K"] = p.k()
		a["L"] = p.l()
		a["M"] = p.m()
		a["N"] = p.n()
		a["O"] = p.o()
		a["R"] = p.r()
		a["S"] = p.s()
		a["T"] = p.t()
		a["U"] = p.u()
		a["V"] = p.v()
		a["W"] = p.w()
		a["X"] = p.x()
		a["Y"] = p.y()
		a["Z"] = p.z()
		anchors.AddAnchors(layer, a)
	}

	return layer
}

func (p *pn4TorsoFront) String() string {
	return fmt.Sprintf("[PN: %d]", p.Details().PieceNumber)
}
