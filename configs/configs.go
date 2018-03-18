package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

type Configs struct {
}

func init() {
	viper.SetConfigName("base")
	viper.AddConfigPath("./configs")
	viper.WatchConfig()
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	fmt.Println(viper.Get("environment"))
}
