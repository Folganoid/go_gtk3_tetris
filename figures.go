package main

import (
	"fmt"
	"github.com/gotk3/gotk3/cairo"
	"strconv"
	"strings"
)

/**
* XXX
*   X
 */
func fig1(cr *cairo.Context) *cairo.Context {

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
	figure1.color = [3]float64 {0,0,255}

	cr.SetSourceRGB(figure1.color[0], figure1.color[1], figure1.color[2])

	fmt.Println(rotate)
	fmt.Println(figure1.coords[rotate].positions)

	for _, cell := range figure1.coords[rotate].positions {

		cellCoords := strings.Split(cell, ".")
		tmpX, _ := strconv.ParseFloat(cellCoords[0], 64)
		tmpY, _ := strconv.ParseFloat(cellCoords[1], 64)
		cr.Rectangle(x*unitSize + ( tmpX * unitSize), y*unitSize + (tmpY * unitSize), unitSize, unitSize)

	}

	cr.Fill()
	return cr
}

/**
*  XXX
*   X
 */
func fig2(cr *cairo.Context) *cairo.Context {

	var coords0 = figureOnePos{}
	coords0.positions[0] = "0.0"
	coords0.positions[1] = "0.1"
	coords0.positions[2] = "0.2"
	coords0.positions[3] = "1.1"

	var coords1 = figureOnePos{}
	coords1.positions[0] = "0.0"
	coords1.positions[1] = "1.0"
	coords1.positions[2] = "2.0"
	coords1.positions[3] = "1.1"

	var coords2 = figureOnePos{}
	coords2.positions[0] = "1.0"
	coords2.positions[1] = "1.1"
	coords2.positions[2] = "1.2"
	coords2.positions[3] = "0.1"

	var coords3 = figureOnePos{}
	coords3.positions[0] = "1.0"
	coords3.positions[1] = "0.1"
	coords3.positions[2] = "1.1"
	coords3.positions[3] = "2.1"

	figure1 := figure{}
	figure1.coords[0] = coords0
	figure1.coords[1] = coords1
	figure1.coords[2] = coords2
	figure1.coords[3] = coords3
	figure1.color = [3]float64 {255,0,0}

	cr.SetSourceRGB(figure1.color[0], figure1.color[1], figure1.color[2])

	fmt.Println(rotate)
	fmt.Println(figure1.coords[rotate].positions)

	for _, cell := range figure1.coords[rotate].positions {

		cellCoords := strings.Split(cell, ".")
		tmpX, _ := strconv.ParseFloat(cellCoords[0], 64)
		tmpY, _ := strconv.ParseFloat(cellCoords[1], 64)
		cr.Rectangle(x*unitSize + ( tmpX * unitSize), y*unitSize + (tmpY * unitSize), unitSize, unitSize)

	}

	cr.Fill()
	return cr
}