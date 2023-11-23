package main

import (
	"context"

	"github.com/alex-appy-love-story/backend/app"
)

func printNumber(n int) int {
	return n
}

func main() {
	app := app.New(app.LoadConfig())
	app.Start(context.Background())
}
