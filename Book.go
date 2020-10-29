package main

import (
	"fmt"
	"time"
)

func main() {
	book := CreateBook(func(builder *bookBuilder) {
		builder.
			SetName("Abai Zholy 4").
			SetBookPageCount(345).
			SetIssueDate(*NewDate(2012,03,23)).
			SetAuthor(*NewAuthor("Abay", "Kunanbayev",*NewDate(1856,04,12),*NewDate(1904,04,9)))
	})


	fmt.Println(book)
}

// Self data struct for 'yyyy-mm-dd'
type date struct {
	date time.Time
}

func NewDate(year int, month int, day int) *date {
	return &date{date: time.Date(year,time.Month(month),day,0,0,0,0, time.Local) }
}

// Author struct
type author struct {
	fName string
	sName string
	birthday date
	deadDay date
}

func NewAuthor(fName string, sName string, birthday date, deadYear date) *author {

	return &author{fName, sName, birthday, deadYear}
}


type buildBook func(builder *bookBuilder)
// Creating book builder
func CreateBook(action buildBook) *book{
	builder := &bookBuilder{}
	action(builder)
	return &builder.book
}

// Book struct
type book struct {
	id int
	name string
	bookPageCount int
	author author
	issueDate date
}

type bookBuilder struct {
	book book
}

//Setters for build book
func (b *bookBuilder) SetName(name string) *bookBuilder{
	b.book.name = name
	return b
}

func (b *bookBuilder) SetBookPageCount(bookPageCount int) *bookBuilder{
	b.book.bookPageCount = bookPageCount
	return b
}

func (b *bookBuilder) SetAuthor(author author) *bookBuilder{
	b.book.author = author
	return b
}

func (b *bookBuilder) SetIssueDate(issueDate date) *bookBuilder{
	b.book.issueDate = issueDate
	return b
}