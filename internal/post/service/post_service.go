package service

import (
	"github.com/Reza-Rayan/twitter-like-app/internal/models"
	"github.com/Reza-Rayan/twitter-like-app/internal/post/repository"
)

type PostService interface {
	Create(post *models.Post) error
	GetAll() ([]models.Post, error)
	GetByID(id int64) (models.Post, error)
	Update(post *models.Post) error
	Delete(id int64) error
	GetPostsWithLikes() ([]models.PostWithLikes, error)

	LikePost(userID, postID int64) error
	UnlikePost(userID, postID int64) error
	IsPostLiked(userID, postID int64) (bool, error)
}

type postService struct {
	repo repository.PostRepository
}

func NewPostService(repo repository.PostRepository) PostService {
	return &postService{repo: repo}
}

func (s *postService) Create(post *models.Post) error {
	return s.repo.Create(post)
}

func (s *postService) GetAll() ([]models.Post, error) {
	return s.repo.FindAll()
}

func (s *postService) GetByID(id int64) (models.Post, error) {
	return s.repo.FindByID(id)
}

func (s *postService) Update(post *models.Post) error {
	return s.repo.Update(post)
}

func (s *postService) Delete(id int64) error {
	return s.repo.Delete(id)
}

func (s *postService) GetPostsWithLikes() ([]models.PostWithLikes, error) {
	return s.repo.GetPostsWithLikes()
}

func (s *postService) LikePost(userID, postID int64) error {
	return s.repo.LikePost(userID, postID)
}

func (s *postService) UnlikePost(userID, postID int64) error {
	return s.repo.UnlikePost(userID, postID)
}

func (s *postService) IsPostLiked(userID, postID int64) (bool, error) {
	return s.repo.IsPostLiked(userID, postID)
}
