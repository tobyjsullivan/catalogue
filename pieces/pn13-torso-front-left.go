package pieces

import (
	"github.com/tailored-style/pattern-generator/geometry"
	"github.com/tailored-style/pattern-generator/pieces"
	"github.com/tailored-style/pattern-generator/symbols"
	"math"
)

type pn13TorsoFrontLeft struct {
	height             float64
	neckCircumference  float64
	chestCircumference float64
	waistCircumference float64
	hipCircumference   float64
}

func NewPN13TorsoFrontLeft(height float64, neck float64, chest float64, waist float64, hip float64) pieces.Piece {
	return &pn13TorsoFrontLeft{
		height:             height,
		neckCircumference:  neck,
		chestCircumference: chest,
		waistCircumference: waist,
		hipCircumference:   hip,
	}
}

func (p *pn13TorsoFrontLeft) OnFold() bool { return false }
func (p *pn13TorsoFrontLeft) Mirrored() bool { return false }
func (p *pn13TorsoFrontLeft) CutCount() int { return 1 }

func (p *pn13TorsoFrontLeft) Details() *pieces.Details {
	return &pieces.Details{
		PieceNumber: 13,
		Description: "Torso Front Left",
	}
}

func (p *pn13TorsoFrontLeft) Stitch() *geometry.Block {
	return p.torsoFront().Stitch().MirrorHorizontally(p.mirrorX())
}

func (p *pn13TorsoFrontLeft) OuterCut() *geometry.Polyline {
	return p.torsoFront().OuterCut().MirrorHorizontally(p.mirrorX())
}

func (p *pn13TorsoFrontLeft) InnerCut() *geometry.Block {
	return p.torsoFront().InnerCut().MirrorHorizontally(p.mirrorX())
}

func (p *pn13TorsoFrontLeft) Ink() *geometry.Block {
	x := p.mirrorX()
	layer := p.torsoFront().Ink().MirrorHorizontally(x)

	for _, btn := range p.torsoFront().buttons() {
		hole := &symbols.ButtonHole{
			Centre: btn.MirrorHorizontally(x),
			Length: BUTTON_DIAMETER + 0.4,
			Angle: &geometry.Angle{Rads: math.Pi / 2.0},
		}
		layer.AddBlock(hole.Block())
	}

	return layer
}

func (p *pn13TorsoFrontLeft) Reference() *geometry.Block {
	return p.torsoFront().Reference().MirrorHorizontally(p.mirrorX())
}

func (p *pn13TorsoFrontLeft) torsoFront() *pn4TorsoFront {
	return &pn4TorsoFront{
		height:             p.height,
		neckCircumference:  p.neckCircumference,
		chestCircumference: p.chestCircumference,
		waistCircumference: p.waistCircumference,
		hipCircumference:   p.hipCircumference,
	}
}

func (p *pn13TorsoFrontLeft) mirrorX() float64 {
	return p.torsoFront().a().X
}

