package pieces

import (
	"fmt"

	"github.com/tailored-style/pattern-generator/geometry"
	"github.com/tailored-style/pattern-generator/pieces"
	"github.com/tobyjsullivan/catalogue/anchors"
	"github.com/tobyjsullivan/catalogue/slopers"
)

type pn7Sleeve struct {
	*slopers.TorsoMeasurements
	cuffDepth            float64
	placketOpeningLength float64
}

func NewPN7Sleeve(m *slopers.TorsoMeasurements, cuffDepth float64, placketOpening float64) pieces.Piece {
	return &pn7Sleeve{
		TorsoMeasurements: 	  m,
		cuffDepth:            cuffDepth,
		placketOpeningLength: placketOpening,
	}
}

func (p *pn7Sleeve) Details() *pieces.Details {
	return &pieces.Details{
		PieceNumber: 7,
		Description: "Sleeve",
	}
}

func (p *pn7Sleeve) CutCount() int {
	return 1
}

func (p *pn7Sleeve) OnFold() bool {
	return false
}

func (p *pn7Sleeve) Mirrored() bool {
	return true
}

func (p *pn7Sleeve) sloper() *slopers.Sleeve {
	return &slopers.Sleeve{
		TorsoMeasurements: p.TorsoMeasurements,
	}
}

func (p *pn7Sleeve) a() *geometry.Point {
	return p.sloper().A()
}

func (p *pn7Sleeve) shoulderSeamLength() float64 {
	return (&pn6Yoke{
		TorsoMeasurements: p.TorsoMeasurements,
	}).shoulderSeamLength()
}

func (p *pn7Sleeve) b() *geometry.Point {
	b := p.sloper().B()
	fmt.Printf("Sleeve B: %v\n", b)
	return b
}

func (p *pn7Sleeve) c() *geometry.Point {
	c := p.sloper().C()
	fmt.Printf("Sleeve C: %v\n", c)

	return c
}

func (p *pn7Sleeve) d() *geometry.Point {
	return p.sloper().D()
}

func (p *pn7Sleeve) e() *geometry.Point {
	e := p.sloper().E()
	fmt.Printf("Sleeve E: %v\n", e)

	return e
}

func (p *pn7Sleeve) f() *geometry.Point {
	return p.sloper().F()
}

func (p *pn7Sleeve) i() *geometry.Point {
	return p.sloper().I()
}

func (p *pn7Sleeve) j() *geometry.Point {
	return p.sloper().J()
}

func (p *pn7Sleeve) l() *geometry.Point {
	return p.sloper().L()
}

func (p *pn7Sleeve) n() *geometry.Point {
	return p.sloper().N()
}

func (p *pn7Sleeve) o() *geometry.Point {
	return p.sloper().O()
}

func (p *pn7Sleeve) s() *geometry.Point {
	return p.sloper().S()
}

func (p *pn7Sleeve) t() *geometry.Point {
	return p.sloper().T()
}

func (p *pn7Sleeve) u() *geometry.Point {
	return p.sloper().U()
}

func (p *pn7Sleeve) v() *geometry.Point {
	return p.sloper().V()
}

func (p *pn7Sleeve) w() *geometry.Point {
	return p.sloper().W()
}

func (p *pn7Sleeve) x() *geometry.Point {
	return p.sloper().X()
}

func (p *pn7Sleeve) y() *geometry.Point {
	return p.b().MidpointTo(p.w())
}

func (p *pn7Sleeve) z() *geometry.Point {
	return p.y().SquareUp(p.placketOpeningLength)
}

func (p *pn7Sleeve) frontArmholeLength() float64 {
	front := &pn4TorsoFront{
		TorsoMeasurements: p.TorsoMeasurements,
	}

	return front.armholeStitch().Length()
}

func (p *pn7Sleeve) backArmholeLength() float64 {
	back := &pn5TorsoBack{
		height:             p.TorsoMeasurements.Height,
		chestCircumference: p.TorsoMeasurements.ChestCircumference,
		waistCircumference: p.TorsoMeasurements.BellyButtonWaistCircumference,
		hipCircumference:   p.TorsoMeasurements.HipCircumference,
	}

	yoke := &pn6Yoke{
		TorsoMeasurements: p.TorsoMeasurements,
	}

	return back.armholeStitch().Length() + yoke.armholeStitch().Length()
}

