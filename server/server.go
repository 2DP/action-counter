package server

import (
	"encoding/json"
	"net/http"
	"log"
		
	"github.com/gorilla/mux"
	"github.com/2DP/action-counter/config"
	"github.com/2DP/action-counter/model"
	"github.com/2DP/action-counter/repository"
	"github.com/satori/go.uuid"
)

type Server struct {
	Router *mux.Router
	Config *config.Config
	Repo repository.Repository;
}



func (server *Server) Initialize(config *config.Config) {
	server.Config = config
	server.Router = mux.NewRouter()
	server.setRouters()
	
	server.Repo.Initialize(config)
}

func (server *Server) setRouters() {
	server.Get("/counter/{uuid:[a-z0-9-]+}", server.GetCounter)
	server.Post("/counter", server.CreateCounter) // body : {"duration-millis" : n}
	server.Post("/counter/{uuid:[a-z0-9-]+}", server.CreateCounter) // body : {"duration-millis" : n}
	server.Put("/counter/{uuid:[a-z0-9-]+}", server.UpdateCounter)
	server.Delete("/counter/{uuid:[a-z0-9-]+}", server.DeleteCounter)
	
	server.Get("/redis/{key:[a-z0-9]+}", server.GetFromRedis)
	server.Put("/redis", server.SetToRedis) // body {"key":"key","value":"value"}
}

func (server *Server) Run(host string) {
	log.Fatal(http.ListenAndServe(host, server.Router))
}



func (server *Server) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	server.Router.HandleFunc(path, f).Methods("GET")
}

func (server *Server) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	server.Router.HandleFunc(path, f).Methods("POST")
}

func (server *Server) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	server.Router.HandleFunc(path, f).Methods("PUT")
}

func (server *Server) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	server.Router.HandleFunc(path, f).Methods("DELETE")
}



func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

func respondError(w http.ResponseWriter, code int, message string) {
	respondJSON(w, code, map[string]string{"error": message})
}

func (server *Server) GetCounter(w http.ResponseWriter, r *http.Request) {
	// TODO:uuid 로 카운터 조회 (증감 안함)
	vars := mux.Vars(r)
	uuid := vars["uuid"]
	
	counter := server.Repo.Get(uuid)

	respondJSON(w, http.StatusOK, counter)
}

func (server *Server) CreateCounter(w http.ResponseWriter, r *http.Request) {
	// TODO: 새로운 카운터 세션을 생성해서 반환
	vars := mux.Vars(r)
	id := vars["uuid"]
	
	if id == "" {
		uuidV4, _ := uuid.NewV4()
		id = uuidV4.String()
	}
	
	counter := model.Counter{UUID:id, Count:1}
	
	err := json.NewDecoder(r.Body).Decode(&counter)
	
	if err != nil {
		log.Panic(err)
	}
		
	server.Repo.Set(id, counter)

	respondJSON(w, http.StatusOK, counter)
}

func (server *Server) UpdateCounter(w http.ResponseWriter, r *http.Request) {
	// TODO: 저장소에서 uuid로 카운터를 조회해서 카운터 증가 후 반환
	vars := mux.Vars(r)
	uuid := vars["uuid"]

	counter := server.Repo.Increse(uuid)
	
	respondJSON(w, http.StatusOK, counter)
}

func (server *Server) DeleteCounter(w http.ResponseWriter, r *http.Request) {
	// TODO 카운터 삭제
	vars := mux.Vars(r)
	uuid := vars["uuid"]

	counter := server.Repo.Delete(uuid)
	
	respondJSON(w, http.StatusOK, counter)
}



func (server *Server) GetFromRedis(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	
	value := server.Repo.GetFromRedis(key)
	
	respondJSON(w, http.StatusOK, value)
}

func (server *Server) SetToRedis(w http.ResponseWriter, r *http.Request) {
	param := model.RedisParam{}
	err := json.NewDecoder(r.Body).Decode(&param)
	
	if err != nil {
		panic(err)
		return
	}
	
	server.Repo.SetToRedis(&param)
	
	respondJSON(w, http.StatusOK, param)
}