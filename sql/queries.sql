-- name: CreateArtist :one
INSERT INTO artists (id, name, genre)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetArtist :one
SELECT id, name, genre
FROM artists
WHERE id = $1
LIMIT 1;

-- name: UpdateArtist :one
UPDATE artists
SET name = $2,
    genre  = $3
WHERE id = $1
RETURNING *;

-- name: DeleteArtist :exec
DELETE
FROM artists
WHERE id = $1;

-- name: ListArtists :many
SELECT *
FROM artists
ORDER BY name;
