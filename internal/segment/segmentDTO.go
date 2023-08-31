package segment

type SegmentDTO struct {
	Name string `json:"segment_name" db:"slug"`
}


type SegmentResponseDTO struct {
	Result bool `json:"result"`
}
