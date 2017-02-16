package styles

import (
	piece_catalogue "github.com/tobyjsullivan/catalogue/pieces"
	"github.com/tailored-style/pattern-generator/pieces"
	"github.com/tailored-style/pattern-generator/styles"
)

type SN11001Shirt struct {
	*pieces.Measurements
	pieces []pieces.Piece
}

func (p *SN11001Shirt) Pieces() []pieces.Piece {
	return []pieces.Piece{
		&piece_catalogue.PN4TorsoFront{
			Measurements: p.Measurements,
		},
		&piece_catalogue.PN5TorsoBack{
			Measurements: p.Measurements,
		},
		&piece_catalogue.PN6Yoke{
			Measurements: p.Measurements,
		},
	}
}

func (p *SN11001Shirt) Details() *styles.Details {
	return &styles.Details{
		Description: "Tailored Shirt - v3.0 TEST",
		StyleNumber: "11001",
		Measurements: p.Measurements,
	}
}
