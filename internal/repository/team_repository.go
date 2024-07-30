package repository

import (
    "myapp/internal/app/entity"
    "gorm.io/gorm"
)

type TeamRepository interface {
    Create(team entity.Team) (entity.Team, error)
    GetByID(id uint) (entity.Team, error)
    GetAll() ([]entity.Team, error)
    Update(team entity.Team) (entity.Team, error)
    Delete(id uint) error
}

type teamRepository struct {
    db *gorm.DB
}

func NewTeamRepository(db *gorm.DB) TeamRepository {
    return &teamRepository{db: db}
}

func (r *teamRepository) Create(team entity.Team) (entity.Team, error) {
    result := r.db.Create(&team)
    return team, result.Error
}

func (r *teamRepository) GetByID(id uint) (entity.Team, error) {
    var team entity.Team
    result := r.db.Preload("Members").First(&team, id)
    return team, result.Error
}

func (r *teamRepository) GetAll() ([]entity.Team, error) {
    var teams []entity.Team
    result := r.db.Preload("Members").Find(&teams)
    return teams, result.Error
}

func (r *teamRepository) Update(team entity.Team) (entity.Team, error) {
    result := r.db.Save(&team)
    return team, result.Error
}

func (r *teamRepository) Delete(id uint) error {
    result := r.db.Delete(&entity.Team{}, id)
    return result.Error
}
