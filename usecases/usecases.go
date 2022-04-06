package usecases

type IBlogGetAll interface {
	BlogGetAll(params GetBlogsParams)
}
type IBlogGet interface {
	BlogGet(params GetBlogParams)
}
type IBlogCreate interface {
	BlogCreate(params CreateBlogParams)
}
type IBlogEdit interface {
	BlogEdit(params EditBlogParams)
}
type IBlogDelete interface {
	BlogDelete(params DeleteBlogParams)
}

type IUserGetAll interface {
	UserGetAll(params GetUsersParams)
}
type IUserGet interface {
	UserGet(params GetUserParams)
}
type IUserCreate interface {
	UserCreate(params CreateUserParams)
}
type IUserEdit interface {
	UserEdit(params EditUserParams)
}
type IUserDelete interface {
	UserDelete(params DeleteUserParams)
}

type ISignUp interface {
	SignUp(params SignUpParams)
}
