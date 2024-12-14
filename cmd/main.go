package main

import (
	"leader-follower-pattern/handler"
)

func main() {
	handler := handler.NewHandler()
	handler.Run()
}
