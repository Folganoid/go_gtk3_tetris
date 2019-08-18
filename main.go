package main

import (
	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

const (
	KEY_LEFT  uint = 65361
	KEY_UP    uint = 65362
	KEY_RIGHT uint = 65363
	KEY_DOWN  uint = 65364
)

var unitSize = 20.0
var x = 5.0
var y = 5.0
var rotate = 0

type figure struct {
	coords [4]figureOnePos
	color  [3]float64
}

type figureOnePos struct {
	positions [4] string
}

func main() {
	gtk.Init(nil)

	// gui boilerplate
	win, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	da, _ := gtk.DrawingAreaNew()
	win.SetDefaultSize(260, 600)
	win.Add(da)
	win.SetTitle("tetris")
	win.Connect("destroy", gtk.MainQuit)
	win.ShowAll()

	// Data
	keyMap := map[uint]func(){
		KEY_LEFT:  func() { x-- },
		KEY_UP:    func() {
			if rotate >= 3 {
				rotate = 0
			} else {
				rotate++
			}
		},
		KEY_RIGHT: func() { x++ },
		KEY_DOWN:  func() {
			if rotate <= 0 {
				rotate = 3
			} else {
				rotate--
			}
		},
	}

	// Event handlers
	da.Connect("draw", func(da *gtk.DrawingArea, cr *cairo.Context) {
		cr = fig2(cr)
	})
	win.Connect("key-press-event", func(win *gtk.Window, ev *gdk.Event) {
		keyEvent := &gdk.EventKey{ev}
		if move, found := keyMap[keyEvent.KeyVal()]; found {
			move()
			win.QueueDraw()
		}
	})

	gtk.Main()
}