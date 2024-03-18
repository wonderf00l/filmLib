package app

import (
	"github.com/go-chi/chi"
	"go.uber.org/zap"

	delivery "github.com/wonderf00l/filmLib/internal/delivery/http/v1"
	"github.com/wonderf00l/filmLib/internal/delivery/http/v1/auth"
	authService "github.com/wonderf00l/filmLib/internal/service/auth"

	"github.com/wonderf00l/filmLib/internal/service/role"
)

type Router struct {
	Mux *chi.Mux
}

func NewRouter() Router {
	return Router{chi.NewMux()}
}

type HandlersHTTP struct {
	auth auth.HandlerHTTP
}

func (r Router) RegisterRoute(h HandlersHTTP, log *zap.SugaredLogger, authService authService.Service, roleService role.Service) {
	r.Mux.Use( /*delivery.RecoverMiddleware,*/ delivery.LoggingMiddleware(log))

	authMW := delivery.AuthMiddleware(authService)
	_ = delivery.CheckRoleMiddleware(roleService)

	r.Mux.Route("/api/v1", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Post("/signup", h.auth.Signup)
			r.Post("/login", h.auth.Login)

			r.With(authMW).Group(func(r chi.Router) {
				r.Put("/update", h.auth.UpdateProfileData)
				r.Delete("/logout", h.auth.Logout)
			})
		})
	})
}
