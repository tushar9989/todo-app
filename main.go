package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
	c "github.com/tushar9989/hullo/controllers"
)

func main() {
	//load creds from env
	//load db
	//load routes with handlers
	router := httprouter.New()
	router.GET("/api/lists", c.Wrapper(c.GetTasksList, nil))
	router.GET("/api/lists/:listId", c.Wrapper(c.GetTaskList, nil))
	router.POST("/api/lists", c.Wrapper(c.AddTaskList, nil))
	router.PUT("/api/lists/:listId", c.Wrapper(c.UpdateTaskList, nil))
	router.DELETE("/api/lists/:listId", c.Wrapper(c.DeleteTaskList, nil))

	router.GET("/api/lists/:listId/tasks", c.Wrapper(c.GetTasks, nil))
	router.GET("/api/lists/:listId/tasks/:taskId", c.Wrapper(c.GetTask, nil))
	router.POST("/api/lists/:listId/tasks", c.Wrapper(c.AddTask, nil))
	router.PUT("/api/lists/:listId/tasks/:taskId", c.Wrapper(c.UpdateTask, nil))
	router.DELETE("/api/lists/:listId/tasks/:taskId", c.Wrapper(c.DeleteTask, nil))
	//start server

	log.Fatal(http.ListenAndServe(":8080", router))
}