func (p *pn7Sleeve) frontArmholeStitch() geometry.Line {
	return &geometry.PolyNCurve{
		Points: []*geometry.Point{
			p.a(),
			p.n(),
			p.l(),
			p.o(),
			p.e(),
		},
		StartAngle: &geometry.Angle{Rads: 0.0},
		EndAngle:   p.w().AngleRelativeTo(p.e()).Perpendicular(),
	}
}

func (p *pn7Sleeve) backArmholeStitch() geometry.Line {
	return &geometry.PolyNCurve{
		Points: []*geometry.Point{
			p.a(),
			p.s(),
			p.t(),
			p.u(),
			p.v(),
			p.f(),
		},
		StartAngle: &geometry.Angle{Rads: 0.0},
		EndAngle:   p.x().AngleRelativeTo(p.f()).Perpendicular(),
	}
}

func (p *pn7Sleeve) underSleeveStitchLeft() geometry.Line {
	return &geometry.StraightLine{
		Start: p.f(),
		End:   p.x(),
	}
}

func (p *pn7Sleeve) underSleeveStitchRight() geometry.Line {
	return &geometry.StraightLine{
		Start: p.e(),
		End:   p.w(),
	}
}

func (p *pn7Sleeve) cuffStitch() geometry.Line {
	return &geometry.StraightLine{
		Start: p.x(),
		End:   p.w(),
	}
}

func (p *pn7Sleeve) OuterCut() *geometry.Polyline {
	return &geometry.Polyline{}
}

func (p *pn7Sleeve) InnerCut() *geometry.Block {
	layer := &geometry.Block{}

	frontArmholeStitch := p.frontArmholeStitch()
	backArmholeStitch := p.backArmholeStitch()

	seamAllowance := pieces.SeamAllowance(true,
		&geometry.ReverseLine{InnerLine: pieces.AddSeamAllowance(frontArmholeStitch, false)},
		pieces.AddSeamAllowance(backArmholeStitch, true),
		pieces.AddSeamAllowance(p.underSleeveStitchLeft(), true),
		pieces.AddSeamAllowance(p.cuffStitch(), true),
		pieces.AddSeamAllowance(p.underSleeveStitchRight(), false),
	)

	placketCut := &geometry.StraightLine{
		Start: p.y().SquareDown(pieces.SEAM_ALLOWANCE),
		End:   p.z(),
	}

	layer.AddLine(
		seamAllowance,
		placketCut,
		pieces.Notch(frontArmholeStitch, 7.6, false),
		pieces.Notch(frontArmholeStitch, frontArmholeStitch.Length()-7.6, false),
		pieces.Notch(backArmholeStitch, 7.6, true),
		pieces.Notch(backArmholeStitch, backArmholeStitch.Length()-7.6, true),
		pieces.Notch(backArmholeStitch, backArmholeStitch.Length()-8.9, true),
	)

	return layer
}

func (p *pn7Sleeve) Stitch() *geometry.Block {
	layer := &geometry.Block{}

	layer.AddLine(
		p.frontArmholeStitch(),
		p.backArmholeStitch(),
		p.underSleeveStitchLeft(),
		p.underSleeveStitchRight(),
		p.cuffStitch(),
	)

	return layer
}

func (p *pn7Sleeve) Ink() *geometry.Block {
	return &geometry.Block{}
}

func (p *pn7Sleeve) Reference() *geometry.Block {
	layer := &geometry.Block{}

	if DEBUG {
		centreOfSleeve := &geometry.StraightLine{
			Start: p.a(),
			End:   p.b(),
		}

		bicepLine := &geometry.StraightLine{
			Start: p.f(),
			End:   p.e(),
		}

		elbowLine := &geometry.StraightLine{
			Start: p.j(),
			End:   p.i(),
		}

		layer.AddLine(
			centreOfSleeve,
			bicepLine,
			elbowLine,
		)

		a := make(map[string]*geometry.Point)
		a["A"] = p.a()
		a["B"] = p.b()
		a["C"] = p.c()
		a["D"] = p.d()
		a["E"] = p.e()
		a["F"] = p.f()
		a["I"] = p.i()
		a["J"] = p.j()
		a["L"] = p.l()
		a["N"] = p.n()
		a["O"] = p.o()
		a["S"] = p.s()
		a["T"] = p.t()
		a["U"] = p.u()
		a["V"] = p.v()
		a["W"] = p.w()
		a["X"] = p.x()
		anchors.AddAnchors(layer, a)
	}

	return layer
}

func (p *pn7Sleeve) String() string {
	return fmt.Sprintf("[PN: %d]", p.Details().PieceNumber)
}
