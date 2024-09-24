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
WHERE password_hash = $4 AND deleted_at IS NULL;

-- name: DeleteUser :exec
UPDATE users
SET 
    deleted_at = now()
WHERE id = $1 AND deleted_at IS NULL;

-- name: UpdatePassword :exec
UPDATE users
SET 
    password_hash = $1
WHERE password_hash = $2 AND deleted_at IS NULL;


