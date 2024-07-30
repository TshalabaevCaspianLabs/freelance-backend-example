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

type MatchHandler struct {
    usecase usecase.MatchUsecase
}

func RegisterMatchRoutes(rg *gin.RouterGroup, db *gorm.DB) {
    matchRepo := repository.NewMatchRepository(db)
    handler := &MatchHandler{usecase: usecase.NewMatchUsecase(matchRepo)}
    rg.POST("/matches", handler.CreateMatch)
    rg.GET("/matches/:id", handler.GetMatchByID)
    rg.GET("/matches", handler.GetAllMatches)
    rg.PUT("/matches/:id", handler.UpdateMatch)
    rg.DELETE("/matches/:id", handler.DeleteMatch)
}

// CreateMatch godoc
// @Summary Create a new match
// @Description Create a new match
// @Tags matches
// @Accept json
// @Produce json
// @Param match body dto.CreateMatchDTO true "Create Match"
// @Success 200 {object} entity.Match
// @Router /matches [post]
func (h *MatchHandler) CreateMatch(c *gin.Context) {
    var createMatchDTO dto.CreateMatchDTO
    if err := c.ShouldBindJSON(&createMatchDTO); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    match, err := h.usecase.CreateMatch(createMatchDTO)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, match)
}

// GetMatchByID godoc
// @Summary Get match by ID
// @Description Get match by ID
// @Tags matches
// @Accept json
// @Produce json
// @Param id path int true "Match ID"
// @Success 200 {object} entity.Match
// @Router /matches/{id} [get]
func (h *MatchHandler) GetMatchByID(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    match, err := h.usecase.GetMatchByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, match)
}

// GetAllMatches godoc
// @Summary Get all matches
// @Description Get all matches
// @Tags matches
// @Accept json
// @Produce json
// @Success 200 {array} entity.Match
// @Router /matches [get]
func (h *MatchHandler) GetAllMatches(c *gin.Context) {
    matches, err := h.usecase.GetAllMatches()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, matches)
}

// UpdateMatch godoc
// @Summary Update match
// @Description Update match
// @Tags matches
// @Accept json
// @Produce json
// @Param id path int true "Match ID"
// @Param match body dto.UpdateMatchDTO true "Update Match"
// @Success 200 {object} entity.Match
// @Router /matches/{id} [put]
func (h *MatchHandler) UpdateMatch(c *gin.Context) {
    var updateMatchDTO dto.UpdateMatchDTO
    if err := c.ShouldBindJSON(&updateMatchDTO); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    idParam := c.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    match, err := h.usecase.UpdateMatch(uint(id), updateMatchDTO)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, match)
}

// DeleteMatch godoc
// @Summary Delete match
// @Description Delete match
// @Tags matches
// @Accept json
// @Produce json
// @Param id path int true "Match ID"
// @Success 200 {object} map[string]interface{} "message: Match deleted"
// @Router /matches/{id} [delete]
func (h *MatchHandler) DeleteMatch(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    if err := h.usecase.DeleteMatch(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Match deleted"})
}
