package auth

import (
	"encoding/json"
	"net/http"
	"time"

	delivery "github.com/wonderf00l/filmLib/internal/delivery/http/v1"
	service "github.com/wonderf00l/filmLib/internal/service/auth"
)

type HandlerHTTP struct {
	serv service.Service
}

func New(s service.Service) HandlerHTTP {
	return HandlerHTTP{serv: s}
}

func (h *HandlerHTTP) Signup(w http.ResponseWriter, r *http.Request) {
	if contentType := r.Header.Get("Content-Type"); contentType != delivery.ApplicationJSON {
		delivery.ResponseError(w, r, &delivery.InvalidContentTypeError{PreferredType: delivery.ApplicationJSON})
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
	} else if err = delivery.ResponseOk(http.StatusOK, w, "registered profile successfully", nil); err != nil {
		delivery.ResponseError(w, r, err)
	}
}

func (h *HandlerHTTP) Login(w http.ResponseWriter, r *http.Request) {
	if contentType := r.Header.Get("Content-Type"); contentType != delivery.ApplicationJSON {
		delivery.ResponseError(w, r, &delivery.InvalidContentTypeError{PreferredType: delivery.ApplicationJSON})
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
	session, err := h.serv.CreateSessionForUser(r.Context(), credentials.Username, credentials.Password)
	if err != nil {
		delivery.ResponseError(w, r, err)
		return
	}

	cookie := &http.Cookie{
		Name:     delivery.CookieName,
		Value:    session.Key,
		HttpOnly: true,
		Path:     "/",
		Expires:  session.Expires,
	}
	http.SetCookie(w, cookie)

	if err = delivery.ResponseOk(http.StatusOK, w, "authenticated successfully", nil); err != nil {
		delivery.ResponseError(w, r, err)
	}
}

func (h *HandlerHTTP) Logout(w http.ResponseWriter, r *http.Request) {
	sessKey, ok := r.Context().Value(delivery.CookieKey).(string)
	if !ok {
		delivery.ResponseError(w, r, &delivery.MiddlewareNotSetError{MWTypes: []delivery.MiddlewareType{delivery.AuthMW}})
		return
	}

	if err := h.serv.Logout(r.Context(), sessKey); err != nil {
		delivery.ResponseError(w, r, err)
		return
	}

	cookie, _ := r.Cookie(delivery.CookieName)
	cookie.Expires = time.Now().UTC().AddDate(0, -1, 0)
	http.SetCookie(w, cookie)

	if err := delivery.ResponseOk(http.StatusOK, w, "logout successfuly", nil); err != nil {
		delivery.ResponseError(w, r, err)
	}
}

func (h *HandlerHTTP) UpdateProfileData(w http.ResponseWriter, r *http.Request) {
	if contentType := r.Header.Get("Content-Type"); contentType != delivery.ApplicationJSON {
		delivery.ResponseError(w, r, &delivery.InvalidContentTypeError{PreferredType: delivery.ApplicationJSON})
		return
	}

	data := updateData{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		delivery.ResponseError(w, r, &delivery.InvalidBodyError{})
		return
	}
	defer r.Body.Close()

	if err := data.Validate(); err != nil {
		delivery.ResponseError(w, r, err)
		return
	}

	userID, ok := r.Context().Value(delivery.UserIDKey).(int)
	if !ok {
		delivery.ResponseError(w, r, &delivery.MiddlewareNotSetError{MWTypes: []delivery.MiddlewareType{delivery.AuthMW}})
		return
	}

	profile, err := h.serv.CheckCredentialsByUserID(r.Context(), userID, *data.OldPassword)
	if err != nil {
		delivery.ResponseError(w, r, err)
		return
	}

	updatedProfile, newRoleToken := updateDataDeliveryToService(data)
	if updatedProfile.Username == "" {
		updatedProfile.Username = profile.Username
	}
	if updatedProfile.Password == "" {
		updatedProfile.Password = profile.Password
	}

	if err = h.serv.UpdateProfileData(r.Context(), updatedProfile, newRoleToken); err != nil {
		delivery.ResponseError(w, r, err)
	} else if err = delivery.ResponseOk(http.StatusOK, w, "Updated profile data successfully", nil); err != nil {
		delivery.ResponseError(w, r, err)
	}
}
