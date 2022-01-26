package formatter

import (
	"clean_architecture/golang/domains"
)

type getGetBlogsResponse struct {
	Blogs []getBlogResponseItem `json:"blogs"`
	Count int                   `json:"count"`
}

func (presenter ResponsePresenter) GetBlogs(blogs domains.BlogCollection) {
	// nilではなく空配列でレスポンスを返せるようにする
	response := getGetBlogsResponse{
		Blogs: []getBlogResponseItem{},
		Count: len(blogs),
	}

	for _, blog := range blogs {
		response.Blogs = append(response.Blogs, BlogItem(blog))
	}

	presenter.Presenter.StatusOK(response)
}
