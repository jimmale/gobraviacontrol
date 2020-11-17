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


type Control struct {
	messageType string // C for Command or E for enquiry
	fourCC string 	   // 4-byte command code
	parameter string   // 16-byte parameter string
}

func (c *Control) getRawMessage() string {
	return fmt.Sprintf("*S%s%s%s\n", c.messageType, c.fourCC, c.parameter)
}

type Answer struct {
	RawContent string
	Timestamp time.Time
}

func (a Answer) isError() bool {
	return ErrorAnswerRegex.MatchString(a.RawContent)
}

func (a Answer) isNotFound() bool {
	return NotFoundAnswerRegex.MatchString(a.RawContent)
}


type PowerStatus string
const (
	POWER_OFF PowerStatus = "0000000000000000"
	POWER_ON PowerStatus = "0000000000000001"
)