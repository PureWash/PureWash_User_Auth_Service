// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package storage

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

const checkIfUserExists = `-- name: CheckIfUserExists :one
SELECT
    COUNT(*)
FROM employees
WHERE username = $1 OR phone_number = $2 AND deleted_at IS NULL
`

type CheckIfUserExistsParams struct {
	Username    string
	PhoneNumber string
}

func (q *Queries) CheckIfUserExists(ctx context.Context, arg CheckIfUserExistsParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, checkIfUserExists, arg.Username, arg.PhoneNumber)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO employees (
    id, 
    username,
    full_name,
    phone_number,
    password_hash,
    role 
) VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id
`

type CreateUserParams struct {
	ID           uuid.UUID
	Username     string
	FullName     string
	PhoneNumber  string
	PasswordHash string
	Role         string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.Username,
		arg.FullName,
		arg.PhoneNumber,
		arg.PasswordHash,
		arg.Role,
	)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const deleteUser = `-- name: DeleteUser :exec
UPDATE employees
SET 
    deleted_at = now()
WHERE id = $1 AND deleted_at IS NULL
`

func (q *Queries) DeleteUser(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT
    id,
    username,
    full_name,
    phone_number,
    role
FROM employees
WHERE
    id = $1 AND deleted_at IS NULL
`

type GetUserRow struct {
	ID          uuid.UUID
	Username    string
	FullName    string
	PhoneNumber string
	Role        string
}

func (q *Queries) GetUser(ctx context.Context, id uuid.UUID) (GetUserRow, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i GetUserRow
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.FullName,
		&i.PhoneNumber,
		&i.Role,
	)
	return i, err
}

const getUserPassword = `-- name: GetUserPassword :one
SELECT
    password_hash
FROM employees
WHERE id = $1
`

func (q *Queries) GetUserPassword(ctx context.Context, id uuid.UUID) (string, error) {
	row := q.db.QueryRowContext(ctx, getUserPassword, id)
	var password_hash string
	err := row.Scan(&password_hash)
	return password_hash, err
}

const loginUser = `-- name: LoginUser :one
SELECT 
    id,
    username,
    password_hash,
    role
FROM employees
WHERE username = $1 AND deleted_at IS NULL
`

type LoginUserRow struct {
	ID           uuid.UUID
	Username     string
	PasswordHash string
	Role         string
}

func (q *Queries) LoginUser(ctx context.Context, username string) (LoginUserRow, error) {
	row := q.db.QueryRowContext(ctx, loginUser, username)
	var i LoginUserRow
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.PasswordHash,
		&i.Role,
	)
	return i, err
}

const updatePassword = `-- name: UpdatePassword :exec
UPDATE employees
SET 
    password_hash = $1,
    updated_at = now()
WHERE id = $2 AND deleted_at IS NULL
`

type UpdatePasswordParams struct {
	PasswordHash string
	ID           uuid.UUID
}

func (q *Queries) UpdatePassword(ctx context.Context, arg UpdatePasswordParams) error {
	_, err := q.db.ExecContext(ctx, updatePassword, arg.PasswordHash, arg.ID)
	return err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE employees
SET 
    username = $1,
    full_name = $2,
    phone_number = $3,
    updated_at = now()
WHERE id = $4 AND deleted_at IS NULL
`

type UpdateUserParams struct {
	Username    string
	FullName    string
	PhoneNumber string
	ID          uuid.UUID
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.ExecContext(ctx, updateUser,
		arg.Username,
		arg.FullName,
		arg.PhoneNumber,
		arg.ID,
	)
	return err
}

const updateUserAdmin = `-- name: UpdateAdminUser :exec
UPDATE employees
SET 
    username = $1,
    full_name = $2,
    phone_number = $3,
	role = $4,
	password_hash = $5,
    updated_at = now()
WHERE id = $6 AND deleted_at IS NULL
`

type UpdateUserAdminParams struct {
	Username    string
	FullName    string
	PhoneNumber string
	Role 		string
	Password 	string
	ID          uuid.UUID
}

func (q *Queries) UpdateAdminUser(ctx context.Context, arg UpdateUserAdminParams) error {
	_, err := q.db.ExecContext(ctx, updateUserAdmin,
		arg.Username,
		arg.FullName,
		arg.PhoneNumber,
		arg.Role,
		arg.Password,
		arg.ID,
	)
	return err
}

var getAllUsers = ` -- name: GetAllUsers :one
SELECT
	id,
	username,
	full_name,
	phone_number,
	role
FROM employees
WHERE deleted_at IS NULL
`
var getAllUsersCount = ` -- name: GetAllUsers :one
SELECT
	COUNT(*)
FROM employees
WHERE deleted_at IS NULL
`

type GetAllUsersParams struct {
	Username 	string 
	FullName 	string
	PhoneNumber string
	Role 		string
	Limit 		int
	Offset 		int
}

type Employees struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	Role        string `json:"role"`
}
type GetAllUsersRow struct {
	Employees  []Employees `json:"employee"`
	Limit      int         `json:"limit"`
	Offset     int         `json:"offset"`
	TotalCount int         `json:"total_count"`
}

func (q *Queries) GetAllUsers(ctx context.Context, arg GetAllUsersParams) (*GetAllUsersRow, error) {
	var (
		filter string
		args  []interface{}
	)

	if arg.FullName != "" {
		filter += fmt.Sprintf(" AND full_name ILIKE $%d", len(args)+1)
		args = append(args, fmt.Sprintf("%%%s%%", arg.FullName))
	}
	if arg.Username != "" {
		filter += fmt.Sprintf(" AND username ILIKE $%d", len(args)+1)
		args = append(args, fmt.Sprintf("%%%s%%", arg.Username))
	}
	if arg.PhoneNumber != "" {
		filter += fmt.Sprintf(" AND phone_number ILIKE $%d", len(args)+1)
		args = append(args, fmt.Sprintf("%%%s%%", arg.PhoneNumber))
	}
	if arg.Role != "" {
		filter += fmt.Sprintf(" AND role ILIKE $%d", len(args)+1)
		args = append(args, fmt.Sprintf("%%%s%%", arg.Role))
	}

	row := q.db.QueryRowContext(ctx, getAllUsersCount+filter, args...)
	var count int
	err := row.Scan(&count)
	if err != nil {
		return nil, err
	}

	filter += fmt.Sprintf(" OFFSET %d LIMIT %d", arg.Offset, arg.Limit)

	rows, err := q.db.QueryContext(ctx, getAllUsers+filter, args...)
	if err != nil {
		return nil, err
	}

	var employees []Employees
	for rows.Next() {
		var i Employees
		
		err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.FullName,
			&i.PhoneNumber,
			&i.Role,
		)
		if err != nil {
			return nil, err
		}
		employees = append(employees, i)
	}

	if err = row.Err(); err != nil {
		return nil, err
	}

	return &GetAllUsersRow{
		Employees: employees,
		Limit: arg.Limit,
		Offset: arg.Offset+1,
		TotalCount: count,
	}, nil
}