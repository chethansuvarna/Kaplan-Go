package config

import "Kaplan-Go/constants"

type Config interface {
	GetPort() string
	GetCountriesEndPoint() string
	GetPopulationEndpoint() string
}

type config struct {
	Util
}

func NewConfig(configuration Util) Config {
	return config{configuration}
}

func (c config) GetPort() string {
	return c.GetString(constants.Port)
}

func (c config) GetCountriesEndPoint() string {
	return c.GetString(constants.GetCountriesEndPoint)
}

func (c config) GetPopulationEndpoint() string {
	return c.GetString(constants.GetPopulationEndpoint)
}
