package foo

var hashtags []string

//func Track(hs ...string) {
//
//	for _, h := range hs {
//		if !sliceContains(hashtags, h) {
//			hashtags = append(hashtags, h)
//		}
//	}
//}

func Untrack(hs ...string) {

	for _, h := range hs {
		if sliceContains(hashtags, h) {
			hashtags = sliceRemove(hashtags, h)
		}
	}
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

func Testing(h string) bool {
	return sliceContains(hashtags, h)
}

func SetHashtags(hs []string) {
	hashtags = hs
}
