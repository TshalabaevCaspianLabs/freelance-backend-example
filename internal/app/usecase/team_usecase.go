package usecase

import (
    "myapp/internal/app/dto"
    "myapp/internal/app/entity"
    "myapp/internal/repository"
)

type TeamUsecase interface {
    CreateTeam(dto dto.CreateTeamDTO) (entity.Team, error)
    GetTeamByID(id uint) (entity.Team, error)
    GetAllTeams() ([]entity.Team, error)
    UpdateTeam(id uint, dto dto.UpdateTeamDTO) (entity.Team, error)
    DeleteTeam(id uint) error
}

type teamUsecase struct {
    teamRepo repository.TeamRepository
    userRepo repository.UserRepository
}

func NewTeamUsecase(teamRepo repository.TeamRepository, userRepo repository.UserRepository) TeamUsecase {
    return &teamUsecase{teamRepo: teamRepo, userRepo: userRepo}
}

func (u *teamUsecase) CreateTeam(dto dto.CreateTeamDTO) (entity.Team, error) {
    team := entity.Team{Name: dto.Name, Members: []entity.User{}}
    for _, memberID := range dto.Members {
        user, err := u.userRepo.GetByID(memberID)
        if err != nil {
            return entity.Team{}, err
        }
        team.Members = append(team.Members, user)
    }
    return u.teamRepo.Create(team)
}

func (u *teamUsecase) GetTeamByID(id uint) (entity.Team, error) {
    return u.teamRepo.GetByID(id)
}

func (u *teamUsecase) GetAllTeams() ([]entity.Team, error) {
    return u.teamRepo.GetAll()
}

func (u *teamUsecase) UpdateTeam(id uint, dto dto.UpdateTeamDTO) (entity.Team, error) {
    team, err := u.teamRepo.GetByID(id)
    if err != nil {
        return entity.Team{}, err
    }

    team.Name = dto.Name
    team.Members = []entity.User{}
    for _, memberID := range dto.Members {
        user, err := u.userRepo.GetByID(memberID)
        if err != nil {
            return entity.Team{}, err
        }
        team.Members = append(team.Members, user)
    }
    return u.teamRepo.Update(team)
}

func (u *teamUsecase) DeleteTeam(id uint) error {
    return u.teamRepo.Delete(id)
}
