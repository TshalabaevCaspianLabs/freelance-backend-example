package repository

import (
    "myapp/internal/app/entity"
    "gorm.io/gorm"
)

type UserRepository interface {
    Create(user entity.User) (entity.User, error)
    GetByID(id uint) (entity.User, error)
    GetAll() ([]entity.User, error)
    Update(user entity.User) (entity.User, error)
    Delete(id uint) error
}

type userRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepository{db: db}
}

func (r *userRepository) Create(user entity.User) (entity.User, error) {
    result := r.db.Create(&user)
    return user, result.Error
}

func (r *userRepository) GetByID(id uint) (entity.User, error) {
    var user entity.User
    result := r.db.First(&user, id)
    return user, result.Error
}

func (r *userRepository) GetAll() ([]entity.User, error) {
    var users []entity.User
    result := r.db.Find(&users)
    return users, result.Error
}

func (r *userRepository) Update(user entity.User) (entity.User, error) {
    result := r.db.Save(&user)
    return user, result.Error
}

func (r *userRepository) Delete(id uint) error {
    result := r.db.Delete(&entity.User{}, id)
    return result.Error
}
