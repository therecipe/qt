package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QLockFile struct {
	ptr unsafe.Pointer
}

type QLockFile_ITF interface {
	QLockFile_PTR() *QLockFile
}

func (p *QLockFile) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QLockFile) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQLockFile(ptr QLockFile_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLockFile_PTR().Pointer()
	}
	return nil
}

func NewQLockFileFromPointer(ptr unsafe.Pointer) *QLockFile {
	var n = new(QLockFile)
	n.SetPointer(ptr)
	return n
}

func (ptr *QLockFile) QLockFile_PTR() *QLockFile {
	return ptr
}

//QLockFile::LockError
type QLockFile__LockError int64

const (
	QLockFile__NoError         = QLockFile__LockError(0)
	QLockFile__LockFailedError = QLockFile__LockError(1)
	QLockFile__PermissionError = QLockFile__LockError(2)
	QLockFile__UnknownError    = QLockFile__LockError(3)
)

func NewQLockFile(fileName string) *QLockFile {
	defer qt.Recovering("QLockFile::QLockFile")

	return NewQLockFileFromPointer(C.QLockFile_NewQLockFile(C.CString(fileName)))
}

func (ptr *QLockFile) Error() QLockFile__LockError {
	defer qt.Recovering("QLockFile::error")

	if ptr.Pointer() != nil {
		return QLockFile__LockError(C.QLockFile_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLockFile) IsLocked() bool {
	defer qt.Recovering("QLockFile::isLocked")

	if ptr.Pointer() != nil {
		return C.QLockFile_IsLocked(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLockFile) Lock() bool {
	defer qt.Recovering("QLockFile::lock")

	if ptr.Pointer() != nil {
		return C.QLockFile_Lock(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLockFile) RemoveStaleLockFile() bool {
	defer qt.Recovering("QLockFile::removeStaleLockFile")

	if ptr.Pointer() != nil {
		return C.QLockFile_RemoveStaleLockFile(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLockFile) SetStaleLockTime(staleLockTime int) {
	defer qt.Recovering("QLockFile::setStaleLockTime")

	if ptr.Pointer() != nil {
		C.QLockFile_SetStaleLockTime(ptr.Pointer(), C.int(staleLockTime))
	}
}

func (ptr *QLockFile) StaleLockTime() int {
	defer qt.Recovering("QLockFile::staleLockTime")

	if ptr.Pointer() != nil {
		return int(C.QLockFile_StaleLockTime(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLockFile) TryLock(timeout int) bool {
	defer qt.Recovering("QLockFile::tryLock")

	if ptr.Pointer() != nil {
		return C.QLockFile_TryLock(ptr.Pointer(), C.int(timeout)) != 0
	}
	return false
}

func (ptr *QLockFile) DestroyQLockFile() {
	defer qt.Recovering("QLockFile::~QLockFile")

	if ptr.Pointer() != nil {
		C.QLockFile_DestroyQLockFile(ptr.Pointer())
	}
}

func (ptr *QLockFile) Unlock() {
	defer qt.Recovering("QLockFile::unlock")

	if ptr.Pointer() != nil {
		C.QLockFile_Unlock(ptr.Pointer())
	}
}
