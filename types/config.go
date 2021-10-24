package types

type RedisConf struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int8   `mapstructure:"port"`
}

type MqConf struct {
	Username string `mapstructure:"username"`
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     string `mapstructure:"port"`
}
type QueueConifg struct {
	Exchange     string   `mapstructure:"exchange"`
	ExchangeType string   `mapstructure:"exchangeType"`
	Name         string   `mapstructure:"name"`
	Consumer     string   `mapstructue:"consumer"`
	Bind         []string `mapstructure:"bind"`
}
