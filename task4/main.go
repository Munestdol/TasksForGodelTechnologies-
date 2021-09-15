package main

import (
	"fmt"
	"math"
)

func main() {
	arr := [6]int{17,18,5,4,6,1}
	check := false

	if len(arr)<int(math.Pow(10,4)) || len(arr)>0{
		for i:=0; i < len(arr); i++{
			if arr[i]>int(math.Pow(10,5)) || arr[i]<1{
				check = true
				break
			}
		}
		if check != true{
			fmt.Println("Input array \t",arr)
			for i := 0; i < len(arr); i++{
				buff := arr[i]
				for j := 0; j+i+1 < len(arr); j++{
					if buff < arr[j+i+1]{
						buff = arr[j+i+1]
					}
				}
				arr[i] = buff
			}
			arr[len(arr)-1] = -1
			fmt.Println("Output array \t",arr)
		}else {
			fmt.Println("Array value is out of range")
		}
	}else {
		fmt.Println("Array length is out of range")
	}
}



