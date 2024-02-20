package main

import (
	"fmt"
	"log"
)
	

func main() {
  envs, error := GetConfig()

  if error != nil {
    log.Fatalf("Error loading .env or ENV: %v", error)
  }

  fmt.Printf("%v", envs)
}
