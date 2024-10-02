package models

import "context"

type NewsResponse struct {
	ID        uint             `json:"id"`
	Title     string           `json:"title"`
	Content   string           `json:"content"`
	Category  CategoryResponse `json:"category"`
	Views     int              `json:"views"`
	ThumbNail string           `json:"thumbnail"`
	CreatedAt string           `json:"createdAt"`
	Author    UserResponse     `json:"author"`
}

type NewsRequest struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	CategoryID int    `json:"categoryId"`
	ThumbNail  string `json:"thumbnail"`
}

type NewsRepository interface {
	Create(c context.Context, news NewsRequest, userID uint) (int, error)
	Edit(c context.Context, news NewsRequest) error
	Delete(c context.Context, newsID int) error
	GetAll(c context.Context, orderByParameter string) ([]NewsResponse, error)
	GetByID(c context.Context, newsID int) (NewsResponse, error)
}
