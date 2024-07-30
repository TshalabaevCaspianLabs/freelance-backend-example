package repository

import (
    "myapp/internal/app/entity"
    "gorm.io/gorm"
)

type MatchRepository interface {
    Create(match entity.Match) (entity.Match, error)
    GetByID(id uint) (entity.Match, error)
    GetAll() ([]entity.Match, error)
    Update(match entity.Match) (entity.Match, error)
    Delete(id uint) error
}

type matchRepository struct {
    db *gorm.DB
}

func NewMatchRepository(db *gorm.DB) MatchRepository {
    return &matchRepository{db: db}
}

func (r *matchRepository) Create(match entity.Match) (entity.Match, error) {
    result := r.db.Create(&match)
    return match, result.Error
}

func (r *matchRepository) GetByID(id uint) (entity.Match, error) {
    var match entity.Match
    result := r.db.Preload("TeamA.Members").Preload("TeamB.Members").First(&match, id)
    return match, result.Error
}

func (r *matchRepository) GetAll() ([]entity.Match, error) {
    var matches []entity.Match
    result := r.db.Preload("TeamA.Members").Preload("TeamB.Members").Find(&matches)
    return matches, result.Error
}

func (r *matchRepository) Update(match entity.Match) (entity.Match, error) {
    result := r.db.Save(&match)
    return match, result.Error
}

func (r *matchRepository) Delete(id uint) error {
    result := r.db.Delete(&entity.Match{}, id)
    return result.Error
}
