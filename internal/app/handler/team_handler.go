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

type TeamHandler struct {
    usecase usecase.TeamUsecase
}

func RegisterTeamRoutes(rg *gin.RouterGroup, db *gorm.DB) {
    teamRepo := repository.NewTeamRepository(db)
    userRepo := repository.NewUserRepository(db)
    handler := &TeamHandler{usecase: usecase.NewTeamUsecase(teamRepo, userRepo)}
    rg.POST("/teams", handler.CreateTeam)
    rg.GET("/teams/:id", handler.GetTeamByID)
    rg.GET("/teams", handler.GetAllTeams)
    rg.PUT("/teams/:id", handler.UpdateTeam)
    rg.DELETE("/teams/:id", handler.DeleteTeam)
}

// CreateTeam godoc
// @Summary Create a new team
// @Description Create a new team
// @Tags teams
// @Accept json
// @Produce json
// @Param team body dto.CreateTeamDTO true "Create Team"
// @Success 200 {object} entity.Team
// @Router /teams [post]
func (h *TeamHandler) CreateTeam(c *gin.Context) {
    var createTeamDTO dto.CreateTeamDTO
    if err := c.ShouldBindJSON(&createTeamDTO); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    team, err := h.usecase.CreateTeam(createTeamDTO)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, team)
}

// GetTeamByID godoc
// @Summary Get team by ID
// @Description Get team by ID
// @Tags teams
// @Accept json
// @Produce json
// @Param id path int true "Team ID"
// @Success 200 {object} entity.Team
// @Router /teams/{id} [get]
func (h *TeamHandler) GetTeamByID(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    team, err := h.usecase.GetTeamByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, team)
}

// GetAllTeams godoc
// @Summary Get all teams
// @Description Get all teams
// @Tags teams
// @Accept json
// @Produce json
// @Success 200 {array} entity.Team
// @Router /teams [get]
func (h *TeamHandler) GetAllTeams(c *gin.Context) {
    teams, err := h.usecase.GetAllTeams()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, teams)
}

// UpdateTeam godoc
// @Summary Update team
// @Description Update team
// @Tags teams
// @Accept json
// @Produce json
// @Param id path int true "Team ID"
// @Param team body dto.UpdateTeamDTO true "Update Team"
// @Success 200 {object} entity.Team
// @Router /teams/{id} [put]
func (h *TeamHandler) UpdateTeam(c *gin.Context) {
    var updateTeamDTO dto.UpdateTeamDTO
    if err := c.ShouldBindJSON(&updateTeamDTO); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    idParam := c.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    team, err := h.usecase.UpdateTeam(uint(id), updateTeamDTO)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, team)
}

// DeleteTeam godoc
// @Summary Delete team
// @Description Delete team
// @Tags teams
// @Accept json
// @Produce json
// @Param id path int true "Team ID"
// @Success 200 {object} map[string]interface{} "message: Team deleted"
// @Router /teams/{id} [delete]
func (h *TeamHandler) DeleteTeam(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    if err := h.usecase.DeleteTeam(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Team deleted"})
}
