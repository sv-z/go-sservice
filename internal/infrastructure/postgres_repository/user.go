package postgres_repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/mitchellh/mapstructure"

	"github.com/sv-z/in-scanner/internal/model"
)

type UserRepositoryPostgres struct {
	db *sql.DB
}

func (self UserRepositoryPostgres) Save(user *model.User) error {
	if user.Id != 0 {
		query := "UPDATE users SET email = $1, encrypted_password = $2 WHERE id = $3"
		if _, err := self.db.Exec(query, user.Email, user.EncryptedPassword, user.Id); err != nil {
			panic(err)
		}

		return nil
	}

	query := "INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id"
	if err := self.db.QueryRow(query, user.Email, user.EncryptedPassword).Scan(&user.Id); err != nil {
		panic(err)
	}

	return nil
}

func (self UserRepositoryPostgres) Delete(user *model.User) {
	query := "DELETE FROM users WHERE id = $1;"
	if _, err := self.db.Exec(query, user.Id); err != nil {
		panic(err)
	}
}

func (self UserRepositoryPostgres) GetById(userId int) *model.User {
	user := model.User{}

	query := "SELECT id, email, encrypted_password FROM users WHERE id=$1;"
	switch err := self.db.QueryRow(query, userId).Scan(&user.Id, &user.Email, &user.EncryptedPassword); err {
	case sql.ErrNoRows:
		panic(fmt.Errorf("user with id %d not found", userId))
	case nil:
		return &user
	default:
		panic(err)
	}
}

func (self UserRepositoryPostgres) FindByEmail(email string) *model.User {
	user := model.User{}

	query := "SELECT id, email, encrypted_password FROM users WHERE email=$1;"
	switch err := self.db.QueryRow(query, email).Scan(&user.Id, &user.Email, &user.EncryptedPassword); err {
	case sql.ErrNoRows:
		return nil
	case nil:
		return &user
	default:
		panic(err)
	}
}

func (self UserRepositoryPostgres) GetAll() []*model.User {
	query := "SELECT * FROM users;"

	rows, _ := self.db.Query(query)
	result := []*model.User{}

	for rows.Next() {
		result = append(result, rowToStruct(rows))
	}

	return result
}

func rowToStruct(rows *sql.Rows) *model.User {
	cols, _ := rows.Columns()

	// Create a slice of interface{}'s to represent each column,
	// and a second slice to contain pointers to each item in the columns slice.
	columns := make([]interface{}, len(cols))
	columnPointers := make([]interface{}, len(cols))
	for i, _ := range columns {
		columnPointers[i] = &columns[i]
	}

	// Scan the result into the column pointers...
	if err := rows.Scan(columnPointers...); err != nil {
		panic(err)
	}

	// Create our map, and retrieve the value for each column from the pointers slice,
	// storing it in the map with the name of the column as the key.
	m := make(map[string]interface{})
	for i, colName := range cols {
		val := columnPointers[i].(*interface{})
		m[toCamelCase(colName)] = *val
	}

	user := model.User{}

	err := mapstructure.Decode(m, &user)
	if err != nil {
		panic(err)
	}

	return &user

}

func toCamelCase(underStr string) (camelCase string) {
	isToUpper := false

	for _, v := range underStr {
		if v == '_' {
			isToUpper = true
		} else if isToUpper {
			camelCase += strings.ToUpper(string(v))
			isToUpper = false
		} else {
			camelCase += string(v)
		}
	}

	return camelCase
}

// NewCargoRepository returns a new instance of a MongoDB cargo repository.
func NewUserRepository(db *sql.DB) model.UserRepository {
	return &UserRepositoryPostgres{db: db}
}
