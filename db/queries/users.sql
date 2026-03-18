-- name: CreateUser :one
INSERT INTO users (name, email)
VALUES ($1, $2)
RETURNING id, name, email;

-- name: SearchUsers :many
SELECT id, name, email
FROM users
WHERE name ILIKE '%' || $1 || '%'
   OR email ILIKE '%' || $1 || '%';