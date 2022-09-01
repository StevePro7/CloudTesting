package foo

import (
	"testing"
)

func TestTrack(t *testing.T) {
	h := "name"
	Track(h)
	a := contains(hashtags, h)
	if !a {
		t.Errorf("not contains")
	}
}

func TestUntrack(t *testing.T) {
	name := "name"
	hashtags = []string{name, "surname"}
	Untrack(name)
	a := contains(hashtags, name)
	if !a {
		t.Errorf("not contains")
	}
}

func TestTrackUntrack(t *testing.T) {
	h := "name"

	Track(h)
	Untrack(h)

	a := contains(hashtags, h)
	if !a {
		t.Errorf("not contains2")
	}
}

func TestSplat(t *testing.T) {
	x := Splat()
	if 4 != x {
		t.Errorf("not equal")
	}
}

// https://stackoverflow.com/questions/34070369/removing-a-string-from-a-slice-in-go
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
