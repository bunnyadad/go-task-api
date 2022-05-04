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

## License
This project is licensed under the MIT license. See the [LICENSE](LICENSE) file for more info.