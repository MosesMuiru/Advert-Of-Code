package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)
	var difference []int
	// var diffMatrix [][]int
	var safe int = 0
	var next int
	var value2 []int
	var check bool
	
func main(){
	data, err := os.ReadFile("test.txt")
	if err != nil{
		fmt.Println(err)
	}

	str := strings.Split(string(data), "\n")

	for _, value := range str[:len(str) - 1]{
		
		value2 = toInt(strings.Split(value, " "))
		if !(slices.IsSorted(value2) || slices.IsSorted(reverse(value2))) {
			checkUnsafe(value2)
			continue	
		}else{
		checkSafety(value2)
		}



	
	}
	fmt.Println("safe reports", safe )
}

func checkSafety(val2 []int){
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
				
				checkUnsafe(val2)
				break
			}
		}

			if check == true{
				safe ++
			}

			check = false
			difference = []int{}

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

func checkUnsafe(slice []int){
	fmt.Println("slice comming in  ---", slice)
	var newSLice []int
	for i := 0; i < len(slice); i ++  {	
		newSLice = removeElementFromSlice(slice, i)
		slice = newSLice
	fmt.Println("slice going out  ---", slice)
	}
}

func removeElementFromSlice(slice []int, index int) []int{
	var newSlice []int
	for i, val := range slice{
		if index != i {
			newSlice = append(newSlice, val)
		fmt.Println("this e new slice ---", newSlice)
		}
	}
	fmt.Println("this should remain the dame  ---", slice)
	return newSlice
}