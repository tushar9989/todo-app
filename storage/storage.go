package storage

import (
	"github.com/tushar9989/hullo/models"
)

type Storage interface {
	addTask(task models.Task) (models.Task, error)
	updateTask(task models.Task) error
	deleteTask(taskID string) error
	getTasks(listID string, pageNo int, pageSize int, search string) ([]models.Task, error)
	getTask(taskID string) (models.Task, error)

	addTaskList(taskList models.TaskList) (models.TaskList, error)
	updateTaskList(taskList models.TaskList) error
	deleteTaskList(taskListID string) error
	getTaskLists(pageNo int, pageSize int, search string) ([]models.TaskList, error)
	getTaskList(taskListID string) (models.TaskList, error)
}