package foo

import (
	"testing"
)

func TestTrack(t *testing.T) {

	doneChan := make(chan struct{}, 1)
	doneFunc = func() { doneChan <- struct{}{} }
	defer func() {
		doneFunc = func() {}
	}()

	h := "name"

	Track(h)
	<-doneChan

	a := sliceContains(hashtags, h)
	if !a {
		t.Errorf("not contains")
	}
}

func TestTrackOld(t *testing.T) {
	h := "name"
	Track(h)
	a := sliceContains(hashtags, h)
	if !a {
		t.Errorf("not contains")
	}
}

func TestUntrack(t *testing.T) {
	name := "name"
	hashtags = []string{name, "surname"}
	Untrack(name)
	a := sliceContains(hashtags, name)
	if !a {
		t.Errorf("not contains")
	}
}

func TestTrackUntrack(t *testing.T) {
	h := "name"

	Track(h)
	Untrack(h)

	a := sliceContains(hashtags, h)
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
