package blogRW

import (
	"clean_architecture/golang/domain"
	"database/sql"
	"log"
	"time"
)

type rw struct {
	store *sql.DB
}

func New(db *sql.DB) *rw {
	return &rw{
		store: db,
	}
}

func (rw rw) GetAll() ([]*domain.Blog, error) {
	var blogs []*domain.Blog
	rows, err := rw.store.Query(`SELECT id, title, body, created_at, updated_at FROM blogs WHERE deleted_at IS NULL`)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var blog domain.Blog
		if err = rows.Scan(
			&blog.ID,
			&blog.Title,
			&blog.Body,
			&blog.CreatedAt,
			&blog.UpdatedAt,
		); err != nil {
			log.Fatal(err)
			return nil, err
		}
		blogs = append(blogs, &blog)
	}
	return blogs, nil
}

func (rw rw) GetById(id int) (*domain.Blog, error) {
	var blog domain.Blog
	result := rw.store.QueryRow(`SELECT id, title, body, created_at, updated_at FROM blogs WHERE id = $1 AND deleted_at IS NULL`, id)
	err := result.Scan(
		&blog.ID,
		&blog.Title,
		&blog.Body,
		&blog.CreatedAt,
		&blog.UpdatedAt,
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &blog, nil
}

func (rw rw) Create(newBlog domain.Blog) (*domain.Blog, error) {
	var id int
	err := rw.store.QueryRow(
		`INSERT INTO blogs (title, body, created_at, updated_at) VALUES($1,$2,$3,$4) RETURNING id`,
		newBlog.Title, newBlog.Body, time.Now(), time.Now()).Scan(&id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	blog := domain.Blog{
		ID:        id,
		Title:     newBlog.Title,
		Body:      newBlog.Body,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return &blog, nil
}

func (rw rw) CreateTx(newBlog domain.Blog, tx *sql.Tx) (*domain.Blog, error) {
	var id int
	err := tx.QueryRow(
		`INSERT INTO blogs (title, body, created_at, updated_at) VALUES($1,$2,$3,$4) RETURNING id`,
		newBlog.Title, newBlog.Body, time.Now(), time.Now()).Scan(&id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	blog := domain.Blog{
		ID:        id,
		Title:     newBlog.Title,
		Body:      newBlog.Body,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return &blog, nil
}

func (rw rw) Update(id int, blog domain.Blog) (*domain.Blog, error) {
	_, err := rw.store.Exec(
		`UPDATE blogs SET title = $2, body = $3, updated_at = $4 WHERE id = $1`,
		id, blog.Title, blog.Body, time.Now())

	if err != nil {
		log.Println(err)
		return nil, err
	}

	updatedBlog := domain.Blog{
		ID:        id,
		Title:     blog.Title,
		Body:      blog.Body,
		UpdatedAt: time.Now(),
	}
	return &updatedBlog, nil
}

func (rw rw) Delete(id int) error {
	if _, err := rw.store.Exec(`
			UPDATE blogs SET updated_at = $2, deleted_at = $3 WHERE id = $1
			`, id, time.Now(), time.Now()); err != nil {
		return err
	}
	return nil
}
