package api

import "OZON_COMMENTS/internal/service"

type Resolver struct {
	postService    *service.PostsService
	commentService *service.CommentService
}

func NewResolver(postService *service.PostsService, commentService *service.CommentService) *Resolver {
	return &Resolver{
		postService:    postService,
		commentService: commentService,
	}
}