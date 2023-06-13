// Copyright 2023 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

/*
	Interface for text-to-speech
*/
package interfaces

import "io"

// Interface for taking text and converting to audio/speech
type Replay interface {
	Start() error
	Read() ([]byte, error)
	Stream(w io.Writer) error
	Mute()
	Unmute()
	Stop() error
}
