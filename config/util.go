package config

type Util interface {
	GetString(property string) string
}

type configUtil struct {
	config map[string]interface{}
}

func NewConfigUtil(config map[string]interface{}) Util {
	return configUtil{config: config}
}

func (c configUtil) GetString(property string) string {
	return c.config[property].(string)
}
