package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QCameraInfo struct {
	ptr unsafe.Pointer
}

type QCameraInfo_ITF interface {
	QCameraInfo_PTR() *QCameraInfo
}

func (p *QCameraInfo) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QCameraInfo) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQCameraInfo(ptr QCameraInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraInfo_PTR().Pointer()
	}
	return nil
}

func NewQCameraInfoFromPointer(ptr unsafe.Pointer) *QCameraInfo {
	var n = new(QCameraInfo)
	n.SetPointer(ptr)
	return n
}

func (ptr *QCameraInfo) QCameraInfo_PTR() *QCameraInfo {
	return ptr
}

func NewQCameraInfo(name core.QByteArray_ITF) *QCameraInfo {
	defer qt.Recovering("QCameraInfo::QCameraInfo")

	return NewQCameraInfoFromPointer(C.QCameraInfo_NewQCameraInfo(core.PointerFromQByteArray(name)))
}

func NewQCameraInfo2(camera QCamera_ITF) *QCameraInfo {
	defer qt.Recovering("QCameraInfo::QCameraInfo")

	return NewQCameraInfoFromPointer(C.QCameraInfo_NewQCameraInfo2(PointerFromQCamera(camera)))
}

func NewQCameraInfo3(other QCameraInfo_ITF) *QCameraInfo {
	defer qt.Recovering("QCameraInfo::QCameraInfo")

	return NewQCameraInfoFromPointer(C.QCameraInfo_NewQCameraInfo3(PointerFromQCameraInfo(other)))
}

func (ptr *QCameraInfo) Description() string {
	defer qt.Recovering("QCameraInfo::description")

	if ptr.Pointer() != nil {
		return C.GoString(C.QCameraInfo_Description(ptr.Pointer()))
	}
	return ""
}

func (ptr *QCameraInfo) DeviceName() string {
	defer qt.Recovering("QCameraInfo::deviceName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QCameraInfo_DeviceName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QCameraInfo) IsNull() bool {
	defer qt.Recovering("QCameraInfo::isNull")

	if ptr.Pointer() != nil {
		return C.QCameraInfo_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCameraInfo) Orientation() int {
	defer qt.Recovering("QCameraInfo::orientation")

	if ptr.Pointer() != nil {
		return int(C.QCameraInfo_Orientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraInfo) Position() QCamera__Position {
	defer qt.Recovering("QCameraInfo::position")

	if ptr.Pointer() != nil {
		return QCamera__Position(C.QCameraInfo_Position(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraInfo) DestroyQCameraInfo() {
	defer qt.Recovering("QCameraInfo::~QCameraInfo")

	if ptr.Pointer() != nil {
		C.QCameraInfo_DestroyQCameraInfo(ptr.Pointer())
	}
}
