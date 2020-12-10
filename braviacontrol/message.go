package braviacontrol

import (
	"fmt"
	"regexp"
	"time"
)

var VALID_MESSAGE_REGEX = regexp.MustCompile("\\*S[CEAN].{20}\\n$")
var CONTROL_MESSAGE_REGEX = regexp.MustCompile("\\*SC.{20}\\n$")
var ENQUIRY_MESSAGE_REGEX = regexp.MustCompile("\\*SE.{20}\\n$")
var ANSWER_MESSAGE_REGEX = regexp.MustCompile("\\*SA.{20}\\n$")
var NOTIFY_MESSAGE_REGEX = regexp.MustCompile("\\*SN.{20}\\n$")

var ErrorAnswerRegex = regexp.MustCompile("\\*SA.{4}F{16}\\n$")
var SuccessAnswerRegex = regexp.MustCompile("\\*SA.{4}0{16}\\n$")
var NotFoundAnswerRegex = regexp.MustCompile("\\*SA.{4}N{16}\\n$")

// Control represents a message sent to the display to issue a command or make an enquiry
type Control struct {
	messageType string // C for Command or E for enquiry
	fourCC      string // 4-byte command code
	parameter   string // 16-byte parameter string
}

// GetRawMessage returns the raw string that can be sent to the display, including a newline at the end
//
// Eg. "*SCPOWR0000000000000001\n"
func (c *Control) GetRawMessage() string {
	return fmt.Sprintf("*S%s%s%s\n", c.messageType, c.fourCC, c.parameter)
}

// Answer represents a message sent from the display in response to a Control message
type Answer struct {
	rawContent string
	timestamp  time.Time
}

// IsError returns if the Answer sent from the display indicates an error condition
func (a Answer) IsError() bool {
	return ErrorAnswerRegex.MatchString(a.rawContent)
}

// IsNotFound returns if the Answer sent from the display indicates a Not Found condition
func (a Answer) IsNotFound() bool {
	return NotFoundAnswerRegex.MatchString(a.rawContent)
}

// GetRawMessage returns the raw string that the display sent
func (a *Answer) GetRawMessage() string {
	return a.rawContent
}

// GetTimeStamp returns the timestamp of when the Answer was received from the display
func (a *Answer) GetTimeStamp() time.Time {
	return a.timestamp
}

type PowerStatus string
const (
	POWER_OFF PowerStatus = "0000000000000000"
	POWER_ON  PowerStatus = "0000000000000001"
	ERROR     PowerStatus = "FFFFFFFFFFFFFFFF"
)

type InputSource string
const (
	HDMI             InputSource = "00000001"
	COMPONENT        InputSource = "00000004"
	SCREEN_MIRRORING InputSource = "00000005"
)
