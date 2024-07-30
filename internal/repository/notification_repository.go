package repository

import (
    "myapp/internal/app/entity"
    "gorm.io/gorm"
)

type NotificationRepository interface {
    Create(notification entity.Notification) (entity.Notification, error)
    GetByID(id uint) (entity.Notification, error)
    GetAll() ([]entity.Notification, error)
    Update(notification entity.Notification) (entity.Notification, error)
    Delete(id uint) error
}

type notificationRepository struct {
    db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) NotificationRepository {
    return &notificationRepository{db: db}
}

func (r *notificationRepository) Create(notification entity.Notification) (entity.Notification, error) {
    result := r.db.Create(&notification)
    return notification, result.Error
}

func (r *notificationRepository) GetByID(id uint) (entity.Notification, error) {
    var notification entity.Notification
    result := r.db.First(&notification, id)
    return notification, result.Error
}

func (r *notificationRepository) GetAll() ([]entity.Notification, error) {
    var notifications []entity.Notification
    result := r.db.Find(&notifications)
    return notifications, result.Error
}

func (r *notificationRepository) Update(notification entity.Notification) (entity.Notification, error) {
    result := r.db.Save(&notification)
    return notification, result.Error
}

func (r *notificationRepository) Delete(id uint) error {
    result := r.db.Delete(&entity.Notification{}, id)
    return result.Error
}
