package main

import (
	"github.com/gotk3/gotk3/cairo"
	"strconv"
	"strings"
)

type figure struct {
	coords [4]figureOnePos
	color int
}

type figureOnePos struct {
	positions [4] string
}

type figureColor struct {
	red float64
	green float64
	blue float64
}

type figureColors struct {
	list[8]figureColor
}

func initColors() figureColors{
	colors := figureColors{}

	color := figureColor{}
	color.red = 255.0
	color.green = 0.0
	color.blue = 0.0
	colors.list[1] = color

	color.red = 0
	color.green = 255
	color.blue = 0
	colors.list[2] = color

	color.red = 0
	color.green = 0
	color.blue = 255
	colors.list[3] = color

	color.red = 0
	color.green = 255
	color.blue = 255
	colors.list[4] = color

	color.red = 255
	color.green = 0
	color.blue = 255
	colors.list[5] = color

	color.red = 255
	color.green = 255
	color.blue = 0
	colors.list[6] = color

	color.red = 100
	color.green = 100
	color.blue = 100
	colors.list[7] = color

	return colors
}


func initFigures(colors figureColors) map[int]figure {

	figures := map[int]figure{}

	// XXX
	// X
	var coords0 = figureOnePos{}
	coords0.positions[0] = "0.0"
	coords0.positions[1] = "0.1"
	coords0.positions[2] = "0.2"
	coords0.positions[3] = "1.0"

	var coords1 = figureOnePos{}
	coords1.positions[0] = "0.0"
	coords1.positions[1] = "1.0"
	coords1.positions[2] = "2.0"
	coords1.positions[3] = "2.1"

	var coords2 = figureOnePos{}
	coords2.positions[0] = "1.0"
	coords2.positions[1] = "1.1"
	coords2.positions[2] = "1.2"
	coords2.positions[3] = "0.2"

	var coords3 = figureOnePos{}
	coords3.positions[0] = "0.0"
	coords3.positions[1] = "0.1"
	coords3.positions[2] = "1.1"
	coords3.positions[3] = "2.1"

	figure1 := figure{}
	figure1.coords[0] = coords0
	figure1.coords[1] = coords1
	figure1.coords[2] = coords2
	figure1.coords[3] = coords3
	figure1.color = 1

	figures[0] = figure1

	// XXX
	//  X
	coords0 = figureOnePos{}
	coords0.positions[0] = "0.0"
	coords0.positions[1] = "0.1"
	coords0.positions[2] = "0.2"
	coords0.positions[3] = "1.1"

	coords1 = figureOnePos{}
	coords1.positions[0] = "0.0"
	coords1.positions[1] = "1.0"
	coords1.positions[2] = "2.0"
	coords1.positions[3] = "1.1"

	coords2 = figureOnePos{}
	coords2.positions[0] = "1.0"
	coords2.positions[1] = "1.1"
	coords2.positions[2] = "1.2"
	coords2.positions[3] = "0.1"

	coords3 = figureOnePos{}
	coords3.positions[0] = "1.0"
	coords3.positions[1] = "0.1"
	coords3.positions[2] = "1.1"
	coords3.positions[3] = "2.1"

	figure1 = figure{}
	figure1.coords[0] = coords0
	figure1.coords[1] = coords1
	figure1.coords[2] = coords2
	figure1.coords[3] = coords3
	figure1.color = 2

	figures[1] = figure1

	// XXXX
	//
	coords0 = figureOnePos{}
	coords0.positions[0] = "1.0"
	coords0.positions[1] = "1.1"
	coords0.positions[2] = "1.2"
	coords0.positions[3] = "1.3"

	coords1 = figureOnePos{}
	coords1.positions[0] = "0.1"
	coords1.positions[1] = "1.1"
	coords1.positions[2] = "2.1"
	coords1.positions[3] = "3.1"

	coords2 = figureOnePos{}
	coords2.positions[0] = "1.0"
	coords2.positions[1] = "1.1"
	coords2.positions[2] = "1.2"
	coords2.positions[3] = "1.3"

	coords3 = figureOnePos{}
	coords3.positions[0] = "0.1"
	coords3.positions[1] = "1.1"
	coords3.positions[2] = "2.1"
	coords3.positions[3] = "3.1"

	figure1 = figure{}
	figure1.coords[0] = coords0
	figure1.coords[1] = coords1
	figure1.coords[2] = coords2
	figure1.coords[3] = coords3
	figure1.color = 3

	figures[2] = figure1

	// XXX
	//   X
	coords0 = figureOnePos{}
	coords0.positions[0] = "0.0"
	coords0.positions[1] = "0.1"
	coords0.positions[2] = "0.2"
	coords0.positions[3] = "1.2"

	coords1 = figureOnePos{}
	coords1.positions[0] = "0.0"
	coords1.positions[1] = "0.1"
	coords1.positions[2] = "1.0"
	coords1.positions[3] = "2.0"

	coords2 = figureOnePos{}
	coords2.positions[0] = "0.0"
	coords2.positions[1] = "1.0"
	coords2.positions[2] = "1.1"
	coords2.positions[3] = "1.2"

	coords3 = figureOnePos{}
	coords3.positions[0] = "0.1"
	coords3.positions[1] = "1.1"
	coords3.positions[2] = "2.1"
	coords3.positions[3] = "2.0"

	figure1 = figure{}
	figure1.coords[0] = coords0
	figure1.coords[1] = coords1
	figure1.coords[2] = coords2
	figure1.coords[3] = coords3
	figure1.color = 4

	figures[3] = figure1

	// XX
	// XX
	coords0 = figureOnePos{}
	coords0.positions[0] = "0.0"
	coords0.positions[1] = "0.1"
	coords0.positions[2] = "1.0"
	coords0.positions[3] = "1.1"

	coords1 = figureOnePos{}
	coords1.positions[0] = "0.0"
	coords1.positions[1] = "0.1"
	coords1.positions[2] = "1.0"
	coords1.positions[3] = "1.1"

	coords2 = figureOnePos{}
	coords2.positions[0] = "0.0"
	coords2.positions[1] = "0.1"
	coords2.positions[2] = "1.0"
	coords2.positions[3] = "1.1"

	coords3 = figureOnePos{}
	coords3.positions[0] = "0.0"
	coords3.positions[1] = "0.1"
	coords3.positions[2] = "1.0"
	coords3.positions[3] = "1.1"

	figure1 = figure{}
	figure1.coords[0] = coords0
	figure1.coords[1] = coords1
	figure1.coords[2] = coords2
	figure1.coords[3] = coords3
	figure1.color = 5

	figures[4] = figure1

	// XX
	//  XX
	coords0 = figureOnePos{}
	coords0.positions[0] = "0.0"
	coords0.positions[1] = "0.1"
	coords0.positions[2] = "1.1"
	coords0.positions[3] = "1.2"

	coords1 = figureOnePos{}
	coords1.positions[0] = "0.1"
	coords1.positions[1] = "1.0"
	coords1.positions[2] = "1.1"
	coords1.positions[3] = "2.0"

	coords2 = figureOnePos{}
	coords2.positions[0] = "0.0"
	coords2.positions[1] = "0.1"
	coords2.positions[2] = "1.1"
	coords2.positions[3] = "1.2"

	coords3 = figureOnePos{}
	coords3.positions[0] = "0.1"
	coords3.positions[1] = "1.0"
	coords3.positions[2] = "1.1"
	coords3.positions[3] = "2.0"

	figure1 = figure{}
	figure1.coords[0] = coords0
	figure1.coords[1] = coords1
	figure1.coords[2] = coords2
	figure1.coords[3] = coords3
	figure1.color = 6

	figures[5] = figure1

	//  XX
	// XX
	coords0 = figureOnePos{}
	coords0.positions[0] = "0.1"
	coords0.positions[1] = "0.2"
	coords0.positions[2] = "1.0"
	coords0.positions[3] = "1.1"

	coords1 = figureOnePos{}
	coords1.positions[0] = "0.0"
	coords1.positions[1] = "1.0"
	coords1.positions[2] = "1.1"
	coords1.positions[3] = "2.1"

	coords2 = figureOnePos{}
	coords2.positions[0] = "0.1"
	coords2.positions[1] = "0.2"
	coords2.positions[2] = "1.0"
	coords2.positions[3] = "1.1"

	coords3 = figureOnePos{}
	coords3.positions[0] = "0.0"
	coords3.positions[1] = "1.0"
	coords3.positions[2] = "1.1"
	coords3.positions[3] = "2.1"

	figure1 = figure{}
	figure1.coords[0] = coords0
	figure1.coords[1] = coords1
	figure1.coords[2] = coords2
	figure1.coords[3] = coords3
	figure1.color = 7

	figures[6] = figure1

	return figures

}

