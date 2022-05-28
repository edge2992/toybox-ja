package main

import (
	"fmt"
)

func main() {
	// 山札を作る
	var all []int = make([]int, 0, 13)

	for n := 2; n <= 14; n++ { // 14がAを表す
		all = append(all, n)
	}

	// 山札を表示させる
	for i, n := range all {
		fmt.Printf("%2d番目: ", i+1)

		switch n {
		case 11:
			fmt.Println("J")
		case 12:
			fmt.Println("Q")
		case 13:
			fmt.Println("K")
		case 14:
			fmt.Println("A")
		default:
			fmt.Println(n)
		}
	}
}
