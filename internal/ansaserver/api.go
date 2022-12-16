package ansaserver

import (
	"LinkCutter/internal/model"
	"LinkCutter/internal/store"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type server struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  store.Store
}

func newServer(store store.Store) *server {
	s := &server{
		config: NewConfig(),
		logger: logrus.New(),
		router: mux.NewRouter(),
		store:  store,
	}

	s.routerConfigure()

	return s
}

func (serv *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	serv.router.ServeHTTP(w, r)
}

func (serv *server) routerConfigure() {
	serv.router.HandleFunc("/users", serv.handleUsersCreate()).Methods("POST")
	serv.router.HandleFunc("/users", serv.handleUsersRemove()).Methods("DELETE")
}

func (serv *server) handleUsersCreate() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			serv.error(w, r, http.StatusBadRequest, err)
			return
		}

		u := &model.User{
			Email:    req.Email,
			Password: req.Password,
		}
		if err := serv.store.User().Create(u); err != nil {
			serv.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		u.Sanitize()
		serv.respond(w, r, http.StatusCreated, u)
	}
}
func (serv *server) handleUsersRemove() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			serv.error(w, r, http.StatusBadRequest, err)
			return
		}

		_, err := serv.store.User().FindByEmail(req.Email)
		if err != nil {
			serv.respond(w, r, http.StatusBadRequest, err)
		}
		serv.respond(w, r, http.StatusCreated, err)
	}
}

func (serv *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	serv.respond(w, r, code, map[string]string{"error": err.Error()})
}
func (serv *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		_ = json.NewEncoder(w).Encode(data)
	}
}
