package usecase

import (
    "myapp/internal/app/dto"
    "myapp/internal/app/entity"
    "myapp/internal/repository"
)

type RatingUsecase interface {
    CreateRating(dto dto.CreateRatingDTO) (entity.Rating, error)
    GetRatingByID(id uint) (entity.Rating, error)
    GetAllRatings() ([]entity.Rating, error)
    UpdateRating(id uint, dto dto.UpdateRatingDTO) (entity.Rating, error)
    DeleteRating(id uint) error
}

type ratingUsecase struct {
    repo repository.RatingRepository
}

func NewRatingUsecase(repo repository.RatingRepository) RatingUsecase {
    return &ratingUsecase{repo: repo}
}

func (u *ratingUsecase) CreateRating(dto dto.CreateRatingDTO) (entity.Rating, error) {
    rating := entity.Rating{UserID: dto.UserID, Value: dto.Value}
    return u.repo.Create(rating)
}

func (u *ratingUsecase) GetRatingByID(id uint) (entity.Rating, error) {
    return u.repo.GetByID(id)
}

func (u *ratingUsecase) GetAllRatings() ([]entity.Rating, error) {
    return u.repo.GetAll()
}

func (u *ratingUsecase) UpdateRating(id uint, dto dto.UpdateRatingDTO) (entity.Rating, error) {
    rating, err := u.repo.GetByID(id)
    if err != nil {
        return entity.Rating{}, err
    }

    rating.Value = dto.Value
    return u.repo.Update(rating)
}

func (u *ratingUsecase) DeleteRating(id uint) error {
    return u.repo.Delete(id)
}
