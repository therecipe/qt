package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QMediaServiceCameraInfoInterface struct {
	ptr unsafe.Pointer
}

type QMediaServiceCameraInfoInterface_ITF interface {
	QMediaServiceCameraInfoInterface_PTR() *QMediaServiceCameraInfoInterface
}

func (p *QMediaServiceCameraInfoInterface) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMediaServiceCameraInfoInterface) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMediaServiceCameraInfoInterface(ptr QMediaServiceCameraInfoInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaServiceCameraInfoInterface_PTR().Pointer()
	}
	return nil
}

func NewQMediaServiceCameraInfoInterfaceFromPointer(ptr unsafe.Pointer) *QMediaServiceCameraInfoInterface {
	var n = new(QMediaServiceCameraInfoInterface)
	n.SetPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QMediaServiceCameraInfoInterface_") {
		n.SetObjectNameAbs("QMediaServiceCameraInfoInterface_" + qt.Identifier())
	}
	return n
}

func (ptr *QMediaServiceCameraInfoInterface) QMediaServiceCameraInfoInterface_PTR() *QMediaServiceCameraInfoInterface {
	return ptr
}

func (ptr *QMediaServiceCameraInfoInterface) CameraOrientation(device core.QByteArray_ITF) int {
	defer qt.Recovering("QMediaServiceCameraInfoInterface::cameraOrientation")

	if ptr.Pointer() != nil {
		return int(C.QMediaServiceCameraInfoInterface_CameraOrientation(ptr.Pointer(), core.PointerFromQByteArray(device)))
	}
	return 0
}

func (ptr *QMediaServiceCameraInfoInterface) CameraPosition(device core.QByteArray_ITF) QCamera__Position {
	defer qt.Recovering("QMediaServiceCameraInfoInterface::cameraPosition")

	if ptr.Pointer() != nil {
		return QCamera__Position(C.QMediaServiceCameraInfoInterface_CameraPosition(ptr.Pointer(), core.PointerFromQByteArray(device)))
	}
	return 0
}

func (ptr *QMediaServiceCameraInfoInterface) DestroyQMediaServiceCameraInfoInterface() {
	defer qt.Recovering("QMediaServiceCameraInfoInterface::~QMediaServiceCameraInfoInterface")

	if ptr.Pointer() != nil {
		C.QMediaServiceCameraInfoInterface_DestroyQMediaServiceCameraInfoInterface(ptr.Pointer())
	}
}

func (ptr *QMediaServiceCameraInfoInterface) ObjectNameAbs() string {
	defer qt.Recovering("QMediaServiceCameraInfoInterface::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaServiceCameraInfoInterface_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMediaServiceCameraInfoInterface) SetObjectNameAbs(name string) {
	defer qt.Recovering("QMediaServiceCameraInfoInterface::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QMediaServiceCameraInfoInterface_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
