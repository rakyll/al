// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package al

/*
#cgo android  CFLAGS: -DGOOS_android -DANDROID
#cgo android  CFLAGS: -DMAX_SOURCES_LOW=4 -DMAX_SOURCES_START=8 -DMAX_SOURCES_HIGH=64
#cgo android  CFLAGS: -marm -DOPENAL_FIXED_POINT -DOPENAL_FIXED_POINT_SHIFT=16 -DPOST_FROYO -DAL_ALEXT_PROTOTYPES -fpic -ffunction-sections -DANDROID -funwind-tables -fstack-protector -DHAVE_GCC_VISIBILITY -O3 -fno-short-enums -g
#cgo android  LDFLAGS: -Wl,--build-id -Bsymbolic -lm -llog -ldl

#ifdef GOOS_android
#include "al.h"
#endif
*/
import "C"

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
