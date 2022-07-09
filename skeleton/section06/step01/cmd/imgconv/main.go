package main

import (
	"flag"
	"fmt"
	"os"

	imgconv "github.com/gohandson/toybox-ja/skeleton/section06/step01"
	// TODO: imgconvパッケージをインポートする
)

var (
	flagTo   = imgconv.PNG  // TODO: パッケージ名をつける
	flagFrom = imgconv.JPEG // TODO: パッケージ名をつける
)

func init() {
	flag.Var(&flagTo, "to", "after format")
	flag.Var(&flagFrom, "from", "before format")
}

func main() {
	// TODO: convertAllはパッケージ名をつけてエクスポートされたものに変える
	if err := imgconv.ConvertAll(os.Args[1], flagFrom, flagTo); err != nil {
		fmt.Fprintln(os.Stderr, "エラー:", err)
		os.Exit(1)
	}
}
