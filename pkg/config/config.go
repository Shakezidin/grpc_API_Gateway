package config

import "github.com/spf13/viper"

type Configure struct {
	APIPORT       string `mapstructure:"APIPORT"`
	GRPCADMINPORT string `mapstructure:"GRPCADMINPORT"`
	GRPCUSERPORT  string `mapstructure:"GRPCUSERPORT"`
	SECRETKEY     string `mapstructure:"SECRETKEY"`
}

func LoadConfigure() (*Configure, error) {
	var cnfg Configure

	viper.SetConfigFile("../../.env")
	err := viper.ReadInConfig()

	err = viper.Unmarshal(&cnfg)

	if err != nil {
		return nil, err
	}

	return &cnfg, nil
}
