package repository

import (
    "myapp/internal/app/entity"
    "gorm.io/gorm"
)

type BanListRepository interface {
    Create(banList entity.BanList) (entity.BanList, error)
    GetByID(id uint) (entity.BanList, error)
    GetAll() ([]entity.BanList, error)
    Update(banList entity.BanList) (entity.BanList, error)
    Delete(id uint) error
}

type banListRepository struct {
    db *gorm.DB
}

func NewBanListRepository(db *gorm.DB) BanListRepository {
    return &banListRepository{db: db}
}

func (r *banListRepository) Create(banList entity.BanList) (entity.BanList, error) {
    result := r.db.Create(&banList)
    return banList, result.Error
}

func (r *banListRepository) GetByID(id uint) (entity.BanList, error) {
    var banList entity.BanList
    result := r.db.First(&banList, id)
    return banList, result.Error
}

func (r *banListRepository) GetAll() ([]entity.BanList, error) {
    var banLists []entity.BanList
    result := r.db.Find(&banLists)
    return banLists, result.Error
}

func (r *banListRepository) Update(banList entity.BanList) (entity.BanList, error) {
    result := r.db.Save(&banList)
    return banList, result.Error
}

func (r *banListRepository) Delete(id uint) error {
    result := r.db.Delete(&entity.BanList{}, id)
    return result.Error
}
