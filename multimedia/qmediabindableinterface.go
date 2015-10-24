package multimedia

//#include "qmediabindableinterface.h"
import "C"
import (
	"unsafe"
)

type QMediaBindableInterface struct {
	ptr unsafe.Pointer
}

type QMediaBindableInterfaceITF interface {
	QMediaBindableInterfacePTR() *QMediaBindableInterface
}

func (p *QMediaBindableInterface) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMediaBindableInterface) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMediaBindableInterface(ptr QMediaBindableInterfaceITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaBindableInterfacePTR().Pointer()
	}
	return nil
}

func QMediaBindableInterfaceFromPointer(ptr unsafe.Pointer) *QMediaBindableInterface {
	var n = new(QMediaBindableInterface)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMediaBindableInterface) QMediaBindableInterfacePTR() *QMediaBindableInterface {
	return ptr
}

func (ptr *QMediaBindableInterface) MediaObject() *QMediaObject {
	if ptr.Pointer() != nil {
		return QMediaObjectFromPointer(unsafe.Pointer(C.QMediaBindableInterface_MediaObject(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QMediaBindableInterface) DestroyQMediaBindableInterface() {
	if ptr.Pointer() != nil {
		C.QMediaBindableInterface_DestroyQMediaBindableInterface(C.QtObjectPtr(ptr.Pointer()))
	}
}
