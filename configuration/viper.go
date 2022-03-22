package configuration

import (
	"fmt"

	"github.com/spf13/viper"
)

var Runmode string

func setUpViper() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("env")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
	viper.SetEnvPrefix("global")
	Runmode = viper.GetString("runmode")
}
