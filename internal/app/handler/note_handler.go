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

type NoteHandler struct {
    usecase usecase.NoteUsecase
}

func RegisterNoteRoutes(rg *gin.RouterGroup, db *gorm.DB) {
    handler := &NoteHandler{usecase: usecase.NewNoteUsecase(repository.NewNoteRepository(db))}
    rg.POST("/notes", handler.CreateNote)
    rg.GET("/notes/:id", handler.GetNoteByID)
    rg.GET("/notes", handler.GetAllNotes)
    rg.PUT("/notes/:id", handler.UpdateNote)
    rg.DELETE("/notes/:id", handler.DeleteNote)
}

// CreateNote godoc
// @Summary Create a new note
// @Description Create a new note
// @Tags notes
// @Accept json
// @Produce json
// @Param note body dto.CreateNoteDTO true "Create Note"
// @Success 200 {object} entity.Note
// @Router /notes [post]
func (h *NoteHandler) CreateNote(c *gin.Context) {
    var createNoteDTO dto.CreateNoteDTO
    if err := c.ShouldBindJSON(&createNoteDTO); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    note, err := h.usecase.CreateNote(createNoteDTO)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, note)
}

// GetNoteByID godoc
// @Summary Get note by ID
// @Description Get note by ID
// @Tags notes
// @Accept json
// @Produce json
// @Param id path int true "Note ID"
// @Success 200 {object} entity.Note
// @Router /notes/{id} [get]
func (h *NoteHandler) GetNoteByID(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    note, err := h.usecase.GetNoteByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, note)
}

// GetAllNotes godoc
// @Summary Get all notes
// @Description Get all notes
// @Tags notes
// @Accept json
// @Produce json
// @Success 200 {array} entity.Note
// @Router /notes [get]
func (h *NoteHandler) GetAllNotes(c *gin.Context) {
    notes, err := h.usecase.GetAllNotes()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, notes)
}

// UpdateNote godoc
// @Summary Update note
// @Description Update note
// @Tags notes
// @Accept json
// @Produce json
// @Param id path int true "Note ID"
// @Param note body dto.UpdateNoteDTO true "Update Note"
// @Success 200 {object} entity.Note
// @Router /notes/{id} [put]
func (h *NoteHandler) UpdateNote(c *gin.Context) {
    var updateNoteDTO dto.UpdateNoteDTO
    if err := c.ShouldBindJSON(&updateNoteDTO); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    idParam := c.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    note, err := h.usecase.UpdateNote(uint(id), updateNoteDTO)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, note)
}

// DeleteNote godoc
// @Summary Delete note
// @Description Delete note
// @Tags notes
// @Accept json
// @Produce json
// @Param id path int true "Note ID"
// @Success 200 {object} map[string]interface{} "message: Note deleted"
// @Router /notes/{id} [delete]
func (h *NoteHandler) DeleteNote(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    if err := h.usecase.DeleteNote(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Note deleted"})
}
