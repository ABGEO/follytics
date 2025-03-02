package helper

import "github.com/abgeo/follytics/internal/model"

func CreateUserReferenceEvents(
	user *model.User,
	referenceUsers []*model.User,
	eventType model.EventType,
) []*model.Event {
	events := make([]*model.Event, len(referenceUsers))
	for i, referenceUser := range referenceUsers {
		events[i] = &model.Event{
			Type:          eventType,
			User:          user,
			ReferenceUser: referenceUser,
		}
	}

	return events
}
