package domains

type Blogs struct {
	// Blogの集合を表現
	value []Blog
}

func NewBlogs(value []Blog) Blogs {
	return Blogs{value}
}

func (blogs Blogs) Value() []Blog {
	return blogs.value
}

func (blogs Blogs) Size() int {
	return len(blogs.value)
}

func EmptyBlogs() Blogs {
	return Blogs{[]Blog{}}
}
func (blogs Blogs) ApplyLimitAndOffset(limit, offset int) []Blog {
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

	return blogs.Value()[min:max]
}
