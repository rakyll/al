// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package al

/*
#cgo darwin  LDFLAGS: -framework OpenAL

#include <stdlib.h>
#ifdef GOOS_darwin
#include <OpenAL/al.h>
#endif
*/

import "C"

type Buffer int32

func GenBuffer() Buffer {
	panic("not yet implemented")
}

func (b Buffer) Valid() bool {
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

func (b Buffer) AppendData(format int32, freq int32, data []byte) {
	panic("not yet implemented")
}

type Source int32

func GenSource() Source {
	panic("not yet implemented")
}

func PlaySources(s []Source) {
	panic("not yet implemented")
}

func StopSources(s []Source) {
	panic("not yet implemented")
}

func PauseSources(s []Source) {
	panic("not yet implemented")
}

func RewindSources(s []Source) {
	panic("not yet implemented")
}

func DopplerFactor() float32 {
	panic("not yet implemented")
}

func DopplerVelocity() float32 {
	panic("not yet implemented")
}

func SpeedOfSound() float32 {
	panic("not yet implemented")
}

func DistanceModel() int32 {
	panic("not yet implemented")
}
