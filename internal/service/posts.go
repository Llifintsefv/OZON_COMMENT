package service

import "OZON_COMMENTS/internal/repository"

type PostsService struct {
	repo repository.PostRepository
}

func NewPostsService(repo repository.PostRepository) *PostsService {
	return &PostsService{
		repo: repo,
	}
}