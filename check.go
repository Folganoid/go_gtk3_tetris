package main

import (
	"strconv"
	"strings"
)

/**
check bottom position
 */
func checkPosDown(xCrd int, yCrd int, fig figure) bool {

	for _, cell := range fig.coords[rotate].positions {

		cellCoords := strings.Split(cell, ".")
		tmpX, _ := strconv.Atoi(cellCoords[0])
		tmpX += xCrd
		tmpY, _ := strconv.Atoi(cellCoords[1])
		tmpY += yCrd

		if tmpY + 1 >= fieldSizeY/unitSize || field[tmpY+1][tmpX] > 0 {
			addFigureToField(xCrd, yCrd, fig, rotate)
			changeSpeed()
			checkFillRows()
			setFigure(figuresArr)
			scores += 10

			return false
		}
	}

	return true
}

/**
* game over
 */
func checkFull(fig figure) bool {

	for _, cell := range fig.coords[rotate].positions {

		cellCoords := strings.Split(cell, ".")
		tmpX, _ := strconv.Atoi(cellCoords[0])
		tmpX += x
		tmpY, _ := strconv.Atoi(cellCoords[1])
		tmpY += y

		if field[tmpY][tmpX] > 0 {
			return false
		}
	}

	return true
}

/**
pre check left right position
 */
func checkLeftRight(xCrd int, yCrd int, fig figure, key string) bool {

	var add int
	if key == "right" {add = 1}
	if key == "left" {add = -1}


	for _, cell := range fig.coords[rotate].positions {

		cellCoords := strings.Split(cell, ".")
		tmpX, _ := strconv.Atoi(cellCoords[0])
		tmpX += xCrd
		tmpY, _ := strconv.Atoi(cellCoords[1])
		tmpY += yCrd

		if tmpX + add >= fieldSizeX/unitSize ||
			tmpX + add < 0 ||
			field[tmpY][tmpX + add] != 0 {
			return false
		}
	}

	return true
}

/**
precheck rotate position
 */
func checkRotate(xCrd int, yCrd int, fig figure, rotateFuture int) bool {

	if rotateFuture >= 5 {rotateFuture = 1}

	for _, cell := range fig.coords[rotateFuture].positions {

		cellCoords := strings.Split(cell, ".")
		tmpX, _ := strconv.Atoi(cellCoords[0])
		tmpX += xCrd
		tmpY, _ := strconv.Atoi(cellCoords[1])
		tmpY += yCrd

		if tmpX > fieldSizeX/unitSize -1 ||
			tmpY > fieldSizeY/unitSize -1 ||
			tmpX < 0 {return false}

		if field[tmpY][tmpX] != 0 {return false}
	}

	return true
}

/**
check filled rows
 */
func checkFillRows() {

	for i := 0 ; i < len(field) ; i++ {
		delete := true
		for z := 0 ; z < len(field[i]) ; z++ {
			if (field[i][z] == 0) {
				delete = false
				break
			}
		}

		if delete {
			deleteRow(i)
		}
	}
}

/**
delete filled row
 */
func deleteRow(row int) {
	for i := row ; i >= 1 ; i-- {
		field[i] = field[i-1]
	}
	changeLevel()
}

func changeLevel() {

	linesCount++
	scores += 100

	if linesCount%1 == 0 {


		gameLevel++
		changeSpeed()
	}
}

func changeSpeed() {
	speedFall = speedStart
	for i := 0; i < gameLevel; i++ {
		speedFall -= speedFall / speedLevelChangePercent
		if speedFall <= 25 {
			speedFall = 25
			break
		}
	}
}
