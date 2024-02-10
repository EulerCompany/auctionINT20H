package main

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

var ErrUserNotExists = errors.New("user not exists")

type User struct {
	ID             int
	Name           string
	HashedPassword []byte
	Created        time.Time
	Active         bool
}

type UserModel struct {
	DB *sql.DB
}

func NewUserModel(db *sql.DB) *UserModel {
	return &UserModel{DB: db}
}

func (m *UserModel) CreateUser(loginData LoginData) error {
	var name string
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(loginData.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// TODO: how to properly handle check + create?
	stmt := `SELECT name FROM user WHERE name = ?`
	row := m.DB.QueryRow(stmt, loginData.Name)
	_ = row.Scan(&name)

	if name != "" {
		return errors.New("user: user already exists")
	}

	stmt = `INSERT INTO user (name, hashed_password, created)
    VALUES (?, ?,  UTC_TIMESTAMP())`

	_, err = m.DB.Exec(stmt, loginData.Name, string(hashedPassword))

	return err
}

func (m *UserModel) Authenticate(loginData LoginData) (int, error) {
	var id int
	var hashedPassword []byte

	stmt := "SELECT id, hashed_password FROM user WHERE name = ? AND active = TRUE"
	row := m.DB.QueryRow(stmt, loginData.Name)

	err := row.Scan(&id, &hashedPassword)
	fmt.Println(err)
	fmt.Println(hashedPassword)
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(loginData.Password))

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (m *UserModel) Get(id int) string {
	var username string
	stmt := "SELECT name FROM user WHERE id = ? AND active = TRUE"
	row := m.DB.QueryRow(stmt, id)
	fmt.Println("errorDB")
	err := row.Scan(&username)
	fmt.Println("errorDB2")
	fmt.Println(err)

	return username
}

func (m *UserModel) GetIdByUsername(username string) (int, error) {
	var id int
	stmt := "SELECT id FROM user WHERE name = ? AND active = TRUE"
	row := m.DB.QueryRow(stmt, username)
	err := row.Scan(&id)
	fmt.Println(err)

	return id, nil
}
