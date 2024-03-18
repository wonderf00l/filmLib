package v1

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type responseJSON struct {
	Status  string      `json:"status" example:"ok"`
	Message string      `json:"message" example:"authorized successfully"`
	Body    interface{} `json:"body"`
}

type errorResponseJSON struct {
	Status  string `json:"status" example:"error"`
	Message string `json:"message" example:"Error description"`
	Code    string `json:"code" example:"no_auth"`
}

func setContentTypeJSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func ResponseOk(statusCode int, w http.ResponseWriter, message string, body any) error {
	res := responseJSON{
		Status:  "ok",
		Message: message,
		Body:    body,
	}
	resBytes, err := json.Marshal(&res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return fmt.Errorf("responseOk: %w", err)
	}

	w.WriteHeader(statusCode)
	setContentTypeJSON(w)
	_, err = w.Write(resBytes)
	return err
}

func ResponseError(w http.ResponseWriter, r *http.Request, serviceError error) {
	serviceLogger, err := getLoggerFromCtx(r.Context())

	fmt.Println("LOGGER: ", serviceLogger, err)

	var defaultLogging bool
	if err != nil {
		defaultLogging = true
		log.Println("response error: use default logging, reason - ", err.Error())
	}

	code, status := getCodeStatusHTTP(serviceError)
	var msg string
	if status == http.StatusInternalServerError {
		warnLogMsg := fmt.Sprintf("unexpected application error: %s, URL - %s, METHOD - %s", serviceError.Error(), r.URL.String(), r.Method)
		if defaultLogging {
			log.Println(warnLogMsg)
		} else {
			serviceLogger.Warnln(warnLogMsg)
		}
		msg = "internal error occured"
	} else {
		logMsg := fmt.Sprintf("got declared error: %s, URL - %s, METHOD - %s", serviceError.Error(), r.URL.String(), r.Method)
		if defaultLogging {
			log.Println(logMsg)
		} else {
			serviceLogger.Infoln(logMsg)
		}
		msg = serviceError.Error()
	}

	res := errorResponseJSON{
		Status:  "error",
		Message: msg,
		Code:    code,
	}
	resBytes, _ := json.Marshal(res)

	w.WriteHeader(status)
	setContentTypeJSON(w)
	_, err = w.Write(resBytes)
	if err != nil {
		errMsg := "couldn't write error response: " + err.Error()
		if defaultLogging {
			log.Println(errMsg)
		} else {
			serviceLogger.Warnln(errMsg)
		}
	}
}
