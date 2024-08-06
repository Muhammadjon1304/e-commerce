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

func (u *UserRepository) GetUserByUsername(username string) models.User {
	query, err := u.DB.Query("SELECT id,username,email,password_hash,role FROM users WHERE username=$1", username)

	if err != nil {
		log.Fatal(err)
		return models.User{}
	}
	var user models.User

	if query != nil {
		for query.Next() {
			var (
				id       uint
				username string
				email    string
				password string
				role     string
			)
			err := query.Scan(&id, &username, &email, &password, &role)
			if err != nil {
				log.Fatal(err)
			}
			user = models.User{id, username, email, password, role}
		}
	}
	return user
}

func (u *UserRepository) GetUserByUsernameForUser(username string) models.User {
	query, err := u.DB.Query("SELECT id,username,email,role FROM users WHERE username=$1", username)

	if err != nil {
		log.Fatal(err)
		return models.User{}
	}
	var user models.User

	if query != nil {
		for query.Next() {
			var (
				id       uint
				username string
				email    string
				role     string
			)
			err := query.Scan(&id, &username, &email, &role)
			if err != nil {
				log.Fatal(err)
			}
			user = models.User{ID: id, Username: username, Email: email, Role: role}
		}
	}
	return user
}

func (u *UserRepository) GetUserIDByUsername(username string) uint {
	query, err := u.DB.Query("SELECT id FROM users WHERE username=$1", username)

	if err != nil {
		log.Fatal(err)
		return 0
	}
	var id uint

	if query != nil {
		for query.Next() {
			var (
				user_id uint
			)
			err := query.Scan(&user_id)
			if err != nil {
				log.Fatal(err)
			}
			id = user_id
		}
	}
	return id
}
