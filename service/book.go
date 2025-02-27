package service

import (
	"gin-rest-api/model"
)

// dbの操作を行うBookService構造体
type BookService struct{}

func (BookService) SetBook(book *model.Book) error {
	_, err := DbEngine.Insert(book)
	if err != nil {
		return err
	}
	return nil
}

func (BookService) GetBookList() []model.Book {

	// Book構造体のスライスを作成
	tests := make([]model.Book, 0)

	// SELECT id, title, content FROM book LIMIT 10 OFFSET 0
	err := DbEngine.Distinct("id", "title", "content").Limit(10, 0).Find(&tests)
	if err != nil {
		panic(err)
	}
	return tests
}

func (BookService) UpdateBook(newBook *model.Book) error {
	_, err := DbEngine.Id(newBook.Id).Update(newBook)
	if err != nil {
		return err
	}
	return nil
}

func (BookService) DeleteBook(id int) error {
	book := new(model.Book)
	_, err := DbEngine.Id(id).Delete(book)
	if err != nil {
		return err
	}
	return nil
}
