package common

import "testing"

func TestGetValue(t *testing.T) {
	want := "running"
    got := GetValue("state: running", "state")
    if got != want {
        t.Errorf("Value of GetValue() = %s; %s ", got, want)
    }
}
