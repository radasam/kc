package managers

var EventManager *Event

type Event struct {
	ListenerMap map[string]func()
	FuncQueue   []string
}

func (e *Event) AddEventListener(event string, listener func(), oneTime bool) {
	if oneTime {
		e.ListenerMap[event] = listener
		e.RemoveEventListener(event)
	} else {
		e.ListenerMap[event] = listener
	}

}

func (e *Event) RemoveEventListener(event string) {
	delete(e.ListenerMap, event)
}

func (e *Event) TriggerEvent(event string) {
	e.FuncQueue = append(e.FuncQueue, event)
}

func (e *Event) Run() {
	for _, v := range e.FuncQueue {
		if f, ok := e.ListenerMap[v]; ok {
			f()
		}
	}

	e.FuncQueue = []string{}
}

func NewEventManager() {
	EventManager = &Event{
		ListenerMap: map[string]func(){},
		FuncQueue:   []string{},
	}
}
