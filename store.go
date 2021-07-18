package main

/*
#include "extensionCallback.h"
*/
import "C"

var (
	Cmd2Action = map[string] func (args ...string) string{
		"init":           extInit,
		"getBans":         getBans,
		"getProfileName": getProfileName,
		"getA3_Time":     getA3Time,
		"getDateRegistr": getDateRegister,
		"getFriends":	  getFriends,
	}
	Inited = false
	extensionCallbackFnc C.extensionCallback
)
