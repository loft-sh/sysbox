//
// Copyright 2022 Nestybox, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package seccomp

import (
	"fmt"
	"unsafe"

	"C"

	"golang.org/x/sys/unix"
)

// File contains memParser specialization logic to allow interaction with seccomp
// tracee's through a scatter-gather (IOvec) interface. This approach is the default
// one in kernels built with 'CONFIG_CROSS_MEMORY_ATTACH' flag enabled -- the usual
// case in most of the linux distros.

type memParserIOvec struct{}

// ReadSyscallBytesArgs reads data from the tracee's process address space to extract
// arguments utilized by the traced syscall.
func (mp *memParserIOvec) ReadSyscallStringArgs(pid uint32, elems []memParserDataElem) ([]string, error) {
	var result []string

	for _, e := range elems {
		if e.size > 0 {
			dataBuf := make([]byte, e.size)

			if err := mp.readProcessMem(pid, dataBuf, e.addr, e.size); err != nil {
				return nil, err
			}

			data := C.GoString((*C.char)(unsafe.Pointer(&dataBuf[0])))
			result = append(result, data)
		}
	}

	return result, nil
}

// ReadSyscallBytesArgs reads arbitrary byte data from the tracee's process address
// space to extract arguments utilized by the traced syscall.
func (mp *memParserIOvec) ReadSyscallBytesArgs(pid uint32, elems []memParserDataElem) ([]string, error) {
	var result []string

	for _, e := range elems {
		if e.size > 0 {
			dataBuf := make([]byte, e.size)

			if err := mp.readProcessMem(pid, dataBuf, e.addr, e.size); err != nil {
				return nil, err
			}

			data := C.GoStringN((*C.char)(unsafe.Pointer(&dataBuf[0])), C.int(e.size))
			result = append(result, data)
		}
	}

	return result, nil
}

// WriteSyscallBytesArgs writes collected state (i.e. syscall responses) into the
// the tracee's address space.
func (mp *memParserIOvec) WriteSyscallBytesArgs(pid uint32, elems []memParserDataElem) error {

	for _, e := range elems {
		if e.size > 0 {
			if err := mp.writeProcessMem(pid, e.addr, e.data, e.size); err != nil {
				return err
			}
		}
	}

	return nil
}

// readsProcessMem reads size bytes at addr from the mem space of process pid,
// and stores the result in the local byte array. Size must be > 0.
func (t *memParserIOvec) readProcessMem(pid uint32, local []byte, addr uint64, size int) error {
	localIovec := make([]unix.Iovec, 1)
	remoteIovec := make([]unix.RemoteIovec, 1)

	localIovec[0].Base = &local[0]
	localIovec[0].Len = uint64(size)

	remoteIovec[0].Base = uintptr(addr)
	remoteIovec[0].Len = size

	// this denotes the end of the read
	if remoteIovec[0].Base == 0 {
		return nil
	}

	// Read from the traced process' memory
	n, err := unix.ProcessVMReadv(int(pid), localIovec, remoteIovec, 0)

	if err != nil {
		return fmt.Errorf("failed to read from mem of pid %d: %s", pid, err)
	} else if n > size {
		return fmt.Errorf("read more bytes (%d) from mem of pid %d than expected (%d)",
			n, pid, size)
	}

	return nil
}

// writeProcessMem writes size bytes in array data to the given address in the
// mem space of process pid.
func (mp *memParserIOvec) writeProcessMem(pid uint32, addr uint64, data []byte, size int) error {
	data = data[:size]

	localIov := []unix.Iovec{
		{
			Base: &data[0],
			Len:  uint64(size),
		},
	}

	remoteIov := []unix.RemoteIovec{
		{
			Base: uintptr(addr),
			Len:  size,
		},
	}

	// Write to the traced process' memory
	n, err := unix.ProcessVMWritev(int(pid), localIov, remoteIov, 0)

	if err != nil {
		return fmt.Errorf("failed to write to mem of pid %d: %s", pid, err)
	} else if n != size {
		return fmt.Errorf("failed to write %d bytes to mem of pid %d: wrote %d bytes only", size, pid, n)
	}

	return nil
}
