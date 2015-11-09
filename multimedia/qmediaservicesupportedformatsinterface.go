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

type QMediaServiceSupportedFormatsInterface_ITF interface {
	QMediaServiceSupportedFormatsInterface_PTR() *QMediaServiceSupportedFormatsInterface
}

func (p *QMediaServiceSupportedFormatsInterface) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMediaServiceSupportedFormatsInterface) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMediaServiceSupportedFormatsInterface(ptr QMediaServiceSupportedFormatsInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaServiceSupportedFormatsInterface_PTR().Pointer()
	}
	return nil
}

func NewQMediaServiceSupportedFormatsInterfaceFromPointer(ptr unsafe.Pointer) *QMediaServiceSupportedFormatsInterface {
	var n = new(QMediaServiceSupportedFormatsInterface)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMediaServiceSupportedFormatsInterface) QMediaServiceSupportedFormatsInterface_PTR() *QMediaServiceSupportedFormatsInterface {
	return ptr
}

func (ptr *QMediaServiceSupportedFormatsInterface) SupportedMimeTypes() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QMediaServiceSupportedFormatsInterface_SupportedMimeTypes(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QMediaServiceSupportedFormatsInterface) DestroyQMediaServiceSupportedFormatsInterface() {
	if ptr.Pointer() != nil {
		C.QMediaServiceSupportedFormatsInterface_DestroyQMediaServiceSupportedFormatsInterface(ptr.Pointer())
	}
}
