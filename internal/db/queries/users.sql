-- name: CreateNewUser :one
INSERT INTO users (email, username, lastname, gender, role, household_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6,  NOW(), NOW())
RETURNING *;
