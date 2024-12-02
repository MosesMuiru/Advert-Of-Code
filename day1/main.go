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
	// read the file

	data, err := os.ReadFile("input.txt")
	if err != nil{
		fmt.Println(err)
	}

	var left []int
	var right []int
	var err1 error
	var val int
	var total int

	e := strings.Fields(string(data))
	for index, value := range e{
			val, err1 = strconv.Atoi(value)
		if index == 0 || index % 2 == 0{
			left = append(left, val)
		}else{
			right = append(right, val)
		}
	}

	if err1 != nil{
		fmt.Println(err)
	}

	//  sort
	slices.Sort(left)
	slices.Sort(right)

	// minus
	for index, value := range left{
		total += int(math.Abs(float64(value) - float64(right[index])))

	}
	fmt.Println(total)
}