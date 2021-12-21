package service

import (
	"github.com/emlog/goexample/dao/note"
	"github.com/emlog/goexample/model/request"
)

type NoteService struct {
	noteDao *note.NoteDao
}

func NewNoteService() *NoteService {
	return &NoteService{
		noteDao: note.NewNoteDao(),
	}
}

func (s *NoteService) CreateNote(req *request.ReqNote) (int64, error) {
	id, err := NewNoteService().noteDao.Create(req)
	return id, err
}
