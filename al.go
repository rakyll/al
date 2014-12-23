package al

/*
#cgo darwin  LDFLAGS: -framework OpenAL
#cgo linux   LDFLAGS: /openal-soft/libs/armeabi/libopenal.so
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

func GenBuffers(n int) []Buffer {
	panic("not yet imeplemented")
}

func DeleteBuffers(b []Buffer) {
	panic("not yet imeplemented")
}

func (b Buffer) Freq() int32 {
	panic("not yet implemented")
}

func (b Buffer) Channels() int32 {
	panic("not yet implemented")
}

func (b Buffer) Size() int32 {
	panic("not yet implemented")
}

func (b Buffer) Bits() int32 {
	panic("not yet implemented")
}

func (b Buffer) CopyData(format int32, data []byte, freq int32) int32 {
	panic("not yet implemented")
}

func (b Buffer) Valid() bool {
	panic("not yet implemented")
}

type Source int32

func Error() uint32 {
	return uint32(C.alGetError())
}
