package population

import (
	"Kaplan-Go/models"
	"encoding/json"
)

var (
	getCountryPopulationURL = "/getPopulationForCountry"

	helloWroldResponse = "hello world"
	expectedResponse   = map[string]interface{}{"countries": []string{
		"Afghanistan",
		"AFRICA",
		"Albania",
	},
	}

	getCountriesGoRequestAscending = models.PopulationRequest{
		Countries: []string{
			"Afghanistan",
			"AFRICA",
			"Albania",
		},
		Sort: "ascending",
	}

	getCountriesGoExpectedResponseAscending = models.PopulationResponse{
		Population: []map[string]interface{}{
			{
				"country":    "Afghanistan",
				"population": 100,
			},
			{
				"country":    "AFRICA",
				"population": 200,
			},
			{
				"country":    "Albania",
				"population": 300,
			},
		},
	}
	getcountryPopulationResponse models.PopulationThirdPartyResponse

	_ = json.Unmarshal([]byte(`{
		total_population: {
		date: "2020-12-17",
		population: 297152
		}
		}`), &getcountryPopulationResponse)

	getCountriesGoExpectedResponseDescending = map[string]interface{}{"Population": []map[string]interface{}{
		{
			"country":    "Afghanistan",
			"population": 100,
		},
		{
			"country":    "AFRICA",
			"population": 200,
		},
		{
			"country":    "Albania",
			"population": 300,
		},
	},
	}
)
