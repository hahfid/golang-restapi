package repository

import "golang-restapi/internal/app/feature/library/book/query"

import "golang-restapi/internal/app/feature/library/book/domain"

type gormRepository struct {
	query *query.BookQuery
}

func NewGormRepository(q *query.BookQuery) BookRepository {
	return &gormRepository{query: q}
}

func (r *gormRepository) FindAll() ([]domain.Book, error) {
	return r.query.FindAll()
}

func (r *gormRepository) FindByID(id uint) (domain.Book, error) {
	return r.query.FindByID(id)
}

func (r *gormRepository) Create(book *domain.Book) error {
	return r.query.Create(book)
}

func (r *gormRepository) Update(book *domain.Book) error {
	return r.query.Update(book)
}

func (r *gormRepository) Delete(book *domain.Book) error {
	return r.query.Delete(book)
}
