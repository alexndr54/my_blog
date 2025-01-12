package entity

type ResponseDataRoute struct {
	Title      string
	IsLogin    bool
	Validation ValidationResponse
	Optional   interface{}
}

type ValidationResponse struct {
	Status              StatusResponse
	ValidationListError map[string]ValidationList
	ValidationValue     map[string]string
}

type StatusResponse struct {
	Label   string
	Message string
}

type ValidationList struct {
	Message string
	Value   string
}
