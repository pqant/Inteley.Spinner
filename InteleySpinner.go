package Inteley_Spinner

import (
	"github.com/Unknwon/log"
	"strings"
	"github.com/pqant/gospinner"
)

var (
	successSymbol         = "✔"
	failureSymbol         = "✖"
	warningSymbol         = "⚠"
	questionMarkRedSymbol = "❓"
	exclamationMarkSymbol = "❗"
	smileySymbol          = "㋡"
	starMarkSymbol        = "★"
	questionMarkSymbol    = "?"
	aliveSymbol           = "෴"
)

func ShowSpin(spinner *gospinner.Spinner,
	desc string, messageType MessageType,
	time int, newLogMode bool,
	logActive bool, colorActive bool) {
	newLineText := "\n"
	detected := strings.Contains(desc, newLineText)
	prefix := "@>>||<<@"

	if !newLogMode {
		if detected {
			desc = strings.Replace(desc, "\n", "", -1)
			desc = prefix + desc
			desc += "\n"
		}
	} else {
		desc = strings.Replace(desc, "\n", "", -1)
		desc = prefix + desc
	}

	if messageType == Info {
		desc = strings.Replace(desc, prefix, "", -1)
		log.Info(desc)
	} else if messageType == Success {
		desc = strings.Replace(desc, prefix, successSymbol+" ", -1)
		log.Info(desc)
	} else if messageType == Error {
		desc = strings.Replace(desc, prefix, failureSymbol+" ", -1)
		log.Error(desc)
	} else if messageType == Warning {
		desc = strings.Replace(desc, prefix, warningSymbol+" ", -1)
		log.Warn(desc)
	} else if messageType == Smiley {
		desc = strings.Replace(desc, prefix, smileySymbol+" ", -1)
		log.Info(desc)
	} else if messageType == Star {
		desc = strings.Replace(desc, prefix, starMarkSymbol+" ", -1)
		log.Info(desc)
	} else if messageType == ExclamationRed {
		desc = strings.Replace(desc, prefix, exclamationMarkSymbol+" ", -1)
		log.Warn(desc)
	} else if messageType == Alive {
		desc = strings.Replace(desc, prefix, aliveSymbol+" ", -1)
		log.Info(desc)
	} else if messageType == QuestionRed {
		desc = strings.Replace(desc, prefix, questionMarkRedSymbol+" ", -1)
		log.Info(desc)
	} else {
		desc = strings.Replace(desc, prefix, "", -1)
		log.Info(desc)
	}
}

type MessageType int

const (
	Success        MessageType = 1 + iota
	Error
	Warning
	Info
	QuestionRed
	Smiley
	Star
	ExclamationRed
	Alive
)

