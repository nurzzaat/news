package models

import "context"

type CategoryResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
}

type CategoryRequest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CategoryRepository interface {
	Create(c context.Context, category CategoryRequest) (int, error)
	Edit(c context.Context, category CategoryRequest) error
	Delete(c context.Context, categoryID int) error
	GetAll(c context.Context) ([]CategoryResponse, error)
	GetByID(c context.Context, categoryID int) (CategoryResponse, error)
}
