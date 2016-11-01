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

	keyCommand map[string]func(c *Controller, current interface{}, previous interface{})
	keyPressed map[string]bool
	keyAnalog  map[string]int
}

func New() Controller {
	width, height := robotgo.GetScreenSize()

	c := Controller{}
	c.screenWidth = int(width)
	c.screenHeight = int(height)

	c.keyCommand = make(map[string]func(c *Controller, current interface{}, previous interface{}))
	c.keyPressed = make(map[string]bool)
	c.keyAnalog = make(map[string]int)

	return c
}

// func (c *Controller) Update() {
// 	c.update()
// }

func (c *Controller) update() {
	for k, v := range c.keyPressed {
		fmt.Println(k, v)
	}

	for k, v := range c.keyAnalog {
		fmt.Println(k, v)
	}
}

func (c *Controller) AddKey(keyName string, keyFunc func(c *Controller, current interface{}, previous interface{})) {
	c.keyCommand[keyName] = keyFunc
}

func (c *Controller) SetKey(key string, pressed bool) {
	c.keyCommand[key](c, pressed, c.keyPressed[key])
	c.keyPressed[key] = pressed
	c.update()
}

func (c *Controller) SetAnalog(key string, value int) {
	c.keyCommand[key](c, value, c.keyAnalog[key])
	c.keyAnalog[key] = value
	c.update()
}
