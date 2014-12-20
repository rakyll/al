package al

/*
#cgo darwin LDFLAGS: -framework OpenAL
#include <stdlib.h>
#include <OpenAL/al.h>
*/
import "C"

func GetError() uint32 {
	return uint32(C.alGetError())
}
