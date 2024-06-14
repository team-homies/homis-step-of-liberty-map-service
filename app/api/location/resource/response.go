package resource

type FindEventResponse struct {
	EventId   uint    `json:"event_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	IsCollect bool    `json:"is_collect"`
}
