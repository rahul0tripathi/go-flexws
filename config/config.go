package config

import (
	"fmt"
	"github.com/rahul0tripathi/fastws/logger"
	"github.com/rahul0tripathi/fastws/types"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

type Config struct {
	Env            string            `mapstructure:"env"`
	Port           string            `mapstructure:"port"`
	Debug          bool              `mapstructure:"debug"`
	Cache          types.RedisConf   `mapstructure:"cache"`
	Mq             types.MqConf      `mapstructure:"amqp"`
	Datastore      types.RedisConf   `mapstructure:"datastore"`
	RoomQueue types.QueueConifg `mapstructure:"roomQueue"`
}

var (
	AppConfig Config
	ConfigDir string
	Debug     = true
)

func ReadAndUnmarshal(configname string, format string, object interface{}, sub interface{}) error {
	conf := viper.New()
	conf.SetConfigName(configname)
	conf.SetConfigType(format)
	conf.AddConfigPath(ConfigDir)
	if err := conf.ReadInConfig(); err != nil {
		fmt.Printf("error reading Sub config %v", err)
		return err
	}
	if sub != nil {
		subItem := conf.Sub(sub.(string))
		return subItem.Unmarshal(object)
	}
	return conf.Unmarshal(object)
}
func setEnv(object map[string]string) {
	for key, value := range object {
		err := os.Setenv(strings.ToUpper(key), value)
		if err != nil {
			fmt.Println("Unable to set env ", key)
		}
	}
}

func init() {
	var err error
	ConfigDir, err = os.Getwd()
	if err != nil {
		log.Fatalln("failed to get workdir", err)
	}
	err = ReadAndUnmarshal("config", "json", &AppConfig, nil)
	if err != nil {
		log.Fatalln("failed to load config", err)
	}
	Debug = AppConfig.Debug
	logger.SetLogLevel(&Debug)
	setEnv(map[string]string{
		"ENV":  AppConfig.Env,
		"PORT": AppConfig.Port,
	})
}
