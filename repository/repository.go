package repository

import (
	"github.com/2DP/action-counter/model"
)


type Repository struct {
	Repo map[string]model.Counter
}


func (repo *Repository) Initialize() {
	repo.Repo = make(map[string]model.Counter)
	
	repo.Repo["test-uuid"] = model.Counter{UUID:"test-uuid", Count:10, DurationMillis:100}
}

func (repo *Repository) Get(uuid string) model.Counter {
	return repo.Repo[uuid]
}

func (repo *Repository) Set(uuid string, counter model.Counter) model.Counter {
	repo.Repo[uuid] = counter
	return counter
}
