# gobraviacontrol
An Unofficial GoLang API for Sony Bravia Professional Displays

This library is not affiliated with Sony, not endorsed by Sony, and not supported by Sony.  

Please see the [LICENSE file](LICENSE). 

It implements the Simple IP control protocol, as publicly documented [here](https://pro-bravia.sony.net/develop/integrate/ssip/overview/index.html).

Due to the current pandemic, I do not have access to an _actual_ display to test this with.

## Goals
- [X] Send Commands
- [x] Receive Answers to the Commands
- [X] Correctly match Answers to the Commands that caused them
- [ ] Allows routing of Notification messages from the TV
- [ ] Full implementation of the command set


## Commands Implemented (aka to do list)
- [x] setIrccCode
- [x] setPowerStatus
- [ ] getPowerStatus
- [x] togglePowerStatus
- [ ] setAudioVolume
- [ ] getAudioVolume
- [ ] setAudioMute
- [ ] getAudioMute
- [x] setInput
- [ ] getInput
- [ ] setPictureMute
- [ ] getPictureMute
- [ ] togglePictureMute
- [ ] getBroadcastAddress
- [ ] getMacAddress
- [ ] setSceneSetting
- [ ] getSceneSetting

## Example Code

```go
package main

import (

"github.com/jimmale/gobraviacontrol/braviacontrol"
"github.com/jimmale/gobraviacontrol/braviacontrol/inputsource"
"github.com/jimmale/gobraviacontrol/braviacontrol/powerstatus"
"net"
"time"
)


func main(){
    display, _ := braviacontrol.NewDisplay(net.IP("192.168.1.42"), 8080)
    _ = display.SetPowerStatus(powerstatus.POWER_ON) // Turn on the display
    _ = display.SetInput(inputsource.HDMI, 2) // Switch to HDMI 2 input
    _ = display.VolumeUp() // Turn the volume up
    _ = display.VolumeUp() // Turn the volume up again
    time.Sleep(8 * time.Hour)
    _ = display.VolumeDown() // Turn the volume down
    _ = display.VolumeDown() // Turn the volume down again
    display.SetPowerStatus(powerstatus.POWER_OFF) // Turn off the display
    display.Close()
}
```
