package foo

// #include "foo.h"
import "C"

func Bar() int {
	bar := int(C.Cbar())
	return bar
}
