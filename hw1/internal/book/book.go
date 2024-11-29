package book

type Book struct {
	title, author string
	id            uint
}

func NewBook(title, author string) Book {
	return Book{title, author, 0}
}

func (book Book) GetTitle() string {
	return book.title
}

func (book Book) GetAuthor() string {
	return book.author
}

func (book Book) GetId() uint {
	return book.id
}

func (book *Book) SetTitle(title string) {
	book.title = title
}

func (book *Book) SetAuthor(author string) {
	book.author = author
}

func (book *Book) SetId(id uint) {
	book.id = id
}
