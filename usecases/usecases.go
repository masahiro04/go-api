package usecases

type UseCases interface {
	BookUseCase
}

type BookUseCase interface {
	// Hoge() error
	BlogCreate(params CreateBlogParams)
}
