package xml

//#include "qdomcharacterdata.h"
import "C"
import (
	"unsafe"
)

type QDomCharacterData struct {
	QDomNode
}

type QDomCharacterDataITF interface {
	QDomNodeITF
	QDomCharacterDataPTR() *QDomCharacterData
}

func PointerFromQDomCharacterData(ptr QDomCharacterDataITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomCharacterDataPTR().Pointer()
	}
	return nil
}

func QDomCharacterDataFromPointer(ptr unsafe.Pointer) *QDomCharacterData {
	var n = new(QDomCharacterData)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDomCharacterData) QDomCharacterDataPTR() *QDomCharacterData {
	return ptr
}

func NewQDomCharacterData() *QDomCharacterData {
	return QDomCharacterDataFromPointer(unsafe.Pointer(C.QDomCharacterData_NewQDomCharacterData()))
}

func NewQDomCharacterData2(x QDomCharacterDataITF) *QDomCharacterData {
	return QDomCharacterDataFromPointer(unsafe.Pointer(C.QDomCharacterData_NewQDomCharacterData2(C.QtObjectPtr(PointerFromQDomCharacterData(x)))))
}

func (ptr *QDomCharacterData) AppendData(arg string) {
	if ptr.Pointer() != nil {
		C.QDomCharacterData_AppendData(C.QtObjectPtr(ptr.Pointer()), C.CString(arg))
	}
}

func (ptr *QDomCharacterData) Data() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDomCharacterData_Data(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QDomCharacterData) Length() int {
	if ptr.Pointer() != nil {
		return int(C.QDomCharacterData_Length(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDomCharacterData) NodeType() QDomNode__NodeType {
	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomCharacterData_NodeType(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDomCharacterData) SetData(v string) {
	if ptr.Pointer() != nil {
		C.QDomCharacterData_SetData(C.QtObjectPtr(ptr.Pointer()), C.CString(v))
	}
}
