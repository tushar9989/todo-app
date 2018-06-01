package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	s "github.com/tushar9989/hullo/storage"
)

type ApiError struct {
	Message string
	Code    int
}

func Wrapper(h func(io.ReadCloser, httprouter.Params, s.Storage, map[string][]string) (interface{}, *ApiError), storage s.Storage) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		defer func() {
			if r := recover(); r != nil {
				w.WriteHeader(500)
				json.NewEncoder(w).Encode(map[string]string{"status": "FAIL", "message": fmt.Sprintf("%v", r)})
			}
		}()
		response, err := h(r.Body, ps, storage, r.URL.Query())
		if err != nil {
			w.WriteHeader(err.Code)
			json.NewEncoder(w).Encode(map[string]string{"status": "FAIL", "message": err.Message})
		} else {
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(map[string]interface{}{"status": "OK", "data": response})
		}
	}
}

func GetStringParamFromMap(getParams map[string][]string, param string, defaultValue *string) *string {
	if value, found := getParams[param]; found && value[0] != "" {
		return &value[0]
	}

	return defaultValue
}

func GetIntParamFromMap(getParams map[string][]string, param string, defaultValue int) (int, error) {
	if stringValue, found := getParams[param]; found {
		value, err := strconv.Atoi(stringValue[0])
		if err != nil {
			return value, err
		}
		return value, nil
	}

	return defaultValue, nil
}
