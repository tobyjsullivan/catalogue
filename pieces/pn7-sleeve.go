package pieces

import (
	"github.com/tailored-style/pattern-generator/pieces"
	"github.com/tailored-style/pattern-generator/geometry"
	"math"
	"fmt"
)

type pn7Sleeve struct {
	height float64
	neckCircumference float64
	chestCircumference float64
	waistCircumference float64
	hipCircumference float64
	sleeveLength float64
}

func NewPN7Sleeve(height float64, neck float64, chest float64, waist float64, hip float64, sleeve float64) pieces.Piece {
	return &pn7Sleeve{
		height:height,
		neckCircumference:neck,
		chestCircumference:chest,
		waistCircumference:waist,
		hipCircumference:hip,
		sleeveLength:sleeve,
	}
}

func (p *pn7Sleeve) Details() *pieces.Details {
	return &pieces.Details{
		PieceNumber: 7,
		Description: "Sleeve",
	}
}

func (p *pn7Sleeve) OnFold() bool {
	return false
}

func (p *pn7Sleeve) a() *geometry.Point {
	return &geometry.Point{
		X: 0.0,
		Y: 0.0,
	}
}

func (p *pn7Sleeve) b() *geometry.Point {
	return p.a().SquareDown(p.sleeveLength - 32.7)
}

func (p *pn7Sleeve) c() *geometry.Point {
	armholeLength := p.frontArmholeLength() + p.backArmholeLength()
	return p.a().SquareDown(armholeLength / 3.0 - 3.0)
}

func (p *pn7Sleeve) d() *geometry.Point {
	return p.c().SquareDown(p.b().DistanceTo(p.c()) / 2.0 - 4.6)
}

func (p *pn7Sleeve) e() *geometry.Point {
	aToE := p.frontArmholeLength() - 0.3
	c := p.c()
	return c.SquareRight(math.Sqrt(math.Pow(aToE, 2.0) - math.Pow(p.a().DistanceTo(c), 2.0)))
}

func (p *pn7Sleeve) f() *geometry.Point {
	aToF := p.backArmholeLength()
	c := p.c()
	return c.SquareLeft(math.Sqrt(math.Pow(aToF, 2.0) - math.Pow(p.a().DistanceTo(c), 2.0)))
}

func (p *pn7Sleeve) g() *geometry.Point {
	return p.e().SquareToHorizontalLine(p.b().Y)
}

func (p *pn7Sleeve) h() *geometry.Point {
	return p.f().SquareToHorizontalLine(p.b().Y)
}

func (p *pn7Sleeve) i() *geometry.Point {
	return p.d().SquareToVerticalLine(p.e().X)
}

func (p *pn7Sleeve) j() *geometry.Point {
	return p.d().SquareToVerticalLine(p.f().X)
}

func (p *pn7Sleeve) k() *geometry.Point {
	a := p.a()
	e := p.e()
	return a.DrawAt(e.AngleRelativeTo(a), a.DistanceTo(e) / 4.0)
}

func (p *pn7Sleeve) l() *geometry.Point {
	a := p.a()
	e := p.e()
	return a.DrawAt(e.AngleRelativeTo(a), a.DistanceTo(e) / 2.0 + 1.0)
}

func (p *pn7Sleeve) m() *geometry.Point {
	a := p.a()
	e := p.e()
	return a.DrawAt(e.AngleRelativeTo(a), a.DistanceTo(e) * 3.0 / 4.0)
}

func (p *pn7Sleeve) n() *geometry.Point {
	k := p.k()
	return k.DrawAt(k.AngleRelativeTo(p.a()).Perpendicular(), 1.6)
}

func (p *pn7Sleeve) o() *geometry.Point {
	m := p.m()
	return m.DrawAt(m.AngleRelativeTo(p.a()).Perpendicular().Opposite(), 1.3)
}

func (p *pn7Sleeve) p() *geometry.Point {
	a := p.a()
	f := p.f()
	return a.DrawAt(f.AngleRelativeTo(a), a.DistanceTo(f) / 4.0)
}

func (p *pn7Sleeve) q() *geometry.Point {
	a := p.a()
	f := p.f()
	return a.DrawAt(f.AngleRelativeTo(a), a.DistanceTo(f) / 2.0)
}

func (p *pn7Sleeve) r() *geometry.Point {
	a := p.a()
	f := p.f()
	return a.DrawAt(f.AngleRelativeTo(a), a.DistanceTo(f) * 3.0 / 4.0)
}

func (p *pn7Sleeve) s() *geometry.Point {
	pa := p.p()
	return pa.DrawAt(pa.AngleRelativeTo(p.a()).Perpendicular().Opposite(), 1.9)
}

func (p *pn7Sleeve) t() *geometry.Point {
	q := p.q()
	return q.DrawAt(q.AngleRelativeTo(p.a()).Perpendicular().Opposite(), 1.0)
}

func (p *pn7Sleeve) u() *geometry.Point {
	a := p.a()
	f := p.f()
	return a.DrawAt(f.AngleRelativeTo(a), a.DistanceTo(f) * 3.0 / 4.0 - 1.9)
}

func (p *pn7Sleeve) v() *geometry.Point {
	f := p.f()
	r := p.r()
	return f.MidpointTo(r).DrawAt(f.AngleRelativeTo(r).Perpendicular(), 0.6)
}

func (p *pn7Sleeve) w() *geometry.Point {
	return p.g().SquareLeft(7.9)
}

func (p *pn7Sleeve) x() *geometry.Point {
	return p.h().SquareRight(5.6)
}

