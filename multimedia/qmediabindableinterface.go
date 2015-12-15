package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QMediaBindableInterface struct {
	ptr unsafe.Pointer
}

type QMediaBindableInterface_ITF interface {
	QMediaBindableInterface_PTR() *QMediaBindableInterface
}

func (p *QMediaBindableInterface) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMediaBindableInterface) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMediaBindableInterface(ptr QMediaBindableInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaBindableInterface_PTR().Pointer()
	}
	return nil
}

func NewQMediaBindableInterfaceFromPointer(ptr unsafe.Pointer) *QMediaBindableInterface {
	var n = new(QMediaBindableInterface)
	n.SetPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QMediaBindableInterface_") {
		n.SetObjectNameAbs("QMediaBindableInterface_" + qt.Identifier())
	}
	return n
}

func (ptr *QMediaBindableInterface) QMediaBindableInterface_PTR() *QMediaBindableInterface {
	return ptr
}

func (ptr *QMediaBindableInterface) MediaObject() *QMediaObject {
	defer qt.Recovering("QMediaBindableInterface::mediaObject")

	if ptr.Pointer() != nil {
		return NewQMediaObjectFromPointer(C.QMediaBindableInterface_MediaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMediaBindableInterface) DestroyQMediaBindableInterface() {
	defer qt.Recovering("QMediaBindableInterface::~QMediaBindableInterface")

	if ptr.Pointer() != nil {
		C.QMediaBindableInterface_DestroyQMediaBindableInterface(ptr.Pointer())
	}
}

func (ptr *QMediaBindableInterface) ObjectNameAbs() string {
	defer qt.Recovering("QMediaBindableInterface::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaBindableInterface_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMediaBindableInterface) SetObjectNameAbs(name string) {
	defer qt.Recovering("QMediaBindableInterface::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QMediaBindableInterface_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
