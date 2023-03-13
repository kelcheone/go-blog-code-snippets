// The aim is to use TextMarshal interfaces to marshal a custom struct to a TOML file.

/*
[song]
name = "The Sound of Silence"
duration = 8m20s
*/

package main

import (
	"os"
	"time"

	"github.com/BurntSushi/toml"
)

type Song struct {
	Name     string
	Duration duration
}

type duration time.Duration

func (d duration) MarshalText() ([]byte, error) {
	return []byte(time.Duration(d).String()), nil
}

func (d *duration) UnmarshalText(text []byte) error {
	dur, err := time.ParseDuration(string(text))
	if err != nil {
		return err
	}
	*d = duration(dur)
	return nil
}

func main() {
	song := Song{
		Name:     "The Sound of Silence",
		Duration: duration(8 * time.Minute),
	}

	f, err := os.Create("song.toml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = toml.NewEncoder(f).Encode(song)
	if err != nil {
		panic(err)
	}

	// decode
	var song2 Song
	_, err = toml.DecodeFile("song.toml", &song2)
	if err != nil {
		panic(err)
	}

	println(song2.Name)
	// convert duration to mm:ss
	println(time.Duration(song2.Duration).String())
}
