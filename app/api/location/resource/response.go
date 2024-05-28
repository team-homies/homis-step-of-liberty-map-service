package resource

type GetPatientResponse struct {
	Patient locationResource `json:"patient"`
}

type GetPatientsResponse struct {
	Patients []locationResource `json:"patients"`
}
