package user

import "time"


const (
	CreateOperation = "CREATE"
	DeleteOperation = "DELETE"
)


type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UserSegment struct {
	SegmentForAdd    []AddSegment `json:"segment_for_adding,omitempty"`
	SegmentForDelete []string `json:"segment_for_deleting,omitempty"`
	User_id          int      `json:"user_id"`
}

type ReportUsers struct {
	ID          int           `db:"id"`
	UserID      int           `db:"user_id"`
	Slug        string        `db:"slug"`
	Operation   string        `db:"operation"`
	TimeOperation   time.Time `db:"execution"`
}

type AddSegment struct {
	Slug string `json:"slug"`
	TTL  string `json:"ttl,omitempty"`
}
