package Singleton

import "testing"

func TestSingleton(t *testing.T) {
	ins1 := GetInstance(1)
	ins2 := GetInstance(2)
	if ins1 != ins2 {
		t.Fatal("instance is not equal")
	}
}
