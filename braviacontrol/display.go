package braviacontrol

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/jimmale/gobraviacontrol/braviacontrol/inputsource"
	"github.com/jimmale/gobraviacontrol/braviacontrol/ircccodes"
	"github.com/jimmale/gobraviacontrol/braviacontrol/powerstatus"
	"net"
	"strconv"
	"sync"
	"time"
)

// This string is used as the command parameter when making an enquiry.
const enquiry string = "################"

type Display struct {

	// We lock access to the display when sending a Control message so that we can match the next
	// Answer message that comes back from the display with that Control message.
	lock sync.Mutex

	// Socket for communication with the display
	connection net.Conn

	// A channel that receives answers from the display
	answers chan *Answer

	// A channel that sends controlMessages to the display
	controlMessages chan *Control

	isClosed bool
}

// NewDisplay initializes a connection to a display over a TCP socket
func NewDisplay(ip net.IP, port uint) (*Display, error) {
	connString := fmt.Sprintf("%s:%d", ip.String(), port)
	conn, err := net.Dial("tcp", connString)
	if err != nil {
		return nil, err
	}

	answerChannel := make(chan *Answer)
	controlChannel := make(chan *Control)

	output := Display{
		lock:            sync.Mutex{},
		connection:      conn,
		answers:         answerChannel,
		controlMessages: controlChannel,
	}

	go output.routeMessagesFromDisplay()
	go output.sendMessagesToDisplay()

	return &output, nil
}

// Close closes communication with the display.
func (d *Display) Close() {
	d.lock.Lock()
	defer d.lock.Unlock()
	_ = d.connection.Close()
	close(d.controlMessages)
	close(d.answers)
	d.isClosed = true
}

// This is intended to be run in a separate goroutine.
// Eventually it will route Answer messages and answer messages to their appropriate destinations
func (d *Display) routeMessagesFromDisplay() {

	// bufio.Scanner to buffer reads from the socket, and split reads on newline chars
	scanner := bufio.NewScanner(d.connection)

	for scanner.Scan() {
		rawmessage := scanner.Text() + "\n" // we need to add the newline back on
		timestamp := time.Now()

		if ANSWER_MESSAGE_REGEX.MatchString(rawmessage) {
			var a = Answer{
				rawContent: rawmessage,
				timestamp:  timestamp,
			}
			d.answers <- &a
		}
	}
}

// This is intended to be run in a separate goroutine.
func (d *Display) sendMessagesToDisplay() {
	for {
		controlMessage := <-d.controlMessages
		_, _ = d.connection.Write([]byte(controlMessage.GetRawMessage()))
	}
}

// This is a convenience wrapper that sends a Control message and gets its matching Answer
func (d *Display) sendControlMessage(message *Control) (*Answer, error) {
	d.lock.Lock()
	defer d.lock.Unlock()
	if d.isClosed {
		err := errors.New("the connection to the display has been closed")
		return nil, err
	}
	d.controlMessages <- message
	answer := <-d.answers
	return answer, nil
}

// ╔═════════════════════════════════════════════════════════════════════════╗
// ║                     Commands provided by the display                    ║
// ╚═════════════════════════════════════════════════════════════════════════╝

func (d *Display) SetIrccCode(code ircccodes.IRCommand) error {
	zeroPaddedCode := fmt.Sprintf("%016d", code)
	c := Control{
		messageType: "C",
		fourCC:      "IRCC",
		parameter:   zeroPaddedCode,
	}
	ans, err := d.sendControlMessage(&c)
	if err != nil {
		return err
	}
	if ans.IsError() {
		return errors.New("the display returned an error")
	}
	return nil
}

// SetPowerStatus can be used to turn the display on or off (aka standby mode)
func (d *Display) SetPowerStatus(status powerstatus.PowerStatus) error {
	c := Control{
		messageType: "C",
		fourCC:      "POWR",
		parameter:   string(status),
	}
	ans, err := d.sendControlMessage(&c)
	if err != nil {
		return err
	}
	if ans.IsError() {
		return errors.New("the display returned an error")
	}
	return nil
}

