package controllers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	s "github.com/tushar9989/hullo/storage"
)

func AddTask(_ http.ResponseWriter, _ *http.Request, _ httprouter.Params, _ s.Storage) (interface{}, *ApiError) {
	return "string", nil
}

func UpdateTask(_ http.ResponseWriter, _ *http.Request, _ httprouter.Params, _ s.Storage) (interface{}, *ApiError) {
	return "string", nil
}

func DeleteTask(_ http.ResponseWriter, _ *http.Request, _ httprouter.Params, _ s.Storage) (interface{}, *ApiError) {
	return "string", nil
}

func GetTasks(_ http.ResponseWriter, _ *http.Request, _ httprouter.Params, _ s.Storage) (interface{}, *ApiError) {
	return "string", nil
}

func GetTask(_ http.ResponseWriter, _ *http.Request, _ httprouter.Params, _ s.Storage) (interface{}, *ApiError) {
	return "", &ApiError{nil, "Can't display record", 500}
}