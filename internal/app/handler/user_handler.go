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

type UserHandler struct {
    usecase usecase.UserUsecase
}

func RegisterUserRoutes(rg *gin.RouterGroup, db *gorm.DB) {
    handler := &UserHandler{usecase: usecase.NewUserUsecase(repository.NewUserRepository(db))}
    rg.POST("/users", handler.CreateUser)
    rg.GET("/users/:id", handler.GetUserByID)
    rg.GET("/users", handler.GetAllUsers)
    rg.PUT("/users/:id", handler.UpdateUser)
    rg.DELETE("/users/:id", handler.DeleteUser)
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body dto.CreateUserDTO true "Create User"
// @Success 200 {object} entity.User
// @Router /users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
    var createUserDTO dto.CreateUserDTO
    if err := c.ShouldBindJSON(&createUserDTO); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := h.usecase.CreateUser(createUserDTO)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, user)
}

// GetUserByID godoc
// @Summary Get user by ID
// @Description Get user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} entity.User
// @Router /users/{id} [get]
func (h *UserHandler) GetUserByID(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    user, err := h.usecase.GetUserByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, user)
}

// GetAllUsers godoc
// @Summary Get all users
// @Description Get all users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} entity.User
// @Router /users [get]
func (h *UserHandler) GetAllUsers(c *gin.Context) {
    users, err := h.usecase.GetAllUsers()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, users)
}

// UpdateUser godoc
// @Summary Update user
// @Description Update user
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body dto.UpdateUserDTO true "Update User"
// @Success 200 {object} entity.User
// @Router /users/{id} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
    var updateUserDTO dto.UpdateUserDTO
    if err := c.ShouldBindJSON(&updateUserDTO); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    idParam := c.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    user, err := h.usecase.UpdateUser(uint(id), updateUserDTO)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, user)
}

// DeleteUser godoc
// @Summary Delete user
// @Description Delete user
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} map[string]interface{} "message: User deleted"
// @Router /users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    if err := h.usecase.DeleteUser(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
