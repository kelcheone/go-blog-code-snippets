package main

import (
"fmt"
"github.com/joho/godotenv"
"os"
)

func main() {
  err := godotenv.Load()
  if err != nil{
   fmt.Println("Error loading the .env file")
  }

  api_key := os.Getenv("API_KEY")
  fmt.Println(api_key)
}
