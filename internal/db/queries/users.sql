-- name: CreateNewUser :one
INSERT INTO users (email, username, lastname, gender, role, household_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6,  NOW(), NOW())
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1;

-- name: GetUsersByHouseholdID :many
SELECT * FROM users
WHERE household_id = $1
ORDER BY created_at DESC;
