package core

//#include "core.h"
import "C"
import (
	"unsafe"
)

type QModulesPrivate struct {
	ptr unsafe.Pointer
}

type QModulesPrivate_ITF interface {
	QModulesPrivate_PTR() *QModulesPrivate
}

func (p *QModulesPrivate) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QModulesPrivate) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQModulesPrivate(ptr QModulesPrivate_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QModulesPrivate_PTR().Pointer()
	}
	return nil
}

func NewQModulesPrivateFromPointer(ptr unsafe.Pointer) *QModulesPrivate {
	var n = new(QModulesPrivate)
	n.SetPointer(ptr)
	return n
}

func (ptr *QModulesPrivate) QModulesPrivate_PTR() *QModulesPrivate {
	return ptr
}

//QModulesPrivate::Names
type QModulesPrivate__Names int64

const (
	QModulesPrivate__Core         = QModulesPrivate__Names(0)
	QModulesPrivate__Gui          = QModulesPrivate__Names(1)
	QModulesPrivate__Widgets      = QModulesPrivate__Names(2)
	QModulesPrivate__Unknown      = QModulesPrivate__Names(3)
	QModulesPrivate__ModulesCount = QModulesPrivate__Names(4)
)
