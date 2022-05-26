package storage

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type TaskRepository struct {
	db *sql.DB
}

func newTaskRepository(db *sql.DB) *TaskRepository {
	taskRepository := &TaskRepository{
		db: db,
	}
	db.Exec(`CREATE TABLE  IF NOT EXISTS tasks(
    	id SERIAL PRIMARY KEY,
    	name VARCHAR(255),
		description TEXT,
		user_id INT REFERENCES users (id),
		solved BOOLEAN);`)
	return taskRepository
}
