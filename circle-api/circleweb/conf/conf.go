package conf

import (
	"github.com/spf13/viper"
)

func init()  {
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("circle")
	viper.AutomaticEnv()
	err:=viper.ReadInConfig()
	if err!=nil {
	panic(err)
	}

}

func GetInt(key string)int  {
	return viper.GetInt(key)
}
func GetString(key string)string  {
	return viper.GetString(key)
}
func GetBool(key string) bool  {
	return viper.GetBool(key)
}
