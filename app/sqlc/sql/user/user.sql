-- name: UserGetMailByID :one
SELECT mail FROM "user" WHERE id = $1;
