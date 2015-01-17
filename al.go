// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package al

// docker run -v $GOPATH/src:/src mobile /bin/bash -c 'cd /src/github.com/rakyll/al && CGO_ENABLED=1 GOARCH=arm GOARM=7  GOOS=android  go build  .'

/*
#cgo darwin   CFLAGS: -DGOOS_darwin
#cgo android  CFLAGS: -DGOOS_android -DANDROID -DAL_SOURCE_BUILD

#cgo android  CFLAGS: -DMAX_SOURCES_LOW=4 -DMAX_SOURCES_START=8 -DMAX_SOURCES_HIGH=64
#cgo android  CFLAGS: -marm -DOPENAL_FIXED_POINT -DOPENAL_FIXED_POINT_SHIFT=16 -DPOST_FROYO -DAL_ALEXT_PROTOTYPES -fpic -ffunction-sections -DANDROID -funwind-tables -fstack-protector -DHAVE_GCC_VISIBILITY -O3 -fno-short-enums -g
#cgo android  LDFLAGS: -Wl,--build-id -Bsymbolic -lm -llog -ldl

#ifdef GOOS_android
#include <stdlib.h>
#include "al.h"
#endif

#cgo darwin  LDFLAGS: -framework OpenAL

#ifdef GOOS_darwin
#include <stdlib.h>
#include <OpenAL/al.h>
#include <OpenAL/alc.h>
#endif

ALCdevice *castAlcOpenDevice(const char *devicename) {
	return alcOpenDevice(devicename);
}

void castAlSourceQueueBuffers(ALuint sid, ALsizei numEntries, const void *bids) {
	return alSourceQueueBuffers(sid, numEntries, bids);
}

void castAlSourcePlayv( ALsizei ns, const void *sids ) {
	return alSourcePlayv(ns, sids);
}

*/
import "C"
import "unsafe"

func Enable(capability int32) {
	C.alEnable(C.ALenum(capability))
}

func Disable(capability int32) {
	C.alDisable(C.ALenum(capability))
}

func Enabled(capability int32) bool {
	return toBoolAL(C.alIsEnabled(C.ALenum(capability)))
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

func toBoolAL(v C.ALboolean) bool {
	return int32(v) == kTrue
}

func toBoolALC(v C.ALCboolean) bool {
	return int32(v) == kTrue
}

type Source int32

func GenSources(n int) []Source {
	s := make([]C.ALuint, n)
	C.alGenSources(C.ALsizei(n), (*C.ALuint)(unsafe.Pointer(&s[0])))
	r := make([]Source, n)
	for i, v := range s {
		r[i] = Source(v)
	}
	return r
}

func PlaySources(source ...Source) {
	C.castAlSourcePlayv(C.ALsizei(len(source)), unsafe.Pointer(&source[0]))
}

func PauseSources(source ...Source) {
	panic("not yet implemented")
}

func StopSources(source ...Source) {
	panic("not yet implemented")
}

func RewindSources(source ...Source) {
	panic("not yet implemented")
}

func DeleteSources(source ...Source) {
	panic("not yet implemented")
}

func (s Source) Gain() float32 {
	panic("not yet implemented")
}

func (s Source) SetGain(v float32) {
	panic("not yet implemented")
}

func (s Source) MaxGain() float32 {
	panic("not yet implemented")
}

func (s Source) SetMaxGain(v float32) {
	panic("not yet implemented")
}

func (s Source) MinGain() float32 {
	panic("not yet implemented")
}

func (s Source) SetMinGain(v float32) {
	panic("not yet implemented")
}

func (s Source) Position() Vector {
	panic("not yet implemented")
}

func (s Source) SetPosition(v Vector) {
	panic("not yet implemented")
}

func (s Source) QueueBuffers(buffer ...Buffer) {
	n := len(buffer)
	C.castAlSourceQueueBuffers(C.ALuint(s), C.ALsizei(n), unsafe.Pointer(&buffer[0]))
}

// TODO(jbd): Add SetPosition.
//
/*
 * LISTENER
 * Listener represents the location and orientation of the
 * 'user' in 3D-space.
 *
 * Properties include: -
 *
 * Gain         AL_GAIN         ALfloat
 * Position     AL_POSITION     ALfloat[3]
 * Velocity     AL_VELOCITY     ALfloat[3]
 * Orientation  AL_ORIENTATION  ALfloat[6] (Forward then Up vectors)
 */

type Vector struct {
	X, Y, Z float32
}

type Orientation struct {
	Forward Vector
	Up      Vector
}

type Listener struct{}

func (l Listener) Gain() float32 {
	panic("not yet implemented")
}

func (l Listener) Position() Vector {
	panic("not yet implemented")
}

func (l Listener) Velocity() Vector {
	panic("not yet implemented")
}

func (l Listener) Orientation() Orientation {
	panic("not yet implemented")
}

func (l Listener) SetGain(v float32) {
	panic("not yet implemented")
}

func (l Listener) SetPosition(v Vector) {
	panic("not yet implemented")
}

func (l Listener) SetVelocity(v Vector) {
	panic("not yet implemented")
}

func (l Listener) SetOrientation(o Orientation) {
	panic("not yet implemented")
}

type Buffer int32

func GenBuffers(n int) []Buffer {
	s := make([]C.ALuint, n)
	C.alGenBuffers(C.ALsizei(n), (*C.ALuint)(unsafe.Pointer(&s[0])))
	r := make([]Buffer, n)
	for i, v := range s {
		r[i] = Buffer(v)
	}
	return r
}

func DeleteBuffers(buffer ...Buffer) {
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

func (b Buffer) BufferData(format int32, data []int32, freq int32) {
	C.alBufferData(C.ALuint(b), C.ALenum(format), unsafe.Pointer(&data[0]), C.ALsizei(len(data)), C.ALsizei(freq))
}

func (b Buffer) Valid() bool {
	panic("not yet implemented")
}

// ------------------------ ALC (Move to the alc package)

type Device struct {
	d *C.ALCdevice
}

type Context struct {
	c *C.ALCcontext
}

func OpenDevice(name string) *Device {
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	d := C.castAlcOpenDevice(n)
	return &Device{d: d}
}

func CreateContext(d *Device, attrs []int32) *Context {
	// TODO: handle attributes
	c := C.alcCreateContext(d.d, nil)
	return &Context{c: c}
}

func MakeContextCurrent(c *Context) bool {
	return toBoolALC(C.alcMakeContextCurrent(c.c))
}

func AlcError(d *Device) uint32 {
	return uint32(C.alcGetError(d.d))
}
