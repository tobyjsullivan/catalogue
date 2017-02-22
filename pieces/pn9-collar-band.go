package pieces

import (
	"github.com/tailored-style/pattern-generator/pieces"
	"fmt"
	"github.com/tailored-style/pattern-generator/geometry"
	"math"
)

type pn9CollarBand struct {
	bandHeight float64
	height float64
	neckCircumference float64
	chestCircumference float64
	waistCircumference float64
	hipCircumference float64
}

func NewPN9CollarBand(height float64, neck float64, chest float64, waist float64, hip float64) pieces.Piece {
	return &pn9CollarBand{
		bandHeight: 3.2,
		height: height,
		neckCircumference: neck,
		chestCircumference: chest,
		waistCircumference: waist,
		hipCircumference: hip,
	}
}

func (p *pn9CollarBand) OnFold() bool {
	return true
}
func (p *pn9CollarBand) StitchLayer() *geometry.Block {
	layer := &geometry.Block{}

	layer.AddLine(
		p.bottomStitch(),
		p.topStitch(),
	)

	return layer
}

func (p *pn9CollarBand) CutLayer() *geometry.Block {
	layer := &geometry.Block{}

	layer.AddLine(
		p.centreBack(),
		pieces.AddSeamAllowance(p.bottomStitch(), true),
		pieces.AddSeamAllowance(p.topStitch(), false),
	)

	return layer
}

func (p *pn9CollarBand) NotationLayer() *geometry.Block {
	layer := &geometry.Block{}

	layer.AddLine(
		p.centreFront(),
	)

	// Draw all points (DEBUG)
	anchors := make(map[string]*geometry.Point)
	anchors["A"] = p.a()
	anchors["B"] = p.b()
	//anchors["C"] = p.c()
	anchors["D"] = p.d()
	//anchors["E"] = p.e()
	anchors["H"] = p.h()
	anchors["I"] = p.i()
	//anchors["J"] = p.j()
	anchors["K"] = p.k()
	//anchors["L"] = p.l()
	anchors["M"] = p.m()
	anchors["N"] = p.n()
	//anchors["O"] = p.o()
	addAnchors(layer, anchors)

	return layer
}

func (p *pn9CollarBand) Details() *pieces.Details {
	return &pieces.Details{
		PieceNumber: 9,
		Description: "Collar Band",
	}
}

func (p *pn9CollarBand) String() string {
	return fmt.Sprintf("[PN: %d]", p.Details().PieceNumber)
}

func (p *pn9CollarBand) backNeckLength() float64 {
	return (&pn6Yoke{
		height: p.height,
		neckCircumference: p.neckCircumference,
		chestCircumference: p.chestCircumference,
		waistCircumference: p.waistCircumference,
		hipCircumference: p.hipCircumference,
	}).backNeckLine().Length()
}

func (p *pn9CollarBand) frontNeckLength() float64 {
	return (&pn4TorsoFront{
		height: p.height,
		neckCircumference: p.neckCircumference,
		chestCircumference: p.chestCircumference,
		waistCircumference: p.waistCircumference,
		hipCircumference: p.hipCircumference,
	}).frontNeckLine().Length()
}

func (p *pn9CollarBand) a() *geometry.Point {
	return &geometry.Point{X: 0.0, Y: 0.0}
}

func (p *pn9CollarBand) b() *geometry.Point {
	return p.a().SquareRight(p.backNeckLength())
}

func (p *pn9CollarBand) c() *geometry.Point {
	return p.b().SquareRight(p.frontNeckLength())
}

func (p *pn9CollarBand) d() *geometry.Point {
	return p.a().SquareUp(p.bandHeight)
}

func (p *pn9CollarBand) e() *geometry.Point {
	return p.d().SquareUp(1.0)
}

func (p *pn9CollarBand) l() *geometry.Point {
	return p.c().SquareUp(p.a().DistanceTo(p.e()))
}

func (p *pn9CollarBand) h() *geometry.Point {
	return p.c().SquareUp(1.3)
}

func (p *pn9CollarBand) frontBottomLine() geometry.Line {
	return &geometry.ParabolaCurve{
		Start: p.b(),
		End: p.h(),
		StartingAngle: &geometry.Angle{Rads: 0.0},
		ArcAngle: &geometry.Angle{Rads: math.Pi / 10.0},
	}
}

func (p *pn9CollarBand) j() *geometry.Point {
	bottomStitch := p.frontBottomLine()
	angle := bottomStitch.AngleAt(bottomStitch.Length() - 0.01).Perpendicular()
	length := p.h().DistanceTo(p.l()) / angle.Sin()
	return p.h().DrawAt(angle, length)
}

func (p *pn9CollarBand) k() *geometry.Point {
	return p.j().SquareLeft(0.3).SquareDown(0.3)
}

func (p *pn9CollarBand) i() *geometry.Point {
	angle := p.h().AngleRelativeTo(p.k())

	return p.l().DrawAt(angle, p.k().DistanceTo(p.h()))
}

func (p *pn9CollarBand) m() *geometry.Point {
	return p.d().SquareRight(p.a().DistanceTo(p.b()))
}

func (p *pn9CollarBand) n() *geometry.Point {
	return p.l().DrawAt(&geometry.Angle{Rads: -math.Pi * 3.0/4.0}, 0.3)
}

func (p *pn9CollarBand) o() *geometry.Point {
	return (&geometry.StraightLine{
		Start: p.l(),
		End: p.i(),
	}).Resize(1.3).End
}

func (p *pn9CollarBand) bottomStitch() geometry.Line {
	line := &geometry.Polyline{}

	end := &geometry.StraightLine{
		Start: p.h(),
		End: p.i(),
	}

	line.AddLine(
		&geometry.StraightLine{
			Start: p.a(),
			End: p.b(),
		},
		&geometry.ParabolaCurve{
			Start: p.b(),
			End: p.h(),
			StartingAngle: &geometry.Angle{Rads: 0.0},
			ArcAngle: end.AngleAt(0.0),
		},
		end,
	)

	return line
}

func (p *pn9CollarBand) topStitch() geometry.Line {
	line := &geometry.Polyline{}

	end := &geometry.StraightLine{
		Start: p.o(),
		End: p.i(),
	}

	angleAtK := (&geometry.StraightLine{Start: p.k(), End: p.l()}).AngleAt(0.0)

	line.AddLine(
		&geometry.StraightLine{
			Start: p.d(),
			End: p.m(),
		},
		&geometry.ParabolaCurve{
			Start: p.m(),
			End: p.k(),
			StartingAngle: &geometry.Angle{Rads: 0.0},
			ArcAngle: angleAtK,
		},
		&geometry.EllipseCurve{
			Start: p.k(),
			End: p.o(),
			StartingAngle: angleAtK.Subtract(&geometry.Angle{Rads: math.Pi / 2.0}),
			ArcAngle: end.AngleAt(0.0).Subtract(angleAtK),
		},
		end,
	)

	return line
}

func (p *pn9CollarBand) centreBack() geometry.Line {
	return &geometry.StraightLine{
		Start: p.a(),
		End: p.d(),
	}
}

func (p *pn9CollarBand) centreFront() geometry.Line {
	return &geometry.StraightLine{
		Start: p.h(),
		End: p.k(),
	}
}
