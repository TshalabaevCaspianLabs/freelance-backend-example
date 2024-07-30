package usecase

import (
    "myapp/internal/app/dto"
    "myapp/internal/app/entity"
    "myapp/internal/repository"
)

type NotificationUsecase interface {
    CreateNotification(dto dto.CreateNotificationDTO) (entity.Notification, error)
    GetNotificationByID(id uint) (entity.Notification, error)
    GetAllNotifications() ([]entity.Notification, error)
    UpdateNotification(id uint, dto dto.UpdateNotificationDTO) (entity.Notification, error)
    DeleteNotification(id uint) error
}

type notificationUsecase struct {
    repo repository.NotificationRepository
}

func NewNotificationUsecase(repo repository.NotificationRepository) NotificationUsecase {
    return &notificationUsecase{repo: repo}
}

func (u *notificationUsecase) CreateNotification(dto dto.CreateNotificationDTO) (entity.Notification, error) {
    notification := entity.Notification{UserID: dto.UserID, Message: dto.Message}
    return u.repo.Create(notification)
}

func (u *notificationUsecase) GetNotificationByID(id uint) (entity.Notification, error) {
    return u.repo.GetByID(id)
}

func (u *notificationUsecase) GetAllNotifications() ([]entity.Notification, error) {
    return u.repo.GetAll()
}

func (u *notificationUsecase) UpdateNotification(id uint, dto dto.UpdateNotificationDTO) (entity.Notification, error) {
    notification, err := u.repo.GetByID(id)
    if err != nil {
        return entity.Notification{}, err
    }

    notification.Message = dto.Message
    return u.repo.Update(notification)
}

func (u *notificationUsecase) DeleteNotification(id uint) error {
    return u.repo.Delete(id)
}
