package blogRW

import (
	"clean_architecture/golang/domains"
	blogModel "clean_architecture/golang/domains/blog"
	"time"

	"database/sql"
	"log"
)

type rw struct {
	store *sql.DB
}

func New(db *sql.DB) *rw {
	return &rw{
		store: db,
	}
}

type BlogDto struct {
	ID        int
	Title     string
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (rw rw) GetAll() (*domains.Blogs, error) {
	var blogs []domains.Blog
	rows, err := rw.store.Query(GetAllSql)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var blogDto BlogDto

		if err = rows.Scan(
			&blogDto.ID,
			&blogDto.Title,
			&blogDto.Body,
			&blogDto.CreatedAt,
			&blogDto.UpdatedAt,
		); err != nil {
			log.Fatal(err)
			return nil, err
		}

		id, _ := blogModel.NewId(blogDto.ID)
		title, _ := blogModel.NewTitle(blogDto.Title)
		body, _ := blogModel.NewBody(blogDto.Body)
		newBlog := domains.BuildBlog(id, title, body)

		blogs = append(blogs, newBlog)
	}
	blogsData := domains.NewBlogs(blogs)
	return &blogsData, nil
}

func (rw rw) GetById(id int) (*domains.Blog, error) {
	var blogDto BlogDto

	result := rw.store.QueryRow(GetByIdSql, id)
	err := result.Scan(
		&blogDto.ID,
		&blogDto.Title,
		&blogDto.Body,
		&blogDto.CreatedAt,
		&blogDto.UpdatedAt,
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	_id, _ := blogModel.NewId(blogDto.ID)
	title, _ := blogModel.NewTitle(blogDto.Title)
	body, _ := blogModel.NewBody(blogDto.Body)
	newBlog := domains.BuildBlog(_id, title, body)
	return &newBlog, nil
}

//
func (rw rw) Create(newBlog domains.Blog) (*domains.Blog, error) {
	var id int
	err := rw.store.QueryRow(
		`INSERT INTO blogs (title, body, created_at, updated_at) VALUES($1,$2,$3,$4) RETURNING id`,
		newBlog.Title, newBlog.Body, time.Now(), time.Now()).Scan(&id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	_id, _ := blogModel.NewId(newBlog.ID().Value())
	title, _ := blogModel.NewTitle(newBlog.Title().Value())
	body, _ := blogModel.NewBody(newBlog.Body().Value())
	blog := domains.BuildBlog(_id, title, body)
	return &blog, nil
}

func (rw rw) CreateTx(newBlog domains.Blog, tx *sql.Tx) (*domains.Blog, error) {
	var id int
	err := tx.QueryRow(
		`INSERT INTO blogs (title, body, created_at, updated_at) VALUES($1,$2,$3,$4) RETURNING id`,
		newBlog.Title, newBlog.Body, time.Now(), time.Now()).Scan(&id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	_id, _ := blogModel.NewId(newBlog.ID().Value())
	title, _ := blogModel.NewTitle(newBlog.Title().Value())
	body, _ := blogModel.NewBody(newBlog.Body().Value())
	blog := domains.BuildBlog(_id, title, body)
	return &blog, nil
}

//
func (rw rw) Update(id int, blog domains.Blog) (*domains.Blog, error) {
	_, err := rw.store.Exec(
		UpdateSql,
		id, blog.Title().Value(), blog.Body().Value(), time.Now())

	if err != nil {
		log.Println(err)
		return nil, err
	}

	_id, _ := blogModel.NewId(id)
	title, _ := blogModel.NewTitle(blog.Title().Value())
	body, _ := blogModel.NewBody(blog.Body().Value())
	newBlog := domains.BuildBlog(_id, title, body)
	return &newBlog, nil
}

func (rw rw) Delete(id int) error {
	if _, err := rw.store.Exec(`
			UPDATE blogs SET updated_at = $2, deleted_at = $3 WHERE id = $1
			`, id, time.Now(), time.Now()); err != nil {
		return err
	}
	return nil
}
