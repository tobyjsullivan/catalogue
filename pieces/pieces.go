package pieces

import "github.com/tailored-style/pattern-generator/geometry"

func addAnchors(b *geometry.Block, anchors map[string]*geometry.Point) {
	for k, p := range anchors {
		addAnchor(b, k, p)
	}
}

func addAnchor(b *geometry.Block, label string, p *geometry.Point) {
	b.AddPoint(p)
	b.AddText(&geometry.Text{
		Content:  label,
		Position: p.Move(-1.5, -1.0),
	})
}
