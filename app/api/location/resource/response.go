package resource

type FindEventResponse struct {
	Id        uint    `json:"id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	IsCollect bool    `json:"is_collect"`
}
