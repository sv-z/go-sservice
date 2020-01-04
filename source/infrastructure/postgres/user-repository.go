package postgres

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

type PostgresOrmUserRepository struct {
	db *pg.DB
}

func (repository *PostgresOrmUserRepository) Save(user *User) error {
	if user.Id == 0 {
		return repository.db.Insert(user)
	}

	return repository.db.Update(user)
}

func (repository *PostgresOrmUserRepository) GetById(userId int64) (*User, error) {
	user := &User{Id: userId}
	return user, repository.db.Select(user)
}

func (repository *PostgresOrmUserRepository) FindByEmail(email string) (*User, error) {
	user := &User{}
	return user, repository.db.Model(user).Where("email = ?", email).Select()
}

func (repository *PostgresOrmUserRepository) Delete(user *User) error {
	return repository.db.Delete(user)
}

func (repository *PostgresOrmUserRepository) getAll() ([]User, error) {
	var users []User
	return users, repository.db.Model(&users).WhereIn("id IN (?)", []int{1, 5, 11}).Select()
}

type User struct {
	Id       int64
	Email    string `pg:",notnull,unique"`
	Password string `pg:",notnull"`
}

func (u User) String() string {
	return fmt.Sprintf("User<%d %s>", u.Id, u.Email)
}

func main() {
	// https://github.com/go-pg/pg
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "pass_postgres",
	})
	defer db.Close()

	//user1 := &User{
	//	Email: "14@admin",
	//	Password: md5hash("a!A2"),
	//}
	//createSchema(db, (*User)(nil))

	repository := PostgresOrmUserRepository{db: db}
	fmt.Println(repository.getAll())
	//err := repository.Save(user1)
	//if err != nil {
	//	panic(err)
	//}
	//
	//user2, err1 := repository.GetById(7)
	//if err1 != nil {
	//	panic(err)
	//}
	//
	//user1.Email = "xxxxxxxx2"
	//err = repository.Save(user2)
	//if err != nil {
	//	panic(err)
	//}
	//
	//user3, _ := repository.FindByEmail("admin10@admin")
	//
	//repository.Delete(user3)
	//
	//
	//fmt.Println(user1)
	//fmt.Println(user2)
	//fmt.Println(user3)

}

func md5hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func createSchema(db *pg.DB, model interface{}) error {
	err := db.CreateTable(model, &orm.CreateTableOptions{
		Temp: false,
	})
	return err
}
