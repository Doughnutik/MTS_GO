package idgenerator

import "hw1/internal/book"

type Generator interface {
	Generate(book *book.Book) uint
}

type HashGenerator struct {
	mod  uint
	base uint
}

type BasicGenerator struct {
	id uint
}

func NewDefaultHashGenerator() HashGenerator {
	return HashGenerator{1e9 + 7, 3001}
}

func NewHashGenerator(mod, base uint) HashGenerator {
	return HashGenerator{mod, base}
}

func NewBasicGenerator() BasicGenerator {
	return BasicGenerator{0}
}

func (hashgen *HashGenerator) Generate(book *book.Book) uint {
	var hash uint
	title := book.GetTitle()
	author := book.GetAuthor()
	for i := 0; i < len(title); i++ {
		hash = (hash + uint(title[i])*hashgen.base) % hashgen.mod
	}
	for i := 0; i < len(author); i++ {
		hash = (hash + uint(author[i])*hashgen.base) % hashgen.mod
	}
	return hash
}

func (basicgen *BasicGenerator) Generate(book *book.Book) uint {
	basicgen.id++
	return basicgen.id
}
