package main

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

type Config struct {
	Database struct {
		Host     string `ini:"host"`
		Port     int    `ini:"port"`
		Username string `ini:"username"`
		Password string `ini:"password"`
	} `ini:"database"`
	Options struct {
		SSLMode string `ini:"sslmode"`
	} `ini:"database.options"`
}

func main() {
	inidata, err := ini.Load("config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	var config Config

	err = inidata.MapTo(&config)
	if err != nil {
		fmt.Printf("Fail to map file: %v", err)
		os.Exit(1)
	}

	fmt.Println(config.Database.Host)
	fmt.Println(config.Database.Port)
	fmt.Println(config.Database.Username)

	fmt.Println(config.Options.SSLMode)
}
