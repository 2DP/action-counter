package repository

import (
	"github.com/2DP/action-counter/model"
)


type Repository struct {
	Repo map[string]model.Counter
}


func (repo *Repository) Initialize() {
	repo.Repo = make(map[string]model.Counter)
}

func (repo *Repository) Get(uuid string) model.Counter {
	return repo.Repo[uuid]
}

func (repo *Repository) Set(uuid string, counter model.Counter) model.Counter {
	repo.Repo[uuid] = counter
	return counter
}

func (repo *Repository) Increse(uuid string) model.Counter {
	counter, contains := repo.Repo[uuid]
	
	if contains {
		counter.Count++
		repo.Repo[uuid] = counter
	}
		
	return counter
}

func (repo *Repository) Delete(uuid string) model.Counter {
	counter, contains := repo.Repo[uuid]

	if contains {
		delete(repo.Repo, uuid)
	}
	
	return counter
}