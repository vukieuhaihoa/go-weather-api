-- name: CreateLocation :one
INSERT INTO location (
  name, longitude, latitude
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetLocation :one
SELECT * FROM location
WHERE id = $1 LIMIT 1;


-- name: GetListLocation :many
SELECT * FROM location
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateLocation :exec
UPDATE location SET count = count + 1
WHERE id = $1;


