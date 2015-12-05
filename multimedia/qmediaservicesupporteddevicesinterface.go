package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
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
	return n
}

func (ptr *QMediaServiceSupportedDevicesInterface) QMediaServiceSupportedDevicesInterface_PTR() *QMediaServiceSupportedDevicesInterface {
	return ptr
}

func (ptr *QMediaServiceSupportedDevicesInterface) DeviceDescription(service core.QByteArray_ITF, device core.QByteArray_ITF) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaServiceSupportedDevicesInterface::deviceDescription")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaServiceSupportedDevicesInterface_DeviceDescription(ptr.Pointer(), core.PointerFromQByteArray(service), core.PointerFromQByteArray(device)))
	}
	return ""
}

func (ptr *QMediaServiceSupportedDevicesInterface) DestroyQMediaServiceSupportedDevicesInterface() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaServiceSupportedDevicesInterface::~QMediaServiceSupportedDevicesInterface")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaServiceSupportedDevicesInterface_DestroyQMediaServiceSupportedDevicesInterface(ptr.Pointer())
	}
}
