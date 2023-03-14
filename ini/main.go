package main

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

func main() {
	inidata, err := ini.Load("config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	section := inidata.Section("database")

	fmt.Println(section.Key("host").String())
	fmt.Println(section.Key("port").String())
	fmt.Println(section.Key("username").String())
	fmt.Println(section.Key("password").String())

	// access sub-section
	section = inidata.Section("database.options")
	fmt.Println(section.Key("sslmode").Bool())
}
