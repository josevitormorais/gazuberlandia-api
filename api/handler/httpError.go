package handler

import (
	"encoding/json"
	"gazuberlandia"
	"net/http"
)

type responseError struct {
	Errors string `json:"errors"`
}

func ResponseError(w http.ResponseWriter, err error) {

	code, message := gazuberlandia.ErrorCode(err), gazuberlandia.ErrorMessage(err)

	w.WriteHeader(ErrorStatusCode(code))

	json.NewEncoder(w).Encode(&responseError{Errors: message})
}

var codes = map[string]int{
	gazuberlandia.CONFLICT:            http.StatusConflict,
	gazuberlandia.INTERNAL:            http.StatusInternalServerError,
	gazuberlandia.INVALID:             http.StatusBadRequest,
	gazuberlandia.NOTFOUND:            http.StatusNotFound,
	gazuberlandia.NOTIMPLEMENTED:      http.StatusNotImplemented,
	gazuberlandia.UNAUTHORIZED:        http.StatusUnauthorized,
	gazuberlandia.UNPROCESSABLEENTITY: http.StatusUnprocessableEntity,
}

func ErrorStatusCode(code string) int {
	if v, ok := codes[code]; ok {
		return v
	}
	return http.StatusUnprocessableEntity
}
