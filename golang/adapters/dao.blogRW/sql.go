package blogRW

const GetAllSql = "SELECT id, title, body, created_at, updated_at FROM blogs WHERE deleted_at IS NULL"
const GetByIdSql = "SELECT id, title, body, created_at, updated_at FROM blogs WHERE id = $1 AND deleted_at IS NULL"
const UpdateSql = "UPDATE blogs SET title = $2, body = $3, updated_at = $4 WHERE id = $1"
