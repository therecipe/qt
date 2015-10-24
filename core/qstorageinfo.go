package core

//#include "qstorageinfo.h"
import "C"
import (
	"unsafe"
)

type QStorageInfo struct {
	ptr unsafe.Pointer
}

type QStorageInfoITF interface {
	QStorageInfoPTR() *QStorageInfo
}

func (p *QStorageInfo) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QStorageInfo) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQStorageInfo(ptr QStorageInfoITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStorageInfoPTR().Pointer()
	}
	return nil
}

func QStorageInfoFromPointer(ptr unsafe.Pointer) *QStorageInfo {
	var n = new(QStorageInfo)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStorageInfo) QStorageInfoPTR() *QStorageInfo {
	return ptr
}

func NewQStorageInfo() *QStorageInfo {
	return QStorageInfoFromPointer(unsafe.Pointer(C.QStorageInfo_NewQStorageInfo()))
}

func NewQStorageInfo3(dir QDirITF) *QStorageInfo {
	return QStorageInfoFromPointer(unsafe.Pointer(C.QStorageInfo_NewQStorageInfo3(C.QtObjectPtr(PointerFromQDir(dir)))))
}

func NewQStorageInfo4(other QStorageInfoITF) *QStorageInfo {
	return QStorageInfoFromPointer(unsafe.Pointer(C.QStorageInfo_NewQStorageInfo4(C.QtObjectPtr(PointerFromQStorageInfo(other)))))
}

func NewQStorageInfo2(path string) *QStorageInfo {
	return QStorageInfoFromPointer(unsafe.Pointer(C.QStorageInfo_NewQStorageInfo2(C.CString(path))))
}

func (ptr *QStorageInfo) DisplayName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QStorageInfo_DisplayName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QStorageInfo) IsReadOnly() bool {
	if ptr.Pointer() != nil {
		return C.QStorageInfo_IsReadOnly(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QStorageInfo) IsReady() bool {
	if ptr.Pointer() != nil {
		return C.QStorageInfo_IsReady(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QStorageInfo) IsRoot() bool {
	if ptr.Pointer() != nil {
		return C.QStorageInfo_IsRoot(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QStorageInfo) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QStorageInfo_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QStorageInfo) Name() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QStorageInfo_Name(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QStorageInfo) Refresh() {
	if ptr.Pointer() != nil {
		C.QStorageInfo_Refresh(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QStorageInfo) RootPath() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QStorageInfo_RootPath(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QStorageInfo) SetPath(path string) {
	if ptr.Pointer() != nil {
		C.QStorageInfo_SetPath(C.QtObjectPtr(ptr.Pointer()), C.CString(path))
	}
}

func (ptr *QStorageInfo) Swap(other QStorageInfoITF) {
	if ptr.Pointer() != nil {
		C.QStorageInfo_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQStorageInfo(other)))
	}
}

func (ptr *QStorageInfo) DestroyQStorageInfo() {
	if ptr.Pointer() != nil {
		C.QStorageInfo_DestroyQStorageInfo(C.QtObjectPtr(ptr.Pointer()))
	}
}
