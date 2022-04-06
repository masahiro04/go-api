package models

type Blogs struct {
	// Blogの集合を表現
	Value []Blog
}

func NewBlogs(value []Blog) Blogs {
	return Blogs{Value: value}
}

func (blogs Blogs) Size() int {
	return len(blogs.Value)
}

func EmptyBlogs() Blogs {
	return Blogs{Value: []Blog{}}
}
func (blogs *Blogs) ApplyLimitAndOffset(limit, offset int) []Blog {
	if limit <= 0 {
		return []Blog{}
	}

	blogsSize := blogs.Size()
	min := offset
	if min < 0 {
		min = 0
	}

	if min > blogsSize {
		return []Blog{}
	}

	max := min + limit
	if max > blogsSize {
		max = blogsSize
	}

	return blogs.Value[min:max]
}
