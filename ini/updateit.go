// Updating INI file

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

	section.Key("host").SetValue("127.0.0.0")
	section.Key("port").SetValue("3306")

	// save to file
	err = inidata.SaveTo("config.ini")
	if err != nil {
		panic(err)
	}
}
