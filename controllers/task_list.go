package controllers

import (
	"io"

	"github.com/julienschmidt/httprouter"
	s "github.com/tushar9989/hullo/storage"
)

func AddTaskList(_ io.ReadCloser, _ httprouter.Params, _ s.Storage, _ map[string][]string) (interface{}, *ApiError) {
	return "string", nil
}

func UpdateTaskList(_ io.ReadCloser, _ httprouter.Params, _ s.Storage, _ map[string][]string) (interface{}, *ApiError) {
	return "string", nil
}

func DeleteTaskList(_ io.ReadCloser, _ httprouter.Params, _ s.Storage, _ map[string][]string) (interface{}, *ApiError) {
	return "string", nil
}

func GetTasksList(_ io.ReadCloser, _ httprouter.Params, _ s.Storage, _ map[string][]string) (interface{}, *ApiError) {
	return "string", nil
}

func GetTaskList(_ io.ReadCloser, _ httprouter.Params, _ s.Storage, _ map[string][]string) (interface{}, *ApiError) {
	return "", &ApiError{"Can't display record", 404}
}
