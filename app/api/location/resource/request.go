package resource

type FindEventRequest struct {
	Latitude  float64 `json:"latitude" query:"latitude"`
	Longitude float64 `json:"longitude" query:"longitude"`
}
