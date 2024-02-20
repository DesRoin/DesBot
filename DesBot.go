package main

import (
	"fmt"
	"log"
	"context"
	"github.com/mattn/go-mastodon"
)
	

func main() {
  envs, error := GetConfig()

  if error != nil {
    log.Fatalf("Error loading .env or ENV: %v", error)
  }

  fmt.Printf("%v", envs)
  
   c := mastodon.NewClient(&mastodon.Config{
    Server:       envs["MASTODON_SERVER"],
    ClientID:     envs["APP_CLIENT_ID"],
    ClientSecret: envs["APP_CLIENT_SECRET"],
  })
  err := c.Authenticate(context.Background(), envs["APP_USER"], envs["APP_PASSWORD"])
  if err != nil {
    log.Fatal(err)
  }
  timeline, err := c.GetTimelineHome(context.Background(), nil)
  if err != nil {
    log.Fatal(err)
  }
  for i := len(timeline) - 1; i >= 0; i-- {
    fmt.Println(timeline[i])
  }
}
