package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QMediaServiceSupportedDevicesInterface struct {
	ptr unsafe.Pointer
}

type QMediaServiceSupportedDevicesInterface_ITF interface {
	QMediaServiceSupportedDevicesInterface_PTR() *QMediaServiceSupportedDevicesInterface
}

func (p *QMediaServiceSupportedDevicesInterface) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMediaServiceSupportedDevicesInterface) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMediaServiceSupportedDevicesInterface(ptr QMediaServiceSupportedDevicesInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaServiceSupportedDevicesInterface_PTR().Pointer()
	}
	return nil
}

func NewQMediaServiceSupportedDevicesInterfaceFromPointer(ptr unsafe.Pointer) *QMediaServiceSupportedDevicesInterface {
	var n = new(QMediaServiceSupportedDevicesInterface)
	n.SetPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QMediaServiceSupportedDevicesInterface_") {
		n.SetObjectNameAbs("QMediaServiceSupportedDevicesInterface_" + qt.Identifier())
	}
	return n
}

func (ptr *QMediaServiceSupportedDevicesInterface) QMediaServiceSupportedDevicesInterface_PTR() *QMediaServiceSupportedDevicesInterface {
	return ptr
}

func (ptr *QMediaServiceSupportedDevicesInterface) DeviceDescription(service core.QByteArray_ITF, device core.QByteArray_ITF) string {
	defer qt.Recovering("QMediaServiceSupportedDevicesInterface::deviceDescription")

	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaServiceSupportedDevicesInterface_DeviceDescription(ptr.Pointer(), core.PointerFromQByteArray(service), core.PointerFromQByteArray(device)))
	}
	return ""
}

func (ptr *QMediaServiceSupportedDevicesInterface) DestroyQMediaServiceSupportedDevicesInterface() {
	defer qt.Recovering("QMediaServiceSupportedDevicesInterface::~QMediaServiceSupportedDevicesInterface")

	if ptr.Pointer() != nil {
		C.QMediaServiceSupportedDevicesInterface_DestroyQMediaServiceSupportedDevicesInterface(ptr.Pointer())
	}
}

func (ptr *QMediaServiceSupportedDevicesInterface) ObjectNameAbs() string {
	defer qt.Recovering("QMediaServiceSupportedDevicesInterface::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaServiceSupportedDevicesInterface_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMediaServiceSupportedDevicesInterface) SetObjectNameAbs(name string) {
	defer qt.Recovering("QMediaServiceSupportedDevicesInterface::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QMediaServiceSupportedDevicesInterface_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
