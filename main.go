package main

import (
	"fmt"
	"github.com/senior-buddy/buddy"
)

type DrawingController struct{}

func (c *DrawingController) Draw(req *buddy.Request) {
	// business logic goes here

	// save to database for example
	fmt.Println(req.Params)
	fmt.Println("Hello!")
}

func main() {
	buddy.Add("draw_endpoint", DrawingController{}, "Draw")
	buddy.Run()
}
