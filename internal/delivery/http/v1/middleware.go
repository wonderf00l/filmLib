package v1

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	entityRole "github.com/wonderf00l/filmLib/internal/entity/role"
	errPkg "github.com/wonderf00l/filmLib/internal/errors"
	"github.com/wonderf00l/filmLib/internal/service/auth"
	"github.com/wonderf00l/filmLib/internal/service/role"

	"go.uber.org/zap"
)

type Middleware func(http.Handler) http.Handler

type MiddlewareType string

type ctxKey uint8

const (
	SessKey ctxKey = iota + 1
	UserIDKey
	loggerKey
	rolesKey

	CookieName = "sess_id"

	AuthMW MiddlewareType = "auth"
)

func getLoggerFromCtx(ctx context.Context) (*zap.SugaredLogger, error) {
	if log, ok := ctx.Value(loggerKey).(*zap.SugaredLogger); ok {
		return log, nil
	}
	return nil, errors.New("get logger from ctx: logger not found")
}

func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("FATAL ERROR: RECOVERED FROM PANIC -> URL - %s, METHOD - %s\n", r.URL.String(), r.Method)
				ResponseError(w, r, &errPkg.InternalError{Message: "GOT PANIC", Layer: string(errPkg.Delivery)})
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func LoggingMiddleware(serviceLogger *zap.SugaredLogger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqID := uuid.NewString()

			serviceLogger.Infow(
				"Got request: ",
				zap.String("ID", reqID),
				zap.String("method", r.Method),
				zap.String("host", r.Host),
				zap.String("proto", r.Proto),
				zap.String("path", r.URL.String()),
				zap.String("content-type", r.Header.Get("Content-Type")),
				zap.Int64("content_length", r.ContentLength),
				zap.String("address", r.RemoteAddr),
			)

			defer func(t time.Time) {
				serviceLogger.Infow(
					"Responsed:",
					zap.String("ID", reqID),
					zap.Int64("processing_time_us", time.Since(t).Microseconds()),
					zap.String("content-type", w.Header().Get("Content-Type")),
				)
			}(time.Now())

			r = r.WithContext(context.WithValue(r.Context(), loggerKey, serviceLogger))
			next.ServeHTTP(w, r)
		})
	}
}

func ctxWithAuthParams(ctx context.Context, sessKey string, userID int) context.Context {
	withParams := context.WithValue(ctx, SessKey, sessKey)
	withParams = context.WithValue(withParams, UserIDKey, userID)
	return withParams
}

func AuthMiddleware(s auth.Service) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if cookie, err := r.Cookie(CookieName); err == nil {
				session, err := s.GetUserSession(r.Context(), cookie.Value)
				if err != nil {
					ResponseError(w, r, err)
					return
				}

				r = r.WithContext(ctxWithAuthParams(r.Context(), session.Key, session.UserID))
				next.ServeHTTP(w, r)
			} else {
				ResponseError(w, r, &NoAuthCookieError{})
			}
		})
	}
}

func SetRolesMiddleware(roles []entityRole.Role) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(context.WithValue(r.Context(), rolesKey, roles))
			next.ServeHTTP(w, r)
		})
	}
}

func CheckRolesMiddleware(s role.Service) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			needRoles, ok := r.Context().Value(rolesKey).([]entityRole.Role)
			if !ok {
				next.ServeHTTP(w, r)
				return
			}

			userID, ok := r.Context().Value(UserIDKey).(int)
			if !ok {
				ResponseError(w, r, &MiddlewareNotSetError{MWTypes: []MiddlewareType{AuthMW}})
				return
			}

			role, err := s.GetUserRole(r.Context(), userID)
			if err != nil {
				ResponseError(w, r, err)
				return
			}

			if !contains(role, needRoles) {
				ResponseError(w, r, &errPkg.InvalidRoleForActionError{Need: convert(needRoles)})
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func convert(roles []entityRole.Role) []string {
	converted := make([]string, len(roles))
	for i, r := range roles {
		converted[i] = entityRole.RoleMap[r]
	}
	return converted
}

func contains(role entityRole.Role, roles []entityRole.Role) bool {
	for _, r := range roles {
		if role == r {
			return true
		}
	}
	return false
}
