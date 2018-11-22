package repository

import (
	"github.com/2DP/action-counter/config"
	"github.com/2DP/action-counter/model"
	"github.com/go-redis/redis"
)


type Repository struct {
	Repo map[string]model.Counter
	RedisClient *redis.Client
}


func (repo *Repository) Initialize(config *config.Config) {
	repo.Repo = make(map[string]model.Counter)
	
	repo.RedisClient = redis.NewClient(
		&redis.Options{Addr: config.RedisAddr, Password: config.RedisPassword, DB: 0})
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



func (repo *Repository) GetFromRedis(key string) string {
	val, err := repo.RedisClient.Get(key).Result()
	
	if err != nil {
	    panic(err)
	}
	
	return val
}

func (repo *Repository) SetToRedis(param *model.RedisParam) *model.RedisParam {
	err := repo.RedisClient.Set(param.Key, param.Value, 0).Err()
	
	if err != nil {
		panic(err)
	}

	return param
}