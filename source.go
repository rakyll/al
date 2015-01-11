// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package al

type Source int32

func GenSources(n int) []Source {
	panic("not yet implemented")
}

func PlaySources(s []Source) {
	panic("not yet implemented")
}

func PauseSources(s []Source) {
	panic("not yet implemented")
}

func StopSources(s []Source) {
	panic("not yet implemented")
}

func RewindSources(s []Source) {
	panic("not yet implemented")
}

func DeleteSources(s []Source) {
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
