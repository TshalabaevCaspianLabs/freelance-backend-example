package repository

import (
    "myapp/internal/app/entity"
    "gorm.io/gorm"
)

type RatingRepository interface {
    Create(rating entity.Rating) (entity.Rating, error)
    GetByID(id uint) (entity.Rating, error)
    GetAll() ([]entity.Rating, error)
    Update(rating entity.Rating) (entity.Rating, error)
    Delete(id uint) error
}

type ratingRepository struct {
    db *gorm.DB
}

func NewRatingRepository(db *gorm.DB) RatingRepository {
    return &ratingRepository{db: db}
}

func (r *ratingRepository) Create(rating entity.Rating) (entity.Rating, error) {
    result := r.db.Create(&rating)
    return rating, result.Error
}

func (r *ratingRepository) GetByID(id uint) (entity.Rating, error) {
    var rating entity.Rating
    result := r.db.First(&rating, id)
    return rating, result.Error
}

func (r *ratingRepository) GetAll() ([]entity.Rating, error) {
    var ratings []entity.Rating
    result := r.db.Find(&ratings)
    return ratings, result.Error
}

func (r *ratingRepository) Update(rating entity.Rating) (entity.Rating, error) {
    result := r.db.Save(&rating)
    return rating, result.Error
}

func (r *ratingRepository) Delete(id uint) error {
    result := r.db.Delete(&entity.Rating{}, id)
    return result.Error
}
