# go-task-api
[![run tests](https://github.com/bunnyadad/go-task-api/actions/workflows/test.yml/badge.svg)](https://github.com/bunnyadad/go-task-api/actions/workflows/test.yml) [![release](https://github.com/bunnyadad/go-task-api/actions/workflows/docker-publish.yml/badge.svg)](https://github.com/bunnyadad/go-task-api/actions/workflows/docker-publish.yml)

## Introduction
- Implement a Restful task list API.
- Task get/add/update/delete.

## Library Usage
    $ go get github.com/gin-gonic/gin

## Test
	$ go test ./... -coverprofile=./test_results/cover.out && go tool cover -html=./test_results/cover.out -o ./test_results/cover.html

## Installation
    $ docker-compose -f docker-compose.yml up

## Doc
### 1.  GET /tasks (list tasks)
```
{
    "result": [
        {"id": 1, "name": "name", "status": 0}
    ]
}
```

### 2.  POST /task  (create task)
```
request
{
  "name": "買晚餐"
}

response status code 201
{
    "result": {"name": "買晚餐", "status": 0, "id": 1}
}
```

### 3. PUT /task/<id> (update task)
```
request
{
  "name": "買早餐",
  "status": 1
  "id": 1
}

response status code 200
{
  "result":{
    "name": "買早餐",
    "status": 1,
    "id": 1
  }
}
```

### 4. DELETE /task/<id> (delete task)
response status code 200

## License
This project is licensed under the MIT license. See the [LICENSE](LICENSE) file for more info.
