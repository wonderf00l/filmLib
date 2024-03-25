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

// @Description	Creating new profile - user registration
// @Tags			Auth
//
// @Accept			json
// @Produce		json
//
// @Param			username	body		string	true	"profile username"												example(clicker123)
// @Param			password	body		string	true	"profile password"												example(verysafePass)
// @Param			role_token	body		string	false	"token for activating specific role(admin token in example)"	example(admToken)
//
// @Success		200			{object}	responseJSON
// @Failure		400			{object}	errorResponseJSON
// @Failure		500			{object}	errorResponseJSON
//
// @Router			/api/v1/auth/signup [post]
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

// @Description	User login, creating new session
// @Tags			Auth
//
// @Accept			json
// @Produce		json
//
// @Param			username	body		string	true	"Profile username"	example(clicker123)
// @Param			password	body		string	true	"Profile password"	example(helloWorld)
//
// @Success		200			{object}	responseJSON
// @Failure		400			{object}	errorResponseJSON
// @Failure		401			{object}	errorResponseJSON
// @Failure		500			{object}	errorResponseJSON
//
// @Header			200			{string}	sess_key	"Auth cookie with new valid session id(base64)"
//
// @Router			/api/v1/auth/login [post]
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

// @Description	User logout, session deletion
// @Tags			Auth
//
// @Produce		json
//
// @Param			sess_key	header		string	false	"Cookie with session key"	example(k5qmqj507SejnpwJd%2FeO2Q%3D%3D)
//
// @Success		200			{object}	responseJSON
// @Failure		400			{object}	errorResponseJSON
// @Failure		401			{object}	errorResponseJSON
// @Failure		500			{object}	errorResponseJSON
//
// @Header			200			{string}	Session-id	"Auth cookie with expired session id"
//
// @Router			/api/v1/auth/logout [delete]
func (h *HandlerHTTP) Logout(w http.ResponseWriter, r *http.Request) {
	sessKey, ok := r.Context().Value(delivery.SessKey).(string)
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

// @Description	Update profile credentials
// @Tags			Auth
//
// @Accept			json
// @Produce		json
//
// @Param			new_username			body		string	false	"New profile username"							example(clicker123)
// @Param			old_password			body		string	true	"Old profile password for user verification"	example(helloWorld)
// @Param			new_password			body		string	false	"New preferable password"						example(helloWorldNew)
// @Param			new_password_repeated	body		string	false	"New preferable password repeated"				example(helloWorldNew)
// @Param			new_role_token			body		string	false	"New role token for optional role change"		example(moderatorToken)
//
// @Success		200						{object}	responseJSON
// @Failure		400						{object}	errorResponseJSON
// @Failure		401						{object}	errorResponseJSON
// @Failure		500						{object}	errorResponseJSON
//
// @Router			/api/v1/auth/update [put]
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
	updatedProfile.ID = profile.ID

	if err = h.serv.UpdateProfileData(r.Context(), updatedProfile, newRoleToken); err != nil {
		delivery.ResponseError(w, r, err)
	} else if err = delivery.ResponseOk(http.StatusOK, w, "Updated profile data successfully", nil); err != nil {
		delivery.ResponseError(w, r, err)
	}
}

func (h *HandlerHTTP) GetProfileData(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(delivery.UserIDKey).(int)
	if !ok {
		delivery.ResponseError(w, r, &delivery.MiddlewareNotSetError{MWTypes: []delivery.MiddlewareType{delivery.AuthMW}})
		return
	}

	profile, err := h.serv.GetProfileDataByUserID(r.Context(), userID)
	if err != nil {
		delivery.ResponseError(w, r, err)
		return
	}

	if err = delivery.ResponseOk(http.StatusOK, w, "Got profile data successfully", profile); err != nil {
		delivery.ResponseError(w, r, err)
	}
}
