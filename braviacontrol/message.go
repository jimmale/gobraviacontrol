package braviacontrol

import (
	"regexp"
	"time"
)

var VALID_MESSAGE_REGEX = regexp.MustCompile("\\*S[CEAN].{20}\\n$")
var CONTROL_MESSAGE_REGEX = regexp.MustCompile("\\*SC.{20}\\n$")
var ENQUIRY_MESSAGE_REGEX = regexp.MustCompile("\\*SE.{20}\\n$")
var ANSWER_MESSAGE_REGEX = regexp.MustCompile("\\*SA.{20}\\n$")
var NOTIFY_MESSAGE_REGEX = regexp.MustCompile("\\*SN.{20}\\n$")

var ErrorAnswerRegex = regexp.MustCompile("\\*SA {4}F{16}\\n$")
var SuccessAnswerRegex = regexp.MustCompile("\\*SA {4}0{16}\\n$")
var NotFoundAnswerRegex = regexp.MustCompile("\\*SA {4}N{16}\\n$")

type Control struct {
	RawContent string
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
