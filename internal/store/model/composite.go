package model

type BookWithAuthors struct {
	Book    *Book
	Authors []*Author
}

type AuthorWithBooks struct {
	Author *Author
	Books  []*Book
}
