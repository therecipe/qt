package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QTextBlockUserData struct {
	ptr unsafe.Pointer
}

type QTextBlockUserData_ITF interface {
	QTextBlockUserData_PTR() *QTextBlockUserData
}

func (p *QTextBlockUserData) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTextBlockUserData) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTextBlockUserData(ptr QTextBlockUserData_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextBlockUserData_PTR().Pointer()
	}
	return nil
}

func NewQTextBlockUserDataFromPointer(ptr unsafe.Pointer) *QTextBlockUserData {
	var n = new(QTextBlockUserData)
	n.SetPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QTextBlockUserData_") {
		n.SetObjectNameAbs("QTextBlockUserData_" + qt.Identifier())
	}
	return n
}

func (ptr *QTextBlockUserData) QTextBlockUserData_PTR() *QTextBlockUserData {
	return ptr
}

func (ptr *QTextBlockUserData) DestroyQTextBlockUserData() {
	defer qt.Recovering("QTextBlockUserData::~QTextBlockUserData")

	if ptr.Pointer() != nil {
		C.QTextBlockUserData_DestroyQTextBlockUserData(ptr.Pointer())
	}
}

func (ptr *QTextBlockUserData) ObjectNameAbs() string {
	defer qt.Recovering("QTextBlockUserData::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QTextBlockUserData_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTextBlockUserData) SetObjectNameAbs(name string) {
	defer qt.Recovering("QTextBlockUserData::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QTextBlockUserData_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
