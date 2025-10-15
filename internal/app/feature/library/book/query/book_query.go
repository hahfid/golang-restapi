package query

import (
	"golang-restapi/internal/app/feature/library/book/domain"

	"gorm.io/gorm"
)

type BookQuery struct {
	db *gorm.DB
}

func NewBookQuery(db *gorm.DB) *BookQuery {
	return &BookQuery{db: db}
}

func (q *BookQuery) FindAll() ([]domain.Book, error) {
	var books []domain.Book
	err := q.db.Find(&books).Error
	return books, err
}

func (q *BookQuery) FindByID(id uint) (domain.Book, error) {
	var book domain.Book
	err := q.db.First(&book, id).Error
	return book, err
}

func (q *BookQuery) Create(book *domain.Book) error {
	return q.db.Create(book).Error
}

func (q *BookQuery) Update(book *domain.Book) error {
	return q.db.Save(book).Error
}

func (q *BookQuery) Delete(book *domain.Book) error {
	return q.db.Delete(book).Error
}
