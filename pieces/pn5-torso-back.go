package pieces

import (
	"fmt"
	"math"

	"github.com/tailored-style/pattern-generator/geometry"
	"github.com/tailored-style/pattern-generator/pieces"
	"github.com/tobyjsullivan/catalogue/anchors"
	"github.com/tobyjsullivan/catalogue/slopers"
)

type pn5TorsoBack struct {
	*slopers.TorsoMeasurements
}

func NewPN5TorsoBack(m *slopers.TorsoMeasurements) pieces.Piece {
	return &pn5TorsoBack{
		TorsoMeasurements: m,
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

func (p *pn5TorsoBack) torsoSloper() *slopers.Torso {
	return &slopers.Torso{
		TorsoMeasurements: p.TorsoMeasurements,
	}
}

func (p *pn5TorsoBack) a() *geometry.Point {
	return p.torsoSloper().P36()
}

func (p *pn5TorsoBack) b() *geometry.Point {
	return p.torsoSloper().P4()
}

func (p *pn5TorsoBack) c() *geometry.Point {
	return p.torsoSloper().P12().SquareRight(1.0)
}

func (p *pn5TorsoBack) f() *geometry.Point {
	return p.torsoSloper().P5()
}

func (p *pn5TorsoBack) g() *geometry.Point {
	return p.torsoSloper().P39().SquareRight(1.0)
}

func (p *pn5TorsoBack) h() *geometry.Point {
	return p.torsoSloper().P1()
}

func (p *pn5TorsoBack) i() *geometry.Point {
	return p.torsoSloper().P15().SquareRight(1.0)
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
	return p.torsoSloper().P8()
}

func (p *pn5TorsoBack) n() *geometry.Point {
	return p.torsoSloper().P37()
}

func (p *pn5TorsoBack) o() *geometry.Point {
	return p.n().SquareDown(1.1)
}

func (p *pn5TorsoBack) p() *geometry.Point {
	return p.torsoSloper().P21()
}

func (p *pn5TorsoBack) q() *geometry.Point {
	return p.n().SquareLeft(8.4)
}

func (p *pn5TorsoBack) r() *geometry.Point {
	return p.p().SquareLeft(0.5)
}

func (p *pn5TorsoBack) yoke() *pn6Yoke {
	return &pn6Yoke{
		TorsoMeasurements: p.TorsoMeasurements,
	}
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

	if yokeArmholeLength := p.yoke().armholeStitch().Length(); yokeArmholeLength < 7.6 {
		layer.AddLine(
			pieces.Notch(armholeStitch, 7.6 - yokeArmholeLength, false),
		)
	}

	return layer
}

func (p *pn5TorsoBack) Stitch() *geometry.Block {
	layer := &geometry.Block{}

	layer.AddLine(
		p.yokeSeamStitch(),
		p.armholeStitch(),
		p.sideSeamStitch(),
		p.hemLineStitch(),
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
		)

		// Draw all points
		a := make(map[string]*geometry.Point)
		a["A"] = p.a()
		a["B"] = p.b()
		a["C"] = p.c()
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
		anchors.AddAnchors(layer, a)
	}

	return layer
}

func (p *pn5TorsoBack) String() string {
	return fmt.Sprintf("[PN: %d]", p.Details().PieceNumber)
}
