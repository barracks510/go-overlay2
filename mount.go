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
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"runtime"
	"syscall"

	"github.com/docker/docker/pkg/reexec"
)

func init() {
	reexec.Register("docker-mountfrom", mountFromMain)
}

func fatal(err error) {
	fmt.Fprint(os.Stderr, err)
	os.Exit(1)
}

type mountOptions struct {
	Device string
	Target string
	Type   string
	Label  string
	Flag   uint32
}

func mountFrom(dir, device, target, mType, label string) error {
	options := &mountOptions{
		Device: device,
		Target: target,
		Type:   mType,
		Flag:   0,
		Label:  label,
	}

	cmd := reexec.Command("docker-mountfrom", dir)
	w, err := cmd.StdinPipe()
	if err != nil {
		return fmt.Errorf("mountfrom error on pipe creation: %v", err)
	}

	output := bytes.NewBuffer(nil)
	cmd.Stdout = output
	cmd.Stderr = output

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("mountfrom error on re-exec cmd: %v", err)
	}
	//write the options to the pipe for the untar exec to read
	if err := json.NewEncoder(w).Encode(options); err != nil {
		return fmt.Errorf("mountfrom json encode to pipe failed: %v", err)
	}
	w.Close()

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("mountfrom re-exec error: %v: output: %s", err, output)
	}
	return nil
}

// mountfromMain is the entry-point for docker-mountfrom on re-exec.
func mountFromMain() {
	runtime.LockOSThread()
	flag.Parse()

	var options *mountOptions

	if err := json.NewDecoder(os.Stdin).Decode(&options); err != nil {
		fatal(err)
	}

	if err := os.Chdir(flag.Arg(0)); err != nil {
		fatal(err)
	}

	if err := syscall.Mount(options.Device, options.Target, options.Type, uintptr(options.Flag), options.Label); err != nil {
		fatal(err)
	}

	os.Exit(0)
}
