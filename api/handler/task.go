package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/programkingstar/task-management-go.git/api/model"
	repo "github.com/programkingstar/task-management-go.git/api/repository"
	"github.com/programkingstar/task-management-go.git/types"
	"github.com/programkingstar/task-management-go.git/util"
)

type TasksResource struct {
	repo *repo.TaskRepository
}

func New(r *repo.TaskRepository) *TasksResource {
	return &TasksResource{r}
}

func (rs TasksResource) Routes(route chi.Router) {
	route.Get("/", rs.List)
	route.Post("/", rs.Create)
	route.Route("/{id}",
		func(r chi.Router) {
			r.Get("/", rs.Get)
			r.Delete("/", rs.Delete)
			r.Put("/", rs.Update)
		})
}

// Get returns a task by id
// @Summary Get a task
// @Description Get a task by id
// @Tags task
// @Accept  json
// @Produce  json
// @Param id path string true "Task ID"
// @Success 200 {object} types.JSONResult{data=model.Task}
// @Failure 400 {string} string "error: id is invalid"
// @Failure 404 {string} string "error: task not found"
// @Router /tasks/{id} [get]
func (rs TasksResource) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var reqDTO model.TaskURLParams
	err := reqDTO.Parse(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := rs.repo.Get(r.Context(), reqDTO)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		return
	}

	jsonData, err := json.Marshal(data.CreateResponseDto())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

// List returns a list of tasks
// @Summary List tasks

// @Description Get List tasks
// @Tags task
// @Accept  json
// @Produce  json
// @Param offset query string false "string default example" default(0) example(1)
// @Param limit query string false "string default example" default(10) example(20)
// @Success 200 {object} types.JSONResult{data=model.Tasks,paginate=types.Pageable,length=int}
// @Failure 400 {string} string "error: offset or limit is invalid"
// @Router /tasks [get]
func (rs TasksResource) List(w http.ResponseWriter, r *http.Request) {
	offset := r.URL.Query().Get("offset")
	limit := r.URL.Query().Get("limit")
	p := types.Pageable{}.Parse(limit, offset)

	data, err := rs.repo.List(r.Context(), &p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	jsonData, err := json.Marshal(data.CreateResponseDto(&p))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(jsonData)
}

// Delete deletes a task by id
// @Summary Delete a task
// @Description Delete a task by id
// @Tags task
// @Accept  json
// @Produce  json
// @Param id path string true "Task ID"
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "error: id is invalid"
// @Failure 404 {string} string "error: task not found"
// @Router /tasks/{id} [delete]
func (rs TasksResource) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var reqDTO model.TaskURLParams
	err := reqDTO.Parse(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = rs.repo.Delete(r.Context(), reqDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(
			[]byte(
				"error: Cannot delete the task with id " + id + " because it is not a valid",
			),
		)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// Create creates a task
// @Summary Create a task
// @Description Create a task
// @Tags task
// @Accept  json
// @Produce  json
// @Param request body model.TaskRequestPayload true "default"
// @Success 201 {object} types.JSONResult{data=model.Task}
// @Failure 400 {string} string "Bad Request: Invalid payload"
// @Router /tasks [post]
func (rs TasksResource) Create(w http.ResponseWriter, r *http.Request) {
	var payload model.TaskRequestPayload
	if err := util.ParseRequestBody(r, &payload); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error: payload is invalid or missing"))
		return
	}

	if _, err := payload.Validate(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	data, err := rs.repo.Create(r.Context(), payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	jsonData, err := json.Marshal(data.CreateResponseDto())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonData)
}

// Create creates a task
// @Summary Create a task
// @Description Create a task
// @Tags task
// @Accept  json
// @Produce  json
// @Param request body model.TaskRequestPayload true "default"
// @Success 200 {object} types.JSONResult{data=model.Task}
// @Failure 400 {string} string "Bad Request: Invalid payload"
// @Router /tasks/{id} [put]
func (rs TasksResource) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var reqDTO model.TaskURLParams
	err := reqDTO.Parse(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var payload model.TaskRequestPayload
	if err = util.ParseRequestBody(r, &payload); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error: payload is invalid or missing"))
		return
	}

	if _, err = payload.Validate(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	data, err := rs.repo.Update(r.Context(), reqDTO, payload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	jsonData, err := json.Marshal(data.CreateResponseDto())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
