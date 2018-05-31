package controllers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	s "github.com/tushar9989/hullo/storage"
)

func AddTaskList(_ http.ResponseWriter, _ *http.Request, _ httprouter.Params, _ s.Storage) (interface{}, *ApiError) {
	return "string", nil
}

func UpdateTaskList(_ http.ResponseWriter, _ *http.Request, _ httprouter.Params, _ s.Storage) (interface{}, *ApiError) {
	return "string", nil
}

func DeleteTaskList(_ http.ResponseWriter, _ *http.Request, _ httprouter.Params, _ s.Storage) (interface{}, *ApiError) {
	return "string", nil
}

func GetTasksList(_ http.ResponseWriter, _ *http.Request, _ httprouter.Params, _ s.Storage) (interface{}, *ApiError) {
	return "string", nil
}

func GetTaskList(_ http.ResponseWriter, _ *http.Request, _ httprouter.Params, _ s.Storage) (interface{}, *ApiError) {
	return "", &ApiError{nil, "Can't display record", 404}
}