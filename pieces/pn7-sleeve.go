package pieces

import (
	"github.com/tailored-style/pattern-generator/pieces"
	"github.com/tailored-style/pattern-generator/geometry"
	"math"
)

type PN7Sleeve struct {
	*pieces.Measurements
}

func (p *PN7Sleeve) Details() *pieces.Details {
	return &pieces.Details{
		PieceNumber: "7",
		Description: "Sleeve",
	}
}

func (p *PN7Sleeve) OnFold() bool {
	return false
}

func (p *PN7Sleeve) a() *geometry.Point {
	return &geometry.Point{
		X: 0.0,
		Y: 0.0,
	}
}

func (p *PN7Sleeve) b() *geometry.Point {
	return p.a().SquareDown(p.Measurements.SleeveLength - 32.7)
}

func (p *PN7Sleeve) c() *geometry.Point {
	armholeLength := p.frontArmholeLength() + p.backArmholeLength()
	return p.a().SquareDown(armholeLength / 3.0 - 3.0)
}

func (p *PN7Sleeve) d() *geometry.Point {
	return p.c().SquareDown(p.b().DistanceTo(p.c()) / 2.0 - 4.6)
}

func (p *PN7Sleeve) e() *geometry.Point {
	aToE := p.frontArmholeLength() - 0.3
	c := p.c()
	return c.SquareRight(math.Sqrt(math.Pow(aToE, 2.0) - math.Pow(p.a().DistanceTo(c), 2.0)))
}

func (p *PN7Sleeve) f() *geometry.Point {
	aToF := p.backArmholeLength()
	c := p.c()
	return c.SquareLeft(math.Sqrt(math.Pow(aToF, 2.0) - math.Pow(p.a().DistanceTo(c), 2.0)))
}

func (p *PN7Sleeve) g() *geometry.Point {
	return p.e().SquareToHorizontalLine(p.b().Y)
}

func (p *PN7Sleeve) h() *geometry.Point {
	return p.f().SquareToHorizontalLine(p.b().Y)
}

func (p *PN7Sleeve) i() *geometry.Point {
	return p.d().SquareToVerticalLine(p.e().X)
}

func (p *PN7Sleeve) j() *geometry.Point {
	return p.d().SquareToVerticalLine(p.f().X)
}

func (p *PN7Sleeve) k() *geometry.Point {
	a := p.a()
	e := p.e()
	return a.DrawAt(e.AngleRelativeTo(a), a.DistanceTo(e) / 4.0)
}

func (p *PN7Sleeve) l() *geometry.Point {
	a := p.a()
	e := p.e()
	return a.DrawAt(e.AngleRelativeTo(a), a.DistanceTo(e) / 2.0 + 1.0)
}

func (p *PN7Sleeve) m() *geometry.Point {
	a := p.a()
	e := p.e()
	return a.DrawAt(e.AngleRelativeTo(a), a.DistanceTo(e) * 3.0 / 4.0)
}

func (p *PN7Sleeve) n() *geometry.Point {
	k := p.k()
	return k.DrawAt(k.AngleRelativeTo(p.a()).Perpendicular(), 1.6)
}

func (p *PN7Sleeve) o() *geometry.Point {
	m := p.m()
	return m.DrawAt(m.AngleRelativeTo(p.a()).Perpendicular().Opposite(), 1.3)
}

func (p *PN7Sleeve) p() *geometry.Point {
	a := p.a()
	f := p.f()
	return a.DrawAt(f.AngleRelativeTo(a), a.DistanceTo(f) / 4.0)
}

func (p *PN7Sleeve) q() *geometry.Point {
	a := p.a()
	f := p.f()
	return a.DrawAt(f.AngleRelativeTo(a), a.DistanceTo(f) / 2.0)
}

func (p *PN7Sleeve) r() *geometry.Point {
	a := p.a()
	f := p.f()
	return a.DrawAt(f.AngleRelativeTo(a), a.DistanceTo(f) * 3.0 / 4.0)
}

func (p *PN7Sleeve) s() *geometry.Point {
	pa := p.p()
	return pa.DrawAt(pa.AngleRelativeTo(p.a()).Perpendicular().Opposite(), 1.9)
}

func (p *PN7Sleeve) t() *geometry.Point {
	q := p.q()
	return q.DrawAt(q.AngleRelativeTo(p.a()).Perpendicular().Opposite(), 1.0)
}

