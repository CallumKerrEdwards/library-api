package book

import (
	"github.com/shopspring/decimal"
)

// Book - representation of a book
type Book struct {
	ID     string
	Title  string
	Author string
	Series
}

// Series - representation of a series of books
type Series struct {
	Sequence decimal.Decimal
	Title    string
}
