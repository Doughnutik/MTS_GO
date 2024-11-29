package main

import (
	"fmt"
	idgenerator "hw1/internal/idGenerator"
	"hw1/internal/library"
	"hw1/internal/storage"
)

func main() {
	books := [][2]string{
		{"Война и мир", "Лев Толстой"},
		{"Преступление и наказание", "Фёдор Достоевский"},
		{"Анна Каренина", "Лев Толстой"},
		{"Мастер и Маргарита", "Михаил Булгаков"},
		{"1984", "Джордж Оруэлл"},
		{"Убить пересмешника", "Харпер Ли"},
		{"Над пропастью во ржи", "Джером Сэлинджер"},
		{"Тихий Дон", "Михаил Шолохов"},
		{"Собачье сердце", "Михаил Булгаков"},
		{"Мы", "Евгений Замятин"},
	}

	library := library.NewLibrary(storage.NewMapStorage(), idgenerator.NewCRC32())
	library.Add(books[0][0], books[0][1])
	library.Add(books[1][0], books[1][1])
	library.Add(books[2][0], books[2][1])

	book, ok := library.Find(books[0][0], books[0][1])
	if ok {
		fmt.Println(book) // {Война и мир Лев Толстой}
	} else {
		fmt.Println("no book")
	}

	book, ok = library.Find(books[2][0], books[2][1])
	if ok {
		fmt.Println(book) // {Анна Каренина Лев Толстой}
	} else {
		fmt.Println("no book")
	}

	book, ok = library.Find(books[7][0], books[7][1])
	if ok {
		fmt.Println(book)
	} else {
		fmt.Println("no book") // no book
	}

	library.Remove(books[0][0], books[0][1])
	book, ok = library.Find(books[0][0], books[0][1])
	if ok {
		fmt.Println(book)
	} else {
		fmt.Println("no book") // no book
	}

	library.Regenerate(idgenerator.NewDefaultHashGenerator())

	book, ok = library.Find(books[1][0], books[1][1])
	if ok {
		fmt.Println(book) // {Преступление и наказание Фёдор Достоевский}
	} else {
		fmt.Println("no book")
	}

	book, ok = library.Find(books[2][0], books[2][1])
	if ok {
		fmt.Println(book) // {Анна Каренина Лев Толстой}
	} else {
		fmt.Println("no book")
	}

	book, ok = library.Find(books[5][0], books[5][1])
	if ok {
		fmt.Println(book)
	} else {
		fmt.Println("no book") // no book
	}

	library.ChangeStorage(storage.NewSliceStorage())
	library.Add(books[7][0], books[7][1])
	library.Add(books[8][0], books[8][1])
	library.Add(books[9][0], books[9][1])

	book, ok = library.Find(books[0][0], books[0][1])
	if ok {
		fmt.Println(book)
	} else {
		fmt.Println("no book") // no book
	}

	book, ok = library.Find(books[8][0], books[8][1])
	if ok {
		fmt.Println(book) // {"Собачье сердце", "Михаил Булгаков"}
	} else {
		fmt.Println("no book")
	}

	library.Remove(books[8][0], books[8][1])
	book, ok = library.Find(books[8][0], books[8][1])
	if ok {
		fmt.Println(book)
	} else {
		fmt.Println("no book") // no book
	}
}
