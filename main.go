package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	c "github.com/tushar9989/hullo/controllers"
	"github.com/tushar9989/hullo/storage"
)

func main() {
	dataSource := storage.NewMockStorage()

	router := httprouter.New()
	router.GET("/api/lists", c.Wrapper(c.GetTaskLists, dataSource))
	router.GET("/api/lists/:listId", c.Wrapper(c.GetTaskList, dataSource))
	router.POST("/api/lists", c.Wrapper(c.AddTaskList, dataSource))
	router.PUT("/api/lists/:listId", c.Wrapper(c.UpdateTaskList, dataSource))
	router.DELETE("/api/lists/:listId", c.Wrapper(c.DeleteTaskList, dataSource))

	router.GET("/api/lists/:listId/tasks", c.Wrapper(c.GetTasks, dataSource))
	router.GET("/api/lists/:listId/tasks/:taskId", c.Wrapper(c.GetTask, dataSource))
	router.POST("/api/lists/:listId/tasks", c.Wrapper(c.AddTask, dataSource))
	router.PUT("/api/lists/:listId/tasks/:taskId", c.Wrapper(c.UpdateTask, dataSource))
	router.DELETE("/api/lists/:listId/tasks/:taskId", c.Wrapper(c.DeleteTask, dataSource))

	log.Fatal(http.ListenAndServe(":8080", router))
}
