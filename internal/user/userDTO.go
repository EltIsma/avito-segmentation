package user

type UsersActiveSlugsDTO struct {
	User_id int `json:"user_id"`
}

type ReportSegmentRequest struct {
	Period string `json:"period"`
}
