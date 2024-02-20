package main

import (
  "fmt"
  "os"
  "time"

  "github.com/BurntSushi/toml"
)

type Post struct {
  Text       string
  Tags       []string
  Asset      string
  AssetAlt   string
  Credits    string
}

func LoadPost() (Post, error) {
  var post Post
  currentTime := time.Now()
  postFileName := fmt.Sprintf("date-posts/%s/%s.toml", currentTime.Format("2006"), currentTime.Format("2006-01-02"))

  _, fileExistErr := os.Stat("./" + postFileName)

  if fileExistErr == nil {
    fmt.Printf("File exists, processing\n")
  } else {
    fmt.Printf("File does not exist\n")
    return post, fileExistErr
  }

  _, tomlError := toml.DecodeFile(postFileName, &post)
  if tomlError != nil {
    fmt.Printf("TOML file reading/decoding error: %v\n", tomlError)
  }

  return post, nil
}
