# To Do app

## Dependencies

* Go
* github.com/julienschmidt/httprouter. Used to route the request to the right controller.

## Running on a unix environment

* `go get github.com/julienschmidt/httprouter`
* create a directory for the source code `mkdir -p $HOME/go/src/github.com/tushar9989/todo-app`
* copy the source code `cp -r * $HOME/go/src/github.com/tushar9989/todo-app`
* run `go install github.com/tushar9989/todo-app`
* To run the code `/go/bin/todo-app`

OR If you have Docker

* `./docker-runner.sh`

## API

### Add a Task List

* URL: {host}/api/lists
* Method: POST
* Example

Request

```javascript
{
    "Name": "Task List 1"
}
```

Response

```javascript
{
    "data": {
        "Name": "Task List 1",
        "ID": "3"
    },
    "status": "OK"
}
```

### Add a Task into a List

* URL: {host}/api/lists/{listId}/tasks
* Method: POST
* Example

Request

```javascript
{
    "Name": "Task 1",
    "Description": "First task"
}
```

Response

```javascript
{
    "data": {
        "ListID": "3",
        "Name": "Task 1",
        "ID": "4",
        "Description": "First task",
        "Completed": false
    },
    "status": "OK"
}
```

### Delete a Task from a List

* URL: {host}/api/lists/{listId}/tasks/{taskId}
* Method: DELETE
* Example Response

```javascript
{
    "data": [],
    "status": "OK"
}
```

### Update an existing Task

* URL: {host}/api/lists/{listId}/tasks/{taskId}
* Method: PUT
* Example

Request

```javascript
{
    "Completed": true,
    "Description": "First Task description modified."
}
```

Response

```javascript
{
    "data": [],
    "status": "OK"
}
```

### Delete a Task List

* URL: {host}/api/lists/{listId}
* Method: DELETE
* Example Response

```javascript
{
    "data": [],
    "status": "OK"
}
```

### List Tasks in a List

* URL: {host}/api/lists/{listId}/tasks
* Method: GET
* Example Response

```javascript
{
    "data": [
        {
            "ListID": "3",
            "Name": "Task One",
            "ID": "4",
            "Description": "First Task.",
            "Completed": false
        },
        {
            "ListID": "3",
            "Name": "Task Two",
            "ID": "5",
            "Description": "Second Task.",
            "Completed": true
        },
        {
            "ListID": "3",
            "Name": "Task Three",
            "ID": "6",
            "Description": "Third Task.",
            "Completed": false
        }
    ],
    "status": "OK"
}
```