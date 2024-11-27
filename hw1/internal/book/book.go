package book

type Book struct {
	title, author     string
	pages, amount, id uint
}

func NewBook(title, author string, pages uint) Book {
	return Book{title, author, pages, 1, 0}
}

func (book Book) GetTitle() string {
	return book.title
}

func (book Book) GetAuthor() string {
	return book.author
}

func (book Book) GetPages() uint {
	return book.pages
}

func (book Book) GetAmount() uint {
	return book.amount
}

func (book Book) GetId() uint {
	return book.pages
}

func (book *Book) SetTitle(title string) {
	book.title = title
}

func (book *Book) SetAuthor(author string) {
	book.author = author
}

func (book *Book) SetPages(pages uint) {
	book.pages = pages
}

func (book *Book) SetAmount(amount uint) {
	book.amount = amount
}

func (book *Book) SetId(id uint) {
	book.id = id
}
