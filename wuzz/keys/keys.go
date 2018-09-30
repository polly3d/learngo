package keys

import (
	"github.com/jroimartin/gocui"
)

var KEYS = map[string]gocui.Key{
	"Enter": gocui.KeyEnter,
	"CtrlC": gocui.KeyCtrlC,
}

var DefaultKeys = map[string]map[string]string{
	"global": {
		"CtrlC": "quit",
	},
	"url": {
		"Enter": "submit",
	},
}
