package main

import (
	"fmt"
	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	"math/rand"
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

var x = 0
var y = 10
var speed = 500
var rotate = 0
var figureNow figure
var figuresArr  map[int]figure
var colors figureColors

var field = [fieldSizeY/unitSize][fieldSizeX/unitSize]int{}


func checkPosDown(xCrd int, yCrd int, fig figure) bool {

	for _, cell := range fig.coords[rotate].positions {

		cellCoords := strings.Split(cell, ".")
		tmpX, _ := strconv.Atoi(cellCoords[0])
		tmpX += xCrd
		tmpY, _ := strconv.Atoi(cellCoords[1])
		tmpY += yCrd

		if tmpY + 1 >= fieldSizeY/unitSize || field[tmpY+1][tmpX] > 0 {
			addFigureToField(xCrd, yCrd, fig, rotate)
			speed = 500
			setFigure(figuresArr)
			return false
		}
	}

	return true

}

/**
* add figure to field
 */
func addFigureToField(xCrd int, yCrd int, fig figure, rotate int) {
	for _, cell := range fig.coords[rotate].positions {

		cellCoords := strings.Split(cell, ".")
		tmpX, _ := strconv.Atoi(cellCoords[0])
		tmpX += xCrd
		tmpY, _ := strconv.Atoi(cellCoords[1])
		tmpY += yCrd
		fmt.Println(tmpY, tmpX)
		field[tmpY][tmpX] = fig.color
	}

	fmt.Println(figure{})
}

func drawField(cr *cairo.Context) *cairo.Context {


	for i:=0; i<len(field);i++{
		for z:=0; z<len(field[i]);z++{
			if field[i][z] > 0 {
				cr.SetSourceRGB(colors.list[field[i][z]].red, colors.list[field[i][z]].green, colors.list[field[i][z]].blue)
				cr.Rectangle(float64(z*unitSize), float64(i*unitSize), float64(unitSize), float64(unitSize))
				fmt.Println(colors.list[field[i][z]].red, colors.list[field[i][z]].blue, colors.list[field[i][z]].blue)
				cr.Fill()
			}
		}
	}
	return cr
}

func setFigure(figuresArr map[int]figure) {
	figureNow = figuresArr[rand.Intn(len(figuresArr) - 0) + 0]
}

func main() {

	gtk.Init(nil)

	//init colors
	colors = initColors()

	// init figures
	figuresArr = initFigures(colors)

	setFigure(figuresArr)

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
			speed = 25
		},
	}

	// Event handlers
	da.Connect("draw", func(da *gtk.DrawingArea, cr *cairo.Context) {
		cr = drawFigure(cr, figureNow)
		cr = drawField(cr)
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
				x = 0
				y = 10
			}

			time.Sleep(time.Duration(speed) * time.Millisecond)
			y++
			win.QueueDraw()

		}
	}()



	gtk.Main()
}