package controller

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

type Controller struct {
	screenWidth  int
	screenHeight int

	moveLength int
	castLength int

	keyCommand map[string]func(*Controller, bool)
	keyPressed map[string]bool
	keyAnalog  map[string]int
}

func New() Controller {
	width, height := robotgo.GetScreenSize()

	c := Controller{}
	c.screenWidth = int(width)
	c.screenHeight = int(height)

	return c
}

func (c *Controller) update() {
	for k, v := range c.keyPressed {
		fmt.Println(k, v)
	}

	for k, v := range c.keyAnalog {
		fmt.Println(k, v)
	}
}

func (c *Controller) SetKey(key string, pressed bool) {
	c.keyPressed[key] = pressed
	c.keyCommand[key](c, true)
	c.update()
}
