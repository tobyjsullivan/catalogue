package pieces

import (
	"github.com/tailored-style/pattern-generator/pieces"
	"github.com/tailored-style/pattern-generator/geometry"
	"fmt"
)

type pn8Cuff struct {
	cuffHeight float64
	wristCircumference float64
}

func NewPN8Cuff(wristCircumference float64) pieces.Piece {
	return &pn8Cuff{
		cuffHeight: 6.2,
		wristCircumference: wristCircumference,
	}
}

func (p *pn8Cuff) OnFold() bool {
	return false
}

func (p *pn8Cuff) a() *geometry.Point {
	return &geometry.Point{X: 0.0, Y: 0.0}
}

func (p *pn8Cuff) b() *geometry.Point {
	return p.a().SquareRight(p.wristCircumference + 7.6)
}

func (p *pn8Cuff) c() *geometry.Point {
	return p.a().SquareRight(2.5)
}

func (p *pn8Cuff) d() *geometry.Point {
	return p.a().SquareDown(p.cuffHeight)
}

func (p *pn8Cuff) e() *geometry.Point {
	return p.d().SquareToVerticalLine(p.b().X)
}

func (p *pn8Cuff) f() *geometry.Point {
	return p.b().SquareLeft(1.6)
}

func (p *pn8Cuff) g() *geometry.Point {
	return p.a().MirrorVertically(p.d().Y)
}

func (p *pn8Cuff) h() *geometry.Point {
	return p.g().SquareToVerticalLine(p.e().X)
}

func (p *pn8Cuff) topStitch() geometry.Line {
	return &geometry.StraightLine{
		Start: p.a(),
		End: p.b(),
	}
}

func (p *pn8Cuff) leftStitch() geometry.Line {
	return &geometry.StraightLine{
		Start: p.a(),
		End: p.g(),
	}
}

func (p *pn8Cuff) rightStitch() geometry.Line {
	return &geometry.StraightLine{
		Start: p.b(),
		End: p.h(),
	}
}

func (p *pn8Cuff) bottomStitch() geometry.Line {
	return &geometry.StraightLine{
		Start: p.g(),
		End: p.h(),
	}
}

func (p *pn8Cuff) middleFold() geometry.Line {
	return &geometry.StraightLine{
		Start: p.d(),
		End: p.e(),
	}
}

func (p *pn8Cuff) StitchLayer() *geometry.Block {
	layer := &geometry.Block{}

	layer.AddLine(
		p.topStitch(),
		p.leftStitch(),
		p.rightStitch(),
		p.bottomStitch(),
	)

	return layer
}

func (p *pn8Cuff) CutLayer() *geometry.Block {
	layer := &geometry.Block{}

	layer.AddLine(
		pieces.AddSeamAllowance(p.topStitch(), false),
		pieces.AddSeamAllowance(p.leftStitch(), true),
		pieces.AddSeamAllowance(p.rightStitch(), false),
		pieces.AddSeamAllowance(p.bottomStitch(), true),
	)

	return layer
}

func (p *pn8Cuff) NotationLayer() *geometry.Block {
	layer := &geometry.Block{}

	layer.AddLine(
		p.middleFold(),
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
	AddAnchors(layer, anchors)

	return layer
}

func (p *pn8Cuff) Details() *pieces.Details {
	return &pieces.Details{
		PieceNumber: 8,
		Description: "Cuff",
	}
}

func (p *pn8Cuff) String() string {
	return fmt.Sprintf("[PN: %d]", p.Details().PieceNumber)
}