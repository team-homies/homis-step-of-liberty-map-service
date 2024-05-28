package resource

type FindEventResponse struct {
	Id        uint   `json:"id"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	IsCollect bool   `json:"is_collect"`
}
