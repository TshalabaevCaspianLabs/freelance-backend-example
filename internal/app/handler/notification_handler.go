package handler

import (
    "myapp/internal/app/dto"
    "myapp/internal/app/usecase"
    "myapp/internal/repository"
    "github.com/gin-gonic/gin"
    "net/http"
    "gorm.io/gorm"
    "strconv"
)

type NotificationHandler struct {
    usecase usecase.NotificationUsecase
}

func RegisterNotificationRoutes(rg *gin.RouterGroup, db *gorm.DB) {
    handler := &NotificationHandler{usecase: usecase.NewNotificationUsecase(repository.NewNotificationRepository(db))}
    rg.POST("/notifications", handler.CreateNotification)
    rg.GET("/notifications/:id", handler.GetNotificationByID)
    rg.GET("/notifications", handler.GetAllNotifications)
    rg.PUT("/notifications/:id", handler.UpdateNotification)
    rg.DELETE("/notifications/:id", handler.DeleteNotification)
}

// CreateNotification godoc
// @Summary Create a new notification
// @Description Create a new notification
// @Tags notifications
// @Accept json
// @Produce json
// @Param notification body dto.CreateNotificationDTO true "Create Notification"
// @Success 200 {object} entity.Notification
// @Router /notifications [post]
func (h *NotificationHandler) CreateNotification(c *gin.Context) {
    var createNotificationDTO dto.CreateNotificationDTO
    if err := c.ShouldBindJSON(&createNotificationDTO); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    notification, err := h.usecase.CreateNotification(createNotificationDTO)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, notification)
}

// GetNotificationByID godoc
// @Summary Get notification by ID
// @Description Get notification by ID
// @Tags notifications
// @Accept json
// @Produce json
// @Param id path int true "Notification ID"
// @Success 200 {object} entity.Notification
// @Router /notifications/{id} [get]
func (h *NotificationHandler) GetNotificationByID(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    notification, err := h.usecase.GetNotificationByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, notification)
}

// GetAllNotifications godoc
// @Summary Get all notifications
// @Description Get all notifications
// @Tags notifications
// @Accept json
// @Produce json
// @Success 200 {array} entity.Notification
// @Router /notifications [get]
func (h *NotificationHandler) GetAllNotifications(c *gin.Context) {
    notifications, err := h.usecase.GetAllNotifications()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, notifications)
}

// UpdateNotification godoc
// @Summary Update notification
// @Description Update notification
// @Tags notifications
// @Accept json
// @Produce json
// @Param id path int true "Notification ID"
// @Param notification body dto.UpdateNotificationDTO true "Update Notification"
// @Success 200 {object} entity.Notification
// @Router /notifications/{id} [put]
func (h *NotificationHandler) UpdateNotification(c *gin.Context) {
    var updateNotificationDTO dto.UpdateNotificationDTO
    if err := c.ShouldBindJSON(&updateNotificationDTO); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    idParam := c.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    notification, err := h.usecase.UpdateNotification(uint(id), updateNotificationDTO)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, notification)
}

// DeleteNotification godoc
// @Summary Delete notification
// @Description Delete notification
// @Tags notifications
// @Accept json
// @Produce json
// @Param id path int true "Notification ID"
// @Success 200 {object} map[string]interface{} "message: Notification deleted"
// @Router /notifications/{id} [delete]
func (h *NotificationHandler) DeleteNotification(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    if err := h.usecase.DeleteNotification(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Notification deleted"})
}
