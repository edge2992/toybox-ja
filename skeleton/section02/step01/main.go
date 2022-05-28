package main

import "fmt"

func main() {

	// 問題文の表示
	fmt.Println("Q. 以下のコードコンパイルして実行するとどうなるか？")
	fmt.Println("package main")
	fmt.Println("func main() {")
	fmt.Println("	true := false")
	fmt.Println("	println(true == false)")
	fmt.Println("}")

	// 選択肢の表示
	fmt.Println("1: コンパイルエラー")
	fmt.Println("2: trueと表示される")
	fmt.Println("3: falseと表示される")
	fmt.Println("4: パニックが起きる")

	var answer int
	fmt.Println("回答>")
	fmt.Scanln(&answer)

	// 回答の表示
	fmt.Println("あなたの回答:", answer)
}
