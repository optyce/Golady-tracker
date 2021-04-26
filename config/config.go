package config

import "github.com/spf13/viper"

type Config struct {
	Token  string `mapstructure:"token"`
	GithubUsername string `mapstructure:"github_username"`
	GithubTokenPassword string `mapstructure:"github_token_password"`
	GithubAouthToken string `mapstructure:"github_aouth_token"`
}

// Return a config from app.env file
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
