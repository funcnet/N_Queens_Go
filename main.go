package main

import (
	"fmt"
	"math"
	"strconv"
)

const constQueensCount = 12

func main() {
	testEightQueens()
}

func testEightQueens() {
	var queenList [constQueensCount]int
	for i := 0; i < constQueensCount; i++ {
		queenList[i] = 0
	}

	putQueen(constQueensCount, queenList, 0)
}

func checkConflict(queenList [constQueensCount]int, nextY int) bool {
	for positionY := 0; positionY < nextY; positionY++ {
		if math.Abs(float64(queenList[positionY]-queenList[nextY])) == math.Abs(float64(positionY-nextY)) {
			return true
		}

		if queenList[positionY] == queenList[nextY] {
			return true
		}
	}
	return false
}

func joinArray(arr [constQueensCount]int) string {
	var ret = ""
	var length = len(arr)
	for i := 0; i < length-1; i++ {
		ret += strconv.Itoa(arr[i]) + ", "
	}
	ret += strconv.Itoa(arr[length-1])

	return ret
}

var count = 0

func putQueen(queenCount int, queenList [constQueensCount]int, nextY int) {
	for queenList[nextY] = 0; queenList[nextY] < queenCount; queenList[nextY]++ {
		if checkConflict(queenList, nextY) == false {
			nextY++
			if nextY < queenCount {
				putQueen(queenCount, queenList, nextY)
			} else {
				count++
				fmt.Println(strconv.Itoa(count) + ": " + joinArray(queenList))
			}
			nextY--
		}
	}
}
