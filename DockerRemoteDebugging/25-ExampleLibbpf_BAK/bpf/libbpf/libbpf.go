package libbpf

// #cgo CFLAGS: -I${SRCDIR}/../../bpf-gpl/include/libbpf/src -I${SRCDIR}/../../bpf-gpl
// #cgo LDFLAGS: -L${SRCDIR}/../../bpf-gpl/include/libbpf/src -Wl,-rpath -Wl,${SRCDIR}/../../bpf-gpl/include/libbpf/src -lbpf -lelf -lz
// #include "libbpf_api.h"
import "C"

func Bar() int {
	bar := int(C.Cbar())
	//bar := 28
	return bar
}
