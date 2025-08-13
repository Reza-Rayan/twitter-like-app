package service

import "github.com/Reza-Rayan/twitter-like-app/internal/post"
import "github.com/Reza-Rayan/twitter-like-app/internal/post/repository"

type PostService interface {
	CreatePost(p *post.Post) error
	GetPosts(limit, offset int) ([]post.PostWithLikes, int, error)
	GetPost(id int64) (*post.Post, error)
	UpdatePost(p *post.Post) error
	DeletePost(id int64) error
}

type postService struct {
	repo repository.PostRepository
}

func NewPostService(repo repository.PostRepository) PostService {
	return &postService{repo: repo}
}

func (s *postService) CreatePost(p *post.Post) error {
	return s.repo.Save(p)
}

func (s *postService) GetPosts(limit, offset int) ([]post.PostWithLikes, int, error) {
	return s.repo.GetAll(limit, offset)
}

func (s *postService) GetPost(id int64) (*post.Post, error) {
	return s.repo.GetByID(id)
}

func (s *postService) UpdatePost(p *post.Post) error {
	return s.repo.Update(p)
}

func (s *postService) DeletePost(id int64) error {
	return s.repo.Delete(id)
}
