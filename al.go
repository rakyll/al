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

ALCdevice *openDevice(const char *devicename) {
	return alcOpenDevice(devicename);
}

void deleteBuffers(ALsizei n, const void *buffers) {
	return alDeleteBuffers(n, buffers);
}

void sourceQueueBuffers(ALuint sid, ALsizei numEntries, const void *bids) {
	return alSourceQueueBuffers(sid, numEntries, bids);
}

void sourceUnqueueBuffers(ALuint sid, ALsizei numEntries, void *bids) {
	return alSourceUnqueueBuffers(sid, numEntries, bids);
}

void sourcePlayv(ALsizei ns, const void *sids) {
	return alSourcePlayv(ns, sids);
}

void sourcePausev(ALsizei ns, const void *sids) {
	return alSourcePausev(ns, sids);
}

void sourceStopv(ALsizei ns, const void *sids) {
	return alSourceStopv(ns, sids);
}

void sourceRewindv(ALsizei ns, const void *sids) {
	return alSourceRewindv(ns, sids);
}

void deleteSources(ALsizei n, const void *sources) {
	return alDeleteSources(n, sources);
}

void getListenerfv(ALenum param, void* values) {
	return alGetListenerf(param, values);
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

// Global parameters.
const (
	DISTANCE_MODEL            = 0xD000
	INVERSE_DISTANCE          = 0xD001
	INVERSE_DISTANCE_CLAMPED  = 0xD002
	LINEAR_DISTANCE           = 0xD003
	LINEAR_DISTANCE_CLAMPED   = 0xD004
	EXPONENT_DISTANCE         = 0xD005
	EXPONENT_DISTANCE_CLAMPED = 0xD006

	DOPPLER_FACTOR   = 0xC000
	DOPPLER_VELOCITY = 0xC001
	SPEED_OF_SOUND   = 0xC003

	VENDOR     = 0xB001
	VERSION    = 0xB002
	RENDERER   = 0xB003
	EXTENSIONS = 0xB004
)

// DISTANCE_MODEL

func Geti(param int) int32 {
	panic("not implemented")
}

// DOPPLER_FACTOR, DOPPLER_VELOCITY, SPEED_OF_SOUND

func Getf(param int) float32 {
	panic("not implemented")
}

// VENDOR, VERSION, RENDERER, EXTENSIONS

func GetString(param int) string {
	panic("not implemented")
}

// Error returns one of these error codes.
const (
	INVALID_NAME      = 0xA001
	INVALID_ENUM      = 0xA002
	INVALID_VALUE     = 0xA003
	INVALID_OPERATION = 0xA004
	OUT_OF_MEMORY     = 0xA005
)

// Error returns the most recently generated error.
func Error() uint32 {
	return uint32(C.alGetError())
}

type Source int32

func GenSources(n int) []Source {
	s := make([]Source, n)
	C.alGenSources(C.ALsizei(n), (*C.ALuint)(unsafe.Pointer(&s[0])))
	return s
}

func PlaySources(sources []Source) {
	C.sourcePlayv(C.ALsizei(len(sources)), unsafe.Pointer(&sources[0]))
}

func PauseSources(sources []Source) {
	C.sourcePausev(C.ALsizei(len(sources)), unsafe.Pointer(&sources[0]))
}

func StopSources(sources []Source) {
	C.sourceStopv(C.ALsizei(len(sources)), unsafe.Pointer(&sources[0]))
}

func RewindSources(sources []Source) {
	C.sourceRewindv(C.ALsizei(len(sources)), unsafe.Pointer(&sources[0]))
}

func DeleteSources(sources []Source) {
	C.deleteSources(C.ALsizei(len(sources)), unsafe.Pointer(&sources[0]))
}

const (
	GAIN              = 0x100A
	POSITION          = 0x1004
	VELOCITY          = 0x1006
	ORIENTATION       = 0x100F
	MIN_GAIN          = 0x100D
	MAX_GAIN          = 0x100E
	SOURCE_STATE      = 0x1010
	INITIAL           = 0x1011
	PLAYING           = 0x1012
	PAUSED            = 0x1013
	STOPPED           = 0x1014
	BUFFERS_QUEUED    = 0x1015
	BUFFERS_PROCESSED = 0x1016
	SEC_OFFSET        = 0x1024
	SAMPLE_OFFSET     = 0x1025
	BYTE_OFFSET       = 0x1026
)

func QueueBuffers(s Source, buffers []Buffer) {
	C.sourceQueueBuffers(C.ALuint(s), C.ALsizei(len(buffers)), unsafe.Pointer(&buffers[0]))
}

func UnqueueBuffers(s Source, buffers []Buffer) {
	C.sourceUnqueueBuffers(C.ALuint(s), C.ALsizei(len(buffers)), unsafe.Pointer(&buffers[0]))
}

// Listener may have GAIN, POSITION, VELOCITY, ORIENTATION.

func GetListenerf(param int) float32 {
	var v C.ALfloat
	C.alGetListenerf(C.ALenum(param), &v)
	return float32(v)
}

func GetListenerfv(param int, v []float32) {
	C.getListenerfv(C.ALenum(param), unsafe.Pointer(&v[0]))
}

func SetListenerf(param int, v float32) {
	C.alListenerf(C.ALenum(param), C.ALfloat(v))
}

func SetListenerfv(param int, v []float32) {
	panic("not yet implemented")
}

type Buffer int32

func GenBuffers(n int) []Buffer {
	s := make([]Buffer, n)
	C.alGenBuffers(C.ALsizei(n), (*C.ALuint)(unsafe.Pointer(&s[0])))
	return s
}

func DeleteBuffers(buffers []Buffer) {
	C.deleteBuffers(C.ALsizei(len(buffers)), unsafe.Pointer(&buffers[0]))
}

// Bufferf accepts one of these parameters.
const (
	FREQUENCY = 0x2001
	BITS      = 0x2002
	CHANNELS  = 0x2003
	SIZE      = 0x2004
)

// Buffer formats. Buffer.BufferData accepts one of these formats
// as the data format.
const (
	FORMAT_MONO8    = 0x1100
	FORMAT_MONO16   = 0x1101
	FORMAT_STEREO8  = 0x1102
	FORMAT_STEREO16 = 0x1103
)

func GetBufferi(b Buffer, param int) uint32 {
	var v C.ALint
	C.alGetBufferi(C.ALuint(b), C.ALenum(param), &v)
	return uint32(v)
}

func BufferData(b Buffer, format int32, data []byte, freq int32) {
	C.alBufferData(C.ALuint(b), C.ALenum(format), unsafe.Pointer(&data[0]), C.ALsizei(len(data)), C.ALsizei(freq))
}

func IsBuffer(b Buffer) bool {
	return toBoolAL(C.alIsBuffer(C.ALuint(b)))
}

func toBoolAL(v C.ALboolean) bool {
	return int32(v) == kTrue
}

// ------------------------ ALC

type Device struct {
	d *C.ALCdevice
}

func (d *Device) Error() uint32 {
	return uint32(C.alcGetError(d.d))
}

type Context struct {
	c *C.ALCcontext
}

func OpenDevice(name string) *Device {
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	return &Device{d: C.openDevice(n)}
}

func CloseDevice(d *Device) bool {
	return toBoolALC(C.alcCloseDevice(d.d))
}

func CreateContext(d *Device, attrs []int32) *Context {
	// TODO: handle attributes
	c := C.alcCreateContext(d.d, nil)
	return &Context{c: c}
}

func MakeContextCurrent(c *Context) bool {
	return toBoolALC(C.alcMakeContextCurrent(c.c))
}

func toBoolALC(v C.ALCboolean) bool {
	return int32(v) == kTrue
}
