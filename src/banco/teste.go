package main

import (
	"fmt"
)

func main()  {
	array := []int{2,1,6,5,8} 
	array = append(array,500)
	fmt.Println(array)
}


