package pieces

import (
	"github.com/tailored-style/pattern-generator/geometry"
	"github.com/tailored-style/pattern-generator/pieces"
	"github.com/tailored-style/pattern-generator/symbols"
	"math"
	"github.com/tobyjsullivan/catalogue/slopers"
)

type pn13TorsoFrontLeft struct {
	*slopers.TorsoMeasurements
}

func NewPN13TorsoFrontLeft(torso *slopers.TorsoMeasurements) pieces.Piece {
	return &pn13TorsoFrontLeft{
		TorsoMeasurements: torso,
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
		TorsoMeasurements: p.TorsoMeasurements,
	}
}

func (p *pn13TorsoFrontLeft) mirrorX() float64 {
	return p.torsoFront().a().X
}

