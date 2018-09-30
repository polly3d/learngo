package commands

import (
	"fmt"
	"io/ioutil"
	"learngo/wuzz/views"
	"net/http"
	"strings"

	"github.com/jroimartin/gocui"
)

type CommandFunc func(*gocui.Gui, *gocui.View) error

var AllCommands map[string]func() CommandFunc = map[string]func() CommandFunc{
	"submit": func() CommandFunc {
		return submitRequest
	},
	"quit": func() CommandFunc {
		return quit
	},
}

func submitRequest(g *gocui.Gui, _ *gocui.View) error {
	vrh, _ := g.View(views.RESPONSE_HEADER_VIEW)
	vrh.Clear()
	vrb, _ := g.View(views.RESPONSE_BODY_VIEW)
	vrb.Clear()

	//TODO: popup

	//get url values
	v, _ := g.View(views.URL_VIEW)
	urlValue := strings.TrimSpace(v.Buffer())

	//start request
	response, _ := http.Get(urlValue)

	for headerKey, headerValue := range response.Header {
		fmt.Fprintf(vrh, "%s:%s\n", headerKey, headerValue)
	}

	resBody, _ := ioutil.ReadAll(response.Body)
	fmt.Fprintln(vrb, "%s", string(resBody))

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
