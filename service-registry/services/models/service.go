package models

//easyjson:json
type ServiceInfo struct {
	Address string `json:"address"`
	Status  string `json:"status"`
}

//easyjson:json
type Service struct {
	Name      string        `json:"name"`
	Instances []ServiceInfo `json:"instances"`
}

//easyjson:json
type Services struct {
	Services []Service `json:"services"`
}
