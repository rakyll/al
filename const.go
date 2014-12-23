// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package al

const (
	DISTANCE_MODEL            = 0xD000
	INVERSE_DISTANCE          = 0xD001
	INVERSE_DISTANCE_CLAMPED  = 0xD002
	LINEAR_DISTANCE           = 0xD003
	LINEAR_DISTANCE_CLAMPED   = 0xD004
	EXPONENT_DISTANCE         = 0xD005
	EXPONENT_DISTANCE_CLAMPED = 0xD006
)

// Buffer properties

const (
	AL_FREQUENCY = 0x2001
	AL_BITS      = 0x2002
	AL_CHANNELS  = 0x2003
	AL_SIZE      = 0x2004
)
