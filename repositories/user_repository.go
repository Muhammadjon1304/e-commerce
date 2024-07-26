package repositories

import (
	"database/sql"
	"github.com/muhammadjon1304/e-commerce/models"
	"log"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return UserRepository{
		DB: db,
	}
}

func (u *UserRepository) SaveUser(user models.User) bool {
	stmt, err := u.DB.Prepare("INSERT INTO users(username,email,password_hash,role) VALUES ($1,$2,$3,$4)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(user.Username, user.Email, user.Password_hash, user.Role)

	if err2 != nil {
		log.Fatal(err2)
		return false
	}
	return true
}
