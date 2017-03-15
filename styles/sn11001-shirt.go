package styles

import (
	"github.com/tailored-style/pattern-generator/pieces"
	"github.com/tailored-style/pattern-generator/styles"
	piece_catalogue "github.com/tobyjsullivan/catalogue/pieces"
	"github.com/tobyjsullivan/catalogue/slopers"
)

const (
	SN11001_CUFF_DEPTH = 6.2
	SN11001_SLEEVE_PLACKET_OPENING = 10.2
)

type sn11001Shirt struct {
	*slopers.TorsoMeasurements
	pieces             []pieces.Piece
}

func NewSN11001Shirt(m *slopers.TorsoMeasurements) styles.Style {
	return &sn11001Shirt{
		TorsoMeasurements: m,
	}
}

func (p *sn11001Shirt) Pieces() []pieces.Piece {
	return []pieces.Piece{
		//piece_catalogue.NewPN4TorsoFront(
		//	p.height,
		//	p.neckCircumference,
		//	p.chestCircumference,
		//	p.waistCircumference,
		//	p.hipCircumference,
		//),
		piece_catalogue.NewPN12TorsoFrontRight(p.TorsoMeasurements),
		piece_catalogue.NewPN13TorsoFrontLeft(p.TorsoMeasurements),
		piece_catalogue.NewPN5TorsoBack(
			p.TorsoMeasurements.Height,
			p.TorsoMeasurements.ChestCircumference,
			p.TorsoMeasurements.BellyButtonWaistCircumference,
			p.TorsoMeasurements.HipCircumference,
		),
		piece_catalogue.NewPN6Yoke(p.TorsoMeasurements),
		piece_catalogue.NewPN7Sleeve(
			p.TorsoMeasurements,
			SN11001_CUFF_DEPTH,
			SN11001_SLEEVE_PLACKET_OPENING,
		),
		piece_catalogue.NewPN8Cuff(
			p.TorsoMeasurements.WristCircumference,
			SN11001_CUFF_DEPTH,
		),
		piece_catalogue.NewPN10Collar(p.TorsoMeasurements),
		piece_catalogue.NewPN9CollarBand(p.TorsoMeasurements),
		piece_catalogue.NewPN11SleevePlacket(
			SN11001_SLEEVE_PLACKET_OPENING,
		),
	}
}

func (p *sn11001Shirt) Details() *styles.Details {
	return &styles.Details{
		Description: "Tailored Shirt - v3.0 TEST",
		StyleNumber: "11001",
		Measurements: &pieces.Measurements{
			Height:             p.TorsoMeasurements.Height,
			NeckCircumference:  p.TorsoMeasurements.NeckCircumference,
			ChestCircumference: p.TorsoMeasurements.ChestCircumference,
			WaistCircumference: p.TorsoMeasurements.BellyButtonWaistCircumference,
			HipCircumference:   p.TorsoMeasurements.HipCircumference,
			SleeveLength:       p.TorsoMeasurements.ShirtSleeveLength,
			WristCircumference: p.TorsoMeasurements.WristCircumference,
		},
	}
}
