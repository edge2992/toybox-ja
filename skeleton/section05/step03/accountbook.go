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
	f, err := os.Open(ab.file)
	if err != nil {
		fmt.Fprintln(os.Stderr, "エラー：", err)
		os.Exit(1)
	}

	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		ss := strings.Split(s.Text(), ",")
		if len(ss) != 2 {
			fmt.Fprintln(os.Stderr, "ファイル形式が不正です")
			os.Exit(1)
		}

		price, err := strconv.Atoi(ss[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, "エラー：", err)
			os.Exit(1)
		}

		item := &Item{
			Category: ss[0],
			Price:    price,
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

	// 書き込み用でファイルを開き、結果を変数fと変数errに入れる
	f, err := os.Create(ab.file)

	if err != nil {
		fmt.Fprintln(os.Stderr, "エラー：", err)
		os.Exit(1)
	}

	for _, item := range ab.items {
		// 品目,値段の形式でファイルに出力する
		// 2つめの結果だけ変数errに入れる
		_, err = fmt.Fprintf(f, "%s,%d\n", item.Category, item.Price)

		if err != nil {
			fmt.Fprintln(os.Stderr, "エラー：", err)
			os.Exit(1)
		}
	}

	if err := f.Close(); err != nil {
		fmt.Fprintln(os.Stderr, "エラー：", err)
		os.Exit(1)
	}
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
