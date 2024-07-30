package usecase

import (
    "myapp/internal/app/dto"
    "myapp/internal/app/entity"
    "myapp/internal/repository"
)

type UserUsecase interface {
    CreateUser(dto dto.CreateUserDTO) (entity.User, error)
    GetUserByID(id uint) (entity.User, error)
    GetAllUsers() ([]entity.User, error)
    UpdateUser(id uint, dto dto.UpdateUserDTO) (entity.User, error)
    DeleteUser(id uint) error
}

type userUsecase struct {
    repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
    return &userUsecase{repo: repo}
}

func (u *userUsecase) CreateUser(dto dto.CreateUserDTO) (entity.User, error) {
    user := entity.User{Name: dto.Name, Email: dto.Email}
    return u.repo.Create(user)
}

func (u *userUsecase) GetUserByID(id uint) (entity.User, error) {
    return u.repo.GetByID(id)
}

func (u *userUsecase) GetAllUsers() ([]entity.User, error) {
    return u.repo.GetAll()
}

func (u *userUsecase) UpdateUser(id uint, dto dto.UpdateUserDTO) (entity.User, error) {
    user, err := u.repo.GetByID(id)
    if err != nil {
        return entity.User{}, err
    }

    user.Name = dto.Name
    user.Email = dto.Email
    return u.repo.Update(user)
}

func (u *userUsecase) DeleteUser(id uint) error {
    return u.repo.Delete(id)
}
