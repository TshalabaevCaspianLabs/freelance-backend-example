package usecase

import (
    "myapp/internal/app/dto"
    "myapp/internal/app/entity"
    "myapp/internal/repository"
)

type BanListUsecase interface {
    CreateBanList(dto dto.CreateBanListDTO) (entity.BanList, error)
    GetBanListByID(id uint) (entity.BanList, error)
    GetAllBanLists() ([]entity.BanList, error)
    UpdateBanList(id uint, dto dto.UpdateBanListDTO) (entity.BanList, error)
    DeleteBanList(id uint) error
}

type banListUsecase struct {
    repo repository.BanListRepository
}

func NewBanListUsecase(repo repository.BanListRepository) BanListUsecase {
    return &banListUsecase{repo: repo}
}

func (u *banListUsecase) CreateBanList(dto dto.CreateBanListDTO) (entity.BanList, error) {
    banList := entity.BanList{UserID: dto.UserID, EndDate: dto.EndDate}
    return u.repo.Create(banList)
}

func (u *banListUsecase) GetBanListByID(id uint) (entity.BanList, error) {
    return u.repo.GetByID(id)
}

func (u *banListUsecase) GetAllBanLists() ([]entity.BanList, error) {
    return u.repo.GetAll()
}

func (u *banListUsecase) UpdateBanList(id uint, dto dto.UpdateBanListDTO) (entity.BanList, error) {
    banList, err := u.repo.GetByID(id)
    if err != nil {
        return entity.BanList{}, err
    }

    banList.EndDate = dto.EndDate
    return u.repo.Update(banList)
}

func (u *banListUsecase) DeleteBanList(id uint) error {
    return u.repo.Delete(id)
}
