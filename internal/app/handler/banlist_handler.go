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

type BanListHandler struct {
    usecase usecase.BanListUsecase
}

func RegisterBanListRoutes(rg *gin.RouterGroup, db *gorm.DB) {
    handler := &BanListHandler{usecase: usecase.NewBanListUsecase(repository.NewBanListRepository(db))}
    rg.POST("/banlists", handler.CreateBanList)
    rg.GET("/banlists/:id", handler.GetBanListByID)
    rg.GET("/banlists", handler.GetAllBanLists)
    rg.PUT("/banlists/:id", handler.UpdateBanList)
    rg.DELETE("/banlists/:id", handler.DeleteBanList)
}

// CreateBanList godoc
// @Summary Create a new ban list entry
// @Description Create a new ban list entry
// @Tags banlists
// @Accept json
// @Produce json
// @Param banlist body dto.CreateBanListDTO true "Create BanList"
// @Success 200 {object} entity.BanList
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /banlists [post]
func (h *BanListHandler) CreateBanList(c *gin.Context) {
    var createBanListDTO dto.CreateBanListDTO
    if err := c.ShouldBindJSON(&createBanListDTO); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    banList, err := h.usecase.CreateBanList(createBanListDTO)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, banList)
}

// GetBanListByID godoc
// @Summary Get ban list entry by ID
// @Description Get ban list entry by ID
// @Tags banlists
// @Accept json
// @Produce json
// @Param id path int true "BanList ID"
// @Success 200 {object} entity.BanList
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 404 {object} map[string]interface{} "Not Found"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /banlists/{id} [get]
func (h *BanListHandler) GetBanListByID(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    banList, err := h.usecase.GetBanListByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, banList)
}

// GetAllBanLists godoc
// @Summary Get all ban list entries
// @Description Get all ban list entries
// @Tags banlists
// @Accept json
// @Produce json
// @Success 200 {array} entity.BanList
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /banlists [get]
func (h *BanListHandler) GetAllBanLists(c *gin.Context) {
    banLists, err := h.usecase.GetAllBanLists()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, banLists)
}

// UpdateBanList godoc
// @Summary Update ban list entry
// @Description Update ban list entry
// @Tags banlists
// @Accept json
// @Produce json
// @Param id path int true "BanList ID"
// @Param banlist body dto.UpdateBanListDTO true "Update BanList"
// @Success 200 {object} entity.BanList
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 404 {object} map[string]interface{} "Not Found"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /banlists/{id} [put]
func (h *BanListHandler) UpdateBanList(c *gin.Context) {
    var updateBanListDTO dto.UpdateBanListDTO
    if err := c.ShouldBindJSON(&updateBanListDTO); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    idParam := c.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    banList, err := h.usecase.UpdateBanList(uint(id), updateBanListDTO)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, banList)
}

// DeleteBanList godoc
// @Summary Delete ban list entry
// @Description Delete ban list entry
// @Tags banlists
// @Accept json
// @Produce json
// @Param id path int true "BanList ID"
// @Success 200 {object} map[string]interface{} "message: BanList deleted"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 404 {object} map[string]interface{} "Not Found"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /banlists/{id} [delete]
func (h *BanListHandler) DeleteBanList(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    if err := h.usecase.DeleteBanList(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, map[string]interface{}{"message": "BanList deleted"})
}
