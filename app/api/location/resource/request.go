package resource

type FindEventRequest struct {
	Id        uint    `json:"id"`
	Latitude  float64 `json:"latitude" query:"latitude"`
	Longitude float64 `json:"longitude" query:"longitude"`
}
