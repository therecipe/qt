package positioning

//#include "qgeoareamonitorinfo.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QGeoAreaMonitorInfo struct {
	ptr unsafe.Pointer
}

type QGeoAreaMonitorInfoITF interface {
	QGeoAreaMonitorInfoPTR() *QGeoAreaMonitorInfo
}

func (p *QGeoAreaMonitorInfo) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QGeoAreaMonitorInfo) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQGeoAreaMonitorInfo(ptr QGeoAreaMonitorInfoITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoAreaMonitorInfoPTR().Pointer()
	}
	return nil
}

func QGeoAreaMonitorInfoFromPointer(ptr unsafe.Pointer) *QGeoAreaMonitorInfo {
	var n = new(QGeoAreaMonitorInfo)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGeoAreaMonitorInfo) QGeoAreaMonitorInfoPTR() *QGeoAreaMonitorInfo {
	return ptr
}

func NewQGeoAreaMonitorInfo2(other QGeoAreaMonitorInfoITF) *QGeoAreaMonitorInfo {
	return QGeoAreaMonitorInfoFromPointer(unsafe.Pointer(C.QGeoAreaMonitorInfo_NewQGeoAreaMonitorInfo2(C.QtObjectPtr(PointerFromQGeoAreaMonitorInfo(other)))))
}

func NewQGeoAreaMonitorInfo(name string) *QGeoAreaMonitorInfo {
	return QGeoAreaMonitorInfoFromPointer(unsafe.Pointer(C.QGeoAreaMonitorInfo_NewQGeoAreaMonitorInfo(C.CString(name))))
}

func (ptr *QGeoAreaMonitorInfo) Identifier() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoAreaMonitorInfo_Identifier(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QGeoAreaMonitorInfo) IsPersistent() bool {
	if ptr.Pointer() != nil {
		return C.QGeoAreaMonitorInfo_IsPersistent(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGeoAreaMonitorInfo) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QGeoAreaMonitorInfo_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGeoAreaMonitorInfo) Name() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoAreaMonitorInfo_Name(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QGeoAreaMonitorInfo) SetArea(newShape QGeoShapeITF) {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorInfo_SetArea(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGeoShape(newShape)))
	}
}

func (ptr *QGeoAreaMonitorInfo) SetExpiration(expiry core.QDateTimeITF) {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorInfo_SetExpiration(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQDateTime(expiry)))
	}
}

func (ptr *QGeoAreaMonitorInfo) SetName(name string) {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorInfo_SetName(C.QtObjectPtr(ptr.Pointer()), C.CString(name))
	}
}

func (ptr *QGeoAreaMonitorInfo) SetPersistent(isPersistent bool) {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorInfo_SetPersistent(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(isPersistent)))
	}
}

func (ptr *QGeoAreaMonitorInfo) DestroyQGeoAreaMonitorInfo() {
	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorInfo_DestroyQGeoAreaMonitorInfo(C.QtObjectPtr(ptr.Pointer()))
	}
}
