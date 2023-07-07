package utils

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Configure Configure `yaml:"configure"`
	Policy    Policy    `yaml:"policy"`
}

func getEnv() *viper.Viper {
	v := viper.New()

	v.AddConfigPath(".")
	v.SetConfigName("env")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}

	return v
}

/*
configure
*/
type Configure struct {
	User   string `yaml:"user"`
	Region string `yaml:"region"`
}

func GetAgentConfigure() Configure {
	v := getEnv()
	var config Config

	v.Unmarshal(&config)
	return config.Configure
}

/*
policy
*/
type Policy struct {
	S3  string `yaml:"s3"`
	EC2 string `yaml:"ec2"`
}

func GetPolicyConfigure() Policy {
	v := getEnv()
	var config Config

	v.Unmarshal(&config)
	return config.Policy
}
