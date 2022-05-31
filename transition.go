package fsm

type transition struct {
	from   string
	to     string
	guards []Guard
	before Callback
	after  Callback
}

type Guard func() bool

type Callback func()

func (t *transition) Apply(fsm *FSM) {

	if t.before != nil {
		t.before()
	}

	fsm.current = t.to

	if t.after != nil {
		t.after()
	}
}

func (t *transition) Guard(fsm *FSM) bool {
	if fsm.Current() != t.from {
		return false
	}

	for _, guard := range t.guards {
		if !guard() {
			return false
		}
	}

	return true
}
