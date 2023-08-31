package service

import (
	"avito-third/internal/repository"
	"avito-third/internal/segment"
	"avito-third/internal/user"
	"time"
)

type Segment interface{
	Create(segment *segment.SegmentDTO) ( error)
	Delete(segment *segment.SegmentDTO) (error)
}

type User interface{
	CRUDOperation(elem *user.UserSegment) ( error)
	GetActive(userId int) ([]string, error)
	GetReport(period time.Time) ([]user.ReportUsers, error)
}

type Service struct{
	Segment
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Segment: NewSegmentService(repos.Segment),
		User: NewUserService(repos.User),
	}
}