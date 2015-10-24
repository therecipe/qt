package multimedia

//#include "qmediaservicedefaultdeviceinterface.h"
import "C"
import (
	"unsafe"
)

type QMediaServiceDefaultDeviceInterface struct {
	ptr unsafe.Pointer
}

type QMediaServiceDefaultDeviceInterfaceITF interface {
	QMediaServiceDefaultDeviceInterfacePTR() *QMediaServiceDefaultDeviceInterface
}

func (p *QMediaServiceDefaultDeviceInterface) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMediaServiceDefaultDeviceInterface) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMediaServiceDefaultDeviceInterface(ptr QMediaServiceDefaultDeviceInterfaceITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaServiceDefaultDeviceInterfacePTR().Pointer()
	}
	return nil
}

func QMediaServiceDefaultDeviceInterfaceFromPointer(ptr unsafe.Pointer) *QMediaServiceDefaultDeviceInterface {
	var n = new(QMediaServiceDefaultDeviceInterface)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMediaServiceDefaultDeviceInterface) QMediaServiceDefaultDeviceInterfacePTR() *QMediaServiceDefaultDeviceInterface {
	return ptr
}

func (ptr *QMediaServiceDefaultDeviceInterface) DestroyQMediaServiceDefaultDeviceInterface() {
	if ptr.Pointer() != nil {
		C.QMediaServiceDefaultDeviceInterface_DestroyQMediaServiceDefaultDeviceInterface(C.QtObjectPtr(ptr.Pointer()))
	}
}
