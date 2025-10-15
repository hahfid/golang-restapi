package mapper

import (
	"golang-restapi/internal/app/feature/library/book/domain"
	"golang-restapi/internal/app/feature/library/book/dto"
)

func ToBookResponse(book domain.Book) dto.BookResponse {
	return dto.BookResponse{
		ID:     book.ID,
		Title:  book.Title,
		Author: book.Author,
	}
}

func ToBookResponses(books []domain.Book) []dto.BookResponse {
	responses := make([]dto.BookResponse, 0, len(books))
	for _, b := range books {
		responses = append(responses, ToBookResponse(b))
	}
	return responses
}
