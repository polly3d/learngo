package views

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

const (
	URL_VIEW             = "url"
	RESPONSE_HEADER_VIEW = "response-header"
	RESPONSE_BODY_VIEW   = "response-body"
)

var AllViews = map[string]View{
	URL_VIEW: View{
		Pos: Position{
			Point{0.0, 0}, Point{0.0, 0}, Point{1.0, -2}, Point{0.0, 3},
		},
		Prop: Property{
			Title:    "URL",
			Frame:    true,
			Editable: true,
			Wrap:     false,
			Editor:   gocui.DefaultEditor,
		},
	},
	RESPONSE_HEADER_VIEW: View{
		Pos: Position{
			Point{0.0, 0}, Point{0.0, 4}, Point{1.0, -2}, Point{0.25, 2},
		},
		Prop: Property{
			Title:    "Response Headers",
			Frame:    true,
			Editable: false,
			Wrap:     true,
		},
	},
	RESPONSE_BODY_VIEW: View{
		Pos: Position{
			Point{0.0, 0}, Point{0.25, 3}, Point{1.0, -2}, Point{1.0, -1},
		},
		Prop: Property{
			Title:    "Response Body",
			Frame:    true,
			Editable: false,
			Wrap:     true,
		},
	},
}

func Create() *gocui.Gui {
	g, err := gocui.NewGui(gocui.Output256)
	if err != nil {
		log.Panicln(err)
	}

	g.Cursor = true
	g.InputEsc = false
	g.BgColor = gocui.ColorDefault
	g.FgColor = gocui.ColorDefault
	g.SetManagerFunc(layout)
	return g
}

func layout(g *gocui.Gui) error {
	for key, value := range AllViews {
		if v, err := setView(g, key, value.Pos); err != nil {
			if err != gocui.ErrUnknownView {
				return err
			}
			setViewProperty(v, value.Prop)
		}
	}
	g.SetCurrentView("url")
	return nil
}

func setView(g *gocui.Gui, viewName string, pos Position) (*gocui.View, error) {
	maxX, maxY := g.Size()
	return g.SetView(viewName,
		pos.x0.GetCoordinate(maxX),
		pos.y0.GetCoordinate(maxY),
		pos.x1.GetCoordinate(maxX),
		pos.y1.GetCoordinate(maxY),
	)
}

func setViewProperty(v *gocui.View, prop Property) {
	v.Title = prop.Title
	v.Frame = prop.Frame
	v.Editable = prop.Editable
	v.Wrap = prop.Wrap
	v.Editor = prop.Editor
	SetViewTextAndCursor(v, prop.Text)
}

func SetViewTextAndCursor(v *gocui.View, s string) {
	v.Clear()
	fmt.Fprint(v, s)
	v.SetCursor(len(s), 0)
}
