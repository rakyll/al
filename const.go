// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package al

const (
	kFalse = 0
	kTrue  = 1
)

// Error returns one of these error codes.
const (
	INVALID_NAME      = 0xA001
	INVALID_ENUM      = 0xA002
	INVALID_VALUE     = 0xA003
	INVALID_OPERATION = 0xA004
	OUT_OF_MEMORY     = 0xA005
)

// Distance models.
const (
	DISTANCE_MODEL            = 0xD000
	INVERSE_DISTANCE          = 0xD001
	INVERSE_DISTANCE_CLAMPED  = 0xD002
	LINEAR_DISTANCE           = 0xD003
	LINEAR_DISTANCE_CLAMPED   = 0xD004
	EXPONENT_DISTANCE         = 0xD005
	EXPONENT_DISTANCE_CLAMPED = 0xD006
)

// Buffer formats. Buffer.BufferData accepts one of these formats
// as the data format.
const (
	FORMAT_MONO8    = 0x1100
	FORMAT_MONO16   = 0x1101
	FORMAT_STEREO8  = 0x1102
	FORMAT_STEREO16 = 0x1103
)

// Buffer properties
const (
	pFrequency = 0x2001
	pBits      = 0x2002
	pChannels  = 0x2003
	pSize      = 0x2004
)
