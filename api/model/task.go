package model

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/programkingstar/task-management-go.git/types"
)

// Task is a task entity model
type Task struct {
	ID         	int64     `json:"id"          example:"1"`
	Title    	string    `json:"title"     example:"Task1"`
	Description string    `json:"desc"     example:"Test task."`
	Priority   	int64     `json:"priority"   example:"1"`
	DueDate  	time.Time `json:"duedate"  example:"2024-01-01T00:00:00Z"`
}

// CreateResponseDto creates a response dto for a single task
func (q Task) CreateResponseDto() types.JSONResult {
	return types.JSONResult{
		Data:    q,
		Code:    200,
		Message: "success",
	}
}

// Tasks is a collection of Task
type Tasks []Task

// Len is the number of elements in the collection.
func (q Tasks) Len() int {
	return len(q)
}

// Less reports whether the element with
func (q Tasks) CreateResponseDto(pag *types.Pageable) types.JSONResultWithPaginate {
	if pag.Total > 0 {
		return types.JSONResultWithPaginate{
			Message:  "success",
			Code:     200,
			Data:     q,
			Length:   len(q),
			Paginate: pag,
		}
	}

	return types.JSONResultWithPaginate{
		Code:    200,
		Message: "success",
		Data:    q,
		Length:  len(q),
	}
}

// CreateTaskRequest is the request payload for creating a task
type CreateTaskRequest struct {
	Title    	string    `json:"title"     example:"Task1"`
	Description string    `json:"desc"     example:"Test task."`
	Priority   	int64     `json:"priority"   example:"1"`
	DueDate  	time.Time `json:"duedate"  example:"2024-01-01T00:00:00Z"`
}

type TaskRequestPayload struct {
	Title    	string    `json:"title"     example:"Task1"`
	Description string    `json:"desc"     example:"Test task."`
	Priority   	int64     `json:"priority"   example:"1"`
	DueDate  	time.Time `json:"duedate"  example:"2024-01-01T00:00:00Z"`
}

// Validate validates the CreateTaskRequest payload
func (req *TaskRequestPayload) Validate() (bool, error) {
	if req == nil {
		return false, errors.New("error: payload is required")
	}

	if req.Title == "" {
		return false, errors.New("error: field \"title\" is required and must be a string")
	}

	if req.Description == "" {
		return false, errors.New(
			"error: field \"desc\" is required and must be a positive number",
		)
	}

	if req.Priority <= 0 {
		return false, errors.New(
			"error: field \"priority\" is required and must be a positive number",
		)
	}

	return true, nil
}

// Parse parses the CreateTaskRequest payload
func (req *TaskRequestPayload) Parse() Task {
	return Task{
		Title:    		req.Title,
		Description:   	req.Description,
		Priority: 		req.Priority,
		DueDate: 		req.DueDate,
	}
}

// DataTask is the response dto for a single task
type DataTask struct {
	Data Task `json:"data" example:"{...}"`
}

// ListDataTasks is the response dto for a list of tasks
type ListDataTasks struct {
	Data     Tasks          `json:"data"`
	Length   int             `json:"length"`
	Paginate *types.Pageable `json:"paginate"`
}

// CreateTaskResponse is the response dto for creating a task
type CreateTaskResponse *DataTask

// ListTaskResponse is the response dto for a list of tasks
type ListTaskResponse ListDataTasks

// ListTaskRequest is the request payload for a list of tasks
type ListTaskRequest types.Pageable

// TaskURLParams is the url parameters for a single task
type TaskURLParams struct {
	ID int64 `json:"id" example:"1" validate:"required"`
}

// Parse parses the id url parameter
func (req *TaskURLParams) Parse(value string) error {
	p, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return fmt.Errorf("param \"id\" is required and must be a number")
	}

	req.ID = p
	return nil
}

// GetTaskResponse is the response dto for a single task
type GetTaskResponse DataTask

// New creates a new Task model
func (q *Task) New(title string, desc string) Task {
	return Task{
		Title:  title,
		Description: desc,
	}
}

// Validate validates the Task model
func (q *Task) Validate() bool {
	if q == nil {
		return false
	}

	if q.Title == "" {
		return false
	}

	if q.Description == "" {
		return false
	}

	return true
}

// CreateListTaskReponse creates a response dto for a list of tasks
func CreateListTaskReponse(tasks Tasks) ListTaskResponse {
	return ListTaskResponse{
		Data: tasks,
	}
}
