package foo

import (
	"testing"
)

func TestTrack(t *testing.T) {

	doneChan := make(chan struct{}, 1)
	doneFunc = func() {
		doneChan <- struct{}{}
	}
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

func TestUntrack(t *testing.T) {
	doneChan := make(chan struct{}, 1)
	doneFunc = func() {
		doneChan <- struct{}{}
	}
	defer func() {
		doneFunc = func() {}
	}()

	name := "name"
	hashtags = []string{name, "surname"}

	Untrack(name)
	<-doneChan

	a := sliceContains(hashtags, name)
	if a {
		t.Errorf("not contains")
	}
}

func TestTrackUntrack(t *testing.T) {

	doneChan := make(chan struct{}, 1)
	doneFunc = func() {
		doneChan <- struct{}{}
	}
	defer func() {
		doneFunc = func() {}
	}()

	h := "name"

	Track(h)
	<-doneChan

	Untrack(h)
	<-doneChan

	a := sliceContains(hashtags, h)
	if a {
		t.Errorf("not contains2")
	}
}
