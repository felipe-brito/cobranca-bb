package handlers

import (
	"encoding/json"
	"net/http"
	"sync"
)

var (
	responseHandler responseHandlerImpl
	onceResponse    sync.Once
)

type ResponseHandler interface {
	OK(w http.ResponseWriter, data interface{})
	Success(w http.ResponseWriter)
	Created(w http.ResponseWriter, data interface{})
	InternalServerError(w http.ResponseWriter, r *http.Request)
}

type responseHandlerImpl struct {
}

func GetInstanceResponse() ResponseHandler {
	onceResponse.Do(func() {
		responseHandler = responseHandlerImpl{}
	})
	return responseHandler
}

func (h responseHandlerImpl) OK(w http.ResponseWriter, data interface{}) {
	response(w, data, http.StatusOK)
}

func (h responseHandlerImpl) Success(w http.ResponseWriter) {
	response(w, nil, http.StatusOK)
}

func (h responseHandlerImpl) Created(w http.ResponseWriter, data interface{}) {
	response(w, data, http.StatusCreated)
}

func (h responseHandlerImpl) InternalServerError(w http.ResponseWriter, r *http.Request) {
	// TODO montar o objeto de response
	response(w, nil, http.StatusInternalServerError)
}

func response(w http.ResponseWriter, data interface{}, httpStatus int) {
	handlers(w)
	w.WriteHeader(httpStatus)
	if data != nil {
		if byts, e := json.Marshal(data); e != nil {
			w.WriteHeader(http.StatusInternalServerError)
			// TODO modelo de error
			bts, _ := json.Marshal(nil)
			_, _ = w.Write(bts)
		} else {
			_, _ = w.Write(byts)
		}
	}
}

func handlers(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With Content-Type, Accept")
}
