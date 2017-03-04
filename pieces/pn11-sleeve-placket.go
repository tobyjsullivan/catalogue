package pieces

import (
	"github.com/tailored-style/pattern-generator/pieces"
	"github.com/tailored-style/pattern-generator/geometry"
	"math"
)

type pn11SleevePlacket struct {
	placketOpeningLength float64
}

const PN11_PLACKET_WIDTH  = 2.5

func NewPN11SleevePlacket(placketOpeningLength float64) pieces.Piece {
	return &pn11SleevePlacket{
		placketOpeningLength: placketOpeningLength,
	}
}

func (p *pn11SleevePlacket) Details() *pieces.Details {
	return &pieces.Details{
		PieceNumber: 11,
		Description: "Sleeve Placket",
	}
}

func (p *pn11SleevePlacket) OnFold() bool {
	return false
}

func (p *pn11SleevePlacket) Mirrored() bool {
	return true
}

func (p *pn11SleevePlacket) StitchLayer() *geometry.Block {
	layer := &geometry.Block{}

	layer.AddLine(
		p.placketLeftStitch(),
		p.placketTopLeftStitch(),
		p.placketTopRightStitch(),
		p.placketCornerStitch(),
		p.placketRightStitch(),
		p.placketBottomStitch(),
		p.standLeftStitch(),
		p.standTopStitch(),
		p.standRightStitch(),
		p.standBottomStitch(),
	)

	return layer
}

func (p *pn11SleevePlacket) CutLayer() *geometry.Block {
	layer := &geometry.Block{}

	cornerCut := &geometry.Polyline{}
	crevice := p.l().DrawAt(&geometry.Angle{Rads: math.Pi / 4.0}, math.Sqrt(math.Pow(pieces.SEAM_ALLOWANCE, 2.0) + math.Pow(pieces.SEAM_ALLOWANCE, 2.0)))
	cornerCut.AddLine(
		&geometry.StraightLine{
			Start: p.f().DrawAt(p.l().AngleRelativeTo(p.f()).Perpendicular(), pieces.SEAM_ALLOWANCE),
			End: crevice,
		},
		&geometry.StraightLine{
			Start: crevice,
			End: p.d().DrawAt(p.l().AngleRelativeTo(p.d()).Perpendicular().Opposite(), pieces.SEAM_ALLOWANCE),
		},
	)

	placketSeamAllowance := pieces.SeamAllowance(true,
		cornerCut,
		pieces.AddSeamAllowance(p.placketRightStitch(), false),
		pieces.AddSeamAllowance(p.placketBottomStitch(), false),
		pieces.AddSeamAllowance(p.placketLeftStitch(), false),
		pieces.AddSeamAllowance(p.placketTopLeftStitch(), false),
		pieces.AddSeamAllowance(p.placketTopRightStitch(), false),
	)

	standSeamAllowance := pieces.SeamAllowance(true,
		pieces.AddSeamAllowance(p.standLeftStitch(), false),
		pieces.AddSeamAllowance(p.standTopStitch(), false),
		pieces.AddSeamAllowance(p.standRightStitch(), false),
		pieces.AddSeamAllowance(p.standBottomStitch(), false),
	)

	layer.AddLine(
		placketSeamAllowance,
		standSeamAllowance,
	)

	return layer
}

func (p *pn11SleevePlacket) NotationLayer() *geometry.Block {
	layer := &geometry.Block{}

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
		AddAnchors(layer, anchors)
	}

	return layer
}

func (p *pn11SleevePlacket) a() *geometry.Point {
	return &geometry.Point{X: 0.0, Y: 0.0}
}

func (p *pn11SleevePlacket) b() *geometry.Point {
	return p.a().SquareUp(p.placketOpeningLength)
}

func (p *pn11SleevePlacket) c() *geometry.Point {
	return p.a().SquareRight(PN11_PLACKET_WIDTH * 2.0)
}

func (p *pn11SleevePlacket) d() *geometry.Point {
	return p.b().SquareToVerticalLine(p.c().X)
}

func (p *pn11SleevePlacket) e() *geometry.Point {
	return p.b().SquareUp(2.5)
}

func (p *pn11SleevePlacket) f() *geometry.Point {
	return p.e().SquareRight(PN11_PLACKET_WIDTH)
}

func (p *pn11SleevePlacket) g() *geometry.Point {
	return p.e().MidpointTo(p.f()).SquareUp(1.0)
}

func (p *pn11SleevePlacket) h() *geometry.Point {
	return p.c().SquareRight(2.5)
}

func (p *pn11SleevePlacket) i() *geometry.Point {
	return p.h().SquareUp(p.placketOpeningLength)
}

func (p *pn11SleevePlacket) j() *geometry.Point {
	return p.i().SquareRight(PN11_PLACKET_WIDTH)
}

func (p *pn11SleevePlacket) k() *geometry.Point {
	return p.j().SquareToHorizontalLine(p.h().Y)
}

func (p *pn11SleevePlacket) l() *geometry.Point {
	return p.b().MidpointTo(p.d())
}
func (p *pn11SleevePlacket) placketLeftStitch() geometry.Line {
	return &geometry.StraightLine{
		Start: p.a(),
		End: p.e(),
	}
}

func (p *pn11SleevePlacket) placketRightStitch() geometry.Line {
	return &geometry.StraightLine{
		Start: p.d(),
		End: p.c(),
	}
}

func (p *pn11SleevePlacket) placketBottomStitch() geometry.Line {
	return &geometry.StraightLine{
		Start: p.c(),
		End: p.a(),
	}
}

func (p *pn11SleevePlacket) placketTopLeftStitch() geometry.Line {
	return &geometry.StraightLine{
		Start: p.e(),
		End: p.g(),
	}
}

func (p *pn11SleevePlacket) placketTopRightStitch() geometry.Line {
	return &geometry.StraightLine{
		Start: p.g(),
		End: p.f(),
	}
}

func (p *pn11SleevePlacket) placketCornerStitch() geometry.Line {
	poly := &geometry.Polyline{}

	poly.AddLine(
		&geometry.StraightLine{
			Start: p.f(),
			End: p.l(),
		},
		&geometry.StraightLine{
			Start: p.l(),
			End: p.d(),
		},
	)

	return poly
}

func (p *pn11SleevePlacket) standLeftStitch() geometry.Line {
	return &geometry.StraightLine{
		Start: p.h(),
		End: p.i(),
	}
}

func (p *pn11SleevePlacket) standTopStitch() geometry.Line {
	return &geometry.StraightLine{
		Start: p.i(),
		End: p.j(),
	}
}

func (p *pn11SleevePlacket) standRightStitch() geometry.Line {
	return &geometry.StraightLine{
		Start: p.j(),
		End: p.k(),
	}
}

func (p *pn11SleevePlacket) standBottomStitch() geometry.Line {
	return &geometry.StraightLine{
		Start: p.k(),
		End: p.h(),
	}
}
