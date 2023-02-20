package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.String())
	fmt.Println(now.Format("01 Jan 06 03:04PM"))
	currentTime := time.Now()

	fmt.Println(currentTime.Year())
	fmt.Println(currentTime.Month())
	fmt.Println(currentTime.Day())
	fmt.Println(currentTime.Hour())
	fmt.Println(currentTime.Minute())
	fmt.Println(currentTime.Second())
	fmt.Println(currentTime.Nanosecond())
	fmt.Println(currentTime.Location())

	formatTime := currentTime.Format("01/02/06 03:04PM")
	fmt.Println(formatTime)
	fmt.Println(currentTime.Format(time.RubyDate))

	stringTime := "2020-04-21 23:12"
	parsedTime, _ := time.Parse("2006-01-02 15:04", stringTime)
	fmt.Println(parsedTime)

	newDate := time.Date(2020, 04, 21, 23, 12, 0, 0, time.Local)
	fmt.Println(newDate)

	newTime := currentTime.Add(24 * time.Hour)
	fmt.Println(newTime)

	diff := newTime.Sub(currentTime)
	fmt.Println(diff.Milliseconds())

	compTime := currentTime.Add(30 * time.Minute)
	fmt.Println(compTime.Before(currentTime))
	fmt.Println(compTime.After(currentTime))
	fmt.Println(compTime.Equal(currentTime))

	fmt.Println(currentTime.Local())

	unixTime := currentTime.UnixMicro()
	fmt.Println(unixTime)
}
