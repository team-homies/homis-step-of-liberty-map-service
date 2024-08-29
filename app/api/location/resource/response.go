package resource

type MapInfoResponse struct {
	EventId     uint    `json:"event_id"`
	Name        string  `json:"name"`
	Address     string  `json:"address"`
	AddressRoad string  `json:"address_road"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}
type FindEventResponse struct {
	EventId   uint    `json:"event_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	IsCollect bool    `json:"is_collect"`
}

type FindLocationListResponse struct {
	EventId     uint    `json:"event_id"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Distance    string  `json:"distance"`
	Address     string  `json:"address"`
	AddressRoad string  `json:"address_road"`
	IsCollect   bool    `json:"is_collect"`
}
