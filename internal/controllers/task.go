package controllers

import (
	"encoding/json"
	"io"

	"github.com/julienschmidt/httprouter"
	m "github.com/tushar9989/todo-app/internal/models"
	s "github.com/tushar9989/todo-app/internal/storage"
)

func getTaskFromBody(reqBody io.ReadCloser) (*m.Task, error) {
	decoder := json.NewDecoder(reqBody)
	defer reqBody.Close()
	var t m.Task
	err := decoder.Decode(&t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func checkGivenTaskExists(inputTask *m.Task, ps httprouter.Params, storage s.Storage) bool {
	taskId := ps.ByName("taskId")
	listId := ps.ByName("listId")
	existingTask, err := storage.GetTask(taskId)
	if err != nil || *existingTask.ListID != listId {
		return false
	}
	inputTask.ID = &taskId
	inputTask.ListID = &listId
	return true
}

func AddTask(reqBody io.ReadCloser, ps httprouter.Params, storage s.Storage, _ map[string][]string) (interface{}, *ApiError) {
	task, err := getTaskFromBody(reqBody)
	if err != nil {
		return nil, &ApiError{"Unable to parse request", 400}
	}

	if task.ID != nil {
		return nil, &ApiError{"Invalid request parameter ID", 400}
	}

	if task.Completed != nil {
		return nil, &ApiError{"Completed cannot be set when creating.", 400}
	}

	if task.Name == nil || *task.Name == "" {
		return nil, &ApiError{"Invalid name.", 400}
	}

	if task.Description == nil || *task.Description == "" {
		return nil, &ApiError{"Invalid description.", 400}
	}

	listExists := checkGivenTaskListExists(&m.TaskList{}, ps, storage)
	if !listExists {
		return nil, &ApiError{"Invalid List Id.", 400}
	}

	listId := ps.ByName("listId")
	task.ListID = &listId
	task.Completed = new(bool)

	createdTask, err := storage.AddTask(*task)
	if err != nil {
		panic(err)
	}

	return createdTask, nil
}

func UpdateTask(reqBody io.ReadCloser, ps httprouter.Params, storage s.Storage, _ map[string][]string) (interface{}, *ApiError) {
	task, err := getTaskFromBody(reqBody)
	if err != nil {
		return nil, &ApiError{"Unable to parse request", 400}
	}

	exists := checkGivenTaskExists(task, ps, storage)
	if !exists {
		return nil, &ApiError{"Task not found!", 400}
	}

	err = storage.UpdateTask(*task)
	if err != nil {
		panic(err)
	}

	return []string{}, nil
}

func DeleteTask(_ io.ReadCloser, ps httprouter.Params, storage s.Storage, _ map[string][]string) (interface{}, *ApiError) {
	task := &m.Task{}

	exists := checkGivenTaskExists(task, ps, storage)
	if !exists {
		return nil, &ApiError{"Task not found!", 400}
	}

	err := storage.DeleteTasks(*task.ID)
	if err != nil {
		panic(err)
	}

	return []string{}, nil
}

func GetTask(_ io.ReadCloser, ps httprouter.Params, storage s.Storage, _ map[string][]string) (interface{}, *ApiError) {
	taskId := ps.ByName("taskId")
	task, err := storage.GetTask(taskId)
	if err != nil {
		panic(err)
	}
	return task, nil
}

func GetTasks(_ io.ReadCloser, ps httprouter.Params, storage s.Storage, getParams map[string][]string) (interface{}, *ApiError) {
	pageNo, err := GetIntParamFromMap(getParams, "pageNo", 1)
	if err != nil {
		return nil, &ApiError{"Invalid pageNo", 400}
	}

	pageSize, err := GetIntParamFromMap(getParams, "pageSize", 10)
	if err != nil {
		return nil, &ApiError{"Invalid pageSize", 400}
	}

	search := GetStringParamFromMap(getParams, "search", nil)

	tasks, err := storage.GetTasks(ps.ByName("listId"), pageNo, pageSize, search)
	if err != nil {
		panic(err)
	}

	return tasks, nil
}
