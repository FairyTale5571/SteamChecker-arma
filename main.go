package main

/*
#include <stdlib.h>
#include <stdio.h>
#include <string.h>

#include "extensionCallback.h"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

//export goRVExtensionVersion
func goRVExtensionVersion(output *C.char, outputsize C.size_t) {
	result := C.CString("Steam checker 1.0")
	defer C.free(unsafe.Pointer(result))
	var size = C.strlen(result) + 1
	if size > outputsize {
		size = outputsize
	}
	C.memmove(unsafe.Pointer(output), unsafe.Pointer(result), size)
}

//export goRVExtensionArgs
func goRVExtensionArgs(output *C.char, outputsize C.size_t, input *C.char, argv **C.char, argc C.int) {
	action := C.GoString(input)
	clearArgs := cleanInput(argv, int(argc))

	if f, ok := Cmd2Action[action]; ok {
		p2Arma(output, outputsize, f(clearArgs...))
		return
	}

	temp := fmt.Sprintf("Function: %s nb params: %d params: %s!", C.GoString(input), argc, clearArgs)
	p2Arma(output, outputsize, temp)
}

//export goRVExtension
func goRVExtension(output *C.char, outputsize C.size_t, input *C.char) {
	temp := fmt.Sprintf("Use args type", C.GoString(input))
	p2Arma(output, outputsize, temp)
}

//export goRVExtensionRegisterCallback
func goRVExtensionRegisterCallback(fnc C.extensionCallback) {
	extensionCallbackFnc = fnc
}

func main() {}
