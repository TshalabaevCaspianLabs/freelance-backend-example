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

type RatingHandler struct {
    usecase usecase.RatingUsecase
}

func RegisterRatingRoutes(rg *gin.RouterGroup, db *gorm.DB) {
    handler := &RatingHandler{usecase: usecase.NewRatingUsecase(repository.NewRatingRepository(db))}
    rg.POST("/ratings", handler.CreateRating)
    rg.GET("/ratings/:id", handler.GetRatingByID)
    rg.GET("/ratings", handler.GetAllRatings)
    rg.PUT("/ratings/:id", handler.UpdateRating)
    rg.DELETE("/ratings/:id", handler.DeleteRating)
}

// CreateRating godoc
// @Summary Create a new rating
// @Description Create a new rating
// @Tags ratings
// @Accept json
// @Produce json
// @Param rating body dto.CreateRatingDTO true "Create Rating"
// @Success 200 {object} entity.Rating
// @Router /ratings [post]
func (h *RatingHandler) CreateRating(c *gin.Context) {
    var createRatingDTO dto.CreateRatingDTO
    if err := c.ShouldBindJSON(&createRatingDTO); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    rating, err := h.usecase.CreateRating(createRatingDTO)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, rating)
}

// GetRatingByID godoc
// @Summary Get rating by ID
// @Description Get rating by ID
// @Tags ratings
// @Accept json
// @Produce json
// @Param id path int true "Rating ID"
// @Success 200 {object} entity.Rating
// @Router /ratings/{id} [get]
func (h *RatingHandler) GetRatingByID(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    rating, err := h.usecase.GetRatingByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, rating)
}

// GetAllRatings godoc
// @Summary Get all ratings
// @Description Get all ratings
// @Tags ratings
// @Accept json
// @Produce json
// @Success 200 {array} entity.Rating
// @Router /ratings [get]
func (h *RatingHandler) GetAllRatings(c *gin.Context) {
    ratings, err := h.usecase.GetAllRatings()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, ratings)
}

// UpdateRating godoc
// @Summary Update rating
// @Description Update rating
// @Tags ratings
// @Accept json
// @Produce json
// @Param id path int true "Rating ID"
// @Param rating body dto.UpdateRatingDTO true "Update Rating"
// @Success 200 {object} entity.Rating
// @Router /ratings/{id} [put]
func (h *RatingHandler) UpdateRating(c *gin.Context) {
    var updateRatingDTO dto.UpdateRatingDTO
    if err := c.ShouldBindJSON(&updateRatingDTO); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    idParam := c.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    rating, err := h.usecase.UpdateRating(uint(id), updateRatingDTO)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, rating)
}

// DeleteRating godoc
// @Summary Delete rating
// @Description Delete rating
// @Tags ratings
// @Accept json
// @Produce json
// @Param id path int true "Rating ID"
// @Success 200 {object} map[string]interface{} "message: Rating deleted"
// @Router /ratings/{id} [delete]
func (h *RatingHandler) DeleteRating(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    if err := h.usecase.DeleteRating(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Rating deleted"})
}
