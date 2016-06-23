package config

import (
	"flag"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func Init() {
	// load file
	viper.SetConfigName("setting")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Setting file read failed: %s \n", err))
	}

	// Commandline
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
}
