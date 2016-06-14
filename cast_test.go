package gomap

import "testing"

func TestToString(t *testing.T)  {
	if toString([]byte("123")) != "123"{
		t.Errorf("cast bytes error")
	}
	if toString(123) != "123"{
		t.Errorf("cast int error")
	}
	if toString("123") != "123"{
		t.Errorf("cast string error")
	}
	if toString(rune(123)) != "123"{
		t.Errorf("cast rune error")
	}
}
