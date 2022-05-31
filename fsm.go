package fsm

type InvalidEventError struct {
	EventName string
	State     string
}

func (e InvalidEventError) Error() string {
	return "event " + e.EventName + " inappropriate in current state " + e.State
}

type UnknownEventError struct {
	EventName string
}

func (e UnknownEventError) Error() string {
	return "event " + e.EventName + " does not exist"
}

type Event struct {
	Name   string
	From   string
	To     string
	Guards []Guard
	Before Callback
	After  Callback
}

type Events []Event

type eKey struct {
	name string
	from string
}

type FSM struct {
	current     string
	transitions map[eKey]transition
}

func NewFSM(initial string, events []Event) *FSM {
	f := &FSM{
		current:     initial,
		transitions: make(map[eKey]transition),
	}
	for _, e := range events {
		if _, ok := f.transitions[eKey{e.Name, e.From}]; ok {
			panic("event name:" + e.Name + " and transition from: " + e.From + " already exist")
		}
		f.transitions[eKey{e.Name, e.From}] = transition{
			from:   e.From,
			to:     e.To,
			guards: e.Guards,
			before: e.Before,
			after:  e.After,
		}

	}
	return f
}

func (f *FSM) Current() string {
	return f.current
}

func (f *FSM) Event(eventName string) error {

	tr, ok := f.transitions[eKey{eventName, f.current}]
	if !ok {
		for ekey := range f.transitions {
			if ekey.name == eventName {
				return InvalidEventError{eventName, f.current}
			}
		}
		return UnknownEventError{eventName}
	}

	if tr.Guard(f) {
		tr.Apply(f)
	}

	return nil
}