func drawFigure(cr *cairo.Context, figure1 figure) *cairo.Context {

	cr.SetSourceRGB(colors.list[figure1.color].red, colors.list[figure1.color].green, colors.list[figure1.color].blue)

	for _, cell := range figure1.coords[rotate].positions {

		cellCoords := strings.Split(cell, ".")
		tmpX, _ := strconv.Atoi(cellCoords[0])
		tmpY, _ := strconv.Atoi(cellCoords[1])
		cr.Rectangle(float64(x*unitSize + ( tmpX * unitSize)), float64(y*unitSize + (tmpY * unitSize)), float64(unitSize), float64(unitSize))
	}

	cr.Fill()
	return cr
}

func drawFigureNext(cr *cairo.Context, figure1 figure) *cairo.Context {

	cr.SetSourceRGB(colors.list[figure1.color].red, colors.list[figure1.color].green, colors.list[figure1.color].blue)

	for _, cell := range figure1.coords[1].positions {

		cellCoords := strings.Split(cell, ".")
		tmpX, _ := strconv.Atoi(cellCoords[0])
		tmpY, _ := strconv.Atoi(cellCoords[1])
		cr.Rectangle(float64(320 + ( tmpX * unitSize)), float64(320 + (tmpY * unitSize)), float64(unitSize), float64(unitSize))
	}

	cr.Fill()
	return cr
}