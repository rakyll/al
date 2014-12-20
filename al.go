package al

/*
#cgo darwin LDFLAGS: -framework OpenAL
#cgo darwin  CFLAGS: -DGOOS_darwin
#cgo linux   CFLAGS: -DGOOS_linux

#ifdef GOOS_darwin
#include <OpenAL/al.h>
#endif

#ifdef GOOS_linux
#include "OpenAL/al.h"
#endif

*/
import "C"

type Buffer int32

type Source int32

func GetError() uint32 {
	return uint32(C.alGetError())
}
