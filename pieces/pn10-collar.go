package pieces

import (
	"fmt"
	"math"

	"github.com/tailored-style/pattern-generator/geometry"
	"github.com/tailored-style/pattern-generator/pieces"
	"github.com/tobyjsullivan/catalogue/anchors"
	"github.com/tobyjsullivan/catalogue/slopers"
)

type pn10Collar struct {
	*slopers.TorsoMeasurements
}

func NewPN10Collar(m *slopers.TorsoMeasurements) pieces.Piece {
	return &pn10Collar{
		TorsoMeasurements: m,
	}
}

func (p *pn10Collar) CutCount() int {
	return 2
}

func (p *pn10Collar) OnFold() bool {
	return true
}

func (p *pn10Collar) Mirrored() bool {
	return false
}

func (p *pn10Collar) Stitch() *geometry.Block {
	layer := &geometry.Block{}

	layer.AddLine(
		p.bottomStitch(),
		p.rightStitch(),
		p.topStitch(),
	)

	return layer
}

func (p *pn10Collar) OuterCut() *geometry.Polyline {
	return &geometry.Polyline{}
}

func (p *pn10Collar) InnerCut() *geometry.Block {
	layer := &geometry.Block{}

	seamAllowance := pieces.SeamAllowance(false,
		pieces.AddSeamAllowance(p.topStitch(), false),
		pieces.AddSeamAllowance(p.rightStitch(), false),
		&geometry.ReverseLine{InnerLine: pieces.AddSeamAllowance(p.bottomStitch(), true)},
	)

	layer.AddLine(
		seamAllowance,
	)

	return layer
}

func (p *pn10Collar) Ink() *geometry.Block {
	return &geometry.Block{}
}

func (p *pn10Collar) Reference() *geometry.Block {
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
		a["D"] = p.d()
		a["E"] = p.e()
		a["F"] = p.f()
		a["G"] = p.g()
		a["H"] = p.h()
		a["I"] = p.i()
		anchors.AddAnchors(layer, a)
	}

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
		TorsoMeasurements: p.TorsoMeasurements,
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
		End:   p.e(),
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
		End:   p.a(),
	}
}

func (p *pn10Collar) bottomStitch() geometry.Line {
	line := &geometry.Polyline{}

	line.AddLine(
		&geometry.StraightLine{
			Start: p.a(),
			End:   p.g(),
		},
		&geometry.PolyNCurve{
			Points: []*geometry.Point{
				p.g(),
				p.c(),
			},
			StartAngle: &geometry.Angle{Rads: 0.0},
			EndAngle:   p.f().AngleRelativeTo(p.c()).Perpendicular(),
		},
	)

	return line
}

func (p *pn10Collar) rightStitch() geometry.Line {
	return &geometry.StraightLine{
		Start: p.f(),
		End:   p.c(),
	}
}

func (p *pn10Collar) topStitch() geometry.Line {
	line := &geometry.Polyline{}

	line.AddLine(
		&geometry.StraightLine{
			Start: p.b(),
			End:   p.h(),
		},
		&geometry.PolyNCurve{
			Points: []*geometry.Point{
				p.h(),
				p.f(),
			},
			StartAngle: &geometry.Angle{Rads: 0.0},
			EndAngle:   &geometry.Angle{Rads: math.Pi / 24.0},
		},
	)

	return line
}
