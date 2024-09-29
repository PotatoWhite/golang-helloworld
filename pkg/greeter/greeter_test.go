package greeter

import "testing"

func TestNewGreeter(t *testing.T) {
	name := "Potato"
	g := NewGreeter(name)
	if g.Name != name {
		t.Errorf("NewGreeter() = %v, want %v", g.Name, name)
	}
}
