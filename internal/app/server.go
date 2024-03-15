package app

import (
	"net/http"
	"strconv"

	"github.com/wonderf00l/filmLib/internal/configs"
	"go.uber.org/zap"
)

type Server struct {
	http.Server
	logger *zap.SugaredLogger
}

func NewServer(cfg configs.SrvConfig, mux http.Handler, log *zap.SugaredLogger) *Server {
	srv := &Server{
		Server: http.Server{
			Addr:    cfg.Host + ":" + strconv.Itoa(cfg.Port),
			Handler: mux,
		},
		logger: log,
	}

	return srv
}

func (s *Server) Run() {
	s.logger.Infoln("Staring http server at", s.Addr)
	_ = s.Server.ListenAndServe()
}
