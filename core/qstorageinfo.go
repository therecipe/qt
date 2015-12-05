package core

//#include "core.h"
import "C"
import (
	"log"
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStorageInfo::QStorageInfo")
		}
	}()

	return NewQStorageInfoFromPointer(C.QStorageInfo_NewQStorageInfo())
}

func NewQStorageInfo3(dir QDir_ITF) *QStorageInfo {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStorageInfo::QStorageInfo")
		}
	}()

	return NewQStorageInfoFromPointer(C.QStorageInfo_NewQStorageInfo3(PointerFromQDir(dir)))
}

func NewQStorageInfo4(other QStorageInfo_ITF) *QStorageInfo {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStorageInfo::QStorageInfo")
		}
	}()

	return NewQStorageInfoFromPointer(C.QStorageInfo_NewQStorageInfo4(PointerFromQStorageInfo(other)))
}

func NewQStorageInfo2(path string) *QStorageInfo {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStorageInfo::QStorageInfo")
		}
	}()

	return NewQStorageInfoFromPointer(C.QStorageInfo_NewQStorageInfo2(C.CString(path)))
}

func (ptr *QStorageInfo) Device() *QByteArray {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStorageInfo::device")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QStorageInfo_Device(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStorageInfo) DisplayName() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStorageInfo::displayName")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QStorageInfo_DisplayName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QStorageInfo) FileSystemType() *QByteArray {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStorageInfo::fileSystemType")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QStorageInfo_FileSystemType(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStorageInfo) IsReadOnly() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStorageInfo::isReadOnly")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QStorageInfo_IsReadOnly(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStorageInfo) IsReady() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStorageInfo::isReady")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QStorageInfo_IsReady(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStorageInfo) IsRoot() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStorageInfo::isRoot")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QStorageInfo_IsRoot(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStorageInfo) IsValid() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStorageInfo::isValid")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QStorageInfo_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStorageInfo) Name() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStorageInfo::name")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QStorageInfo_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QStorageInfo) Refresh() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStorageInfo::refresh")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStorageInfo_Refresh(ptr.Pointer())
	}
}

func (ptr *QStorageInfo) RootPath() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStorageInfo::rootPath")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QStorageInfo_RootPath(ptr.Pointer()))
	}
	return ""
}

func (ptr *QStorageInfo) SetPath(path string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStorageInfo::setPath")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStorageInfo_SetPath(ptr.Pointer(), C.CString(path))
	}
}

func (ptr *QStorageInfo) Swap(other QStorageInfo_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStorageInfo::swap")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStorageInfo_Swap(ptr.Pointer(), PointerFromQStorageInfo(other))
	}
}

func (ptr *QStorageInfo) DestroyQStorageInfo() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStorageInfo::~QStorageInfo")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStorageInfo_DestroyQStorageInfo(ptr.Pointer())
	}
}
