package xml

//#include "qdomcharacterdata.h"
import "C"
import (
	"unsafe"
)

type QDomCharacterData struct {
	QDomNode
}

type QDomCharacterData_ITF interface {
	QDomNode_ITF
	QDomCharacterData_PTR() *QDomCharacterData
}

func PointerFromQDomCharacterData(ptr QDomCharacterData_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomCharacterData_PTR().Pointer()
	}
	return nil
}

func NewQDomCharacterDataFromPointer(ptr unsafe.Pointer) *QDomCharacterData {
	var n = new(QDomCharacterData)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDomCharacterData) QDomCharacterData_PTR() *QDomCharacterData {
	return ptr
}

func NewQDomCharacterData() *QDomCharacterData {
	return NewQDomCharacterDataFromPointer(C.QDomCharacterData_NewQDomCharacterData())
}

func NewQDomCharacterData2(x QDomCharacterData_ITF) *QDomCharacterData {
	return NewQDomCharacterDataFromPointer(C.QDomCharacterData_NewQDomCharacterData2(PointerFromQDomCharacterData(x)))
}

func (ptr *QDomCharacterData) AppendData(arg string) {
	if ptr.Pointer() != nil {
		C.QDomCharacterData_AppendData(ptr.Pointer(), C.CString(arg))
	}
}

func (ptr *QDomCharacterData) Data() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDomCharacterData_Data(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomCharacterData) Length() int {
	if ptr.Pointer() != nil {
		return int(C.QDomCharacterData_Length(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomCharacterData) NodeType() QDomNode__NodeType {
	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomCharacterData_NodeType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomCharacterData) SetData(v string) {
	if ptr.Pointer() != nil {
		C.QDomCharacterData_SetData(ptr.Pointer(), C.CString(v))
	}
}
