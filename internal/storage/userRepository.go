package storage

import (
	"database/sql"
	"github.com/Dshepett/task-manager/internal/models"
	_ "github.com/lib/pq"
)

type UserRepository struct {
	db *sql.DB
}

func newUserRepository(db *sql.DB) *UserRepository {
	userRepository := &UserRepository{
		db: db,
	}
	db.Exec(`CREATE TABLE  IF NOT EXISTS users(
    	id SERIAL PRIMARY KEY,
    	name VARCHAR(255));`)
	return userRepository
}

func (u *UserRepository) Create(user *models.User) error {
	err := u.db.QueryRow("INSERT INTO users(name) VALUES($1) RETURNING id", user.Name).Scan(&user.Id)
	return err
}

func (u *UserRepository) Edit(user *models.User) error {
	err := u.db.QueryRow("UPDATE users SET name = $1 WHERE id = $2 RETURNING name", user.Name, user.Id).Scan(&user.Name)
	return err
}

func (u *UserRepository) GetById(user *models.User) error {
	err := u.db.QueryRow("SELECT name FROM users WHERE id = $1", user.Id).Scan(&user.Name)
	return err
}

func (u *UserRepository) GetAll() ([]models.User, error) {
	rows, err := u.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	var users []models.User
	for rows.Next() {
		user := models.User{}
		if err := rows.Scan(&user.Id, &user.Name); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *UserRepository) Delete(id int) error {
	_, err := u.db.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}
