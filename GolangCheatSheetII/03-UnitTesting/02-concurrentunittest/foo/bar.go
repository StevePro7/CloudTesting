package foo

var (
	trackChan   = make(chan string)
	untrackChan = make(chan string)
	doneFunc    = func() {}
)

func Track(hs ...string) {
	for _, h := range hs {
		trackChan <- h
	}
}

func Untrack(hs ...string) {
	for _, h := range hs {
		untrackChan <- h
	}
}

var hashtags []string

func init() {
	go func() {
		for {
			select {
			case h := <-trackChan:
				if !sliceContains(hashtags, h) {
					hashtags = append(hashtags, h)
					doneFunc()
				}
			case h := <-untrackChan:
				if sliceContains(hashtags, h) {
					hashtags = sliceRemove(hashtags, h)
					doneFunc()
				}
			}
		}
	}()
}

// https://stackoverflow.com/questions/34070369/removing-a-string-from-a-slice-in-go
func sliceContains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// https://stackoverflow.com/questions/34070369/removing-a-string-from-a-slice-in-go
func sliceRemove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}
