package pieces

import (
	"github.com/tailored-style/pattern-generator/geometry"
	"github.com/tailored-style/pattern-generator/pieces"
	"github.com/tailored-style/pattern-generator/symbols"
)

type pn12TorsoFrontRight struct {
	height             float64
	neckCircumference  float64
	chestCircumference float64
	waistCircumference float64
	hipCircumference   float64
}

func NewPN12TorsoFrontRight(height float64, neck float64, chest float64, waist float64, hip float64) pieces.Piece {
	return &pn12TorsoFrontRight{
		height:             height,
		neckCircumference:  neck,
		chestCircumference: chest,
		waistCircumference: waist,
		hipCircumference:   hip,
	}
}

func (p *pn12TorsoFrontRight) OnFold() bool { return false }
func (p *pn12TorsoFrontRight) Mirrored() bool { return false }
func (p *pn12TorsoFrontRight) CutCount() int { return 1 }

func (p *pn12TorsoFrontRight) Details() *pieces.Details {
	return &pieces.Details{
		PieceNumber: 12,
		Description: "Torso Front Right",
	}
}

func (p *pn12TorsoFrontRight) Stitch() *geometry.Block {
	return p.torsoFront().Stitch()
}

func (p *pn12TorsoFrontRight) OuterCut() *geometry.Polyline {
	return p.torsoFront().OuterCut()
}

func (p *pn12TorsoFrontRight) InnerCut() *geometry.Block {
	return p.torsoFront().InnerCut()
}

func (p *pn12TorsoFrontRight) Ink() *geometry.Block {
	layer := p.torsoFront().Ink()

	for _, btn := range p.torsoFront().buttons() {
		hole := &symbols.Button{
			Centre: btn,
			Diameter: BUTTON_DIAMETER,
		}
		layer.AddBlock(hole.Block())
	}

	return layer
}

func (p *pn12TorsoFrontRight) Reference() *geometry.Block {
	return p.torsoFront().Reference()
}

func (p *pn12TorsoFrontRight) torsoFront() *pn4TorsoFront {
	return &pn4TorsoFront{
		height:             p.height,
		neckCircumference:  p.neckCircumference,
		chestCircumference: p.chestCircumference,
		waistCircumference: p.waistCircumference,
		hipCircumference:   p.hipCircumference,
	}
}

