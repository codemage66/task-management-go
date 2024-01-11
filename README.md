## Start

#### Required\*

- GNU Make [Read Docs](https://www.gnu.org/software/make/)
- Docker [Read Docs](https://docs.docker.com/)
- Go version go1.21.5 [Read Docs](https://go.dev/dl/)

#### Setup project

- ##### Clone the Repository

  Make sure your git is setup using ssh-key and already in your code working directory

  ```
  git clone git@github.com:Kbgjtn/n.git
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

  **Get Quote** (positive int)

  ```
  # GET /quotes/:id
  curl -i -X GET http://127.0.0.1:3000/api/quotes/1
  ```

  **Response**

  ```
  HTTP/1.1 200 OK
  Content-Type: application/json; charset=UTF-8
  Date: Fri, 29 Dec 2023 13:59:04 GMT
  Content-Length: 158

  {
    "data": {
      "id": 1,
      "content": "My first quote",
      "author_id": 1,
      "category_id": 1,
      "created_at": "2023-12-29T13:53:57.33503Z",
      "updated_at": "2023-12-29T13:53:57.33503Z"
      }
    }

  ```

  **List Quotes**

  ```
  # GET /quotes/:id (positive int)
  curl -i -X GET http://127.0.0.1:3000/api/quotes
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
        "content": "My first quote",
        "author_id": 1,
        "category_id": 1,
        "created_at": "2023-12-29T13:53:57.33503Z",
        "updated_at": "2023-12-29T13:53:57.33503Z"
      },
      {
        "id": 2,
        "content": "My second quote",
        "author_id": 2,
        "category_id": 2,
        "created_at": "2023-12-29T13:53:57.33503Z",
        "updated_at": "2023-12-29T13:53:57.33503Z"
      },
      {
        "id": 3,
        "content": "My third quote",
        "author_id": 3,
        "category_id": 3,
        "created_at": "2023-12-29T13:53:57.33503Z",
        "updated_at": "2023-12-29T13:53:57.33503Z"
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

  **Create Quote**

  ```
  curl -v -XPOST -H "Content-type: application/json" \
  -d '{"author_id": 1, "content": "Hello World!", "category_id": 1 }' \
  '127.0.0.1:3000/api/quotes'

  ```

  **Response**

  ```
  Trying 127.0.0.1:3000...*
  Connected to 127.0.0.1 (127.0.0.1) port 3000
  > POST /api/quotes HTTP/1.1
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
  		"id":4,
  		"content":"Hello World!",
  		"author_id":1,
  		"category_id":1,
  		"created_at":"2023-12-29T14:14:25.889183Z",
  		"updated_at":"2023-12-29T14:14:25.889183Z"
  	}
  }
  ```

  **Update Quote** (id: positive int)

  ```
  curl -v -XPUT -H "Content-type: application/json" \
  -d '{"author_id": 1, "content": "Test!", "category_id": 1 }'\  	'127.0.0.1:3000/api/quotes/4'
  ```

  **Response**

  ```
  Trying 127.0.0.1:3000...*
  Connected to 127.0.0.1 (127.0.0.1) port 3000
  > POST /api/quotes/4 HTTP/1.1
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
  		"id":4,
  		"content":"Test!",
  		"author_id":1,
  		"category_id":1,
  		"created_at":"2023-12-29T14:14:25.889183Z",
  		"updated_at":"2023-12-29T14:14:25.889183Z"
  	}
  }
  ```

  **Delete Quote**

  ```
  curl -v -XDELETE '127.0.0.1:3000/api/quotes/4'
  ```

  **Response**

  ```
  *   Trying 127.0.0.1:3000...
  * Connected to 127.0.0.1 (127.0.0.1) port 3000
  > DELETE /api/quotes/4 HTTP/1.1
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
[Read Docs](https://github.com/Kbgjtn/n/tree/main/docs/swagger.json) or if you're running the server: [Go to Open API](http://localhost:3000/openapi)

## TODO

- [ ] Create an Open API Docs
- [ ] Cover all Test
- [ ] Create UI
