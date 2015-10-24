package multimedia

//#include "qmediaservicesupportedformatsinterface.h"
import "C"
import (
	"strings"
	"unsafe"
)

type QMediaServiceSupportedFormatsInterface struct {
	ptr unsafe.Pointer
}

type QMediaServiceSupportedFormatsInterfaceITF interface {
	QMediaServiceSupportedFormatsInterfacePTR() *QMediaServiceSupportedFormatsInterface
}

func (p *QMediaServiceSupportedFormatsInterface) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMediaServiceSupportedFormatsInterface) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMediaServiceSupportedFormatsInterface(ptr QMediaServiceSupportedFormatsInterfaceITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaServiceSupportedFormatsInterfacePTR().Pointer()
	}
	return nil
}

func QMediaServiceSupportedFormatsInterfaceFromPointer(ptr unsafe.Pointer) *QMediaServiceSupportedFormatsInterface {
	var n = new(QMediaServiceSupportedFormatsInterface)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMediaServiceSupportedFormatsInterface) QMediaServiceSupportedFormatsInterfacePTR() *QMediaServiceSupportedFormatsInterface {
	return ptr
}

func (ptr *QMediaServiceSupportedFormatsInterface) SupportedMimeTypes() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QMediaServiceSupportedFormatsInterface_SupportedMimeTypes(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QMediaServiceSupportedFormatsInterface) DestroyQMediaServiceSupportedFormatsInterface() {
	if ptr.Pointer() != nil {
		C.QMediaServiceSupportedFormatsInterface_DestroyQMediaServiceSupportedFormatsInterface(C.QtObjectPtr(ptr.Pointer()))
	}
}
