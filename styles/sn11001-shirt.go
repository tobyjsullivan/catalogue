package styles

import (
	piece_catalogue "github.com/tobyjsullivan/catalogue/pieces"
	"github.com/tailored-style/pattern-generator/pieces"
	"github.com/tailored-style/pattern-generator/styles"
)

type sn11001Shirt struct {
	height float64
	neckCircumference float64
	chestCircumference float64
	waistCircumference float64
	hipCircumference float64
	sleeveLength float64
	wristCircumference float64
	pieces []pieces.Piece
}

func NewSN11001Shirt(height float64, neck float64, chest float64, waist float64, hip float64, sleeve float64, wrist float64) styles.Style {
	return &sn11001Shirt{
		height:height,
		neckCircumference:neck,
		chestCircumference:chest,
		waistCircumference:chest,
		hipCircumference:hip,
		sleeveLength:sleeve,
		wristCircumference: wrist,
	}
}

func (p *sn11001Shirt) Pieces() []pieces.Piece {
	return []pieces.Piece{
		piece_catalogue.NewPN4TorsoFront(
			p.height,
			p.neckCircumference,
			p.chestCircumference,
			p.waistCircumference,
			p.hipCircumference,
		),
		piece_catalogue.NewPN5TorsoBack(
			p.height,
			p.chestCircumference,
			p.waistCircumference,
			p.hipCircumference,
		),
		piece_catalogue.NewPN6Yoke(
			p.height,
			p.neckCircumference,
			p.chestCircumference,
			p.waistCircumference,
			p.hipCircumference,
		),
		piece_catalogue.NewPN7Sleeve(
			p.height,
			p.neckCircumference,
			p.chestCircumference,
			p.waistCircumference,
			p.hipCircumference,
			p.sleeveLength,
		),
		piece_catalogue.NewPN8Cuff(
			p.wristCircumference,
		),
	}
}

func (p *sn11001Shirt) Details() *styles.Details {
	return &styles.Details{
		Description: "Tailored Shirt - v3.0 TEST",
		StyleNumber: "11001",
		Measurements: &pieces.Measurements{
			Height: p.height,
			NeckCircumference: p.neckCircumference,
			ChestCircumference: p.chestCircumference,
			WaistCircumference: p.waistCircumference,
			HipCircumference: p.hipCircumference,
			SleeveLength: p.sleeveLength,
			WristCircumference: p.wristCircumference,
		},
	}
}
