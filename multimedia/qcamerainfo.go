package multimedia

//#include "qcamerainfo.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QCameraInfo struct {
	ptr unsafe.Pointer
}

type QCameraInfoITF interface {
	QCameraInfoPTR() *QCameraInfo
}

func (p *QCameraInfo) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QCameraInfo) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQCameraInfo(ptr QCameraInfoITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraInfoPTR().Pointer()
	}
	return nil
}

func QCameraInfoFromPointer(ptr unsafe.Pointer) *QCameraInfo {
	var n = new(QCameraInfo)
	n.SetPointer(ptr)
	return n
}

func (ptr *QCameraInfo) QCameraInfoPTR() *QCameraInfo {
	return ptr
}

func NewQCameraInfo(name core.QByteArrayITF) *QCameraInfo {
	return QCameraInfoFromPointer(unsafe.Pointer(C.QCameraInfo_NewQCameraInfo(C.QtObjectPtr(core.PointerFromQByteArray(name)))))
}

func NewQCameraInfo2(camera QCameraITF) *QCameraInfo {
	return QCameraInfoFromPointer(unsafe.Pointer(C.QCameraInfo_NewQCameraInfo2(C.QtObjectPtr(PointerFromQCamera(camera)))))
}

func NewQCameraInfo3(other QCameraInfoITF) *QCameraInfo {
	return QCameraInfoFromPointer(unsafe.Pointer(C.QCameraInfo_NewQCameraInfo3(C.QtObjectPtr(PointerFromQCameraInfo(other)))))
}

func (ptr *QCameraInfo) Description() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QCameraInfo_Description(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QCameraInfo) DeviceName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QCameraInfo_DeviceName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QCameraInfo) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QCameraInfo_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QCameraInfo) Orientation() int {
	if ptr.Pointer() != nil {
		return int(C.QCameraInfo_Orientation(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCameraInfo) Position() QCamera__Position {
	if ptr.Pointer() != nil {
		return QCamera__Position(C.QCameraInfo_Position(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCameraInfo) DestroyQCameraInfo() {
	if ptr.Pointer() != nil {
		C.QCameraInfo_DestroyQCameraInfo(C.QtObjectPtr(ptr.Pointer()))
	}
}
