package main

import (
	"errors"
	"fmt"
)

type BookList struct {
	bookLocationMap map[string]string
}

func (bl BookList) findBook(bookName string) string {
	return bl.bookLocationMap[bookName]
}

type LendingList struct {
	lendingMap map[string]bool
}

func (ll LendingList) check(bookName string) bool {
	return ll.lendingMap[bookName]
}

// facade
type Librarian struct {}

func (l Librarian) searchBook(bookName string) (location string, err error) {
	bookList := BookList{
		map[string]string{
			"昆虫図鑑": "2F",
			"国語辞典": "1F",
		},
	}
	location = bookList.findBook(bookName)
	if location != "" {
		lending := LendingList{
			map[string]bool{
				"昆虫図鑑": false,
				"国語辞典": true,
			},
		}
		lend := lending.check(bookName)
		if lend {
			return "", errors.New("貸出中です。")
		} else {
			return location, nil
		}
	}
	return "", errors.New("所蔵されていません。")
}

func main() {
	librarian := Librarian{}
	for _, bookName := range []string{"昆虫図鑑", "国語辞典"} {
		loc, err := librarian.searchBook(bookName)
		if err != nil {
			fmt.Printf("%sは%s\n", bookName, err)
		} else {
			fmt.Printf("%sは%sにあります。\n", bookName, loc)
		}
	}
}