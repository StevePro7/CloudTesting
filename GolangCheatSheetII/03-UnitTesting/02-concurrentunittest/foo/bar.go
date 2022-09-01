package foo

var (
	trackChan   = make(chan string)
	untrackChan = make(chan string)
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
				//				if !sliceContains(h, hashtags) {
				hashtags = append(hashtags, h)
				//				}
			case h := <-untrackChan:
				//				if sliceContains(h, hashtags) {
				hashtags = sliceRemove(hashtags, h)
				//				}
			}
		}
	}()
}

func Splat() int {
	return 7
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
