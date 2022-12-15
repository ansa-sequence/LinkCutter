package ansaserver

import (
	"LinkCutter/internal/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type Server struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

func NewServer(config *Config) *Server {
	return &Server{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (serv *Server) Start() error {
	if err := serv.loggerConfigure(); err != nil {
		return err
	}
	serv.routerConfigure()
	if err := serv.storeConfigure(); err != nil {
		return err
	}

	serv.logger.Debug("Starting server")
	return http.ListenAndServe(serv.config.BindAddr, serv.router)
}

func (serv *Server) loggerConfigure() error {
	level, err := logrus.ParseLevel(serv.config.LogLevel)
	if err != nil {
		return err
	}
	serv.logger.SetLevel(level)
	return nil
}

func (serv *Server) routerConfigure() {
	serv.router.HandleFunc("/debug", serv.handleHello())
}

func (serv *Server) storeConfigure() error {
	st := store.NewStore(serv.config.Store)
	if err := st.Open(); err != nil {
		return err
	}
	return nil
}

func (serv *Server) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Good Luck")
	}
}
