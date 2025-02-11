package other

import (
	"fmt"
	"strconv"
)

const floorNum int = 3

type House struct {
	floor int
}

func NewHouse(floor int) House {
	return House{floor: floor}
}

func (h *House) Print() {
	fmt.Printf("%s\n %d этажей.", buildFloor(h.floor), h.floor)
}

func indent() string {
	temp := ""
	for i := 0; i < floorNum; i++ {
		temp += "="
	}
	return temp + "\n"
}

func buildFloor(num int) string {
	temp := ""
	for i := 1; i < num; i++ {
		for j := 0; j < floorNum; j++ {
			for x := 0; x < floorNum; x++ {
				if x != floorNum/2 {
					temp += "*"
					if x == floorNum-1 {
						temp += "\n"
					}
				} else {
					if j != floorNum/2 {
						temp += "*"
					} else {
						temp += strconv.Itoa(i)
					}
				}
			}
		}
		temp += indent()
	}
	return temp
}
