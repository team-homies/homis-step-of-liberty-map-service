package resource

type GetPatientResponse struct {
	Patient MapResource `json:"patient"`
}

type GetPatientsResponse struct {
	Patients []MapResource `json:"patients"`
}
