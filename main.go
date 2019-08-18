package main

import (
	"fmt"
	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	"strconv"
	"strings"
	"time"
)

const (
	KEY_LEFT  uint = 65361
	KEY_UP    uint = 65362
	KEY_RIGHT uint = 65363
	KEY_DOWN  uint = 65364
)

const unitSize = 20
const fieldSizeX = 260
const fieldSizeY = 600
var x = 5
var y = 15
var rotate = 0
var figureNow figure

type figure struct {
	coords [4]figureOnePos
	color  [3]float64
}

type figureOnePos struct {
	positions [4] string
}

var field = [fieldSizeX/unitSize][fieldSizeY/unitSize]bool{}


func checkPosDown(xCrd int, yCrd int, fig figure) bool {

	for _, cell := range fig.coords[rotate].positions {

		cellCoords := strings.Split(cell, ".")
		tmpX, _ := strconv.Atoi(cellCoords[0])
		tmpX += xCrd
		tmpY, _ := strconv.Atoi(cellCoords[1])
		tmpY += yCrd

		if tmpY + 1 >= fieldSizeY/unitSize || field[tmpX][tmpY] == true {
			return false
		}
	}

	return true

}

func main() {

	gtk.Init(nil)

	// gui boilerplate
	win, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	da, _ := gtk.DrawingAreaNew()
	win.SetDefaultSize(fieldSizeX, fieldSizeY)
	win.Add(da)
	win.SetTitle("Tetris")
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

	//moving down
	go func() {
		for {
			if figureNow == (figure{}) {
				continue
			}

			if  !checkPosDown(x,y,figureNow) {
				fmt.Println("+++++++++++++")
				break
			}

			time.Sleep(time.Second)
			y++
			win.QueueDraw()

		}
	}()



	gtk.Main()
}