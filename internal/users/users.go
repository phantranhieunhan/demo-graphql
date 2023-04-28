package users

import (
	database "github.com/phantranhieunhan/demo-graphql/internal/pkg/db/migrations/postgresql"
	"golang.org/x/crypto/bcrypt"

	"log"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"name"`
	Password string `json:"password"`
}

func (user *User) Create() int64 {
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		log.Fatal(err)
	}

	row := database.Db.QueryRow("INSERT INTO Users(Username,Password) VALUES($1,$2) RETURNING id", user.Username, hashedPassword)

	var id int64
	err = row.Scan(&id)
	if err != nil {
		log.Fatal(err)
	}

	return id
}

// GetUserIdByUsername check if a user exists in database by given username
func GetUserIdByUsername(username string) (int, error) {
	row := database.Db.QueryRow("select id from users WHERE username = $1", username)
	var id int
	err := row.Scan(&id)
	if err != nil {
		log.Fatal(err)
	}

	return id, nil
}

func (user *User) Authenticate() bool {
	row := database.Db.QueryRow("select password from users WHERE username = $1", user.Username)
	var hashedPassword string
	err := row.Scan(&hashedPassword)
	if err != nil {
		log.Fatal(err)
	}

	return CheckPasswordHash(user.Password, hashedPassword)
}

// HashPassword hashes given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPassword hash compares raw password with it's hashed values
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
