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
