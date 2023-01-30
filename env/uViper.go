package main

import (
    "fmt"
//    "log"

    "github.com/spf13/viper"
)

func main() {
    viper.AutomaticEnv()

    api_Key := viper.Get("API_KEY")
    fmt.Println(api_Key)
}
