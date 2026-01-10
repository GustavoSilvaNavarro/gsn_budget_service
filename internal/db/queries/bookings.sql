-- name: CreateNewBooking :one
INSERT INTO bookings (amount, user_id, booking_platform, free_cancel_before, booking_start, booking_end, description, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), NOW())
RETURNING *;