func (p *PN7Sleeve) u() *geometry.Point {
	a := p.a()
	f := p.f()
	return a.DrawAt(f.AngleRelativeTo(a), a.DistanceTo(f) * 3.0 / 4.0 - 1.9)
}

func (p *PN7Sleeve) v() *geometry.Point {
	f := p.f()
	r := p.r()
	return f.MidpointTo(r).DrawAt(f.AngleRelativeTo(r).Perpendicular(), 0.6)
}

func (p *PN7Sleeve) w() *geometry.Point {
	return p.g().SquareLeft(7.9)
}

func (p *PN7Sleeve) x() *geometry.Point {
	return p.h().SquareRight(5.6)
}

func (p *PN7Sleeve) frontArmholeLength() float64 {
	front := &PN4TorsoFront{
		p.Measurements,
	}

	return front.armholeStitch().Length()
}

func (p *PN7Sleeve) backArmholeLength() float64 {
	back := &PN5TorsoBack{
		p.Measurements,
	}

	yoke := &PN6Yoke{
		p.Measurements,
	}

	return back.armholeStitch().Length() + yoke.armholeStitch().Length()
}

func (p *PN7Sleeve) frontArmholeStitch() geometry.Line {
	a := p.a()

	shoulderAngle := p.e().AngleRelativeTo(a)

	partA := &geometry.EllipseCurve{
		Start: a,
		End: p.n(),
		StartingAngle: &geometry.Angle{Rads: math.Pi * 3.0 / 2.0},
		ArcAngle: shoulderAngle.Perpendicular(),
	}

	partB := &geometry.ThreePointCurve{
		Start: p.n(),
		Middle: p.l(),
		End: p.o(),
		Rotation: shoulderAngle,
	}

	startAngle := p.w().AngleRelativeTo(p.e()).Perpendicular()
	partC := &geometry.ParabolaCurve{
		Start: p.e(),
		End: p.o(),
		StartingAngle: startAngle,
		ArcAngle: shoulderAngle.Subtract(startAngle),
	}

	//partC := &geometry.EllipseCurve{
	//	Start: p.e(),
	//	End: p.o(),
	//	StartingAngle: p.w().AngleRelativeTo(p.e()).Opposite(),
	//	ArcAngle: shoulderAngle.Perpendicular().Neg(),
	//}

	line := &geometry.Polyline{}

	line.AddLine(
		partA,
		partB,
		&geometry.ReverseLine{InnerLine: partC},
	)

	return line
}

func (p *PN7Sleeve) backArmholeStitch() geometry.Line {
	shoulderAngle := p.f().AngleRelativeTo(p.a())

	partA := &geometry.EllipseCurve{
		Start: p.a(),
		End: p.s(),
		StartingAngle: &geometry.Angle{Rads: math.Pi / 2.0},
		ArcAngle: shoulderAngle.Perpendicular(),
	}

	partB := &geometry.ThreePointCurve{
		Start: p.s(),
		Middle: p.u(),
		End: p.v(),
		Rotation: shoulderAngle,
	}

	startAngle := p.x().AngleRelativeTo(p.f())
	partC := &geometry.EllipseCurve{
		Start: p.f(),
		End: p.v(),
		StartingAngle: startAngle,
		ArcAngle: p.a().AngleRelativeTo(p.f()).Subtract(startAngle).Perpendicular().Opposite(),
	}

	line := &geometry.Polyline{}

	line.AddLine(
		partA,
		partB,
		&geometry.ReverseLine{InnerLine: partC},
	)

	return line
}

func (p *PN7Sleeve) underSleeveStitchLeft() geometry.Line {
	return &geometry.StraightLine{
		Start: p.f(),
		End: p.x(),
	}
}

func (p *PN7Sleeve) underSleeveStitchRight() geometry.Line {
	return &geometry.StraightLine{
		Start: p.e(),
		End: p.w(),
	}
}

func (p *PN7Sleeve) cuffStitch() geometry.Line {
	return &geometry.StraightLine{
		Start: p.x(),
		End: p.w(),
	}
}

func (p *PN7Sleeve) CutLayer() *geometry.Block {
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

func (p *PN7Sleeve) StitchLayer() *geometry.Block {
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

func (p *PN7Sleeve) NotationLayer() *geometry.Block {
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
