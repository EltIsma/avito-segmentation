package service

import (
	"avito-third/internal/repository"
	"avito-third/internal/segment"
)


type SegmentService struct {
	repo repository.Segment
}

func NewSegmentService(repo repository.Segment) *SegmentService {
	return &SegmentService{repo: repo}
}

func (s *SegmentService) Create(segment *segment.SegmentDTO) ( error) {
	return s.repo.Create(segment)
}

func (s *SegmentService) Delete(segment *segment.SegmentDTO) ( error) {
	return s.repo.Delete(segment)
}