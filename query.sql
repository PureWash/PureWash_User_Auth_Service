-- name: CreateUser :one
INSERT INTO employees (
    id, 
    username,
    full_name,
    phone_number,
    password_hash,
    role 
) VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id;

-- name: LoginUser :one
SELECT 
    id,
    username,
    password_hash,
    role
FROM employees
WHERE username = $1 AND deleted_at IS NULL;

-- name: CheckIfUserExists :one
SELECT
    COUNT(*)
FROM employees
WHERE username = $1 OR phone_number = $2 AND deleted_at IS NULL;

-- name: GetUser :one
SELECT
    id,
    username,
    full_name,
    phone_number,
    role
FROM employees
WHERE
    id = $1 AND deleted_at IS NULL;

-- name: UpdateUser :exec
UPDATE employees
SET 
    username = $1,
    full_name = $2,
    phone_number = $3,
    updated_at = now()
WHERE id = $4 AND deleted_at IS NULL;

-- name: DeleteUser :exec
UPDATE employees
SET 
    deleted_at = now()
WHERE id = $1 AND deleted_at IS NULL;

-- name: UpdatePassword :exec
UPDATE employees
SET 
    password_hash = $1,
    updated_at = now()
WHERE id = $2 AND deleted_at IS NULL;

-- name: GetUserPassword :one
SELECT
    password_hash
FROM employees
WHERE id = $1;