package json

import (
	"go-api/domains"
)

type getGetBlogsResponse struct {
	Blogs []getBlogResponseItem `json:"response"`
	Count int                   `json:"count"`
}

func (presenter ResponsePresenter) GetBlogs(blogs *domains.Blogs) {
	// nilではなく空配列でレスポンスを返せるようにする
	response := getGetBlogsResponse{
		Blogs: []getBlogResponseItem{},
		Count: blogs.Size(),
	}

	for _, blog := range blogs.Value {
		response.Blogs = append(response.Blogs, BlogItem(&blog))
	}

	presenter.Presenter.StatusOK(response)
}
