package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

func main() {
	// Reading config file
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
		return
	}

	students := viper.GetStringSlice("students")
	for i, student := range students {
		log.Println(i, student)
	}

	viper.SetConfigName("another_config")
	viper.SetDefault("project_name", "learning_viper_go")
	viper.SetDefault("list", []string{"Gary", "Harold", "Quora"})

	//err2 := viper.SafeWriteConfig()
	err2 := viper.WriteConfig()
	if err2 != nil {
		log.Fatalln(err2)
		return
	}

	//environment variables
	os.Setenv("TEST_VALUE", "this is a test")

	err3 := viper.BindEnv("test_value")
	if err3 != nil {
		log.Fatalln(err3)
		return
	}
	fmt.Println(viper.GetString("test_value"))
}
