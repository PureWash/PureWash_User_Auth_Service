-- name: CreateUser :one
INSERT INTO users (
    id, 
    username,
    full_name,
    phone_number,
    password_hash
) VALUES ($1, $2, $3, $4, $5)
RETURNING id;

-- name: LoginUser :one
SELECT 
    id,
    username,
    password_hash,
    role
FROM users
WHERE username = $1 AND deleted_at IS NULL;

-- name: CheckIfUserExists :one
SELECT
    COUNT(*)
FROM users
WHERE username = $1 OR phone_number = $2 AND deleted_at IS NULL;

-- name: GetUser :one
SELECT
    id,
    username,
    full_name,
    phone_number,
    role
FROM users
WHERE
    id = $1 AND deleted_at IS NULL;

-- name: UpdateUser :exec
UPDATE users
SET 
    username = $1,
    full_name = $2,
    phone_number = $3,
    updated_at = now()
WHERE id = $4 AND password_hash = $5 AND deleted_at IS NULL;

-- name: DeleteUser :exec
UPDATE users
SET 
    deleted_at = now()
WHERE id = $1 AND deleted_at IS NULL;

-- name: UpdatePassword :exec
UPDATE users
SET 
    password_hash = $1
WHERE id = $2 AND password_hash = $3 AND deleted_at IS NULL;


