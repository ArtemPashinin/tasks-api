package db

import (
	"context"
	"os"
	"tasks-api/models"

	"github.com/jackc/pgx/v5"
)

type Repository interface {
	FindAll() ([]models.Task, error)
	CreateOne(task models.Task) (int, error)
	UpdateOne(id int, task models.Task) error
	DeleteOne(id int) error
}

type PostgresService struct {
	pool *pgx.Conn
}

func NewPostgresService() (*PostgresService, error) {
	pool, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	return &PostgresService{pool: pool}, nil
}

func ClosePostgresService(p *PostgresService) {
	p.pool.Close(context.Background())
}

func (p *PostgresService) FindAll() ([]models.Task, error) {
	rows, err := p.pool.Query(context.Background(), "SELECT id, title, description, status, created_at, updated_at FROM tasks")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {
		var t models.Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (p *PostgresService) CreateOne(task models.Task) (int, error) {
	var id int
	err := p.pool.QueryRow(context.Background(),
		"INSERT INTO tasks (title, description, status) VALUES ($1, $2, $3) RETURNING id",
		task.Title, task.Description, task.Status).Scan(&id)
	return id, err
}

func (r *PostgresService) UpdateOne(id int, task models.Task) error {
	_, err := r.pool.Exec(context.Background(),
		"UPDATE tasks SET title=$1, description=$2, status=$3, updated_at=now() WHERE id=$4",
		task.Title, task.Description, task.Status, id)
	return err
}

func (r *PostgresService) DeleteOne(id int) error {
	_, err := r.pool.Exec(context.Background(), "DELETE FROM tasks WHERE id=$1", id)
	return err
}