func (d *Display) GetPowerStatus() (powerstatus.PowerStatus, error) {
	c := Control{
		messageType: "E",
		fourCC:      "POWR",
		parameter:   enquiry,
	}

	ans, err := d.sendControlMessage(&c)
	if err != nil {
		return powerstatus.ERROR, err
	}
	if ans.IsError() {
		return powerstatus.ERROR, errors.New("the display returned an error")
	}
	if ans.GetParameter() == string(powerstatus.POWER_ON) {
		return powerstatus.POWER_ON, nil
	}
	if ans.GetParameter() == string(powerstatus.POWER_OFF) {
		return powerstatus.POWER_OFF, nil
	}
	return powerstatus.ERROR, errors.New("the display returned a malformed response")
}

func (d *Display) TogglePowerStatus() error {
	c := Control{
		messageType: "C",
		fourCC:      "TPOW",
		parameter:   "################",
	}

	ans, err := d.sendControlMessage(&c)
	if err != nil {
		return err
	}
	if ans.IsError() {
		return errors.New("the display returned an error")
	}
	return nil
}

func (d *Display) SetAudioVolume(volume uint) error {
	zeroPaddedVolumeLevel := fmt.Sprintf("%016d", volume)

	c := Control{
		messageType: "C",
		fourCC:      "VOLU",
		parameter:   zeroPaddedVolumeLevel,
	}

	ans, err := d.sendControlMessage(&c)
	if err != nil {
		return err
	}

	if ans.IsError() {
		return errors.New("the display returned an error")
	}
	return nil
}

func (d *Display) GetAudioVolume() (uint, error) {
	c := Control{
		messageType: "C",
		fourCC:      "VOLU",
		parameter:   enquiry,
	}

	ans, err := d.sendControlMessage(&c)
	if err != nil {
		return 0, err
	}

	if ans.IsError() {
		return 0, errors.New("the display returned an error")
	}

	result, err := strconv.Atoi(ans.GetParameter())
	return uint(result), err
}

func (d *Display) SetAudioMute(mute bool) error {

	var parameter string
	if mute {
		parameter = "0000000000000001"
	} else {
		parameter = "0000000000000000"
	}

	c := Control{
		messageType: "C",
		fourCC:      "AMUT",
		parameter:   parameter,
	}
	ans, err := d.sendControlMessage(&c)
	if err != nil {
		return err
	}

	if ans.IsError() {
		return errors.New("the display returned an error")
	}
	return nil
}

func (d *Display) GetAudioMute() (bool, error) {
	c := Control{
		messageType: "E",
		fourCC:      "AMUT",
		parameter:   "################",
	}
	ans, err := d.sendControlMessage(&c)
	if err != nil {
		return false, err
	}
	if ans.IsError() {
		return false, errors.New("the display returned an error")
	}

	if ans.GetParameter() == "0000000000000001" {
		return true, nil
	}
	if ans.GetParameter() == "0000000000000000" {
		return false, nil
	}
	return false, nil
}

func (d *Display) SetInput(source inputsource.InputSource, number uint) error {

	parameter := fmt.Sprintf("%s%06d", source, number)

	c := Control{
		messageType: "C",
		fourCC:      "INPT",
		parameter:   parameter,
	}
	ans, err := d.sendControlMessage(&c)
	if err != nil {
		return err
	}
	if ans.IsError() {
		return errors.New("the display returned an error")
	}
	return nil
}

// ╔═════════════════════════════════════════════════════════════════════════╗
// ║                        Convenience Wrapper Methods                      ║
// ╚═════════════════════════════════════════════════════════════════════════╝

func (d *Display) VolumeUp() error {
	// TODO implement this
	return nil
}

func (d *Display) VolumeDown() error {
	// TODO implement this
	return nil
}
