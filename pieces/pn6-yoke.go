package pieces

import (
	"fmt"
	"math"

	"github.com/tailored-style/pattern-generator/geometry"
	"github.com/tailored-style/pattern-generator/pieces"
	"github.com/tobyjsullivan/catalogue/anchors"
	"github.com/tobyjsullivan/catalogue/slopers"
)

type pn6Yoke struct {
	*slopers.TorsoMeasurements
}

func NewPN6Yoke(m *slopers.TorsoMeasurements) pieces.Piece {
	return &pn6Yoke{
		TorsoMeasurements: m,
	}
}

func (p *pn6Yoke) Details() *pieces.Details {
	return &pieces.Details{
		PieceNumber: 6,
		Description: "Yoke",
	}
}

func (p *pn6Yoke) OnFold() bool {
	return true
}

func (p *pn6Yoke) CutCount() int {
	return 2
}

func (p *pn6Yoke) Mirrored() bool {
	return false
}

func (p *pn6Yoke) torso() *slopers.Torso {
	return &slopers.Torso{
		TorsoMeasurements: p.TorsoMeasurements,
	}
}

func (p *pn6Yoke) a() *geometry.Point {
	return p.torso().P0()
}

func (p *pn6Yoke) b() *geometry.Point {
	return p.torso().P36()
}

func (p *pn6Yoke) c() *geometry.Point {
	return p.torso().P37()
}

func (p *pn6Yoke) e() *geometry.Point {
	return p.torso().P33()
}

func (p *pn6Yoke) f() *geometry.Point {
	return p.torso().P17()
}

func (p *pn6Yoke) g() *geometry.Point {
	return p.torso().P19()
}

func (p *pn6Yoke) h() *geometry.Point {
	return p.torso().P34()
}

func (p *pn6Yoke) shoulderSeamLength() float64 {
	return (&pn4TorsoFront{
		TorsoMeasurements: p.TorsoMeasurements,
	}).shoulderStitch().Length()
}

func (p *pn6Yoke) backNeckLine() geometry.Line {
	return &geometry.EllipseCurve{
		Start:         p.a(),
		End:           p.f(),
		StartingAngle: &geometry.Angle{Rads: math.Pi * (3.0 / 2.0)},
		ArcAngle:      &geometry.Angle{Rads: math.Pi * (7.0 / 16.0)},
	}
}

func (p *pn6Yoke) necklineStitch() geometry.Line {
	return p.backNeckLine()
}

func (p *pn6Yoke) frontStitch() geometry.Line {
	return &geometry.StraightLine{
		Start: p.f(),
		End:   p.g(),
	}
}

func (p *pn6Yoke) armholeStitch() geometry.Line {
	return &geometry.QuadraticBezierCurve{
		P0: p.c(),
		P1: p.h(),
		P2: p.g(),
	}
}

func (p *pn6Yoke) backStitch() geometry.Line {
	return &geometry.StraightLine{
		Start: p.b(),
		End:   p.c(),
	}
}

func (p *pn6Yoke) centreBack() geometry.Line {
	return &geometry.StraightLine{
		Start: p.a(),
		End:   p.b(),
	}
}

func (p *pn6Yoke) Stitch() *geometry.Block {

	layer := &geometry.Block{}
	layer.AddLine(
		p.necklineStitch(),
		p.frontStitch(),
		p.armholeStitch(),
		p.backStitch(),
	)

	return layer
}

func (p *pn6Yoke) OuterCut() *geometry.Polyline {
	return &geometry.Polyline{}
}

func (p *pn6Yoke) InnerCut() *geometry.Block {
	layer := &geometry.Block{}

	armholeStitch := &geometry.ReverseLine{InnerLine: p.armholeStitch()}
	seamAllowance := pieces.SeamAllowance(false,
		pieces.AddSeamAllowance(p.necklineStitch(), false),
		pieces.AddSeamAllowance(p.frontStitch(), false),
		pieces.AddSeamAllowance(armholeStitch, false),
		pieces.AddSeamAllowance(p.backStitch(), true),
	)

	layer.AddLine(
		seamAllowance,
	)
	if armholeStitch.Length() > 7.6 {
		layer.AddLine(
			pieces.Notch(armholeStitch, 7.6, false),
		)
	}

	return layer
}

func (p *pn6Yoke) Ink() *geometry.Block {
	return &geometry.Block{}
}

func (p *pn6Yoke) Reference() *geometry.Block {
	layer := &geometry.Block{}

	layer.AddLine(
		p.centreBack(),
	)

	// Draw all points (DEBUG)
	if DEBUG {
		a := make(map[string]*geometry.Point)
		a["A"] = p.a()
		a["B"] = p.b()
		a["C"] = p.c()
		a["E"] = p.e()
		a["F"] = p.f()
		a["G"] = p.g()
		a["H"] = p.h()
		anchors.AddAnchors(layer, a)
	}

	return layer
}

func (p *pn6Yoke) String() string {
	return fmt.Sprintf("[PN: %d]", p.Details().PieceNumber)
}
