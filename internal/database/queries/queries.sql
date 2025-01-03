-- name: CreateUser :one
INSERT INTO users (username, email, hashed_password)
VALUES ($1, $2, $3)
RETURNING id, username, email, created_at, updated_at;

-- name: GetUserByID :one
SELECT id, username, email, hashed_password, created_at, updated_at
FROM users
WHERE id = $1;

-- name: CreateGroup :one
INSERT INTO groups (groupname)
VALUES ($1)
RETURNING id, groupname;

-- name: GetGroupByID :one
SELECT id, groupname
FROM groups
WHERE id = $1;

-- name: CreatePermission :one
INSERT INTO permissions (title)
VALUES ($1)
RETURNING id, title, created_at, updated_at;

-- name: GetPermissionByID :one
SELECT id, title
FROM permissions
WHERE id = $1;
