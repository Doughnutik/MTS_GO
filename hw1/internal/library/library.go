package library

import (
	"hw1/internal/book"
	idgenerator "hw1/internal/idGenerator"
	"hw1/internal/storage"
)

type Library interface {
	Find(title string, author string) (book.Book, bool)
	Add(title string, author string)
	Remove(title string, author string)
	Regenerate(generator idgenerator.Generator)
	ChangeStorage(storage storage.Storage)
}

type MyLibrary struct {
	storage   storage.Storage
	generator idgenerator.Generator
}

func NewLibrary(storage storage.Storage, generator idgenerator.Generator) MyLibrary {
	return MyLibrary{storage, generator}
}

func (lib MyLibrary) Find(title string, author string) (book.Book, bool) {
	book := book.NewBook(title, author)
	book.SetId(lib.generator.Generate(book))
	return lib.storage.Find(book.GetId())
}

func (lib *MyLibrary) Add(title string, author string) {
	book := book.NewBook(title, author)
	book.SetId(lib.generator.Generate(book))
	lib.storage.Add(book)
}

func (lib *MyLibrary) Remove(title string, author string) {
	book := book.NewBook(title, author)
	book.SetId(lib.generator.Generate(book))
	lib.storage.Remove(book.GetId())
}

func (lib *MyLibrary) Regenerate(generator idgenerator.Generator) {
	lib.generator = generator
	lib.storage.Regenerate(lib.generator)
}

func (lib *MyLibrary) ChangeStorage(storage storage.Storage) {
	books := lib.storage.GetBooks()
	lib.storage = storage
	for _, book := range books {
		lib.Add(book.GetTitle(), book.GetAuthor())
	}
}