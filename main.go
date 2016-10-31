package main

import (
	"fmt"

	"github.com/cyrusroshan/leguendos/controller"
	"github.com/cyrusroshan/leguendos/utils"
)

func main() {
	fmt.Println(utils.ASCIIART)

	controller := controller.New()
	fmt.Println(controller)
}
