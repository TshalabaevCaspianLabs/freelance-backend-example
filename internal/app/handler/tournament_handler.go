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

type TournamentHandler struct {
    usecase usecase.TournamentUsecase
}

func RegisterTournamentRoutes(rg *gin.RouterGroup, db *gorm.DB) {
    handler := &TournamentHandler{usecase: usecase.NewTournamentUsecase(repository.NewTournamentRepository(db))}
    rg.POST("/tournaments", handler.CreateTournament)
    rg.GET("/tournaments/:id", handler.GetTournamentByID)
    rg.GET("/tournaments", handler.GetAllTournaments)
    rg.PUT("/tournaments/:id", handler.UpdateTournament)
    rg.DELETE("/tournaments/:id", handler.DeleteTournament)
}

// CreateTournament godoc
// @Summary Create a new tournament
// @Description Create a new tournament
// @Tags tournaments
// @Accept json
// @Produce json
// @Param tournament body dto.CreateTournamentDTO true "Create Tournament"
// @Success 200 {object} entity.Tournament
// @Router /tournaments [post]
func (h *TournamentHandler) CreateTournament(c *gin.Context) {
    var createTournamentDTO dto.CreateTournamentDTO
    if err := c.ShouldBindJSON(&createTournamentDTO); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    tournament, err := h.usecase.CreateTournament(createTournamentDTO)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, tournament)
}

// GetTournamentByID godoc
// @Summary Get tournament by ID
// @Description Get tournament by ID
// @Tags tournaments
// @Accept json
// @Produce json
// @Param id path int true "Tournament ID"
// @Success 200 {object} entity.Tournament
// @Router /tournaments/{id} [get]
func (h *TournamentHandler) GetTournamentByID(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    tournament, err := h.usecase.GetTournamentByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, tournament)
}

// GetAllTournaments godoc
// @Summary Get all tournaments
// @Description Get all tournaments
// @Tags tournaments
// @Accept json
// @Produce json
// @Success 200 {array} entity.Tournament
// @Router /tournaments [get]
func (h *TournamentHandler) GetAllTournaments(c *gin.Context) {
    tournaments, err := h.usecase.GetAllTournaments()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, tournaments)
}

// UpdateTournament godoc
// @Summary Update tournament
// @Description Update tournament
// @Tags tournaments
// @Accept json
// @Produce json
// @Param id path int true "Tournament ID"
// @Param tournament body dto.UpdateTournamentDTO true "Update Tournament"
// @Success 200 {object} entity.Tournament
// @Router /tournaments/{id} [put]
func (h *TournamentHandler) UpdateTournament(c *gin.Context) {
    var updateTournamentDTO dto.UpdateTournamentDTO
    if err := c.ShouldBindJSON(&updateTournamentDTO); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    idParam := c.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    tournament, err := h.usecase.UpdateTournament(uint(id), updateTournamentDTO)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, tournament)
}

// DeleteTournament godoc
// @Summary Delete tournament
// @Description Delete tournament
// @Tags tournaments
// @Accept json
// @Produce json
// @Param id path int true "Tournament ID"
// @Success 200 {object} map[string]interface{} "message: Tournament deleted"
// @Router /tournaments/{id} [delete]
func (h *TournamentHandler) DeleteTournament(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    if err := h.usecase.DeleteTournament(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Tournament deleted"})
}
