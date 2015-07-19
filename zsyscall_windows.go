// MACHINE GENERATED BY 'go generate' COMMAND; DO NOT EDIT

package winapi

import "unsafe"
import "syscall"

var _ unsafe.Pointer

var (
	modkernel32 = syscall.NewLazyDLL("kernel32.dll")

	procGlobalMemoryStatusEx = modkernel32.NewProc("GlobalMemoryStatusEx")
)

func GlobalMemoryStatusEx(buf *MEMORYSTATUSEX) (err error) {
	r1, _, e1 := syscall.Syscall(procGlobalMemoryStatusEx.Addr(), 1, uintptr(unsafe.Pointer(buf)), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}