package repository

import (
    "myapp/internal/app/entity"
    "gorm.io/gorm"
)

type NoteRepository interface {
    Create(note entity.Note) (entity.Note, error)
    GetByID(id uint) (entity.Note, error)
    GetAll() ([]entity.Note, error)
    Update(note entity.Note) (entity.Note, error)
    Delete(id uint) error
}

type noteRepository struct {
    db *gorm.DB
}

func NewNoteRepository(db *gorm.DB) NoteRepository {
    return &noteRepository{db: db}
}

func (r *noteRepository) Create(note entity.Note) (entity.Note, error) {
    result := r.db.Create(&note)
    return note, result.Error
}

func (r *noteRepository) GetByID(id uint) (entity.Note, error) {
    var note entity.Note
    result := r.db.First(&note, id)
    return note, result.Error
}

func (r *noteRepository) GetAll() ([]entity.Note, error) {
    var notes []entity.Note
    result := r.db.Find(&notes)
    return notes, result.Error
}

func (r *noteRepository) Update(note entity.Note) (entity.Note, error) {
    result := r.db.Save(&note)
    return note, result.Error
}

func (r *noteRepository) Delete(id uint) error {
    result := r.db.Delete(&entity.Note{}, id)
    return result.Error
}
