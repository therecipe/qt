package core

//#include "qlockfile.h"
import "C"
import (
	"unsafe"
)

type QLockFile struct {
	ptr unsafe.Pointer
}

type QLockFileITF interface {
	QLockFilePTR() *QLockFile
}

func (p *QLockFile) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QLockFile) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQLockFile(ptr QLockFileITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLockFilePTR().Pointer()
	}
	return nil
}

func QLockFileFromPointer(ptr unsafe.Pointer) *QLockFile {
	var n = new(QLockFile)
	n.SetPointer(ptr)
	return n
}

func (ptr *QLockFile) QLockFilePTR() *QLockFile {
	return ptr
}

//QLockFile::LockError
type QLockFile__LockError int

var (
	QLockFile__NoError         = QLockFile__LockError(0)
	QLockFile__LockFailedError = QLockFile__LockError(1)
	QLockFile__PermissionError = QLockFile__LockError(2)
	QLockFile__UnknownError    = QLockFile__LockError(3)
)

func NewQLockFile(fileName string) *QLockFile {
	return QLockFileFromPointer(unsafe.Pointer(C.QLockFile_NewQLockFile(C.CString(fileName))))
}

func (ptr *QLockFile) Error() QLockFile__LockError {
	if ptr.Pointer() != nil {
		return QLockFile__LockError(C.QLockFile_Error(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLockFile) IsLocked() bool {
	if ptr.Pointer() != nil {
		return C.QLockFile_IsLocked(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QLockFile) Lock() bool {
	if ptr.Pointer() != nil {
		return C.QLockFile_Lock(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QLockFile) RemoveStaleLockFile() bool {
	if ptr.Pointer() != nil {
		return C.QLockFile_RemoveStaleLockFile(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QLockFile) SetStaleLockTime(staleLockTime int) {
	if ptr.Pointer() != nil {
		C.QLockFile_SetStaleLockTime(C.QtObjectPtr(ptr.Pointer()), C.int(staleLockTime))
	}
}

func (ptr *QLockFile) StaleLockTime() int {
	if ptr.Pointer() != nil {
		return int(C.QLockFile_StaleLockTime(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLockFile) TryLock(timeout int) bool {
	if ptr.Pointer() != nil {
		return C.QLockFile_TryLock(C.QtObjectPtr(ptr.Pointer()), C.int(timeout)) != 0
	}
	return false
}

func (ptr *QLockFile) DestroyQLockFile() {
	if ptr.Pointer() != nil {
		C.QLockFile_DestroyQLockFile(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QLockFile) Unlock() {
	if ptr.Pointer() != nil {
		C.QLockFile_Unlock(C.QtObjectPtr(ptr.Pointer()))
	}
}
