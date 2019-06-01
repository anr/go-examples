// Package cgo is an example of using CGo to wrap a C++ class.
package cgo

//#include <stdlib.h> // For free()
//#include "wrap.h"
import "C"

import (
	"fmt"
	"unsafe"
)

// Cook holds the wrapped C++ cook.
type Cook struct {
	cook C.wcook
}

// NewCook creates a Cook.
func NewCook(name string, shouldFail bool) (*Cook, error) {
	csName := C.CString(name)
	defer C.free(unsafe.Pointer(csName))

	var cFail C.int
	if shouldFail {
		cFail = 1
	}

	var csErr *C.char
	cook := C.create_wcook(csName, cFail, &csErr)
	if csErr != nil {
		defer C.free(unsafe.Pointer(csErr))
		return nil, fmt.Errorf("failed to create app: %v", C.GoString(csErr))
	}

	return &Cook{
		cook: cook,
	}, nil
}

// Greet talks to the cook.
func (c *Cook) Greet(start string) string {
	csStart := C.CString(start)
	defer C.free(unsafe.Pointer(csStart))

	csRes := C.wcook_greet(c.cook, csStart)
	defer C.free(unsafe.Pointer(csRes))

	return C.GoString(csRes)
}

// Close releases the Cook.
func (c *Cook) Close() {
	C.destroy_wcook(c.cook)
}
