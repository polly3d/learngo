package views

import (
	"github.com/jroimartin/gocui"
)

type Point struct {
	Pect float32
	Abs  int
}

type Position struct {
	x0, y0, x1, y1 Point
}

type Property struct {
	Title    string
	Frame    bool
	Editable bool
	Wrap     bool
	Editor   gocui.Editor
	Text     string
}

type View struct {
	Pos  Position
	Prop Property
}

func (p *Point) GetCoordinate(max int) int {
	return int(p.Pect*float32(max)) + p.Abs
}
