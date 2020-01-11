package infrastructure

import (
	"fmt"

	repositoryPostgres "github.com/sv-z/in-scaner/internal/infrastructure/repository"
	"github.com/sv-z/in-scaner/internal/model"
)

type repositoryKey string

type RepositoryManager struct {
	connectionHolder *ConnectionHolder
	repositories     map[repositoryKey]interface{}
}

// NewRepositoryManager ...
func NewRepositoryManager(con *ConnectionHolder) *RepositoryManager {

	var repositories = make(map[repositoryKey]interface{})
	repositories[repositoryKey("PostgresUserRepository")] = repositoryPostgres.NewUserRepository(con.postgresDB)

	return &RepositoryManager{
		connectionHolder: con,
		repositories:     repositories,
	}
}

// return user repository
func (rm *RepositoryManager) User() model.UserRepository {
	key := repositoryKey("PostgresUserRepository")

	return rm.getRepository(key).(model.UserRepository)
}

// fetch user repo
func (rm *RepositoryManager) getRepository(key repositoryKey) interface{} {
	repository, ok := rm.repositories[key]

	if !ok {
		panic(fmt.Errorf("the key %s not register yet", key))
	}

	return repository
}
