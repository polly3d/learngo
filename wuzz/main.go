package main

import (
	"learngo/wuzz/keys"
	"learngo/wuzz/views"
	"log"

	"github.com/jroimartin/gocui"
)

func main() {
	g := views.Create()
	defer g.Close()
	keys.BindingKeys(g)

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
