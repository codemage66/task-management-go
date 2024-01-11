## Start

#### Required\*

- GNU Make [Read Docs](https://www.gnu.org/software/make/)
- Docker [Read Docs](https://docs.docker.com/)
- Go version go1.21.5 [Read Docs](https://go.dev/dl/)

#### Setup project

- ##### Clone the Repository

  Make sure your git is setup using ssh-key and already in your code working directory

  ```
  git clone git@github.com:programkingstar/task-management-go.git
  cd n
  ```

- ##### Setup and Run server

  Make sure you're already have GNU make to run the command bellow:

  ```
  make setup
  make run
  ```

- ##### Test Server
  This will run all test code if you wanna check the tests
  ```
  make setup
  make run
  ```
- ##### Test API

  **Get Task** (positive int)

  ```
  # GET /tasks/:id
  curl -i -X GET http://127.0.0.1:3000/api/tasks/1
  ```

  **Response**

  ```
  HTTP/1.1 200 OK
  Content-Type: application/json; charset=UTF-8
  Date: Fri, 29 Dec 2023 13:59:04 GMT
  Content-Length: 158

  {
    "data": 
    {
      "id":1,
      "title":"My first task",
      "desc":"Test",
      "priority":1,
      "duedate":"2024-01-11T23:02:24.317453Z"
    }
  }

  ```

  **List Tasks**

  ```
  # GET /tasks/:id (positive int)
  curl -i -X GET http://127.0.0.1:3000/api/tasks
  ```

  **Response**

  ```
  Content-Type: application/json; charset=UTF-8
  Date: Fri, 29 Dec 2023 14:06:35 GMT
  Content-Length: 569

  {
    "data": [
        {
            "id": 1,
            "title": "My first task",
            "desc": "Test",
            "priority": 1,
            "duedate": "2024-01-11T23:02:24.317453Z"
        },
        {
            "id": 2,
            "title": "My second task",
            "desc": "Test",
            "priority": 2,
            "duedate": "2024-01-11T23:02:24.317453Z"
        },
        {
            "id": 3,
            "title": "My third task",
            "desc": "Test",
            "priority": 3,
            "duedate": "2024-01-11T23:02:24.317453Z"
        }
    ],
    "length": 3,
    "paginate": {
        "offset": 0,
        "limit": 10,
        "total": 3,
        "prev": 0,
        "next": 3,
        "has_next": false,
        "has_prev": false
    }
  }

  ```

  **Create Task**

  ```
  curl -v -XPOST -H "Content-type: application/json" \
  -d '{"title": "Task4", "desc": "Hello, Task" ,"priority": 1 }' \
  '127.0.0.1:3000/api/tasks'

  ```

  **Response**

  ```
  Trying 127.0.0.1:3000...*
  Connected to 127.0.0.1 (127.0.0.1) port 3000
  > POST /api/tasks HTTP/1.1
  > Host: 127.0.0.1:3000
  > User-Agent: curl/8.4.0
  > Accept: */*
  > Content-type: application/json
  > Content-Length: 76

  < HTTP/1.1 201 Created
  < Content-Type: application/json; charset=UTF-8
  < Date: Fri, 29 Dec 2023 14:14:25 GMT
  < Content-Length: 158
  <* Connection #0 to host 127.0.0.1 left intact

  {
    "data": {
        "id": 4,
        "title": "Task4",
        "desc": "Hello, Task",
        "priority": 1,
        "duedate": "2024-01-11T23:07:42.478092Z"
    }
  }
  ```

  **Update Task** (id: positive int)

  ```
  curl -v -XPUT -H "Content-type: application/json" \
  -d '{"id": 4, "title": "NewTask", "desc": "Hello", "priority": 2 }'\  	
  '127.0.0.1:3000/api/tasks/4'
  ```

  **Response**

  ```
  Trying 127.0.0.1:3000...*
  Connected to 127.0.0.1 (127.0.0.1) port 3000
  > POST /api/tasks/4 HTTP/1.1
  > Host: 127.0.0.1:3000
  > User-Agent: curl/8.4.0
  > Accept: */*
  > Content-type: application/json
  > Content-Length: 69
  >
  < HTTP/1.1 200 OK
  < Content-Type: application/json; charset=UTF-8
  < Date: Fri, 29 Dec 2023 14:24:39 GMT
  < Content-Length: 151
  <* Connection #0 to host 127.0.0.1 left intact

  {
    "data": {
        "id": 4,
        "title": "NewTask",
        "desc": "Hello",
        "priority": 2,
        "duedate": "2024-01-11T23:11:59.742527Z"
    }
  }
  ```

  **Delete Task**

  ```
  curl -v -XDELETE '127.0.0.1:3000/api/tasks/4'
  ```

  **Response**

  ```
  *   Trying 127.0.0.1:3000...
  * Connected to 127.0.0.1 (127.0.0.1) port 3000
  > DELETE /api/tasks/4 HTTP/1.1
  > Host: 127.0.0.1:3000
  > User-Agent: curl/8.4.0
  > Accept: */*
  >
  < HTTP/1.1 200 OK
  < Date: Fri, 29 Dec 2023 14:28:17 GMT
  < Content-Length: 0
  <
  * Connection #0 to host 127.0.0.1 left intact
  ```

## OpenAPI Doc

See Documentation REST API in here:
[Read Docs](https://github.com/programkingstar/task-management-go/tree/main/docs/swagger.json) or if you're running the server: [Go to Open API](http://localhost:3000/openapi)

## TODO

- [ ] Create an Open API Docs
- [ ] Cover all Test
- [ ] Create UI
