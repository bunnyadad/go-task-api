package rest

type TaskResponse struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Status bool   `json:"status"`
}

type TaskRequest struct {
	Name   string `json:"name"`
	Status bool   `json:"status"`
}

type TasksResponse struct {
	Tasks []TaskResponse `json:"result"`
}
