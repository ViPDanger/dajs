package buisness

import (
	"DAJ/Server/internal/models"
)

func Result(command string, Event models.Event) models.Result {
	var result models.Result
	switch command {
	case "create_event":
		result.Event = Event
		result.Action = "Create Event!"
	case "update_event":
		result.Event = Event
		result.Action = "Update Event!"
	case "events_for_day":
		result.Event = Event
		result.Action = "Events for Day!"
	case "events_for_week":
		result.Event = Event
		result.Action = "Events for Week!"
	case "events_for_mont":
		result.Event = Event
		result.Action = "Events for Month!"
	default:
	}
	return result
}
