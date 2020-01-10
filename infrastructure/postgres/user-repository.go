package postgres

import (
	"github.com/go-pg/pg/v9"
	model "github.com/sv-z/sservice"
)

type PostgresOrmUserRepository struct {
	db *pg.DB
}

func (repository *PostgresOrmUserRepository) Store(user *model.User) error {
	if user.Id == 0 {
		return repository.db.Insert(user)
	}

	return repository.db.Update(user)
}

func (repository *PostgresOrmUserRepository) GetById(userId int64) (*model.User, error) {
	user := &model.User{Id: userId}
	return user, repository.db.Select(user)
}

func (repository *PostgresOrmUserRepository) FindByEmail(email string) (*model.User, error) {
	user := &model.User{}
	return user, repository.db.Model(user).Where("email = ?", email).Select()
}

func (repository *PostgresOrmUserRepository) Delete(user *model.User) error {
	return repository.db.Delete(user)
}

func (repository *PostgresOrmUserRepository) GetAll() ([]model.User, error) {
	var users []model.User
	return users, repository.db.Model(&users).WhereIn("id IN (?)", []int{1, 5, 11}).Select()
}

// NewCargoRepository returns a new instance of a MongoDB cargo repository.
func NewUserRepository() (model.UserRepository, error) {
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

	return &PostgresOrmUserRepository{db: db}, nil
}
//
//func main() {
//	// https://github.com/go-pg/pg
//	db := pg.Connect(&pg.Options{
//		User:     "postgres",
//		Password: "pass_postgres",
//	})
//	defer db.Close()
//
//	//user1 := &User{
//	//	Email: "14@admin",
//	//	Password: md5hash("a!A2"),
//	//}
//	//createSchema(db, (*User)(nil))
//
//	repository := PostgresOrmUserRepository{db: db}
//	fmt.Println(repository.getAll())
//	//err := repository.Store(user1)
//	//if err != nil {
//	//	panic(err)
//	//}
//	//
//	//user2, err1 := repository.GetById(7)
//	//if err1 != nil {
//	//	panic(err)
//	//}
//	//
//	//user1.Email = "xxxxxxxx2"
//	//err = repository.Store(user2)
//	//if err != nil {
//	//	panic(err)
//	//}
//	//
//	//user3, _ := repository.FindByEmail("admin10@admin")
//	//
//	//repository.Delete(user3)
//	//
//	//
//	//fmt.Println(user1)
//	//fmt.Println(user2)
//	//fmt.Println(user3)
//
//}
//
//func md5hash(text string) string {
//	hash := md5.Sum([]byte(text))
//	return hex.EncodeToString(hash[:])
//}
//
//func createSchema(db *pg.DB, model interface{}) error {
//	err := db.CreateTable(model, &orm.CreateTableOptions{
//		Temp: false,
//	})
//	return err
//}
