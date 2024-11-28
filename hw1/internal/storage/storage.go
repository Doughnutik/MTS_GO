package storage

import (
	"hw1/internal/book"
	idgenerator "hw1/internal/idGenerator"
)

type Storage interface {
	Find(id uint) (book.Book, bool)
	Add(book book.Book)
	AddNum(id uint, num uint)
	RemoveNum(id uint, num uint)
	RemoveAll(id uint)
	Regenerate(generator idgenerator.Generator)
	GetBooks() []*book.Book
}

type MapStorage struct {
	storage map[uint]*book.Book
}

type SliceStorage struct {
	storage []book.Book
}

func NewMapStorage() MapStorage {
	return MapStorage{map[uint]*book.Book{}}
}

func (mapstore *MapStorage) Find(id uint) (book.Book, bool) {
	book, ok := mapstore.storage[id]
	return *book, ok
}

func (mapstore *MapStorage) Add(book book.Book) {
	if _, ok := mapstore.storage[book.GetId()]; !ok {
		mapstore.storage[book.GetId()] = &book
	}
}

func (mapstore *MapStorage) AddNum(id uint, num uint) {
	if _, ok := mapstore.storage[id]; ok {
		mapstore.storage[id].SetAmount(mapstore.storage[id].GetAmount() + num)
	}
}

func (mapstore *MapStorage) RemoveNum(id uint, num uint) {
	if _, ok := mapstore.storage[id]; ok {
		mapstore.storage[id].SetAmount(max(mapstore.storage[id].GetAmount()-num, 0))
		if mapstore.storage[id].GetAmount() == 0 {
			delete(mapstore.storage, id)
		}
	}
}

func (mapstore *MapStorage) RemoveAll(id uint) {
	delete(mapstore.storage, id)
}

func (mapstore *MapStorage) Regenerate(generator idgenerator.Generator) {
	newstorage := map[uint]*book.Book{}
	for _, value := range mapstore.storage {
		id := generator.Generate(value)
		newstorage[id] = value
	}
	mapstore.storage = newstorage
}

func (slicestore *SliceStorage) Find(id uint) (book.Book, bool) {
	for _, value := range slicestore.storage {
		if value.GetId() == id {
			return value, true
		}
	}
	return book.Book{}, false
}

func (slicestore *SliceStorage) Add(book book.Book) {
	for _, value := range slicestore.storage {
		if value == book {
			return
		}
	}
	slicestore.storage = append(slicestore.storage, book)
}

func (slicestore *SliceStorage) AddNum(id uint, num uint) {
	for _, value := range slicestore.storage {
		if value.GetId() == id {
			value.SetAmount(value.GetAmount() + num)
		}
	}
}

func (slicestore *SliceStorage) RemoveNum(id uint, num uint) {
	for i, value := range slicestore.storage {
		if value.GetId() == id {
			value.SetAmount(max(value.GetAmount()-num, 0))
			if value.GetAmount() == 0 {
				slicestore.storage[i] = slicestore.storage[len(slicestore.storage)-1]
				slicestore.storage = slicestore.storage[:len(slicestore.storage)-1]
			}
		}
	}
}

func (slicestore *SliceStorage) RemoveAll(id uint) {
	for i, value := range slicestore.storage {
		if value.GetId() == id {
			slicestore.storage[i] = slicestore.storage[len(slicestore.storage)-1]
			slicestore.storage = slicestore.storage[:len(slicestore.storage)-1]
		}
	}
}

func (slicestore *SliceStorage) Regenerate(generator idgenerator.Generator) {
	newstorage := []book.Book{}
	for _, value := range slicestore.storage {
		id := generator.Generate(&value)
		value.SetId(id)
		newstorage = append(newstorage, value)
	}
	slicestore.storage = newstorage
}

func (slicestore *SliceStorage) GetBooks() []*book.Book {
	books := []*book.Book{}
	for _, value := range slicestore.storage {
		books = append(books, &value)
	}
	return books
}

func (mapstore *MapStorage) GetBooks() []*book.Book {
	books := []*book.Book{}
	for _, value := range mapstore.storage {
		books = append(books, value)
	}
	return books
}
