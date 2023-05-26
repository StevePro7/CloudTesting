package libbpf

// #cgo CFLAGS: -I${SRCDIR}/../../bpf-gpl/include/libbpf/src -I${SRCDIR}/../../bpf-gpl
// #cgo amd64 LDFLAGS: -L${SRCDIR}/../../bpf-gpl/include/libbpf/src/amd64 -lbpf -lelf -lz
// #cgo arm64 LDFLAGS: -L${SRCDIR}/../../bpf-gpl/include/libbpf/src/arm64 -lbpf -lelf -lz
// #cgo armv7 LDFLAGS: -L${SRCDIR}/../../bpf-gpl/include/libbpf/src/armv7 -lbpf -lelf -lz
// #include "libbpf_api.h"
import "C"

func Bar() int {
	bar := int(C.Cbar())
	return bar
}
