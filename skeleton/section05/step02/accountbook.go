package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 家計簿の項目
type Item struct {
	Category string
	Price    int
}

// 家計簿の処理を行う型
type AccountBook struct {
	file  string
	items []*Item
}

// 新しいAccountBookを作成する
func NewAccountBook(file string) *AccountBook {
	// AccountBook構造体を作成する
	ab := &AccountBook{
		file: file,
	}

	ab.readItems()

	// AccountBookのポインタを返す
	return ab
}

func (ab *AccountBook) readItems() {
	// ファイルを開き、結果を変数fと変数errに入れる
	f, err := os.Open(ab.file)

	if err != nil {
		fmt.Fprintln(os.Stderr, "エラー：", err)
		os.Exit(1)
	}

	// 関数終了時にファイルを閉じる
	defer f.Close()

	s := bufio.NewScanner(f)
	/* TODO: スキャナーを読み込む */
	for s.Scan() {
		// 1行分を,で分割する
		ss := strings.Split(s.Text(), ",")

		if len(ss) != 2 {
			fmt.Fprintln(os.Stderr, "ファイル形式が不正です")
			os.Exit(1)
		}

		// ss[1]をint型に変換し、結果を変数priceと変数errに入れる
		price, err := strconv.Atoi(ss[1])

		if err != nil {
			fmt.Fprintln(os.Stderr, "エラー：", err)
			os.Exit(1)
		}

		item := &Item{
			Category: ss[0],
			// Priceフィールドをセットする
			Price: price,
		}

		ab.AddItem(item)
	}

	if err := s.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "エラー：", err)
		os.Exit(1)
	}
}

// 新しいItemを追加する
func (ab *AccountBook) AddItem(item *Item) {
	ab.items = append(ab.items, item)
}

// 最近追加したものを最大limit件だけItemを取得する
func (ab *AccountBook) GetItems(limit int) []*Item {
	n := len(ab.items) - limit
	if n < 0 {
		n = 0
	}
	return ab.items[n:]
}

// 件数を取得する
func (ab *AccountBook) NumItems() int {
	return len(ab.items)
}
