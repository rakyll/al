package al

/*
#cgo darwin LDFLAGS: -framework OpenAL
#cgo linux LDFLAGS: /openal-soft/libs/armeabi/libopenal.so
#cgo darwin  CFLAGS: -DGOOS_darwin
#cgo linux   CFLAGS: -DGOOS_linux

#ifdef GOOS_darwin
#include <OpenAL/al.h>
#endif

#ifdef GOOS_linux
#include "/openal-soft/jni/OpenAL/include/AL/al.h"
#endif
*/
import "C"

type Buffer int32

type Source int32

func GetError() uint32 {
	return uint32(C.alGetError())
}
