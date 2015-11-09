package multimedia

//#include "qmediaservicecamerainfointerface.h"
import "C"
import (
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
	return n
}

func (ptr *QMediaServiceCameraInfoInterface) QMediaServiceCameraInfoInterface_PTR() *QMediaServiceCameraInfoInterface {
	return ptr
}

func (ptr *QMediaServiceCameraInfoInterface) CameraOrientation(device core.QByteArray_ITF) int {
	if ptr.Pointer() != nil {
		return int(C.QMediaServiceCameraInfoInterface_CameraOrientation(ptr.Pointer(), core.PointerFromQByteArray(device)))
	}
	return 0
}

func (ptr *QMediaServiceCameraInfoInterface) CameraPosition(device core.QByteArray_ITF) QCamera__Position {
	if ptr.Pointer() != nil {
		return QCamera__Position(C.QMediaServiceCameraInfoInterface_CameraPosition(ptr.Pointer(), core.PointerFromQByteArray(device)))
	}
	return 0
}

func (ptr *QMediaServiceCameraInfoInterface) DestroyQMediaServiceCameraInfoInterface() {
	if ptr.Pointer() != nil {
		C.QMediaServiceCameraInfoInterface_DestroyQMediaServiceCameraInfoInterface(ptr.Pointer())
	}
}
