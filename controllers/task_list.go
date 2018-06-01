package controllers

import (
	"encoding/json"
	"io"

	"github.com/julienschmidt/httprouter"
	m "github.com/tushar9989/hullo/models"
	s "github.com/tushar9989/hullo/storage"
)

func getTaskListFromBody(reqBody io.ReadCloser) (*m.TaskList, error) {
	decoder := json.NewDecoder(reqBody)
	defer reqBody.Close()
	var t m.TaskList
	err := decoder.Decode(&t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func checkGivenTaskListExists(inputTaskList *m.TaskList, ps httprouter.Params, storage s.Storage) bool {
	listId := ps.ByName("listId")
	_, err := storage.GetTaskList(listId)
	if err != nil {
		return false
	}
	inputTaskList.ID = &listId
	return true
}

func AddTaskList(reqBody io.ReadCloser, _ httprouter.Params, storage s.Storage, _ map[string][]string) (interface{}, *ApiError) {
	taskList, err := getTaskListFromBody(reqBody)
	if err != nil {
		return nil, &ApiError{"Unable to parse request", 400}
	}

	if taskList.ID != nil {
		return nil, &ApiError{"Invalid request parameter ID", 400}
	}

	if taskList.Name == nil || *taskList.Name == "" {
		return nil, &ApiError{"Invalid name.", 400}
	}

	createdTaskList, err := storage.AddTaskList(*taskList)
	if err != nil {
		panic(err)
	}

	return createdTaskList, nil
}

func UpdateTaskList(reqBody io.ReadCloser, ps httprouter.Params, storage s.Storage, _ map[string][]string) (interface{}, *ApiError) {
	taskList, err := getTaskListFromBody(reqBody)
	if err != nil {
		return nil, &ApiError{"Unable to parse request", 400}
	}

	exists := checkGivenTaskListExists(taskList, ps, storage)
	if !exists {
		return nil, &ApiError{"Task List not found!", 400}
	}

	err = storage.UpdateTaskList(*taskList)
	if err != nil {
		panic(err)
	}

	return []string{}, nil
}

func deleteAllTasksForList(listId string, storage s.Storage) {
	pageNo := 1
	pageSize := 100
	var tasks []m.Task
	for ok := true; ok; ok = (len(tasks) == pageSize) {
		tasks, err := storage.GetTasks(listId, pageNo, pageSize, nil)
		if err != nil {
			panic(err)
		}
		ids := make([]string, len(tasks))

		for i, task := range tasks {
			ids[i] = *task.ID
		}

		err = storage.DeleteTasks(ids...)
		if err != nil {
			panic(err)
		}
		pageNo++
	}
}

func DeleteTaskList(_ io.ReadCloser, ps httprouter.Params, storage s.Storage, _ map[string][]string) (interface{}, *ApiError) {
	taskList := &m.TaskList{}

	exists := checkGivenTaskListExists(taskList, ps, storage)
	if !exists {
		return nil, &ApiError{"Task List not found!", 400}
	}

	deleteAllTasksForList(*taskList.ID, storage)

	err := storage.DeleteTaskList(*taskList.ID)
	if err != nil {
		panic(err)
	}

	return []string{}, nil
}

func GetTaskLists(_ io.ReadCloser, _ httprouter.Params, storage s.Storage, getParams map[string][]string) (interface{}, *ApiError) {
	pageNo, err := GetIntParamFromMap(getParams, "pageNo", 1)
	if err != nil {
		return nil, &ApiError{"Invalid pageNo", 400}
	}

	pageSize, err := GetIntParamFromMap(getParams, "pageSize", 10)
	if err != nil {
		return nil, &ApiError{"Invalid pageSize", 400}
	}

	search := GetStringParamFromMap(getParams, "search", nil)

	taskLists, err := storage.GetTaskLists(pageNo, pageSize, search)
	if err != nil {
		panic(err)
	}

	return taskLists, nil
}

func GetTaskList(_ io.ReadCloser, ps httprouter.Params, storage s.Storage, _ map[string][]string) (interface{}, *ApiError) {
	listId := ps.ByName("listId")
	taskList, err := storage.GetTaskList(listId)
	if err != nil {
		panic(err)
	}
	return taskList, nil
}
