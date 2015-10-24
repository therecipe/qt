package gui

//#include "qtextblockuserdata.h"
import "C"
import (
	"unsafe"
)

type QTextBlockUserData struct {
	ptr unsafe.Pointer
}

type QTextBlockUserDataITF interface {
	QTextBlockUserDataPTR() *QTextBlockUserData
}

func (p *QTextBlockUserData) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTextBlockUserData) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTextBlockUserData(ptr QTextBlockUserDataITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextBlockUserDataPTR().Pointer()
	}
	return nil
}

func QTextBlockUserDataFromPointer(ptr unsafe.Pointer) *QTextBlockUserData {
	var n = new(QTextBlockUserData)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTextBlockUserData) QTextBlockUserDataPTR() *QTextBlockUserData {
	return ptr
}

func (ptr *QTextBlockUserData) DestroyQTextBlockUserData() {
	if ptr.Pointer() != nil {
		C.QTextBlockUserData_DestroyQTextBlockUserData(C.QtObjectPtr(ptr.Pointer()))
	}
}
