// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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

// TODO(jbd): alc?
//

func Enable(capability int32) {
	panic("not implemented")
}

func Disable(capability int32) {
	panic("not implemented")
}

func Enabled(capability int32) bool {
	panic("not implemented")
}

func DistanceModel() int32 {
	panic("not implemented")
}

func DopplerFactor() float32 {
	panic("not implemented")
}

func DopplerVelocity() float32 {
	panic("not implemented")
}

func SpeedOfSound() float32 {
	panic("not implemented")
}

func SetDopplerFactor(v float32) {
	panic("not implemented")
}

func SetDopplerVelocity(v float32) {
	panic("not implemented")
}

func SetSpeedOfSound(v float32) {
	panic("not implemented")
}

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

func Error() uint32 {
	return uint32(C.alGetError())
}
