package main

import (
	"fmt"
	"strings"

	"github.com/GeertJohan/go.hid"
	"github.com/cyrusroshan/leguendos/controller"
	"github.com/cyrusroshan/leguendos/utils"
)

// need to run with GODEBUG=cgocheck=0

func main() {
	manufacturer := "Logitech"
	timeout := 100
	fmt.Println(utils.ASCIIART)

	var hidController *hid.Device
	keyController := controller.New()
	addBindings(keyController)

	devices, _ := hid.Enumerate(0x0, 0x0)
	for _, device := range devices {
		if strings.Contains(device.Product, manufacturer) {
			var err error
			hidController, err = hid.Open(device.VendorId, device.ProductId, device.SerialNumber)
			if err != nil {
				panic(err)
			}

			break
		}
	}

	if hidController == nil {
		fmt.Println("Controller not found, exiting")
	}

	data := make([]byte, 8)
	hidController.SetReadWriteNonBlocking(true)
	for {
		_, err := hidController.ReadTimeout(data, timeout)
		if err != nil {
			panic(err)
		}

		interpretData(keyController, data)
	}
}

func interpretData(keyController controller.Controller, data []byte) {
	keyController.SetAnalog("LeftX", int(data[0]))
	keyController.SetAnalog("LeftY", int(data[1]))

	keyController.SetAnalog("RightX", int(data[2]))
	keyController.SetAnalog("RightY", int(data[3]))

	remainder := controller.UpdateCompoundedValues(keyController, int(data[4]), []controller.KeyData{
		controller.KeyData{"Y", 128},
		controller.KeyData{"B", 64},
		controller.KeyData{"A", 32},
		controller.KeyData{"X", 16},
	})

	// could make this a cleaner function, but I think the logic is easier to understand like this
	up, down, left, right := false, false, false, false
	switch remainder {
	case 8:
		//do nothing
	case 7:
		up = true
		left = true
	case 6:
		left = true
	case 5:
		left = true
		down = true
	case 4:
		down = true
	case 3:
		down = true
		right = true
	case 2:
		right = true
	case 1:
		right = true
		up = true
	case 0:
		up = true
	default:
	}
	keyController.SetKey("Up", up)
	keyController.SetKey("Down", down)
	keyController.SetKey("Left", left)
	keyController.SetKey("Right", right)

	controller.UpdateCompoundedValues(keyController, int(data[5]), []controller.KeyData{
		controller.KeyData{"RightAnalogButton", 128},
		controller.KeyData{"LeftAnalogButton", 64},
		controller.KeyData{"Start", 32},
		controller.KeyData{"Back", 16},
		controller.KeyData{"RT", 8},
		controller.KeyData{"LT", 4},
		controller.KeyData{"RB", 2},
		controller.KeyData{"LB", 1},
	})
}

func addBindings(keyController controller.Controller) {
	sampleFunc := func(arg1 *controller.Controller, arg2 interface{}, arg3 interface{}) {}
	keyController.AddKey("LeftX", sampleFunc)
	keyController.AddKey("LeftY", sampleFunc)

	keyController.AddKey("RightX", sampleFunc)
	keyController.AddKey("RightY", sampleFunc)

	keyController.AddKey("RightAnalogButton", sampleFunc)
	keyController.AddKey("LeftAnalogButton", sampleFunc)
	keyController.AddKey("Start", sampleFunc)
	keyController.AddKey("Back", sampleFunc)
	keyController.AddKey("RT", sampleFunc)
	keyController.AddKey("LT", sampleFunc)
	keyController.AddKey("RB", sampleFunc)
	keyController.AddKey("LB", sampleFunc)

	keyController.AddKey("Y", sampleFunc)
	keyController.AddKey("B", sampleFunc)
	keyController.AddKey("A", sampleFunc)
	keyController.AddKey("X", sampleFunc)
	keyController.AddKey("Up", sampleFunc)
	keyController.AddKey("Down", sampleFunc)
	keyController.AddKey("Left", sampleFunc)
	keyController.AddKey("Right", sampleFunc)
}
