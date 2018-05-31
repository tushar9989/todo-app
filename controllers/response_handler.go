package controllers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
	s "github.com/tushar9989/hullo/storage"
)

type ApiError struct {
    Error   error
    Message string
    Code    int
}

//type ResponseHandler func(http.ResponseWriter, *http.Request, ps httprouter.Params) (interface{}, *ApiError)

func Wrapper(h func(http.ResponseWriter, *http.Request, httprouter.Params, s.Storage) (interface{}, *ApiError), storage s.Storage) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		response, err := h(w, r, ps, storage)
		if err != nil {
			w.WriteHeader(err.Code)
			json.NewEncoder(w).Encode(map[string]string{"message": err.Message})
		} else {
			w.WriteHeader(http.StatusCreated)
    		json.NewEncoder(w).Encode(response)
		}
	}
}