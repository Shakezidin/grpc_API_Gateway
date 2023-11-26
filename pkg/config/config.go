package config

import "github.com/spf13/viper"

type Configure struct {
	APIPORT       string
	GRPCADMINPORT string
	GRPCUSERPORT  string
	SECRETKEY     string
}

func LoadConfigure() (*Configure, error) {
	var cnfg Configure

	viper.SetConfigFile("../../.env")
	err := viper.ReadInConfig()

	err = viper.Unmarshal(&cnfg)

	if err != nil {
		return &Configure{}, nil
	}

	return &cnfg, nil
}
