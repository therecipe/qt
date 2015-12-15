package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
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
	for len(n.ObjectNameAbs()) < len("QMediaServiceSupportedFormatsInterface_") {
		n.SetObjectNameAbs("QMediaServiceSupportedFormatsInterface_" + qt.Identifier())
	}
	return n
}

func (ptr *QMediaServiceSupportedFormatsInterface) QMediaServiceSupportedFormatsInterface_PTR() *QMediaServiceSupportedFormatsInterface {
	return ptr
}

func (ptr *QMediaServiceSupportedFormatsInterface) SupportedMimeTypes() []string {
	defer qt.Recovering("QMediaServiceSupportedFormatsInterface::supportedMimeTypes")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QMediaServiceSupportedFormatsInterface_SupportedMimeTypes(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QMediaServiceSupportedFormatsInterface) DestroyQMediaServiceSupportedFormatsInterface() {
	defer qt.Recovering("QMediaServiceSupportedFormatsInterface::~QMediaServiceSupportedFormatsInterface")

	if ptr.Pointer() != nil {
		C.QMediaServiceSupportedFormatsInterface_DestroyQMediaServiceSupportedFormatsInterface(ptr.Pointer())
	}
}

func (ptr *QMediaServiceSupportedFormatsInterface) ObjectNameAbs() string {
	defer qt.Recovering("QMediaServiceSupportedFormatsInterface::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaServiceSupportedFormatsInterface_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMediaServiceSupportedFormatsInterface) SetObjectNameAbs(name string) {
	defer qt.Recovering("QMediaServiceSupportedFormatsInterface::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QMediaServiceSupportedFormatsInterface_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
