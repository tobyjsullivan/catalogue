package pieces

import (
	"fmt"
	"math"

	"github.com/tailored-style/pattern-generator/geometry"
	"github.com/tailored-style/pattern-generator/pieces"
	"github.com/tobyjsullivan/catalogue/anchors"
	"github.com/tobyjsullivan/catalogue/slopers"
)

type pn9CollarBand struct {
	*slopers.TorsoMeasurements
}

func NewPN9CollarBand(m *slopers.TorsoMeasurements) pieces.Piece {
	return &pn9CollarBand{
		TorsoMeasurements: m,
	}
}

func (p *pn9CollarBand) CutCount() int {
	return 2
}

func (p *pn9CollarBand) OnFold() bool {
	return true
}

func (p *pn9CollarBand) Mirrored() bool {
	return false
}

func (p *pn9CollarBand) Stitch() *geometry.Block {
	layer := &geometry.Block{}

	layer.AddLine(
		p.bottomStitch(),
		p.topStitch(),
	)

	return layer
}

func (p *pn9CollarBand) OuterCut() *geometry.Polyline {
	return &geometry.Polyline{}
}

func (p *pn9CollarBand) InnerCut() *geometry.Block {
	layer := &geometry.Block{}

	seamAllowance := pieces.SeamAllowance(false,
		pieces.AddSeamAllowance(p.topStitch(), false),
		&geometry.ReverseLine{InnerLine: pieces.AddSeamAllowance(p.bottomStitch(), true)},
	)

	layer.AddLine(
		seamAllowance,
	)

	return layer
}

func (p *pn9CollarBand) Ink() *geometry.Block {
	return &geometry.Block{}
}

func (p *pn9CollarBand) Reference() *geometry.Block {
	layer := &geometry.Block{}

	layer.AddLine(
		p.centreBack(),
	)

	// Draw all points (DEBUG)
	if DEBUG {
		layer.AddLine(
			p.centreFront(),
		)

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
		a["J"] = p.j()
		a["K"] = p.k()
		a["L"] = p.l()
		a["M"] = p.m()
		anchors.AddAnchors(layer, a)
	}

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
		TorsoMeasurements: p.TorsoMeasurements,
	}).backNeckLine().Length()
}

func (p *pn9CollarBand) frontNeckLength() float64 {
	return (&pn4TorsoFront{
		TorsoMeasurements: p.TorsoMeasurements,
	}).frontNeckLine().Length()
}

func (p *pn9CollarBand) bandHeight() float64 {
	return 3.2
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
	return p.a().SquareUp(p.bandHeight())
}

func (p *pn9CollarBand) e() *geometry.Point {
	return p.d().SquareUp(1.0)
}

func (p *pn9CollarBand) f() *geometry.Point {
	return p.c().SquareUp(p.a().DistanceTo(p.e()))
}

func (p *pn9CollarBand) g() *geometry.Point {
	return p.c().SquareUp(1.3)
}

func (p *pn9CollarBand) frontBottomLine() geometry.Line {
	return &geometry.ParabolaCurve{
		Start:         p.b(),
		End:           p.g(),
		StartingAngle: &geometry.Angle{Rads: 0.0},
		ArcAngle:      &geometry.Angle{Rads: math.Pi / 10.0},
	}
}

func (p *pn9CollarBand) h() *geometry.Point {
	bottomStitch := p.frontBottomLine()
	angle := bottomStitch.AngleAt(bottomStitch.Length() - 0.01).Perpendicular()
	length := p.g().DistanceTo(p.f()) / angle.Sin()
	return p.g().DrawAt(angle, length)
}

func (p *pn9CollarBand) i() *geometry.Point {
	return p.h().SquareLeft(0.3).SquareDown(0.3)
}

func (p *pn9CollarBand) j() *geometry.Point {
	angle := p.g().AngleRelativeTo(p.i())

	return p.f().DrawAt(angle, p.i().DistanceTo(p.g()))
}

func (p *pn9CollarBand) k() *geometry.Point {
	return p.d().SquareRight(p.a().DistanceTo(p.b()))
}

func (p *pn9CollarBand) l() *geometry.Point {
	return p.f().DrawAt(&geometry.Angle{Rads: -math.Pi * 3.0 / 4.0}, 0.3)
}

func (p *pn9CollarBand) m() *geometry.Point {
	return (&geometry.StraightLine{
		Start: p.f(),
		End:   p.j(),
	}).Resize(1.3).End
}

func (p *pn9CollarBand) bottomStitch() geometry.Line {
	line := &geometry.Polyline{}

	end := &geometry.StraightLine{
		Start: p.g(),
		End:   p.j(),
	}

	line.AddLine(
		&geometry.StraightLine{
			Start: p.a(),
			End:   p.b(),
		},
		&geometry.PolyNCurve{
			Points: []*geometry.Point{
				p.b(),
				p.g(),
			},
			StartAngle: &geometry.Angle{Rads: 0.0},
			EndAngle:   end.AngleAt(0.0),
		},
		end,
	)

	return line
}

func (p *pn9CollarBand) topStitch() geometry.Line {
	line := &geometry.Polyline{}

	end := &geometry.StraightLine{
		Start: p.m(),
		End:   p.j(),
	}

	angleAtI := (&geometry.StraightLine{Start: p.i(), End: p.f()}).AngleAt(0.0)

	line.AddLine(
		&geometry.StraightLine{
			Start: p.d(),
			End:   p.k(),
		},
		&geometry.PolyNCurve{
			Points: []*geometry.Point{
				p.k(),
				p.i(),
			},
			StartAngle: &geometry.Angle{Rads: 0.0},
			EndAngle:   angleAtI,
		},
		&geometry.PolyNCurve{
			Points: []*geometry.Point{
				p.i(),
				p.l(),
				p.m(),
			},
			StartAngle: angleAtI,
			EndAngle:   end.AngleAt(0.0),
		},
		end,
	)

	return line
}

func (p *pn9CollarBand) centreBack() geometry.Line {
	return &geometry.StraightLine{
		Start: p.a(),
		End:   p.d(),
	}
}

func (p *pn9CollarBand) centreFront() geometry.Line {
	return &geometry.StraightLine{
		Start: p.g(),
		End:   p.i(),
	}
}
