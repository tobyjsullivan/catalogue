package pieces

import (
	"github.com/tailored-style/pattern-generator/geometry"
	"math"
	"github.com/tailored-style/pattern-generator/pieces"
	"fmt"
)

type pn6Yoke struct {
	height float64
	neckCircumference float64
	chestCircumference float64
	waistCircumference float64
	hipCircumference float64
}

func NewPN6Yoke(height float64, neck float64, chest float64, waist float64, hip float64) pieces.Piece {
	return &pn6Yoke{
		height:height,
		neckCircumference:neck,
		chestCircumference:chest,
		waistCircumference:waist,
		hipCircumference:hip,
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

func (p *pn6Yoke) a() *geometry.Point {
	return &geometry.Point{X: 0.0, Y: 0.0}
}

func (p *pn6Yoke) b() *geometry.Point {
	return p.a().SquareDown(9.5)
}

func (p *pn6Yoke) c() *geometry.Point {
	return p.b().SquareRight(p.chestCircumference/6.0 + 6.2)
}

func (p *pn6Yoke) d() *geometry.Point {
	return p.c().SquareToHorizontalLine(p.a().Y)
}

func (p *pn6Yoke) e() *geometry.Point {
	return p.a().SquareRight(p.neckCircumference/8.0 + 3.7)
}

func (p *pn6Yoke) f() *geometry.Point {
	e := p.e()
	return e.SquareUp(p.a().DistanceTo(e)/2.0 + 0.3)
}

func (p *pn6Yoke) g() *geometry.Point {
	return  (&geometry.StraightLine{Start: p.f(), End: p.d()}).Resize(p.shoulderSeamLength()).End
}

func (p *pn6Yoke) shoulderSeamLength() float64 {
	return (&pn4TorsoFront{
		height: p.height,
		neckCircumference: p.neckCircumference,
		chestCircumference: p.chestCircumference,
		waistCircumference: p.waistCircumference,
		hipCircumference: p.hipCircumference,
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
	return &geometry.EllipseCurve{
		Start: p.c(),
		End:   p.g(),
		StartingAngle: &geometry.Angle{Rads: math.Pi},
		ArcAngle: p.g().AngleRelativeTo(p.d()).Opposite().Subtract(&geometry.Angle{Rads: math.Pi}),
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

func (p *pn6Yoke) StitchLayer() *geometry.Block {

	layer := &geometry.Block{}
	layer.AddLine(
		p.necklineStitch(),
		p.frontStitch(),
		p.armholeStitch(),
		p.backStitch(),
	)

	return layer
}

func (p *pn6Yoke) CutLayer() *geometry.Block {
	layer := &geometry.Block{}
	layer.AddLine(
		p.centreBack(),
		pieces.AddSeamAllowance(p.necklineStitch(), false),
		pieces.AddSeamAllowance(p.frontStitch(), false),
		pieces.AddSeamAllowance(p.armholeStitch(), true),
		pieces.AddSeamAllowance(p.backStitch(), true),
	)

	return layer
}

func (p *pn6Yoke) NotationLayer() *geometry.Block {
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
	AddAnchors(layer, anchors)

	return layer
}

func (p *pn6Yoke) String() string {
	return fmt.Sprintf("[PN: %d]", p.Details().PieceNumber)
}
