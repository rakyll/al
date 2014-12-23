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

func Enable(capability int32) {
	C.alEnable(C.ALenum(capability))
}

func Disable(capability int32) {
	C.alDisable(C.ALenum(capability))
}

func Enabled(capability int32) bool {
	return toBool(C.alIsEnabled(C.ALenum(capability)))
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

// Error returns the most recently generated error.
func Error() uint32 {
	return uint32(C.alGetError())
}

func toBool(v C.ALboolean) bool {
	return int32(v) == kTrue
}
