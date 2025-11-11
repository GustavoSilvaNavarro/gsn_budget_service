-- name: CreateHousehold :one
INSERT INTO households (name, address, created_at, updated_at)
VALUES ($1, $2, NOW(), NOW())
RETURNING *;

-- name: GetHouseholdByID :one
SELECT * FROM households
WHERE id = $1;

-- name: GetHouseholdByName :one
SELECT * FROM households
WHERE name = $1;
