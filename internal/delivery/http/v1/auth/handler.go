package auth

import (
	"encoding/json"
	"net/http"
	"time"

	delivery "github.com/wonderf00l/filmLib/internal/delivery/http/v1"
	service "github.com/wonderf00l/filmLib/internal/service/auth"
)

type AuthHandlerHTTP struct {
	serv service.Service
}

func NewHandler(s service.Service) *AuthHandlerHTTP {
	return &AuthHandlerHTTP{serv: s}
}

// signup
// login post
// logout

// username pass token

/*
 1. check content type
 2. try to unmarshal
 3. validate unmarshaled data
    username,password - content, len
 4. call service
 5. response
*/

// change credentials
func (h *AuthHandlerHTTP) Signup(w http.ResponseWriter, r *http.Request) {
	if contentType := r.Header.Get("Content-Type"); contentType != delivery.ApplicationJson {
		delivery.ResponseError(w, r, &delivery.InvalidContentTypeError{PreferredType: delivery.ApplicationJson})
		return
	}

	data := signupData{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		delivery.ResponseError(w, r, &delivery.InvalidBodyError{})
		return
	}
	defer r.Body.Close()

	if err := data.Validate(); err != nil {
		delivery.ResponseError(w, r, err)
		return
	}

	profile, roleToken := signupDataDeliveryToService(data)
	if err := h.serv.Signup(r.Context(), profile, roleToken); err != nil {
		delivery.ResponseError(w, r, err)
	} else if err := delivery.ResponseOk(http.StatusOK, w, "registered profile successfully", nil); err != nil {
		delivery.ResponseError(w, r, err)
	}
}

func (h *AuthHandlerHTTP) Login(w http.ResponseWriter, r *http.Request) {
	if contentType := r.Header.Get("Content-Type"); contentType != delivery.ApplicationJson {
		delivery.ResponseError(w, r, &delivery.InvalidContentTypeError{PreferredType: delivery.ApplicationJson})
		return
	}

	data := loginData{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		delivery.ResponseError(w, r, &delivery.InvalidBodyError{})
		return
	}
	defer r.Body.Close()

	if err := data.Validate(); err != nil {
		delivery.ResponseError(w, r, err)
		return
	}

	credentials := loginDataDeliveryToService(data)
	session, err := h.serv.CheckCredentials(r.Context(), credentials.Username, credentials.Password)
	if err != nil {
		delivery.ResponseError(w, r, err)
		return
	}

	cookie := &http.Cookie{
		Name:     delivery.CookieKey,
		Value:    session.Key,
		HttpOnly: true,
		Path:     "/",
		Expires:  session.Expires,
	}
	http.SetCookie(w, cookie)

	if err := delivery.ResponseOk(http.StatusOK, w, "authenticated successfully", nil); err != nil {
		delivery.ResponseError(w, r, err)
	}
}

// under MW
func (h *AuthHandlerHTTP) Logout(w http.ResponseWriter, r *http.Request) {
	sessKey, ok := r.Context().Value(delivery.CookieKey).(string)
	if !ok {
		delivery.ResponseError(w, r, &delivery.MiddlewareNotSetError{MWTypes: []delivery.MiddlewareType{delivery.AuthMW}})
		return
	}

	if err := h.serv.Logout(r.Context(), sessKey); err != nil {
		delivery.ResponseError(w, r, err)
		return
	}

	cookie, _ := r.Cookie(delivery.CookieKey)
	cookie.Expires = time.Now().UTC().AddDate(0, -1, 0)
	http.SetCookie(w, cookie)

	if err := delivery.ResponseOk(http.StatusOK, w, "logout successfuly", nil); err != nil {
		delivery.ResponseError(w, r, err)
	}
}
