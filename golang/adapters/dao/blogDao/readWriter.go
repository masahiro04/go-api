package blogDao

import (
	"clean_architecture/golang/domains"
	blogModel "clean_architecture/golang/domains/blog"

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

func (rw rw) GetAll() (*domains.Blogs, error) {
	var dtos []BlogDto
	rw.db.Find(&dtos)
	var blogs []domains.Blog
	for _, blogDto := range dtos {
		id, _ := blogModel.NewId(blogDto.ID)
		title, _ := blogModel.NewTitle(blogDto.Title)
		body, _ := blogModel.NewBody(blogDto.Body)
		newBlog := domains.BuildBlog(id, title, body, blogDto.CreatedAt, blogDto.UpdatedAt)

		blogs = append(blogs, newBlog)
	}

	blogsData := domains.NewBlogs(blogs)
	return &blogsData, nil
}

func (rw rw) GetById(id int) (*domains.Blog, error) {
	var dto BlogDto

	if err := rw.db.Where("id = ?", id).First(&dto).Error; gorm.ErrRecordNotFound != nil {
		return nil, err
	}

	_id, _ := blogModel.NewId(dto.ID)
	title, _ := blogModel.NewTitle(dto.Title)
	body, _ := blogModel.NewBody(dto.Body)
	newBlog := domains.BuildBlog(_id, title, body, dto.CreatedAt, dto.UpdatedAt)
	return &newBlog, nil
}

//
func (rw rw) Create(newBlog domains.Blog) (*domains.Blog, error) {
	dto := BlogDto{
		Title: newBlog.Title.Value,
		Body:  newBlog.Body.Value,
	}
	rw.db.Create(&dto)

	_id, _ := blogModel.NewId(dto.ID)
	title, _ := blogModel.NewTitle(newBlog.Title.Value)
	body, _ := blogModel.NewBody(newBlog.Body.Value)
	blog := domains.BuildBlog(_id, title, body, newBlog.CreatedAt, newBlog.UpdatedAt)
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
func (rw rw) Update(id int, blog domains.Blog) (*domains.Blog, error) {
	dto := BlogDto{}

	rw.db.Where("id = ?", id).First(&dto).Updates(BlogDto{
		ID:    id,
		Title: blog.Title.Value,
		Body:  blog.Body.Value,
	})

	_id, _ := blogModel.NewId(id)
	title, _ := blogModel.NewTitle(blog.Title.Value)
	body, _ := blogModel.NewBody(blog.Body.Value)
	newBlog := domains.BuildBlog(_id, title, body, blog.CreatedAt, blog.UpdatedAt)
	return &newBlog, nil
}

func (rw rw) Delete(id int) error {
	dto := BlogDto{}
	rw.db.Where("id = ?", id).Delete(&dto)
	return nil
}
