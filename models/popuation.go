package models

import "time"

// PopulationResponse ...
type PopulationResponse struct {
	Population []map[string]interface{} `json:"Population"`
	Errors     []map[string]interface{} `json:"Errors"`
}

type Data struct {
	Population   int    `json:"Population`
	ErrorMessage string `json:"ErrorMessage"`
}

type PopulationClientResponse struct {
	Population map[string]interface{} `json:"Population"`
	Erorr      string                 `json:"Erorr"`
}

// type Population struct {
// 	A int `json:"a"`
// 	B int `json:"b"`
// }

type PopulationRequest struct {
	Countries []string  `json:"countries"`
	Date      time.Time `json:"Date"`
	Sort      string    `json:"sort"`
}

type PopulationThirdPartyResponse struct {
	TotalPopulation TotalPopulation `json:"total_population"`
}

type TotalPopulation struct {
	Date       string `json:"date"`
	Population int    `json:"population"`
}
