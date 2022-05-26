package storage

import (
	"database/sql"
	"fmt"
	"github.com/Dshepett/task-manager/internal/config"
	_ "github.com/lib/pq"
	"log"
)

type Storage struct {
	config         *config.Config
	UserRepository *UserRepository
	TaskRepository *TaskRepository
}

func New(config *config.Config) *Storage {
	return &Storage{config: config}
}

func (s *Storage) Open() {
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		s.config.DBUser, s.config.DBPassword, s.config.DBName)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	s.UserRepository = newUserRepository(db)
	s.TaskRepository = newTaskRepository(db)
}
