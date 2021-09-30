package formatter

import (
	"clean_architecture/golang/domain"
)

type getBlogResponse struct {
	Blog getBlogResponseItem `json:"blog"`
}

type getBlogResponseItem struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func (presenter ResponsePresenter) GetBlog(blog *domain.Blog) {
	response := getBlogResponse{Blog: BlogItem(blog)}
	presenter.Presenter.StatusOK(response)
}

func (presenter ResponsePresenter) CreateBlog(blog *domain.Blog) {
	response := getBlogResponse{Blog: BlogItem(blog)}
	presenter.Presenter.Created(response)
}

func BlogItem(blog *domain.Blog) getBlogResponseItem {
	return getBlogResponseItem{
		Id:        blog.ID,
		Title:     blog.Title,
		Body:      blog.Body,
		CreatedAt: blog.CreatedAt.UTC().Format(dateLayout),
		UpdatedAt: blog.UpdatedAt.UTC().Format(dateLayout),
	}
}
