package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/nurzzaat/news/internal/models"
)

type CategoryRepository struct {
	db *pgxpool.Pool
}

func NewCategoryRepository(db *pgxpool.Pool) models.CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) Create(c context.Context, category models.CategoryRequest, userID uint) (int, error) {
	var id int
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	query := `INSERT INTO category (name, created_at, author_id) VALUES ($1, $2, $3) RETURNING id;`
	if err := r.db.QueryRow(c, query, category.Name, currentTime, userID).Scan(&id); err != nil {
		return id, err
	}
	return id, nil
}
func (r *CategoryRepository) Edit(c context.Context, category models.CategoryRequest) error {
	query := `UPDATE category SET name = $1 WHERE id = $2  and author_id != $2;`
	if _, err := r.db.Exec(c, query, category.Name, category.ID); err != nil {
		return err
	}
	return nil
}
func (r *CategoryRepository) Delete(c context.Context, categoryID int) error {
	query := `DELETE FROM category WHERE id = $1 and author_id != $2;`
	if _, err := r.db.Exec(c, query, categoryID); err != nil {
		return err
	}
	return nil
}
func (r *CategoryRepository) GetAll(c context.Context) ([]models.CategoryResponse, error) {
	categories := []models.CategoryResponse{}
	query := `SELECT id, name, created_at FROM category order by created_at desc;`
	rows, err := r.db.Query(c, query)
	if err != nil {
		return categories, err
	}
	for rows.Next() {
		category := models.CategoryResponse{}
		if err := rows.Scan(&category.ID, &category.Name, &category.CreatedAt); err != nil {
			return categories, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}
func (r *CategoryRepository) GetByID(c context.Context, categoryID int) (models.CategoryResponse, error) {
	category := models.CategoryResponse{}
	query := `SELECT id, name, created_at FROM category where id = $1`
	if err := r.db.QueryRow(c, query, categoryID).Scan(&category.ID, &category.Name, &category.CreatedAt); err != nil {
		return category, err
	}
	return category, nil
}
