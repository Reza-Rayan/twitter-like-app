package repository

import (
	"github.com/Reza-Rayan/twitter-like-app/db"
	"github.com/Reza-Rayan/twitter-like-app/internal/models"
	"gorm.io/gorm"
)

type PostRepository interface {
	Create(post *models.Post) error
	FindAll() ([]models.Post, error)
	FindByID(id int64) (models.Post, error)
	Update(post *models.Post) error
	Delete(id int64) error
	GetPostsWithLikes() ([]models.PostWithLikes, error)

	LikePost(userID, postID int64) error
	UnlikePost(userID, postID int64) error
	IsPostLiked(userID, postID int64) (bool, error)
}

type postRepository struct{}

func NewPostRepository() PostRepository {
	return &postRepository{}
}

// Create -> POST method
func (r *postRepository) Create(post *models.Post) error {
	return db.DB.Create(post).Error
}

// FindAll -> GET method
func (r *postRepository) FindAll() ([]models.Post, error) {
	var posts []models.Post
	err := db.DB.Preload("User").Preload("Likes").Find(&posts).Error
	return posts, err
}

// FindByID -> GET method
func (r *postRepository) FindByID(id int64) (models.Post, error) {
	var post models.Post
	err := db.DB.Preload("User").Preload("Likes").First(&post, id).Error
	return post, err
}

// Update -> PUT method
func (r *postRepository) Update(post *models.Post) error {
	return db.DB.Save(post).Error
}

func (r *postRepository) Delete(id int64) error {
	return db.DB.Delete(&models.Post{}, id).Error
}

func (r *postRepository) GetPostsWithLikes() ([]models.PostWithLikes, error) {
	var result []models.PostWithLikes
	err := db.DB.Table("posts").
		Select("posts.id, posts.title, posts.content, posts.created_at, posts.user_id, posts.image, COUNT(likes.id) as likes_count").
		Joins("LEFT JOIN likes ON likes.post_id = posts.id").
		Group("posts.id").
		Scan(&result).Error
	return result, err
}

// LikePost -> POST method
func (r *postRepository) LikePost(userID, postID int64) error {
	like := models.Like{
		UserID: userID,
		PostID: postID,
	}
	return db.DB.Create(&like).Error
}

// UnlikePost -> DELETE method
func (r *postRepository) UnlikePost(userID, postID int64) error {
	return db.DB.Where("user_id = ? AND post_id = ?", userID, postID).Delete(&models.Like{}).Error
}

func (r *postRepository) IsPostLiked(userID, postID int64) (bool, error) {
	var like models.Like
	err := db.DB.Where("user_id = ? AND post_id = ?", userID, postID).First(&like).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
