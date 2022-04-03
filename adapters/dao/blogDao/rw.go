package blogDao

import (
	"errors"
	"go-api/domains/models"
	"go-api/domains/models/blog"

	"gorm.io/gorm"
)

type rw struct {
	db *gorm.DB
}

func New(db *gorm.DB) *rw {
	return &rw{
		db: db,
	}
}

type BlogDto struct {
	gorm.Model
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

// NOTE(okubo): table名を指定
func (BlogDto) TableName() string {
	return "blogs"
}

func (rw rw) GetAll() (*models.Blogs, error) {
	var dtos []BlogDto
	rw.db.Find(&dtos)
	var blogs []models.Blog
	for _, blogDto := range dtos {
		id, _ := blog.NewId(blogDto.ID)
		title, _ := blog.NewTitle(blogDto.Title)
		body, _ := blog.NewBody(blogDto.Body)
		newBlog := models.BuildBlog(id, title, body, blogDto.CreatedAt, blogDto.UpdatedAt)

		blogs = append(blogs, newBlog)
	}

	blogsData := models.NewBlogs(blogs)
	return &blogsData, nil
}

func (rw rw) GetById(id int) (*models.Blog, error) {
	var dto BlogDto

	err := rw.db.First(&dto, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	_id, _ := blog.NewId(dto.ID)
	title, _ := blog.NewTitle(dto.Title)
	body, _ := blog.NewBody(dto.Body)
	newBlog := models.BuildBlog(_id, title, body, dto.CreatedAt, dto.UpdatedAt)
	return &newBlog, nil
}

//
func (rw rw) Create(newBlog models.Blog) (*models.Blog, error) {
	dto := BlogDto{
		Title: newBlog.Title.Value,
		Body:  newBlog.Body.Value,
	}
	rw.db.Create(&dto)

	_id, _ := blog.NewId(dto.ID)
	title, _ := blog.NewTitle(newBlog.Title.Value)
	body, _ := blog.NewBody(newBlog.Body.Value)
	blog := models.BuildBlog(_id, title, body, newBlog.CreatedAt, newBlog.UpdatedAt)
	return &blog, nil
}

// func (rw rw) CreateTx(newBlog domains.Blog, tx *sql.Tx) (*domains.Blog, error) {
// 	var id int
// 	err := tx.QueryRow(
// 		CreateSql,
// 		newBlog.Title, newBlog.Body, time.Now(), time.Now()).Scan(&id)
// 	if err != nil {
// 		log.Println(err)
// 		return nil, err
// 	}
//
// 	_id, _ := blogModel.NewId(newBlog.ID.Value)
// 	title, _ := blogModel.NewTitle(newBlog.Title.Value)
// 	body, _ := blogModel.NewBody(newBlog.Body.Value)
// 	blog := domains.BuildBlog(_id, title, body, newBlog.CreatedAt, newBlog.UpdatedAt)
// 	return &blog, nil
// }

//
func (rw rw) Update(id int, newBlog models.Blog) (*models.Blog, error) {
	dto := BlogDto{}

	rw.db.First(&dto, id).Updates(BlogDto{
		ID:    id,
		Title: newBlog.Title.Value,
		Body:  newBlog.Body.Value,
	})

	_id, _ := blog.NewId(id)
	title, _ := blog.NewTitle(newBlog.Title.Value)
	body, _ := blog.NewBody(newBlog.Body.Value)
	updatedBlog := models.BuildBlog(_id, title, body, newBlog.CreatedAt, newBlog.UpdatedAt)
	return &updatedBlog, nil
}

func (rw rw) Delete(id int) error {
	dto := BlogDto{}
	rw.db.First(&dto, id).Delete(&dto)
	return nil
}
