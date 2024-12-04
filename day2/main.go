package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main(){
	// read file
	data, err := os.ReadFile("input.txt")
	if err != nil{
		fmt.Println(err)
	}

	// str := strings.Fields(string(data))
	str := strings.Split(string(data), "\n")
	// var total string
	var difference []int
	// var diffMatrix [][]int
	var safe int = 0
	var next int
	var val2 []int
	// var valInt int
	// var valNext int
	var check bool
	// var safety []string

	// str = slices.DeleteFunc(str, func(i []any) bool {
	// 	return len(i) == 0
	// })
	for _, value := range str[:len(str) - 1]{
		// fmt.Println(strings.Split(value, ""))

		
		val2 = toInt(strings.Split(value, " "))

		fmt.Print(val2)
		if !(slices.IsSorted(val2) || slices.IsSorted(reverse(val2))) {
		fmt.Println("val2 that is invalid", val2)
			continue	
		}
		for i:= 0; i < len(val2) - 1; i ++{
			check = false
			next = i + 1
			difference = append(difference, int(math.Abs(float64(val2[i]) - float64(val2[next]))))
			}
			for _, val := range difference{
				if val < 4 && val > 0 {
					check = true
				}else {
				check = false
				break
			}
		}

			if check == true{
				safe ++
			}
			check = false
			difference = []int{}
	}
	fmt.Println("safe reports", safe )
}

func reverse(data []int) []int{
	var newData []int
	for i := len(data) - 1; i >= 0; i --{
		newData = append(newData, data[i])
	}
	return newData
}

func toInt(data []string) []int{
	var newData []int
	for i := 0; i < len(data); i ++{
		val, _:= strconv.Atoi(data[i])
		newData = append(newData, val)
		
	}
	return newData
}