func (p *pn7Sleeve) frontArmholeLength() float64 {
	front := &pn4TorsoFront{
		height: p.height,
		neckCircumference: p.neckCircumference,
		chestCircumference: p.chestCircumference,
		waistCircumference: p.waistCircumference,
		hipCircumference: p.hipCircumference,
	}

	return front.armholeStitch().Length()
}

func (p *pn7Sleeve) backArmholeLength() float64 {
	back := &pn5TorsoBack{
		height: p.height,
		chestCircumference: p.chestCircumference,
		waistCircumference: p.waistCircumference,
		hipCircumference: p.hipCircumference,
	}

	yoke := &pn6Yoke{
		height: p.height,
		neckCircumference: p.neckCircumference,
		chestCircumference: p.chestCircumference,
		waistCircumference: p.waistCircumference,
		hipCircumference: p.hipCircumference,
	}

	return back.armholeStitch().Length() + yoke.armholeStitch().Length()
}

func (p *pn7Sleeve) frontArmholeStitch() geometry.Line {
	a := p.a()

	shoulderAngle := p.e().AngleRelativeTo(a)

	partA := &geometry.Poly3Curve{
		P0: a,
		P1: p.n(),
		A0: &geometry.Angle{Rads: 0.0},
		A1: shoulderAngle,
	}

	partB := &geometry.ThreePointCurve{
		Start: p.n(),
		Middle: p.l(),
		End: p.o(),
		Rotation: shoulderAngle,
	}

	startAngle := p.w().AngleRelativeTo(p.e()).Perpendicular()
	partC := &geometry.Poly3Curve{
		P0: p.o(),
		P1: p.e(),
		A0: shoulderAngle,
		A1: startAngle,
	}

	line := &geometry.Polyline{}

	line.AddLine(
		partA,
		partB,
		partC,
	)

	return line
}

func (p *pn7Sleeve) backArmholeStitch() geometry.Line {
	shoulderAngle := p.f().AngleRelativeTo(p.a())

	partA := &geometry.Poly3Curve{
		P0: p.a(),
		P1: p.s(),
		A0: &geometry.Angle{Rads: 0.0},
		A1: shoulderAngle,
	}

	partB := &geometry.ThreePointCurve{
		Start: p.s(),
		Middle: p.u(),
		End: p.v(),
		Rotation: shoulderAngle,
	}

	startAngle := p.x().AngleRelativeTo(p.f())
	partC := &geometry.Poly3Curve{
		P0: p.f(),
		P1: p.v(),
		A0: startAngle.Perpendicular(),
		A1: shoulderAngle,
	}

	line := &geometry.Polyline{}

	line.AddLine(
		partA,
		partB,
		&geometry.ReverseLine{InnerLine: partC},
	)

	return line
}

func (p *pn7Sleeve) underSleeveStitchLeft() geometry.Line {
	return &geometry.StraightLine{
		Start: p.f(),
		End: p.x(),
	}
}

func (p *pn7Sleeve) underSleeveStitchRight() geometry.Line {
	return &geometry.StraightLine{
		Start: p.e(),
		End: p.w(),
	}
}

func (p *pn7Sleeve) cuffStitch() geometry.Line {
	return &geometry.StraightLine{
		Start: p.x(),
		End: p.w(),
	}
}

func (p *pn7Sleeve) CutLayer() *geometry.Block {
	layer := &geometry.Block{}

	frontArmholeCut := pieces.AddSeamAllowance(p.frontArmholeStitch(), false)
	backArmholeCut := pieces.AddSeamAllowance(p.backArmholeStitch(), true)

	layer.AddLine(
		frontArmholeCut,
		pieces.Notch(frontArmholeCut, 7.6),
		pieces.Notch(frontArmholeCut, frontArmholeCut.Length() - 7.6),
		pieces.Notch(frontArmholeCut, frontArmholeCut.Length() - 8.9),
		backArmholeCut,
		pieces.Notch(backArmholeCut, 7.6),
		pieces.Notch(backArmholeCut, backArmholeCut.Length() - 7.6),
		pieces.AddSeamAllowance(p.underSleeveStitchLeft(), true),
		pieces.AddSeamAllowance(p.underSleeveStitchRight(), false),
		pieces.AddSeamAllowance(p.cuffStitch(), true),
	)

	return layer
}

func (p *pn7Sleeve) StitchLayer() *geometry.Block {
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

func (p *pn7Sleeve) NotationLayer() *geometry.Block {
	layer := &geometry.Block{}

	centreOfSleeve := &geometry.StraightLine{
		Start: p.a(),
		End: p.b(),
	}

	bicepLine := &geometry.StraightLine{
		Start: p.f(),
		End: p.e(),
	}

	elbowLine := &geometry.StraightLine{
		Start: p.j(),
		End: p.i(),
	}

	layer.AddLine(
		centreOfSleeve,
		bicepLine,
		elbowLine,
	)

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
	anchors["M"] = p.m()
	anchors["N"] = p.n()
	anchors["O"] = p.o()
	anchors["P"] = p.p()
	anchors["Q"] = p.q()
	anchors["R"] = p.r()
	anchors["S"] = p.s()
	anchors["T"] = p.t()
	anchors["U"] = p.u()
	anchors["V"] = p.v()
	anchors["W"] = p.w()
	anchors["X"] = p.x()
	addAnchors(layer, anchors)

	return layer
}

func (p *pn7Sleeve) String() string {
	return fmt.Sprintf("[PN: %d]", p.Details().PieceNumber)
}