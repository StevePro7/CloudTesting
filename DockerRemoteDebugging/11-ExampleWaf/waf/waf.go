package waf

// #include "waf.h"
import "C"

func ProcessModSec() int {

	inter := int(C.ProcessModSec())
	return inter
}
