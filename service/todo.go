package service

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/bukidai/go-stations/model"
)

// A TODOService implements CRUD of TODO entities.
type TODOService struct {
	db *sql.DB
}

// NewTODOService returns new TODOService.
func NewTODOService(db *sql.DB) *TODOService {
	return &TODOService{
		db: db,
	}
}

// CreateTODO creates a TODO on DB.
func (s *TODOService) CreateTODO(ctx context.Context, subject, description string) (*model.TODO, error) {
	const (
		insert  = `INSERT INTO todos(subject, description) VALUES(?, ?)`
		confirm = `SELECT subject, description, created_at, updated_at FROM todos WHERE id = ?`
	)
	stmt, err := s.db.PrepareContext(ctx, insert)
	if err != nil {
		return nil, err
	}
	res, err := stmt.ExecContext(ctx, subject, description)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	var todo model.TODO
	err = s.db.QueryRowContext(ctx, confirm, int32(id)).Scan(&todo.Subject, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		return nil, err
	}
	todo.ID = id
	return &todo, nil
}

// ReadTODO reads TODOs on DB.
func (s *TODOService) ReadTODO(ctx context.Context, prevID, size int64) ([]*model.TODO, error) {
	const (
		read       = `SELECT id, subject, description, created_at, updated_at FROM todos ORDER BY id DESC LIMIT ?`
		readWithID = `SELECT id, subject, description, created_at, updated_at FROM todos WHERE id < ? ORDER BY id DESC LIMIT ?`
	)
	var rows *sql.Rows
	var err error
	if prevID == 0 {
		rows, err = s.db.QueryContext(ctx, read, size)
		if err != nil {
			return nil, err
		}
	} else {
		rows, err = s.db.QueryContext(ctx, readWithID, prevID, size)
		if err != nil {
			return nil, err
		}
	}
	defer rows.Close()
	var todos []*model.TODO
	for rows.Next() {
		var todo model.TODO
		if err := rows.Scan(&todo.ID, &todo.Subject, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
			return nil, err
		}
		todos = append(todos, &todo)
	}
	if len(todos) == 0 {
		return []*model.TODO{}, nil
	}
	return todos, nil
}

// UpdateTODO updates the TODO on DB.
func (s *TODOService) UpdateTODO(ctx context.Context, id int64, subject, description string) (*model.TODO, error) {
	const (
		update  = `UPDATE todos SET subject = ?, description = ? WHERE id = ?`
		confirm = `SELECT subject, description, created_at, updated_at FROM todos WHERE id = ?`
	)
	stmt, err := s.db.PrepareContext(ctx, update)
	if err != nil {
		return nil, err
	}
	res, err := stmt.ExecContext(ctx, subject, description, id)
	if err != nil {
		return nil, err
	}
	row, err := res.RowsAffected()
	if err != nil {
		return nil, err
	} else if row == 0 {
		return nil, &model.ErrNotFound{}
	}
	var todo model.TODO
	err = s.db.QueryRowContext(ctx, confirm, id).Scan(&todo.Subject, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		return nil, err
	}
	todo.ID = id

	return &todo, nil
}

// DeleteTODO deletes TODOs on DB by ids.
func (s *TODOService) DeleteTODO(ctx context.Context, ids []int64) error {
	const deleteFmt = `DELETE FROM todos WHERE id IN (%s)`

	if len(ids) == 0 {
		return nil
	}

	prepares := make([]string, len(ids))
	args := make([]interface{}, len(ids))
	for i, v := range ids {
		prepares[i] = "?"
		args[i] = v
	}
	stmt, err := s.db.PrepareContext(ctx, fmt.Sprintf(deleteFmt, strings.Join(prepares, ",")))
	if err != nil {
		return err
	}
	res, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	} else if rows == 0 {
		return &model.ErrNotFound{}
	}
	return nil
}
