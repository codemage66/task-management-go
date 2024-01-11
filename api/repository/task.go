package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/programkingstar/task-management-go.git/api/model"
	"github.com/programkingstar/task-management-go.git/types"
)

type TaskRepository struct {
	store *sql.DB
}

func NewTaskRepo(s *sql.DB) *TaskRepository {
	return &TaskRepository{s}
}

func (db TaskRepository) List(
	c context.Context, arg *types.Pageable,
) (model.Tasks, error) {
	query := `SELECT *, COUNT(*) OVER() AS total FROM "tasks" ORDER by "id" LIMIT $1 OFFSET $2`
	rows, err := db.store.QueryContext(c, query, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tasks model.Tasks

	for rows.Next() {
		var task model.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Priority, &task.DueDate, &arg.Total); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	arg.Calc()
	return tasks, nil
}

func (db TaskRepository) Create(
	c context.Context,
	payload model.TaskRequestPayload,
) (model.Task, error) {
	now := time.Now().UTC()
	query := ` INSERT INTO "tasks" ("title", "desc", "priority", "duedate")
		VALUES ($1, $2, $3, $4)
		RETURNING * `

	row := db.store.QueryRowContext(
		c, query,
		payload.Title, payload.Description, payload.Priority, now,
	)
	var task model.Task

	err := row.Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Priority,
		&task.DueDate,
	)
	if err != nil {
		return task, err
	}
	return task, nil
}

func (db TaskRepository) Update(
	c context.Context,
	arg model.TaskURLParams,
	payload model.TaskRequestPayload,
) (model.Task, error) {
	query := ` UPDATE "tasks" 
		SET "title" = $1, "desc" = $2, "priority" = $3, "duedate" = $4
		WHERE "id" = $5
		RETURNING * `

	row := db.store.QueryRowContext(
		c, query,
		payload.Title, payload.Description, payload.Priority, time.Now().UTC(), arg.ID,
	)
	var task model.Task

	err := row.Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Priority,
		&task.DueDate,
	)
	if err != nil {
		return task, fmt.Errorf("error: task with \"id\" %d not found", arg.ID)
	}
	return task, nil
}

func (db TaskRepository) Get(c context.Context, arg model.TaskURLParams) (model.Task, error) {
	query := ` SELECT * FROM "tasks" WHERE "id" = $1 LIMIT 1 `
	row := db.store.QueryRowContext(c, query, arg.ID)
	var task model.Task

	if err := row.Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Priority,
		&task.DueDate,
	); err != nil {
		return task, err
	}

	return task, nil
}

func (db TaskRepository) Delete(c context.Context, arg model.TaskURLParams) error {
	query := ` DELETE FROM "tasks" WHERE "id" = $1 `
	_, err := db.store.ExecContext(c, query, arg.ID)
	if err != nil {
		return err
	}

	return nil
}
