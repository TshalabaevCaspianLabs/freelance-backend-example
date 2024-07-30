package repository

import (
    "myapp/internal/app/entity"
    "gorm.io/gorm"
)

type TournamentRepository interface {
    Create(tournament entity.Tournament) (entity.Tournament, error)
    GetByID(id uint) (entity.Tournament, error)
    GetAll() ([]entity.Tournament, error)
    Update(tournament entity.Tournament) (entity.Tournament, error)
    Delete(id uint) error
}

type tournamentRepository struct {
    db *gorm.DB
}

func NewTournamentRepository(db *gorm.DB) TournamentRepository {
    return &tournamentRepository{db: db}
}

func (r *tournamentRepository) Create(tournament entity.Tournament) (entity.Tournament, error) {
    result := r.db.Create(&tournament)
    return tournament, result.Error
}

func (r *tournamentRepository) GetByID(id uint) (entity.Tournament, error) {
    var tournament entity.Tournament
    result := r.db.Preload("Participants").Preload("Matches").First(&tournament, id)
    return tournament, result.Error
}

func (r *tournamentRepository) GetAll() ([]entity.Tournament, error) {
    var tournaments []entity.Tournament
    result := r.db.Preload("Participants").Preload("Matches").Find(&tournaments)
    return tournaments, result.Error
}

func (r *tournamentRepository) Update(tournament entity.Tournament) (entity.Tournament, error) {
    result := r.db.Save(&tournament)
    return tournament, result.Error
}

func (r *tournamentRepository) Delete(id uint) error {
    result := r.db.Delete(&entity.Tournament{}, id)
    return result.Error
}
