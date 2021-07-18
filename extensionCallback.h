#pragma once

#include <stdlib.h>

typedef int (*extensionCallback)(char const *name, char const *function, char const *data);

static inline int runExtensionCallback(extensionCallback fnc, char const *name, char const *function, char const *data) {
	return fnc(name, function, data);
}