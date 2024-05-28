package resource

type FindEventRequest struct {
	Id        uint   `json:"id"`
	Latitude  string `json:"latitude" query:"latitude"`
	Longitude string `json:"longitude" query:"longitude"`
}
