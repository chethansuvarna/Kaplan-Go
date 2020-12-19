package models

// PopulationResponse ...
type PopulationResponse struct {
	Population []map[string]interface{} `json:"Population"`
	Errors     []map[string]interface{} `json:"Errors"`
}

type PopulationRequest struct {
	Countries []string `json:"countries"`
	Date      string   `json:"Date"`
	Sort      string   `json:"sort"`
}

type PopulationThirdPartyResponse struct {
	TotalPopulation TotalPopulation `json:"total_population"`
}

type TotalPopulation struct {
	Date       string `json:"date"`
	Population int    `json:"population"`
}
