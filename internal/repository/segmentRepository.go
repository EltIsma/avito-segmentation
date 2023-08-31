package repository

import (
	"avito-third/internal/segment"
	"avito-third/pkg/client/postgresql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type SegmentRepository struct {
	db *sqlx.DB
}

func NewSegmentDB(db *sqlx.DB) *SegmentRepository {
	return &SegmentRepository{db: db}
}

func (r *SegmentRepository) Create(segmentDTO *segment.SegmentDTO) error {
	var id string
	query := fmt.Sprintf("INSERT INTO %s (slug) values ($1) ON CONFLICT DO NOTHING RETURNING id", postgresql.SegmentTable)
	row := r.db.QueryRow(query, segmentDTO.Name)
	if err := row.Scan(&id); err != nil {
		return err
	}

	return nil

}

func (r *SegmentRepository) Delete(segmentDTO *segment.SegmentDTO) error {
	res, err := r.db.Exec(`DELETE FROM storage where slug = $1`, segmentDTO.Name)
	if err != nil {
		return err
	}
	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAfected == 0 {
		return nil
	}

	return nil

}
