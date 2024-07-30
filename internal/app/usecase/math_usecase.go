package usecase

import (
    "myapp/internal/app/dto"
    "myapp/internal/app/entity"
    "myapp/internal/repository"
)

type MatchUsecase interface {
    CreateMatch(dto dto.CreateMatchDTO) (entity.Match, error)
    GetMatchByID(id uint) (entity.Match, error)
    GetAllMatches() ([]entity.Match, error)
    UpdateMatch(id uint, dto dto.UpdateMatchDTO) (entity.Match, error)
    DeleteMatch(id uint) error
}

type matchUsecase struct {
    repo repository.MatchRepository
}

func NewMatchUsecase(repo repository.MatchRepository) MatchUsecase {
    return &matchUsecase{repo: repo}
}

func (u *matchUsecase) CreateMatch(dto dto.CreateMatchDTO) (entity.Match, error) {
    match := entity.Match{
        TournamentID: dto.TournamentID,
        TeamAID:      dto.TeamAID,
        TeamBID:      dto.TeamBID,
        Schedule:     dto.Schedule,
    }
    return u.repo.Create(match)
}

func (u *matchUsecase) GetMatchByID(id uint) (entity.Match, error) {
    return u.repo.GetByID(id)
}

func (u *matchUsecase) GetAllMatches() ([]entity.Match, error) {
    return u.repo.GetAll()
}

func (u *matchUsecase) UpdateMatch(id uint, dto dto.UpdateMatchDTO) (entity.Match, error) {
    match, err := u.repo.GetByID(id)
    if err != nil {
        return entity.Match{}, err
    }

    match.TeamAID = dto.TeamAID
    match.TeamBID = dto.TeamBID
    match.Schedule = dto.Schedule
    match.ScoreA = dto.ScoreA
    match.ScoreB = dto.ScoreB
    return u.repo.Update(match)
}

func (u *matchUsecase) DeleteMatch(id uint) error {
    return u.repo.Delete(id)
}
