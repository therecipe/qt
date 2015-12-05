package xml

//#include "xml.h"
import "C"
import (
	"log"
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDomCharacterData::QDomCharacterData")
		}
	}()

	return NewQDomCharacterDataFromPointer(C.QDomCharacterData_NewQDomCharacterData())
}

func NewQDomCharacterData2(x QDomCharacterData_ITF) *QDomCharacterData {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDomCharacterData::QDomCharacterData")
		}
	}()

	return NewQDomCharacterDataFromPointer(C.QDomCharacterData_NewQDomCharacterData2(PointerFromQDomCharacterData(x)))
}

func (ptr *QDomCharacterData) AppendData(arg string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDomCharacterData::appendData")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDomCharacterData_AppendData(ptr.Pointer(), C.CString(arg))
	}
}

func (ptr *QDomCharacterData) Data() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDomCharacterData::data")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QDomCharacterData_Data(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomCharacterData) Length() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDomCharacterData::length")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QDomCharacterData_Length(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomCharacterData) NodeType() QDomNode__NodeType {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDomCharacterData::nodeType")
		}
	}()

	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomCharacterData_NodeType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomCharacterData) SetData(v string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDomCharacterData::setData")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDomCharacterData_SetData(ptr.Pointer(), C.CString(v))
	}
}
