package configs

import "github.com/spf13/viper"

type Config struct {
	Port         string `mapstructure:"PORT"`
	DbUrl        string `mapstructure:"DB_URL"`
	ParserSrvUrl string `mapstructure:"PARSER_SERVICE_URL"`
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)

	return
}
