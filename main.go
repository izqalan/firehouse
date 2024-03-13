/*
Copyright © 2024 Izqalan Nor'Izad <izqalan@gmail.com>
*/
package main

import (
	"fmt"
	"log"

	"github.com/izqalan/firehouse/cmd"
	"github.com/izqalan/firehouse/utils"
	"github.com/spf13/viper"
)

func init() {
	log.Print("Loading configuration from config.json")
	viper.SetConfigType("json")
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	// if file not found, create a new one
	if err != nil {
		fmt.Println("Config file not found, creating a new one")
		viper.WriteConfig()
	}

	utils.NewFirebaseClient()
}

func main() {
	cmd.Execute()

}
