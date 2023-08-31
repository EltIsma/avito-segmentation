package repository

import (
	"avito-third/internal/user"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserDB(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CRUDOperation(elem *user.UserSegment) error {
	tx, err := r.db.Beginx()

	defer func() {
		if err != nil {
			errRB := tx.Rollback()
			if errRB != nil {
				return
			}
			return
		}
		err = tx.Commit()
	}()
	var exist bool
	var res sql.Result
	query := "SELECT EXISTS (SELECT * FROM userssegment where user_id = $1)"
	err = r.db.QueryRow(query, elem.User_id).Scan(&exist)
	if err != nil {
		return err
	}
	slugName := make([]string, 0)
	for _, name := range elem.SegmentForAdd {
		slugName = append(slugName, name.Slug)
	}

	if exist == true {
		res, err = tx.Exec(`UPDATE userssegment SET slug_name = slug_name || $1 WHERE user_id = $2;`, pq.Array(slugName), elem.User_id)
		if err != nil {
			return err
		}
	} else {
		res, err = tx.Exec(`INSERT INTO UsersSegment (user_id, slug_name) VALUES ($1, $2);`, elem.User_id, pq.Array(slugName))
		if err != nil {
			return err
		}
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected > 0 {
		for _, slug := range elem.SegmentForAdd {
			if slug.TTL == "" {
				_, err := tx.Exec(`INSERT INTO Report (user_id, slug, operation, execution) VALUES ($1, $2, $3, $4);`, elem.User_id, slug.Slug, user.CreateOperation, time.Now().UTC().Format(time.RFC3339))
				if err != nil {
					return err
				}
			} else {
				ttl, _ := time.Parse(time.DateTime, slug.TTL)
				_, err := tx.Exec(`INSERT INTO Report (user_id, slug, operation, execution, ttl) VALUES ($1, $2, $3, $4, $5::timestamp);`, elem.User_id, slug.Slug, user.CreateOperation, time.Now().UTC().Format(time.RFC3339), ttl)
				if err != nil {
					return err
				}
			}

		}

	}

	res2, err := tx.Exec(`UPDATE userssegment SET slug_name = ARRAY(SELECT UNNEST(slug_name) EXCEPT SELECT UNNEST($1))	WHERE $1 && slug_name AND user_id= $2;`, pq.Array(elem.SegmentForDelete), elem.User_id)
	if err != nil {
		return err
	}
	rowsAffected2, err := res2.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected2 > 0 {
		for _, slug := range elem.SegmentForDelete {
			_, err := tx.Exec(`INSERT INTO Report (user_id, slug, operation, execution) VALUES ($1, $2, $3, $4);`, elem.User_id, slug, user.DeleteOperation, time.Now().UTC().Format(time.RFC3339))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *UserRepository) GetActive(userId int) ([]string, error) {
	segments := make([]string, 0)
	err := r.db.Select(&segments, "select  ARRAY(SELECT UNNEST(slug_name) from userssegment where user_id=$1 INTERSECT SELECT slug from storage)", userId)
	if err != nil {
		return segments, nil
	}
	segments[0] = segments[0][1 : len(segments[0])-1]
	words := strings.Split(segments[0], ",")
	var final []string
	for _, word := range words {
		final = append(final, word)
	}
	return final, nil
}

func (r *UserRepository) GetReport(period time.Time) ([]user.ReportUsers, error) {
	var report []user.ReportUsers

	err := r.db.Select(&report, "select * from report where execution <= $1::timestamp ORDER BY execution DESC", period.UTC().Format(time.RFC3339))
	if err != nil {
		return report, nil
	}

	return report, nil

}
func (r *UserRepository) DeleteExpiredUser() error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
		//fmt.Println("dfgjdsfdsdgfhsjkdfhsdhkfasdhjf")
	}
	defer tx.Rollback()
	timeNow := time.Now().UTC().Format(time.RFC3339)
	query := "select user_id, slug FROM report where operation = 'CREATE' AND ttl < $1::timestamp"
	rows, _ := tx.Query(query, timeNow)

	if rows.Next() {
		for rows.Next() {
			var userID int
			var slug string
			err := rows.Scan(&userID, &slug)
			if err != nil {

				log.Fatal(err)
			}
			fmt.Println(userID, slug)
			query := "UPDATE userssegment SET slug_name = array_remove(slug_name, $1) WHERE user_id = $2 "
			r.db.QueryRow(query, slug, userID)
			r.db.QueryRow("UPDATE report SET ttl = '9999-01-01 00:00:00'::timestamp WHERE slug = $1 and  user_id= $2 and operation='CREATE'", slug, userID)

		}
	}

	return tx.Commit()
}
