package xml

//#include "qdomdocument.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QDomDocument struct {
	QDomNode
}

type QDomDocumentITF interface {
	QDomNodeITF
	QDomDocumentPTR() *QDomDocument
}

func PointerFromQDomDocument(ptr QDomDocumentITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomDocumentPTR().Pointer()
	}
	return nil
}

func QDomDocumentFromPointer(ptr unsafe.Pointer) *QDomDocument {
	var n = new(QDomDocument)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDomDocument) QDomDocumentPTR() *QDomDocument {
	return ptr
}

func NewQDomDocument() *QDomDocument {
	return QDomDocumentFromPointer(unsafe.Pointer(C.QDomDocument_NewQDomDocument()))
}

func NewQDomDocument4(x QDomDocumentITF) *QDomDocument {
	return QDomDocumentFromPointer(unsafe.Pointer(C.QDomDocument_NewQDomDocument4(C.QtObjectPtr(PointerFromQDomDocument(x)))))
}

func NewQDomDocument3(doctype QDomDocumentTypeITF) *QDomDocument {
	return QDomDocumentFromPointer(unsafe.Pointer(C.QDomDocument_NewQDomDocument3(C.QtObjectPtr(PointerFromQDomDocumentType(doctype)))))
}

func NewQDomDocument2(name string) *QDomDocument {
	return QDomDocumentFromPointer(unsafe.Pointer(C.QDomDocument_NewQDomDocument2(C.CString(name))))
}

func (ptr *QDomDocument) NodeType() QDomNode__NodeType {
	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomDocument_NodeType(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDomDocument) SetContent7(dev core.QIODeviceITF, errorMsg string, errorLine int, errorColumn int) bool {
	if ptr.Pointer() != nil {
		return C.QDomDocument_SetContent7(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQIODevice(dev)), C.CString(errorMsg), C.int(errorLine), C.int(errorColumn)) != 0
	}
	return false
}

func (ptr *QDomDocument) SetContent3(dev core.QIODeviceITF, namespaceProcessing bool, errorMsg string, errorLine int, errorColumn int) bool {
	if ptr.Pointer() != nil {
		return C.QDomDocument_SetContent3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQIODevice(dev)), C.int(qt.GoBoolToInt(namespaceProcessing)), C.CString(errorMsg), C.int(errorLine), C.int(errorColumn)) != 0
	}
	return false
}

func (ptr *QDomDocument) SetContent8(source QXmlInputSourceITF, reader QXmlReaderITF, errorMsg string, errorLine int, errorColumn int) bool {
	if ptr.Pointer() != nil {
		return C.QDomDocument_SetContent8(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlInputSource(source)), C.QtObjectPtr(PointerFromQXmlReader(reader)), C.CString(errorMsg), C.int(errorLine), C.int(errorColumn)) != 0
	}
	return false
}

func (ptr *QDomDocument) SetContent4(source QXmlInputSourceITF, namespaceProcessing bool, errorMsg string, errorLine int, errorColumn int) bool {
	if ptr.Pointer() != nil {
		return C.QDomDocument_SetContent4(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlInputSource(source)), C.int(qt.GoBoolToInt(namespaceProcessing)), C.CString(errorMsg), C.int(errorLine), C.int(errorColumn)) != 0
	}
	return false
}

func (ptr *QDomDocument) SetContent6(buffer core.QByteArrayITF, errorMsg string, errorLine int, errorColumn int) bool {
	if ptr.Pointer() != nil {
		return C.QDomDocument_SetContent6(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(buffer)), C.CString(errorMsg), C.int(errorLine), C.int(errorColumn)) != 0
	}
	return false
}

func (ptr *QDomDocument) SetContent(data core.QByteArrayITF, namespaceProcessing bool, errorMsg string, errorLine int, errorColumn int) bool {
	if ptr.Pointer() != nil {
		return C.QDomDocument_SetContent(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(data)), C.int(qt.GoBoolToInt(namespaceProcessing)), C.CString(errorMsg), C.int(errorLine), C.int(errorColumn)) != 0
	}
	return false
}

func (ptr *QDomDocument) SetContent5(text string, errorMsg string, errorLine int, errorColumn int) bool {
	if ptr.Pointer() != nil {
		return C.QDomDocument_SetContent5(C.QtObjectPtr(ptr.Pointer()), C.CString(text), C.CString(errorMsg), C.int(errorLine), C.int(errorColumn)) != 0
	}
	return false
}

func (ptr *QDomDocument) SetContent2(text string, namespaceProcessing bool, errorMsg string, errorLine int, errorColumn int) bool {
	if ptr.Pointer() != nil {
		return C.QDomDocument_SetContent2(C.QtObjectPtr(ptr.Pointer()), C.CString(text), C.int(qt.GoBoolToInt(namespaceProcessing)), C.CString(errorMsg), C.int(errorLine), C.int(errorColumn)) != 0
	}
	return false
}

func (ptr *QDomDocument) ToString(indent int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDomDocument_ToString(C.QtObjectPtr(ptr.Pointer()), C.int(indent)))
	}
	return ""
}

func (ptr *QDomDocument) DestroyQDomDocument() {
	if ptr.Pointer() != nil {
		C.QDomDocument_DestroyQDomDocument(C.QtObjectPtr(ptr.Pointer()))
	}
}
