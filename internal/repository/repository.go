package repository

import (
	"avito-third/internal/segment"
	"avito-third/internal/user"
	"time"

	"github.com/jmoiron/sqlx"
)

type Segment interface {
	Create(segment *segment.SegmentDTO) ( error)
	Delete(segment *segment.SegmentDTO) ( error)
}

type User interface {
	CRUDOperation(elem *user.UserSegment) ( error)
	GetActive(userId int) ([]string, error)
	GetReport(period time.Time) ([]user.ReportUsers, error)
	DeleteExpiredUser() ( error) 
}

type Repository struct {
	Segment
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Segment: NewSegmentDB(db),
		User:    NewUserDB(db),
	}
}
