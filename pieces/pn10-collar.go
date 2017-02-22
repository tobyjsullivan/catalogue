package pieces

import (
	"github.com/tailored-style/pattern-generator/pieces"
	"fmt"
	"github.com/tailored-style/pattern-generator/geometry"
	"math"
)

type pn10Collar struct {
	height float64
	neckCircumference float64
	chestCircumference float64
	waistCircumference float64
	hipCircumference float64
}

func NewPN10Collar(height float64, neck float64, chest float64, waist float64, hip float64) pieces.Piece {
	return &pn10Collar{
		height: height,
		neckCircumference: neck,
		chestCircumference: chest,
		waistCircumference: waist,
		hipCircumference: hip,
	}
}
func (p *pn10Collar) OnFold() bool {
	return true
}
func (p *pn10Collar) StitchLayer() *geometry.Block {
	layer := &geometry.Block{}

	layer.AddLine(
		p.bottomStitch(),
		p.rightStitch(),
		p.topStitch(),
	)

	return layer
}

func (p *pn10Collar) CutLayer() *geometry.Block {
	layer := &geometry.Block{}

	layer.AddLine(
		p.centreBack(),
	)

	return layer
}

func (p *pn10Collar) NotationLayer() *geometry.Block {
	layer := &geometry.Block{}

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
	addAnchors(layer, anchors)

	return layer
}

func (p *pn10Collar) Details() *pieces.Details {
	return &pieces.Details{
		PieceNumber: 10,
		Description: "Collar",
	}
}

func (p *pn10Collar) String() string {
	return fmt.Sprintf("[PN: %d]", p.Details().PieceNumber)
}

func (p *pn10Collar) collarBand() *pn9CollarBand {
	return &pn9CollarBand{
		height: p.height,
		neckCircumference: p.neckCircumference,
		chestCircumference: p.chestCircumference,
		waistCircumference: p.waistCircumference,
	}
}

func (p *pn10Collar) backNeckLength() float64 {
	return p.collarBand().backNeckLength()
}

func (p *pn10Collar) frontNeckLength() float64 {
	return p.collarBand().frontNeckLength()
}

func (p *pn10Collar) a() *geometry.Point {
	return p.collarBand().e().SquareUp(1.9)
}

func (p *pn10Collar) b() *geometry.Point {
	return p.a().SquareUp(p.collarBand().bandHeight() + 1.3)
}

func (p *pn10Collar) c() *geometry.Point {
	return p.collarBand().i()
}

func (p *pn10Collar) d() *geometry.Point {
	return p.c().SquareUp(p.collarBand().e().DistanceTo(p.b()))
}

func (p *pn10Collar) e() *geometry.Point {
	return p.d().SquareRight(2.5)
}

func (p *pn10Collar) f() *geometry.Point {
	return (&geometry.StraightLine{
		Start: p.c(),
		End: p.e(),
	}).Resize(p.c().DistanceTo(p.e()) + 0.6).End
}

func (p *pn10Collar) g() *geometry.Point {
	return p.collarBand().k().SquareToHorizontalLine(p.a().Y)
}

func (p *pn10Collar) h() *geometry.Point {
	return p.g().SquareToHorizontalLine(p.b().Y)
}

func (p *pn10Collar) i() *geometry.Point {
	return p.c().MidpointTo(p.g()).DrawAt(p.g().AngleRelativeTo(p.c()).Perpendicular(), 0.3)
}

func (p *pn10Collar) centreBack() geometry.Line {
	return &geometry.StraightLine{
		Start: p.b(),
		End: p.a(),
	}
}

func (p *pn10Collar) bottomStitch() geometry.Line {
	line := &geometry.Polyline{}

	line.AddLine(
		&geometry.StraightLine{
			Start: p.a(),
			End: p.g(),
		},
		&geometry.ParabolaCurve{
			Start: p.g(),
			End: p.c(),
			StartingAngle: &geometry.Angle{Rads: 0.0},
			ArcAngle: p.f().AngleRelativeTo(p.c()).Subtract(&geometry.Angle{Rads: math.Pi / 2.0}),
		},
	)

	return line
}

func (p *pn10Collar) rightStitch() geometry.Line {
	return &geometry.StraightLine{
		Start: p.f(),
		End: p.c(),
	}
}

func (p *pn10Collar) topStitch() geometry.Line {
	line := &geometry.Polyline{}

	line.AddLine(
		&geometry.StraightLine{
			Start: p.b(),
			End: p.h(),
		},
		&geometry.ParabolaCurve{
			Start: p.h(),
			End: p.f(),
			StartingAngle: &geometry.Angle{Rads: 0.0},
			ArcAngle: &geometry.Angle{Rads: math.Pi / 16.0},
		},
	)

	return line
}