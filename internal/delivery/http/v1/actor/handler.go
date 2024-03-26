package actor

import (
	"encoding/json"
	"net/http"
	"strconv"

	delivery "github.com/wonderf00l/filmLib/internal/delivery/http/v1"
	service "github.com/wonderf00l/filmLib/internal/service/actor"
)

type HandlerHTTP struct {
	serv service.Service
}

func New(s service.Service) HandlerHTTP {
	return HandlerHTTP{serv: s}
}

// @Description	Add information about new actor
// @Tags			Actor
//
// @Accept			json
// @Produce		json
//
// @Param			name			body		string	true	"name of the actor"		example(Ryan Gosling)
// @Param			gender			body		string	true	"gender of the actor"	example(male)
// @Param			date_of_bitrh	body		string	false	"actor's date of birth"	example(2002-11-10)
//
// @Success		200				{object}	responseJSON
// @Failure		400				{object}	errorResponseJSON
// @Failure		500				{object}	errorResponseJSON
//
// @Router			/api/v1/actor/create [post]
func (h *HandlerHTTP) CreateActor(w http.ResponseWriter, r *http.Request) {
	if contentType := r.Header.Get("Content-Type"); contentType != delivery.ApplicationJSON {
		delivery.ResponseError(w, r, &delivery.InvalidContentTypeError{PreferredType: delivery.ApplicationJSON})
		return
	}

	data := actorData{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		delivery.ResponseError(w, r, &delivery.InvalidBodyError{})
		return
	}
	defer r.Body.Close()

	actor, err := actorDataDeliveryToService(data)
	if err != nil {
		delivery.ResponseError(w, r, err)
		return
	}

	if err := h.serv.AddActor(r.Context(), *actor); err != nil {
		delivery.ResponseError(w, r, err)
	} else if err = delivery.ResponseOk(http.StatusOK, w, "add actor successfully", nil); err != nil {
		delivery.ResponseError(w, r, err)
	}
}

// @Description	Get information about existing actor
// @Tags			Actor
//
// @Produce		json
//
// @Param			id	query		int	true	"id of the actor to get"
//
// @Success		200	{object}	responseJSON
// @Failure		404	{object}	errorResponseJSON
// @Failure		500	{object}	errorResponseJSON
//
// @Router			/api/v1/actor/get [get]
func (h *HandlerHTTP) GetActor(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	actorID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		delivery.ResponseError(w, r, &delivery.InvalidQueryParamError{Params: map[string]string{"id": id}, Need: "id"})
		return
	}

	actor, err := h.serv.GetActor(r.Context(), int(actorID))
	if err != nil {
		delivery.ResponseError(w, r, err)
		return
	}

	if err = delivery.ResponseOk(http.StatusOK, w, "Got actor successfully", getActorServiceToDelivery(*actor)); err != nil {
		delivery.ResponseError(w, r, err)
	}
}

// @Description	Delete information about existing actor
// @Tags			Actor
//
// @Produce		json
//
// @Param			id	query		int	true	"id of the actor to delete"
//
// @Success		200	{object}	responseJSON
// @Failure		401	{object}	errorResponseJSON
// @Failure		403	{object}	errorResponseJSON
// @Failure		404	{object}	errorResponseJSON
// @Failure		500	{object}	errorResponseJSON
//
// @Router			/api/v1/actor/delete [delete]
func (h *HandlerHTTP) DeleteActor(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	actorID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		delivery.ResponseError(w, r, &delivery.InvalidQueryParamError{Params: map[string]string{"id": id}, Need: "id"})
		return
	}

	if err = h.serv.DeleteActorData(r.Context(), int(actorID)); err != nil {
		delivery.ResponseError(w, r, err)
		return
	}

	if err = delivery.ResponseOk(http.StatusOK, w, "delete actor successfully", nil); err != nil {
		delivery.ResponseError(w, r, err)
	}
}

// @Description	Update info about exising actor
// @Tags			Actor
//
// @Accept			json
// @Produce		json
//
// @Param			id				body		int		true	"ID of the actor to update" example(2)
// @Param			name			body		string	false	"New actor's name"			example(Ryan Gosling)
// @Param			gender			body		string	false	"New actor's gender"		example(male)
// @Param			date_of_birth	body		string	false	"New actor's date of birth"	example(2001-02-02)
//
// @Success		200				{object}	responseJSON
// @Failure		400				{object}	errorResponseJSON
// @Failure		401				{object}	errorResponseJSON
// @Failure		403				{object}	errorResponseJSON
// @Failure		404				{object}	errorResponseJSON
// @Failure		500				{object}	errorResponseJSON
//
// @Router			/api/v1/actor/update [put]
func (h *HandlerHTTP) UpdateActorData(w http.ResponseWriter, r *http.Request) {
	if contentType := r.Header.Get("Content-Type"); contentType != delivery.ApplicationJSON {
		delivery.ResponseError(w, r, &delivery.InvalidContentTypeError{PreferredType: delivery.ApplicationJSON})
		return
	}

	data := actorData{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		delivery.ResponseError(w, r, &delivery.InvalidBodyError{})
		return
	}
	defer r.Body.Close()

	if data.ID == nil {
		delivery.ResponseError(w, r, &actorIDnotProvidedError{})
		return
	}

	updData, err := updateActorDeliveryToService(data)
	if err != nil {
		delivery.ResponseError(w, r, err)
		return
	}

	if err := h.serv.UpdateActorData(r.Context(), *data.ID, updData); err != nil {
		delivery.ResponseError(w, r, err)
	} else if err = delivery.ResponseOk(http.StatusOK, w, "updated actor successfully", nil); err != nil {
		delivery.ResponseError(w, r, err)
	}
}
