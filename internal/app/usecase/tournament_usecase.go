package usecase

import (
    "myapp/internal/app/dto"
    "myapp/internal/app/entity"
    "myapp/internal/repository"
)

type TournamentUsecase interface {
    CreateTournament (dto dto.CreateTournamentDTO) (entity.Tournament, error)
    GetTournamentByID(id uint) (entity.Tournament, error)
    GetAllTournaments() ([]entity.Tournament, error)
    UpdateTournament (id uint, dto dto.UpdateTournamentDTO) (entity.Tournament, error)
    DeleteTournament (id uint) error
}

type tournamentUsecase struct {
    repo repository.TournamentRepository
}

func NewTournamentUsecase(repo repository.TournamentRepository) TournamentUsecase {
    return &tournamentUsecase{repo: repo}
}

func (u *tournamentUsecase) CreateTournament (dto dto.CreateTournamentDTO) (entity.Tournament, error) {
    tournament := entity.Tournament{Name: dto.Name, StartDate: dto.StartDate, EndDate: dto.EndDate}
    return u.repo.Create(tournament)
}

func (u *tournamentUsecase) GetTournamentByID(id uint) (entity.Tournament, error) {
    return u.repo.GetByID(id)
}

func (u *tournamentUsecase) GetAllTournaments() ([]entity.Tournament, error) {
    return u.repo.GetAll()
}

func (u *tournamentUsecase) UpdateTournament (id uint, dto dto.UpdateTournamentDTO) (entity.Tournament, error) {
    tournament, err := u.repo.GetByID(id)
    if err != nil {
        return entity.Tournament{}, err
    }

    tournament.Name = dto.Name
    tournament.StartDate = dto.StartDate
    tournament.EndDate = dto.EndDate
    return u.repo.Update(tournament)
}

func (u *tournamentUsecase) DeleteTournament (id uint) error {
    return u.repo.Delete(id)
}
