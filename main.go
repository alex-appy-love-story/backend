package main

import (
  "context"
	"fmt"

  "github.com/alex-appy-love-story/backend/app"
)

func main() {
  app := application.New()

  err := app.Start(context.TODO())
  if err != nil {
    fmt.Println("failed to start app:", err)
  }
}
