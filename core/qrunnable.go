package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QRunnable struct {
	ptr unsafe.Pointer
}

type QRunnable_ITF interface {
	QRunnable_PTR() *QRunnable
}

func (p *QRunnable) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QRunnable) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQRunnable(ptr QRunnable_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRunnable_PTR().Pointer()
	}
	return nil
}

func NewQRunnableFromPointer(ptr unsafe.Pointer) *QRunnable {
	var n = new(QRunnable)
	n.SetPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QRunnable_") {
		n.SetObjectNameAbs("QRunnable_" + qt.Identifier())
	}
	return n
}

func (ptr *QRunnable) QRunnable_PTR() *QRunnable {
	return ptr
}

func (ptr *QRunnable) AutoDelete() bool {
	defer qt.Recovering("QRunnable::autoDelete")

	if ptr.Pointer() != nil {
		return C.QRunnable_AutoDelete(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QRunnable) Run() {
	defer qt.Recovering("QRunnable::run")

	if ptr.Pointer() != nil {
		C.QRunnable_Run(ptr.Pointer())
	}
}

func (ptr *QRunnable) SetAutoDelete(autoDelete bool) {
	defer qt.Recovering("QRunnable::setAutoDelete")

	if ptr.Pointer() != nil {
		C.QRunnable_SetAutoDelete(ptr.Pointer(), C.int(qt.GoBoolToInt(autoDelete)))
	}
}

func (ptr *QRunnable) DestroyQRunnable() {
	defer qt.Recovering("QRunnable::~QRunnable")

	if ptr.Pointer() != nil {
		C.QRunnable_DestroyQRunnable(ptr.Pointer())
	}
}

func (ptr *QRunnable) ObjectNameAbs() string {
	defer qt.Recovering("QRunnable::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QRunnable_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QRunnable) SetObjectNameAbs(name string) {
	defer qt.Recovering("QRunnable::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QRunnable_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
