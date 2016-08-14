// +build linux

// Copyright 2016 Dennis Chen <barracks510@gmail.com>
// Copyright 2013-2016 Docker, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package overlay2

import (
	"os"
	"runtime"
	"syscall"
)

func mountFrom(dir, device, target, mType, label string) error {
	runtime.LockOSThread()

	// We want to store the original directory so we can re-enter after a
	// successful mount. This solves the problem of a process living
	// in the mounted directory when we want to unmount it. We do this
	// without invoking the reexec chain.
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	if err := os.Chdir(dir); err != nil {
		return err
	}
	if err := syscall.Mount(device, target, mType, uintptr(0), label); err != nil {
		return err
	}
	if err := os.Chdir(cwd); err != nil {
		return err
	}

	runtime.UnlockOSThread()

	return nil
}
