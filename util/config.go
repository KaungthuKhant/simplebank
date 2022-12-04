package util

import "github.com/spf13/viper"

// This Config struct store all configuration of the application
// The values are read by viper from a config file or environment variables.
type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

// user unmarshall (viper) to specify the name of each config fields
func LoadConfig(path string) (config Config, err error) {
	// tell the viper the location of the config files
	viper.AddConfigPath(path)
	// tell viper for config file with the name app. ours is app.env so it is app
	viper.SetConfigName("app")
	// tell the viper the type of the config file
	viper.SetConfigType("env")

	// read value from environment variable
	// overwrite the values of the config file with the value of the corresponding environment variable if they exist
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
