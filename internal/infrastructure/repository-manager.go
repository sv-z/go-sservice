package infrastructure

import (
	"database/sql"
	"fmt"

	postgresRepository "github.com/sv-z/in-scanner/internal/infrastructure/postgres_repository"
	"github.com/sv-z/in-scanner/internal/model"
)

type RepositoryManagerInterface interface {
	User() model.UserRepository
}

type repositoryKey string

type repositoryManager struct {
	repositories map[repositoryKey]interface{}
}

// NewRepositoryManager ...
func NewRepositoryManager(postgresDB *sql.DB) RepositoryManagerInterface {

	var repositories = make(map[repositoryKey]interface{})
	repositories[repositoryKey("PostgresUserRepository")] = postgresRepository.NewUserRepository(postgresDB)

	return &repositoryManager{
		repositories: repositories,
	}
}

// return user repository
func (rm *repositoryManager) User() model.UserRepository {
	key := repositoryKey("PostgresUserRepository")

	return rm.getRepository(key).(model.UserRepository)
}

// fetch user repo
func (rm *repositoryManager) getRepository(key repositoryKey) interface{} {
	repository, ok := rm.repositories[key]

	if !ok {
		panic(fmt.Errorf("the key %s not register yet", key))
	}

	return repository
}
