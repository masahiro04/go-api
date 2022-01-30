package json

import (
	"clean_architecture/golang/domains"
)

type getBlogResponse struct {
	Blog getBlogResponseItem `json:"response"`
}

type getBlogResponseItem struct {
	Id        interface{} `json:"id"`
	Title     interface{} `json:"title"`
	Body      interface{} `json:"body"`
	CreatedAt interface{} `json:"createdAt"`
	UpdatedAt interface{} `json:"updatedAt"`
}

func (presenter ResponsePresenter) GetBlog(blog *domains.Blog) {
	response := getBlogResponse{Blog: BlogItem(blog)}
	presenter.Presenter.StatusOK(response)
}

func (presenter ResponsePresenter) CreateBlog(blog *domains.Blog) {
	response := getBlogResponse{Blog: BlogItem(blog)}
	presenter.Presenter.Created(response)
}

func BlogItem(blog *domains.Blog) getBlogResponseItem {
	return getBlogResponseItem{
		Id:        blog.ID.Value,
		Title:     blog.Title.Value,
		Body:      blog.Body.Value,
		CreatedAt: blog.CreatedAt.UTC().Format(dateLayout),
		UpdatedAt: blog.UpdatedAt.UTC().Format(dateLayout),
	}
}
