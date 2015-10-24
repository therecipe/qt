package multimedia

//#include "qmediaservicesupporteddevicesinterface.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QMediaServiceSupportedDevicesInterface struct {
	ptr unsafe.Pointer
}

type QMediaServiceSupportedDevicesInterfaceITF interface {
	QMediaServiceSupportedDevicesInterfacePTR() *QMediaServiceSupportedDevicesInterface
}

func (p *QMediaServiceSupportedDevicesInterface) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMediaServiceSupportedDevicesInterface) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMediaServiceSupportedDevicesInterface(ptr QMediaServiceSupportedDevicesInterfaceITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaServiceSupportedDevicesInterfacePTR().Pointer()
	}
	return nil
}

func QMediaServiceSupportedDevicesInterfaceFromPointer(ptr unsafe.Pointer) *QMediaServiceSupportedDevicesInterface {
	var n = new(QMediaServiceSupportedDevicesInterface)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMediaServiceSupportedDevicesInterface) QMediaServiceSupportedDevicesInterfacePTR() *QMediaServiceSupportedDevicesInterface {
	return ptr
}

func (ptr *QMediaServiceSupportedDevicesInterface) DeviceDescription(service core.QByteArrayITF, device core.QByteArrayITF) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaServiceSupportedDevicesInterface_DeviceDescription(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(service)), C.QtObjectPtr(core.PointerFromQByteArray(device))))
	}
	return ""
}

func (ptr *QMediaServiceSupportedDevicesInterface) DestroyQMediaServiceSupportedDevicesInterface() {
	if ptr.Pointer() != nil {
		C.QMediaServiceSupportedDevicesInterface_DestroyQMediaServiceSupportedDevicesInterface(C.QtObjectPtr(ptr.Pointer()))
	}
}
