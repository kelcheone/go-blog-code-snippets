package main

import (
	"gopkg.in/ini.v1"
)

// Writing INI files
func main() {
	cfg := ini.Empty()
	// add a section
	sec, err := cfg.NewSection("database")
	if err != nil {
		panic(err)
	}
	// add key-value pairs
	_, err = sec.NewKey("host", "localhost")
	if err != nil {
		panic(err)
	}

	// add a sub-section
	sec, err = cfg.NewSection("database.options")
	if err != nil {
		panic(err)
	}

	// add key-value pairs
	_, err = sec.NewKey("sslmode", "disable")
	if err != nil {
		panic(err)
	}

	// save to file
	err = cfg.SaveTo("config2.ini")
	if err != nil {
		panic(err)
	}

}
