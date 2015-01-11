// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package al

type Source int32

func GenSources(n int) []Source {
	panic("not yet implemented")
}

func PlaySources(source ...Source) {
	panic("not yet implemented")
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
	panic("not yet implemented")
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
