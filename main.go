package main

import (
	"github.com/jimmale/gobraviacontrol/braviacontrol"
	"net"
)

func main() {
	display, _ := braviacontrol.NewDisplay(net.IP("192.168.1.42"), 8080)

	_ = display.SetPowerStatus(braviacontrol.POWER_ON)
	_ = display.TogglePowerStatus()

}
