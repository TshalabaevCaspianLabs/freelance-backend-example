package usecase

import (
    "myapp/internal/app/dto"
    "myapp/internal/app/entity"
    "myapp/internal/repository"
)

type NoteUsecase interface {
    CreateNote (dto dto.CreateNoteDTO) (entity.Note, error)
    GetNoteByID(id uint) (entity.Note, error)
    GetAllNotes() ([]entity.Note, error)
    UpdateNote (id uint, dto dto.UpdateNoteDTO) (entity.Note, error)
    DeleteNote (id uint) error
}

type noteUsecase struct {
    repo repository.NoteRepository
}

func NewNoteUsecase(repo repository.NoteRepository) NoteUsecase {
    return &noteUsecase{repo: repo}
}

func (u *noteUsecase) CreateNote(dto dto.CreateNoteDTO) (entity.Note, error) {
    note := entity.Note{Title: dto.Title, Content: dto.Content, UserID: dto.UserID}
    return u.repo.Create(note)
}

func (u *noteUsecase) GetNoteByID(id uint) (entity.Note, error) {
    return u.repo.GetByID(id)
}

func (u *noteUsecase) GetAllNotes() ([]entity.Note, error) {
    return u.repo.GetAll()
}

func (u *noteUsecase) UpdateNote (id uint, dto dto.UpdateNoteDTO) (entity.Note, error) {
    note, err := u.repo.GetByID(id)
    if err != nil {
        return entity.Note{}, err
    }

    note.Title = dto.Title
    note.Content = dto.Content
    return u.repo.Update(note)
}

func (u *noteUsecase) DeleteNote (id uint) error {
    return u.repo.Delete(id)
}
