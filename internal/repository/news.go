package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/nurzzaat/news/internal/models"
)

type NewsRepository struct {
	db *pgxpool.Pool
}

const (
	adminID = 2
)

func NewNewsRepository(db *pgxpool.Pool) models.NewsRepository {
	return &NewsRepository{db: db}
}

func (r *NewsRepository) Create(c context.Context, news models.NewsRequest, userID uint) (int, error) {
	var id int
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	query := `INSERT INTO news (title, content, category_id, thumbnail, created_at, author_id)
				VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;`
	if err := r.db.QueryRow(c, query, news.Title, news.Content, news.CategoryID, news.ThumbNail, currentTime, userID).Scan(&id); err != nil {
		return id, err
	}
	return id, nil
}
func (r *NewsRepository) Edit(c context.Context, news models.NewsRequest) error {
	query := `UPDATE news SET title = $1, content = $2, category_id = $3, thumbnail = $4 WHERE id = $5 and author_id != $6;`
	if _, err := r.db.Exec(c, query, news.Title, news.Content, news.CategoryID, news.ThumbNail, news.ID, adminID); err != nil {
		return err
	}
	return nil
}
func (r *NewsRepository) Delete(c context.Context, newsID int) error {
	query := `DELETE FROM news WHERE id = $1 and author_id != $2;`
	if _, err := r.db.Exec(c, query, newsID, adminID); err != nil {
		return err
	}
	return nil
}
func (r *NewsRepository) GetAll(c context.Context, orderByParameter string) ([]models.NewsResponse, error) {
	news := []models.NewsResponse{}
	var query string
	query = `SELECT id, title, content, category_id, thumbnail, views, created_at, author_id FROM news order by created_at desc;`
	if orderByParameter == models.Newer {
		query = `SELECT id, title, content, category_id, thumbnail, views, created_at, author_id FROM news order by created_at;`
	}
	if orderByParameter == models.Popular {
		query = `SELECT id, title, content, category_id, thumbnail, views, created_at, author_id FROM news order by views desc;`
	}
	rows, err := r.db.Query(c, query)
	if err != nil {
		return news, err
	}
	for rows.Next() {
		new := models.NewsResponse{}
		if err := rows.Scan(&new.ID, &new.Title, &new.Content, &new.Category.ID, &new.ThumbNail, &new.Views, &new.CreatedAt, &new.Author.ID); err != nil {
			return news, err
		}
		news = append(news, new)
	}
	return news, nil
}
func (r *NewsRepository) GetByID(c context.Context, newsID int) (models.NewsResponse, error) {
	new := models.NewsResponse{}
	query := `SELECT id, title, content, category_id, thumbnail, views, created_at, author_id FROM news where id = $1`
	if err := r.db.QueryRow(c, query, newsID).Scan(&new.ID, &new.Title, &new.Content, &new.Category.ID, &new.ThumbNail, &new.Views, &new.CreatedAt, &new.Author.ID); err != nil {
		return new, err
	}
	queryy := `SELECT id, name, created_at FROM category where id = $1`
	if err := r.db.QueryRow(c, queryy, new.Category.ID).Scan(&new.Category.ID, &new.Category.Name, &new.Category.CreatedAt); err != nil {
		return new, err
	}
	queryyy := `SELECT id, email, name FROM users where id = $1`
	row := r.db.QueryRow(c, queryyy, new.Author.ID)
	if err := row.Scan(&new.Author.ID, &new.Author.Email, &new.Author.Name); err != nil {
		return new, err
	}

	return new, nil
}

func ClearNews(c context.Context, db *pgxpool.Pool) error {
	query := `DELETE FROM news WHERE author_id != $1;`
	if _, err := db.Exec(c, query, adminID); err != nil {
		return err
	}
	return nil
}
