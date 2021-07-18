package main

/*
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
*/
import "C"
import (
	"fmt"
	"regexp"
	"unicode"
	"unsafe"
)

func cleanInput(argv **C.char, argc int) []string {
	newArgs := make([]string, argc)
	offset := unsafe.Sizeof(uintptr(0))
	i := 0
	for i < argc {
		_arg := (**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + offset*uintptr(i)))
		arg := C.GoString(*_arg)
		arg = arg[1:len(arg)-1]

		reArg := regexp.MustCompile(`""`)
		arg = reArg.ReplaceAllString(arg, `"`)

		newArgs[i] = arg
		i += 1
	}

	return newArgs
}

func p2Arma(output *C.char, outputsize C.size_t, input string) {
	result := C.CString(input)
	defer C.free(unsafe.Pointer(result))
	var size = C.strlen(result) + 1
	if size > outputsize {
		size = outputsize
	}
	C.memmove(unsafe.Pointer(output), unsafe.Pointer(result), size)
}

func cleanWmci(val string) string {
	normalVal := ""
	for _, ch := range val {
		if !unicode.IsLetter(ch) && !unicode.IsDigit(ch) && !unicode.IsPunct(ch) {
			continue
		}
		normalVal = fmt.Sprint(normalVal, string(ch))
	}
	return normalVal
}