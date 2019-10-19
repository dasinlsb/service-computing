package main

import (
  "agenda/cmd"
  "agenda/model/user"
  "log"
)

func main() {
  if err := user.Load("data"); err != nil {
    log.Fatal(err)
  }
  cmd.Execute()
}
