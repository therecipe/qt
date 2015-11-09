package core

//#include "qstorageinfo.h"
import "C"
import (
	"unsafe"
)

type QStorageInfo struct {
	ptr unsafe.Pointer
}

type QStorageInfo_ITF interface {
	QStorageInfo_PTR() *QStorageInfo
}

func (p *QStorageInfo) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QStorageInfo) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQStorageInfo(ptr QStorageInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStorageInfo_PTR().Pointer()
	}
	return nil
}

func NewQStorageInfoFromPointer(ptr unsafe.Pointer) *QStorageInfo {
	var n = new(QStorageInfo)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStorageInfo) QStorageInfo_PTR() *QStorageInfo {
	return ptr
}

func NewQStorageInfo() *QStorageInfo {
	return NewQStorageInfoFromPointer(C.QStorageInfo_NewQStorageInfo())
}

func NewQStorageInfo3(dir QDir_ITF) *QStorageInfo {
	return NewQStorageInfoFromPointer(C.QStorageInfo_NewQStorageInfo3(PointerFromQDir(dir)))
}

func NewQStorageInfo4(other QStorageInfo_ITF) *QStorageInfo {
	return NewQStorageInfoFromPointer(C.QStorageInfo_NewQStorageInfo4(PointerFromQStorageInfo(other)))
}

func NewQStorageInfo2(path string) *QStorageInfo {
	return NewQStorageInfoFromPointer(C.QStorageInfo_NewQStorageInfo2(C.CString(path)))
}

func (ptr *QStorageInfo) Device() *QByteArray {
	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QStorageInfo_Device(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStorageInfo) DisplayName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QStorageInfo_DisplayName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QStorageInfo) FileSystemType() *QByteArray {
	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QStorageInfo_FileSystemType(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStorageInfo) IsReadOnly() bool {
	if ptr.Pointer() != nil {
		return C.QStorageInfo_IsReadOnly(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStorageInfo) IsReady() bool {
	if ptr.Pointer() != nil {
		return C.QStorageInfo_IsReady(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStorageInfo) IsRoot() bool {
	if ptr.Pointer() != nil {
		return C.QStorageInfo_IsRoot(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStorageInfo) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QStorageInfo_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStorageInfo) Name() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QStorageInfo_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QStorageInfo) Refresh() {
	if ptr.Pointer() != nil {
		C.QStorageInfo_Refresh(ptr.Pointer())
	}
}

func (ptr *QStorageInfo) RootPath() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QStorageInfo_RootPath(ptr.Pointer()))
	}
	return ""
}

func (ptr *QStorageInfo) SetPath(path string) {
	if ptr.Pointer() != nil {
		C.QStorageInfo_SetPath(ptr.Pointer(), C.CString(path))
	}
}

func (ptr *QStorageInfo) Swap(other QStorageInfo_ITF) {
	if ptr.Pointer() != nil {
		C.QStorageInfo_Swap(ptr.Pointer(), PointerFromQStorageInfo(other))
	}
}

func (ptr *QStorageInfo) DestroyQStorageInfo() {
	if ptr.Pointer() != nil {
		C.QStorageInfo_DestroyQStorageInfo(ptr.Pointer())
	}
}
