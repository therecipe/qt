package multimedia

//#include "qmediabindableinterface.h"
import "C"
import (
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
	return n
}

func (ptr *QMediaBindableInterface) QMediaBindableInterface_PTR() *QMediaBindableInterface {
	return ptr
}

func (ptr *QMediaBindableInterface) MediaObject() *QMediaObject {
	if ptr.Pointer() != nil {
		return NewQMediaObjectFromPointer(C.QMediaBindableInterface_MediaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMediaBindableInterface) DestroyQMediaBindableInterface() {
	if ptr.Pointer() != nil {
		C.QMediaBindableInterface_DestroyQMediaBindableInterface(ptr.Pointer())
	}
}
