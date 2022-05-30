-- name: InsertUserAndReturnID :one
INSERT INTO users (name) VALUES ($1)
  RETURNING id;

-- name: InsertUserAndReturnUser :one
INSERT INTO users (name) VALUES ($1)
  RETURNING *;

-- name: UpdateUserAndReturnID :one
UPDATE users SET name = $1
  WHERE name = $2
  RETURNING id;

-- na