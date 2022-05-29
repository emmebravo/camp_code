package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

func main(){
	//sort given
	s := []int{345, 78, 123, 10, 76, 2, 567, 5}
	sort.Ints(s)
	fmt.Println("sorted slice: ", s)

	//find index
	fmt.Println("index val: ", strings.Index("The foundations of decay", "ay"))

	//find time
	fmt.Println("time: ", time.Now().Unix())
}