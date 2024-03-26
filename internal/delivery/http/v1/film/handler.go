package film

import (
	"encoding/json"
	"net/http"

	delivery "github.com/wonderf00l/filmLib/internal/delivery/http/v1"
	service "github.com/wonderf00l/filmLib/internal/service/film"
)

type HandlerHTTP struct {
	serv service.Service
}

func New(s service.Service) HandlerHTTP {
	return HandlerHTTP{serv: s}
}

func (h *HandlerHTTP) CreateFilm(w http.ResponseWriter, r *http.Request) {
	if contentType := r.Header.Get("Content-Type"); contentType != delivery.ApplicationJSON {
		delivery.ResponseError(w, r, &delivery.InvalidContentTypeError{PreferredType: delivery.ApplicationJSON})
		return
	}

	data := filmData{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		delivery.ResponseError(w, r, &delivery.InvalidBodyError{})
		return
	}
	defer r.Body.Close()

	film, err := filmDataDeliveryToService(data)
	if err != nil {
		delivery.ResponseError(w, r, err)
		return
	}

	if err := h.serv.AddFilm(r.Context(), film); err != nil {
		delivery.ResponseError(w, r, err)
	} else if err = delivery.ResponseOk(http.StatusOK, w, "add film successfully", nil); err != nil {
		delivery.ResponseError(w, r, err)
	}
}
