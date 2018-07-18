package storage

import (
	"errors"
	"sort"
	"strconv"
	"strings"

	m "github.com/tushar9989/todo-app/internal/models"
)

type inMemoryStorage struct {
	taskListsMap     map[string]m.TaskList
	taskMap          map[string]m.Task
	listIdToTasksMap map[string]map[string]m.Task
	nextTaskId       int
	nextListId       int
}

func NewMockStorage() *inMemoryStorage {
	p := new(inMemoryStorage)
	p.taskListsMap = make(map[string]m.TaskList)
	p.taskMap = make(map[string]m.Task)
	p.listIdToTasksMap = make(map[string]map[string]m.Task)
	p.nextTaskId = 1
	p.nextListId = 1
	return p
}

func (iMS *inMemoryStorage) AddTask(task m.Task) (m.Task, error) {
	task.ID = new(string)
	*task.ID = strconv.Itoa(iMS.nextTaskId)
	iMS.nextTaskId++
	iMS.taskMap[*task.ID] = task
	listMap, ok := iMS.listIdToTasksMap[*task.ListID]
	if !ok {
		listMap = make(map[string]m.Task)
		iMS.listIdToTasksMap[*task.ListID] = listMap
	}
	iMS.listIdToTasksMap[*task.ListID][*task.ID] = task
	return task, nil
}

func (iMS *inMemoryStorage) UpdateTask(task m.Task) error {
	currentTask, ok := iMS.taskMap[*task.ID]
	if !ok {
		return errors.New("Invalid task id")
	}
	if task.Name != nil {
		currentTask.Name = task.Name
	}
	if task.Description != nil {
		currentTask.Description = task.Description
	}
	if task.Completed != nil {
		currentTask.Completed = task.Completed
	}
	iMS.taskMap[*task.ID] = currentTask
	iMS.listIdToTasksMap[*currentTask.ListID][*task.ID] = currentTask
	return nil
}

func (iMS *inMemoryStorage) DeleteTasks(taskIDs ...string) error {
	for _, taskID := range taskIDs {
		currentTask, ok := iMS.taskMap[taskID]
		if !ok {
			return errors.New("Invalid task id")
		}

		delete(iMS.listIdToTasksMap[*currentTask.ListID], *currentTask.ID)
		delete(iMS.taskMap, *currentTask.ID)
	}
	return nil
}

func (iMS *inMemoryStorage) GetTask(taskID string) (m.Task, error) {
	task, ok := iMS.taskMap[taskID]
	if !ok {
		return m.Task{}, errors.New("Invalid task id")
	}
	return task, nil
}

func (iMS *inMemoryStorage) GetTasks(listID string, pageNo int, pageSize int, search *string) ([]m.Task, error) {
	filtered := make([]m.Task, 0)
	listMap, ok := iMS.listIdToTasksMap[listID]
	if !ok {
		return filtered, nil
	}

	var keys []string
	for k, v := range listMap {
		if search != nil && !strings.Contains(*v.Name, *search) {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		filtered = append(filtered, listMap[k])
	}
	pageStart := (pageNo - 1) * pageSize

	if pageStart >= len(filtered) {
		return make([]m.Task, 0), nil
	}

	pageEnd := pageStart + pageSize

	if pageEnd > len(filtered) {
		pageEnd = len(filtered)
	}
	return filtered[pageStart:pageEnd], nil
}

func (iMS *inMemoryStorage) AddTaskList(taskList m.TaskList) (m.TaskList, error) {
	taskList.ID = new(string)
	*taskList.ID = strconv.Itoa(iMS.nextListId)
	iMS.nextListId++
	iMS.taskListsMap[*taskList.ID] = taskList
	iMS.listIdToTasksMap[*taskList.ID] = make(map[string]m.Task)
	return taskList, nil
}

func (iMS *inMemoryStorage) UpdateTaskList(taskList m.TaskList) error {
	currentTaskList, ok := iMS.taskListsMap[*taskList.ID]
	if !ok {
		return errors.New("Invalid task list id")
	}
	if taskList.Name != nil {
		currentTaskList.Name = taskList.Name
	}
	iMS.taskListsMap[*taskList.ID] = currentTaskList
	return nil
}

func (iMS *inMemoryStorage) DeleteTaskList(listId string) error {
	currentTaskList, ok := iMS.taskListsMap[listId]
	if !ok {
		return errors.New("Invalid task list id")
	}
	delete(iMS.taskListsMap, *currentTaskList.ID)
	return nil
}

func (iMS *inMemoryStorage) GetTaskList(taskListID string) (m.TaskList, error) {
	taskList, ok := iMS.taskListsMap[taskListID]
	if !ok {
		return m.TaskList{}, errors.New("Invalid task list id")
	}
	return taskList, nil
}

func (iMS *inMemoryStorage) GetTaskLists(pageNo int, pageSize int, search *string) ([]m.TaskList, error) {
	filtered := make([]m.TaskList, 0)

	var keys []string
	for k, v := range iMS.taskListsMap {
		if search != nil && !strings.Contains(*v.Name, *search) {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		filtered = append(filtered, iMS.taskListsMap[k])
	}
	pageStart := (pageNo - 1) * pageSize

	if pageStart >= len(filtered) {
		return make([]m.TaskList, 0), nil
	}

	pageEnd := pageStart + pageSize

	if pageEnd > len(filtered) {
		pageEnd = len(filtered)
	}
	return filtered[pageStart:pageEnd], nil
}
