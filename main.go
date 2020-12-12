package main

import (
	"github.com/jimmale/gobraviacontrol/braviacontrol"
	"github.com/jimmale/gobraviacontrol/braviacontrol/powerstatus"
	"net"
	"time"
)

func main() {
	display, _ := braviacontrol.NewDisplay(net.IP("192.168.1.42"), 8080)

	_ = display.SetPowerStatus(powerstatus.POWER_ON)

	time.Sleep(30 * time.Minute)

	_ = display.SetPowerStatus(powerstatus.POWER_OFF)
	display.Close()

}
