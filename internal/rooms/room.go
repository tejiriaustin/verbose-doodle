package rooms

import "github.com/tejiriaustin/verbose-doodle/internal/model"

var listOfEvents []model.Event

func CreateRoom(key string, participants int, sources int) {
	NewEvent := model.MakeNewEvent(key, participants, sources)
	listOfEvents = append(listOfEvents, NewEvent)

	NewEvent.Mutex.Lock()
	defer NewEvent.Mutex.Unlock()
}


func DeleteRoom(key string) {
	for i, j := range listOfEvents {
		if j.
		listOfEvents = append(listOfEvents[:1], listOfEvents[2:]...)
	}

}
