package keys

import (
	"learngo/wuzz/commands"

	"github.com/jroimartin/gocui"
)

func BindingKeys(g *gocui.Gui) {
	for viewName, keys := range DefaultKeys {
		if viewName == "global" {
			viewName = ""
		}

		for keyStr, commandStr := range keys {
			gKey := KEYS[keyStr]
			handler := commands.AllCommands[commandStr]
			g.SetKeybinding(viewName, gKey, gocui.ModNone, handler())
		}

	}
}
