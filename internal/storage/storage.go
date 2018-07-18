package storage

import (
	"github.com/tushar9989/todo-app/internal/models"
)

type Storage interface {
	AddTask(task models.Task) (models.Task, error)
	UpdateTask(task models.Task) error
	DeleteTasks(taskIDs ...string) error
	GetTasks(listID string, pageNo int, pageSize int, search *string) ([]models.Task, error)
	GetTask(taskID string) (models.Task, error)

	AddTaskList(taskList models.TaskList) (models.TaskList, error)
	UpdateTaskList(taskList models.TaskList) error
	DeleteTaskList(taskListID string) error
	GetTaskLists(pageNo int, pageSize int, search *string) ([]models.TaskList, error)
	GetTaskList(taskListID string) (models.TaskList, error)
}
