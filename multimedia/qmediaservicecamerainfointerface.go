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

type QMediaServiceCameraInfoInterfaceITF interface {
	QMediaServiceCameraInfoInterfacePTR() *QMediaServiceCameraInfoInterface
}

func (p *QMediaServiceCameraInfoInterface) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMediaServiceCameraInfoInterface) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMediaServiceCameraInfoInterface(ptr QMediaServiceCameraInfoInterfaceITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaServiceCameraInfoInterfacePTR().Pointer()
	}
	return nil
}

func QMediaServiceCameraInfoInterfaceFromPointer(ptr unsafe.Pointer) *QMediaServiceCameraInfoInterface {
	var n = new(QMediaServiceCameraInfoInterface)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMediaServiceCameraInfoInterface) QMediaServiceCameraInfoInterfacePTR() *QMediaServiceCameraInfoInterface {
	return ptr
}

func (ptr *QMediaServiceCameraInfoInterface) CameraOrientation(device core.QByteArrayITF) int {
	if ptr.Pointer() != nil {
		return int(C.QMediaServiceCameraInfoInterface_CameraOrientation(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(device))))
	}
	return 0
}

func (ptr *QMediaServiceCameraInfoInterface) CameraPosition(device core.QByteArrayITF) QCamera__Position {
	if ptr.Pointer() != nil {
		return QCamera__Position(C.QMediaServiceCameraInfoInterface_CameraPosition(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(device))))
	}
	return 0
}

func (ptr *QMediaServiceCameraInfoInterface) DestroyQMediaServiceCameraInfoInterface() {
	if ptr.Pointer() != nil {
		C.QMediaServiceCameraInfoInterface_DestroyQMediaServiceCameraInfoInterface(C.QtObjectPtr(ptr.Pointer()))
	}
}
