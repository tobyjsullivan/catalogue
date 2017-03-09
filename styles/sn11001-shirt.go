package styles

import (
	"github.com/tailored-style/pattern-generator/pieces"
	"github.com/tailored-style/pattern-generator/styles"
	piece_catalogue "github.com/tobyjsullivan/catalogue/pieces"
)

const (
	SN11001_CUFF_DEPTH = 6.2
	SN11001_SLEEVE_PLACKET_OPENING = 10.2
)

type sn11001Shirt struct {
	height             float64
	neckCircumference  float64
	chestCircumference float64
	waistCircumference float64
	hipCircumference   float64
	sleeveLength       float64
	wristCircumference float64
	pieces             []pieces.Piece
}

type SN11001Measurements struct {
	Height float64
	Neck   float64
	Chest  float64
	Waist  float64
	Hip    float64
	Sleeve float64
	Wrist  float64
}

func NewSN11001Shirt(m *SN11001Measurements) styles.Style {
	return &sn11001Shirt{
		height:             m.Height,
		neckCircumference:  m.Neck,
		chestCircumference: m.Chest,
		waistCircumference: m.Waist,
		hipCircumference:   m.Hip,
		sleeveLength:       m.Sleeve,
		wristCircumference: m.Wrist,
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
		piece_catalogue.NewPN12TorsoFrontRight(
			p.height,
			p.neckCircumference,
			p.chestCircumference,
			p.waistCircumference,
			p.hipCircumference,
		),
		piece_catalogue.NewPN13TorsoFrontLeft(
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
			SN11001_CUFF_DEPTH,
			SN11001_SLEEVE_PLACKET_OPENING,
		),
		piece_catalogue.NewPN8Cuff(
			p.wristCircumference,
			SN11001_CUFF_DEPTH,
		),
		piece_catalogue.NewPN10Collar(
			p.height,
			p.neckCircumference,
			p.chestCircumference,
			p.waistCircumference,
			p.hipCircumference,
		),
		piece_catalogue.NewPN9CollarBand(
			p.height,
			p.neckCircumference,
			p.chestCircumference,
			p.waistCircumference,
			p.hipCircumference,
		),
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
			Height:             p.height,
			NeckCircumference:  p.neckCircumference,
			ChestCircumference: p.chestCircumference,
			WaistCircumference: p.waistCircumference,
			HipCircumference:   p.hipCircumference,
			SleeveLength:       p.sleeveLength,
			WristCircumference: p.wristCircumference,
		},
	}
}
