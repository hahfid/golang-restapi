package repository

import "golang-restapi/internal/app/feature/library/book/domain"

type BookRepository interface {
	FindAll() ([]domain.Book, error)
	FindByID(id uint) (domain.Book, error)
	Create(book *domain.Book) error
	Update(book *domain.Book) error
	Delete(book *domain.Book) error
}
