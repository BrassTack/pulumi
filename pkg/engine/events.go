// Copyright 2017, Pulumi Corporation.  All rights reserved.

package engine

import (
	"fmt"

	"github.com/pulumi/pulumi/pkg/diag"
)

// Event represents an event generated by the engine during an operation. The underlying
// type for the `Payload` field will differ depending on the value of the `Type` field
type Event struct {
	Type    EventType
	Payload interface{}
}

// EventType is the kind of event being emitted.
type EventType string

const (
	StdoutColorEvent EventType = "stdoutcolor"
	DiagEvent        EventType = "Diag"
)

// DiagEventPayload is the payload for an event with type `diag`
type DiagEventPayload struct {
	Severity diag.Severity
	UseColor bool
	Message  string
}

func stdOutEventWithColor(s fmt.Stringer) Event {
	return Event{Type: StdoutColorEvent, Payload: s.String()}
}

func diagDebugEvent(useColor bool, msg string) Event {
	return Event{
		Type: DiagEvent,
		Payload: DiagEventPayload{
			Severity: diag.Debug,
			UseColor: useColor,
			Message:  msg,
		},
	}
}

func diagInfoEvent(useColor bool, msg string) Event {
	return Event{
		Type: DiagEvent,
		Payload: DiagEventPayload{
			Severity: diag.Info,
			UseColor: useColor,
			Message:  msg,
		},
	}
}

func diagInfoerrEvent(useColor bool, msg string) Event {
	return Event{
		Type: DiagEvent,
		Payload: DiagEventPayload{
			Severity: diag.Infoerr,
			UseColor: useColor,
			Message:  msg,
		},
	}
}

func diagErrorEvent(useColor bool, msg string) Event {
	return Event{
		Type: DiagEvent,
		Payload: DiagEventPayload{
			Severity: diag.Error,
			UseColor: useColor,
			Message:  msg,
		},
	}
}

func diagWarningEvent(useColor bool, msg string) Event {
	return Event{
		Type: DiagEvent,
		Payload: DiagEventPayload{
			Severity: diag.Warning,
			UseColor: useColor,
			Message:  msg,
		},
	}
}
