package fsm

import (
	"testing"
)

func TestTableFSM(t *testing.T) {
	f := NewFSM(
		"idle",
		Events{
			{
				Name: "tick",
				From: "idle",
				To:   "ready",
				// Guards: []Guard{
				// 	func() bool { return true },
				// },
			},
		},
	)
	if err := f.Event("doTrun"); err != nil {
		t.Errorf("doTrun error: %s", err.Error())
	}
	if err := f.Event("tick"); err != nil {
		t.Errorf("error: %s", err.Error())
	}

	t.Log(f.Current())
}
