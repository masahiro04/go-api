package userDao

const GetAllSql = "SELECT id, name, email, created_at, updated_at FROM users WHERE deleted_at IS NULL"

const GetByIdSql = "SELECT id, name, email, created_at, updated_at FROM users WHERE id = $1 AND deleted_at IS NULL"

const CreateSql = "INSERT INTO users (name, email, created_at, updated_at) VALUES($1,$2,$3,$4) RETURNING id"

const UpdateSql = "UPDATE users SET name = $2, email = $3, updated_at = $4 WHERE id = $1"
const DeleteSql = "UPDATE users SET updated_at = $2, deleted_at = $3 WHERE id = $1"
