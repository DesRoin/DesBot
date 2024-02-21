package main

import (
	"fmt"
	"log"
	"context"
	"github.com/mattn/go-mastodon"
	"os"
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
   post, err := LoadPost()
  if err != nil {
    log.Fatal(err)
  }

  var media mastodon.Media
  file, err := os.Open("./Posts/" + post.Asset)
  if err != nil {
    log.Fatal(err)
  }

  media = mastodon.Media{File: file, Description: post.AssetAlt}
  attachment, err := c.UploadMediaFromMedia(context.Background(), &media)
  var attachmentIDs []mastodon.ID

  attachmentIDs = append(attachmentIDs, attachment.ID)

  finalText := post.Text + "\n"

  for i := 0; i < len(post.Tags); i++ {
    finalText = finalText + "#" + post.Tags[i] + " "
  }

  finalText = finalText + "\n" + "Credits: " + post.Credits

  toot := mastodon.Toot{
    Status:   finalText,
    MediaIDs: attachmentIDs,
  }

  fmt.Printf("About to publish: %#v\n", toot)
  status,err := c.PostStatus(context.Background(), &toot)
  _ = status

  if err != nil {
    log.Fatalf("%#v\n", err)
  }
 
}
