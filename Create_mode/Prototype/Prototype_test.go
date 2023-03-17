package Prototype

import "testing"

func TestClone(t *testing.T) {
	t1 := &Type1{
		name: "type1",
	}

	t2 := t1.Clone()

	if t1 == t2 {
		t.Fatal("error! get clone not working")
	}
}
