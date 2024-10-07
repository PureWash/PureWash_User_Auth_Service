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
WHERE id = $4 AND password_hash = $5 AND deleted_at IS NULL;

-- name: DeleteUser :exec
UPDATE employees
SET 
    deleted_at = now()
WHERE id = $1 AND deleted_at IS NULL;

-- name: UpdatePassword :exec
UPDATE employees
SET 
    password_hash = $1
WHERE id = $2 AND password_hash = $3 AND deleted_at IS NULL;


-- name: CreateClient :exec
INSERT INTO clients (
    full_name,
    phone_number,
    latitude,
    longitude
) VALUES($1, $2, $3, $4)
RETURNING (id, full_name);


-- name: UpdateClient :exec
UPDATE clients
SET
    latitude = $1,
    longitude = $2
WHERE
    id = $3 AND deleted_at IS NULL;

-- name: GetClient :one
SELECT
    id,
    full_name,
    phone_number,
    latitude,
    longitude
FROM clients
WHERE  id = $1 AND deleted_at IS NULL;

-- name: DeleteClient :exec

UPDATE clients
SET
    deleted_at = now()
WHERE deleted_at IS NULL AND id = $1;
