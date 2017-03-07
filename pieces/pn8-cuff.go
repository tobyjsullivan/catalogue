package pieces

import (
	"fmt"

	"github.com/tailored-style/pattern-generator/geometry"
	"github.com/tailored-style/pattern-generator/pieces"
	"github.com/tailored-style/pattern-generator/symbols"
)

type pn8Cuff struct {
	cuffDepth          float64
	wristCircumference float64
}

func NewPN8Cuff(wristCircumference float64, cuffDepth float64) pieces.Piece {
	return &pn8Cuff{
		cuffDepth:          cuffDepth,
		wristCircumference: wristCircumference,
	}
}

func (p *pn8Cuff) CutCount() int {
	return 4
}

func (p *pn8Cuff) OnFold() bool {
	return false
}

func (p *pn8Cuff) Mirrored() bool {
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
	return p.a().SquareDown(p.cuffDepth)
}

func (p *pn8Cuff) e() *geometry.Point {
	return p.d().SquareToVerticalLine(p.b().X)
}

func (p *pn8Cuff) f() *geometry.Point {
	return p.b().SquareLeft(1.6)
}

func (p *pn8Cuff) g() *geometry.Point {
	return p.c().MidpointTo(p.d())
}

func (p *pn8Cuff) h() *geometry.Point {
	return p.f().SquareDown(p.b().DistanceTo(p.e()) / 2.0).SquareLeft(BUTTON_DIAMETER / 2.0)
}

func (p *pn8Cuff) topStitch() geometry.Line {
	return &geometry.StraightLine{
		Start: p.a(),
		End:   p.b(),
	}
}

func (p *pn8Cuff) leftStitch() geometry.Line {
	return &geometry.StraightLine{
		Start: p.d(),
		End:   p.a(),
	}
}

func (p *pn8Cuff) rightStitch() geometry.Line {
	return &geometry.StraightLine{
		Start: p.b(),
		End:   p.e(),
	}
}

func (p *pn8Cuff) bottomStitch() geometry.Line {
	return &geometry.StraightLine{
		Start: p.e(),
		End:   p.d(),
	}
}

func (p *pn8Cuff) Stitch() *geometry.Block {
	layer := &geometry.Block{}

	layer.AddLine(
		p.topStitch(),
		p.leftStitch(),
		p.rightStitch(),
		p.bottomStitch(),
	)

	return layer
}

func (p *pn8Cuff) InnerCut() *geometry.Block {
	layer := &geometry.Block{}

	seamAllowance := pieces.SeamAllowance(true,
		pieces.AddSeamAllowance(p.topStitch(), false),
		pieces.AddSeamAllowance(p.rightStitch(), false),
		pieces.AddSeamAllowance(p.bottomStitch(), false),
		pieces.AddSeamAllowance(p.leftStitch(), false),
	)

	layer.AddLine(
		seamAllowance,
	)

	return layer
}

func (p *pn8Cuff) button() *symbols.Button {
	return &symbols.Button{
		Centre: p.g(),
		Diameter: BUTTON_DIAMETER,
	}
}

func (p *pn8Cuff) buttonHole() *symbols.ButtonHole {
	return &symbols.ButtonHole{
		Centre: p.h(),
		Length: BUTTON_DIAMETER + 0.4,
	}
}

func (p *pn8Cuff) Ink() *geometry.Block {
	layer := &geometry.Block{}

	layer.AddBlock(
		p.button().Block(),
		p.buttonHole().Block(),
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
		AddAnchors(layer, anchors)
	}

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
