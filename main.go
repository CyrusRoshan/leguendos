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
	fmt.Println(data)
}
