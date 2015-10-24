package core

//#include "qxmlstreamreader.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QXmlStreamReader struct {
	ptr unsafe.Pointer
}

type QXmlStreamReaderITF interface {
	QXmlStreamReaderPTR() *QXmlStreamReader
}

func (p *QXmlStreamReader) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlStreamReader) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlStreamReader(ptr QXmlStreamReaderITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlStreamReaderPTR().Pointer()
	}
	return nil
}

func QXmlStreamReaderFromPointer(ptr unsafe.Pointer) *QXmlStreamReader {
	var n = new(QXmlStreamReader)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlStreamReader) QXmlStreamReaderPTR() *QXmlStreamReader {
	return ptr
}

//QXmlStreamReader::Error
type QXmlStreamReader__Error int

var (
	QXmlStreamReader__NoError                     = QXmlStreamReader__Error(0)
	QXmlStreamReader__UnexpectedElementError      = QXmlStreamReader__Error(1)
	QXmlStreamReader__CustomError                 = QXmlStreamReader__Error(2)
	QXmlStreamReader__NotWellFormedError          = QXmlStreamReader__Error(3)
	QXmlStreamReader__PrematureEndOfDocumentError = QXmlStreamReader__Error(4)
)

//QXmlStreamReader::ReadElementTextBehaviour
type QXmlStreamReader__ReadElementTextBehaviour int

var (
	QXmlStreamReader__ErrorOnUnexpectedElement = QXmlStreamReader__ReadElementTextBehaviour(0)
	QXmlStreamReader__IncludeChildElements     = QXmlStreamReader__ReadElementTextBehaviour(1)
	QXmlStreamReader__SkipChildElements        = QXmlStreamReader__ReadElementTextBehaviour(2)
)

//QXmlStreamReader::TokenType
type QXmlStreamReader__TokenType int

var (
	QXmlStreamReader__NoToken               = QXmlStreamReader__TokenType(0)
	QXmlStreamReader__Invalid               = QXmlStreamReader__TokenType(1)
	QXmlStreamReader__StartDocument         = QXmlStreamReader__TokenType(2)
	QXmlStreamReader__EndDocument           = QXmlStreamReader__TokenType(3)
	QXmlStreamReader__StartElement          = QXmlStreamReader__TokenType(4)
	QXmlStreamReader__EndElement            = QXmlStreamReader__TokenType(5)
	QXmlStreamReader__Characters            = QXmlStreamReader__TokenType(6)
	QXmlStreamReader__Comment               = QXmlStreamReader__TokenType(7)
	QXmlStreamReader__DTD                   = QXmlStreamReader__TokenType(8)
	QXmlStreamReader__EntityReference       = QXmlStreamReader__TokenType(9)
	QXmlStreamReader__ProcessingInstruction = QXmlStreamReader__TokenType(10)
)

func (ptr *QXmlStreamReader) NamespaceProcessing() bool {
	if ptr.Pointer() != nil {
		return C.QXmlStreamReader_NamespaceProcessing(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QXmlStreamReader) SetNamespaceProcessing(v bool) {
	if ptr.Pointer() != nil {
		C.QXmlStreamReader_SetNamespaceProcessing(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(v)))
	}
}

func NewQXmlStreamReader() *QXmlStreamReader {
	return QXmlStreamReaderFromPointer(unsafe.Pointer(C.QXmlStreamReader_NewQXmlStreamReader()))
}

func NewQXmlStreamReader2(device QIODeviceITF) *QXmlStreamReader {
	return QXmlStreamReaderFromPointer(unsafe.Pointer(C.QXmlStreamReader_NewQXmlStreamReader2(C.QtObjectPtr(PointerFromQIODevice(device)))))
}

func NewQXmlStreamReader3(data QByteArrayITF) *QXmlStreamReader {
	return QXmlStreamReaderFromPointer(unsafe.Pointer(C.QXmlStreamReader_NewQXmlStreamReader3(C.QtObjectPtr(PointerFromQByteArray(data)))))
}

func NewQXmlStreamReader4(data string) *QXmlStreamReader {
	return QXmlStreamReaderFromPointer(unsafe.Pointer(C.QXmlStreamReader_NewQXmlStreamReader4(C.CString(data))))
}

func NewQXmlStreamReader5(data string) *QXmlStreamReader {
	return QXmlStreamReaderFromPointer(unsafe.Pointer(C.QXmlStreamReader_NewQXmlStreamReader5(C.CString(data))))
}

func (ptr *QXmlStreamReader) AddData(data QByteArrayITF) {
	if ptr.Pointer() != nil {
		C.QXmlStreamReader_AddData(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQByteArray(data)))
	}
}

func (ptr *QXmlStreamReader) AddData2(data string) {
	if ptr.Pointer() != nil {
		C.QXmlStreamReader_AddData2(C.QtObjectPtr(ptr.Pointer()), C.CString(data))
	}
}

func (ptr *QXmlStreamReader) AddData3(data string) {
	if ptr.Pointer() != nil {
		C.QXmlStreamReader_AddData3(C.QtObjectPtr(ptr.Pointer()), C.CString(data))
	}
}

func (ptr *QXmlStreamReader) AddExtraNamespaceDeclaration(extraNamespaceDeclaration QXmlStreamNamespaceDeclarationITF) {
	if ptr.Pointer() != nil {
		C.QXmlStreamReader_AddExtraNamespaceDeclaration(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlStreamNamespaceDeclaration(extraNamespaceDeclaration)))
	}
}

func (ptr *QXmlStreamReader) AtEnd() bool {
	if ptr.Pointer() != nil {
		return C.QXmlStreamReader_AtEnd(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QXmlStreamReader) Clear() {
	if ptr.Pointer() != nil {
		C.QXmlStreamReader_Clear(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QXmlStreamReader) Device() *QIODevice {
	if ptr.Pointer() != nil {
		return QIODeviceFromPointer(unsafe.Pointer(C.QXmlStreamReader_Device(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QXmlStreamReader) EntityResolver() *QXmlStreamEntityResolver {
	if ptr.Pointer() != nil {
		return QXmlStreamEntityResolverFromPointer(unsafe.Pointer(C.QXmlStreamReader_EntityResolver(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QXmlStreamReader) Error() QXmlStreamReader__Error {
	if ptr.Pointer() != nil {
		return QXmlStreamReader__Error(C.QXmlStreamReader_Error(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QXmlStreamReader) ErrorString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlStreamReader_ErrorString(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QXmlStreamReader) HasError() bool {
	if ptr.Pointer() != nil {
		return C.QXmlStreamReader_HasError(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QXmlStreamReader) IsCDATA() bool {
	if ptr.Pointer() != nil {
		return C.QXmlStreamReader_IsCDATA(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QXmlStreamReader) IsCharacters() bool {
	if ptr.Pointer() != nil {
		return C.QXmlStreamReader_IsCharacters(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QXmlStreamReader) IsComment() bool {
	if ptr.Pointer() != nil {
		return C.QXmlStreamReader_IsComment(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QXmlStreamReader) IsDTD() bool {
	if ptr.Pointer() != nil {
		return C.QXmlStreamReader_IsDTD(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QXmlStreamReader) IsEndDocument() bool {
	if ptr.Pointer() != nil {
		return C.QXmlStreamReader_IsEndDocument(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QXmlStreamReader) IsEndElement() bool {
	if ptr.Pointer() != nil {
		return C.QXmlStreamReader_IsEndElement(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QXmlStreamReader) IsEntityReference() bool {
	if ptr.Pointer() != nil {
		return C.QXmlStreamReader_IsEntityReference(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QXmlStreamReader) IsProcessingInstruction() bool {
	if ptr.Pointer() != nil {
		return C.QXmlStreamReader_IsProcessingInstruction(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QXmlStreamReader) IsStandaloneDocument() bool {
	if ptr.Pointer() != nil {
		return C.QXmlStreamReader_IsStandaloneDocument(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QXmlStreamReader) IsStartDocument() bool {
	if ptr.Pointer() != nil {
		return C.QXmlStreamReader_IsStartDocument(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QXmlStreamReader) IsStartElement() bool {
	if ptr.Pointer() != nil {
		return C.QXmlStreamReader_IsStartElement(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QXmlStreamReader) IsWhitespace() bool {
	if ptr.Pointer() != nil {
		return C.QXmlStreamReader_IsWhitespace(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QXmlStreamReader) RaiseError(message string) {
	if ptr.Pointer() != nil {
		C.QXmlStreamReader_RaiseError(C.QtObjectPtr(ptr.Pointer()), C.CString(message))
	}
}

func (ptr *QXmlStreamReader) ReadElementText(behaviour QXmlStreamReader__ReadElementTextBehaviour) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlStreamReader_ReadElementText(C.QtObjectPtr(ptr.Pointer()), C.int(behaviour)))
	}
	return ""
}

func (ptr *QXmlStreamReader) ReadNext() QXmlStreamReader__TokenType {
	if ptr.Pointer() != nil {
		return QXmlStreamReader__TokenType(C.QXmlStreamReader_ReadNext(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QXmlStreamReader) ReadNextStartElement() bool {
	if ptr.Pointer() != nil {
		return C.QXmlStreamReader_ReadNextStartElement(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QXmlStreamReader) SetDevice(device QIODeviceITF) {
	if ptr.Pointer() != nil {
		C.QXmlStreamReader_SetDevice(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQIODevice(device)))
	}
}

func (ptr *QXmlStreamReader) SetEntityResolver(resolver QXmlStreamEntityResolverITF) {
	if ptr.Pointer() != nil {
		C.QXmlStreamReader_SetEntityResolver(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlStreamEntityResolver(resolver)))
	}
}

func (ptr *QXmlStreamReader) SkipCurrentElement() {
	if ptr.Pointer() != nil {
		C.QXmlStreamReader_SkipCurrentElement(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QXmlStreamReader) TokenString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlStreamReader_TokenString(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QXmlStreamReader) TokenType() QXmlStreamReader__TokenType {
	if ptr.Pointer() != nil {
		return QXmlStreamReader__TokenType(C.QXmlStreamReader_TokenType(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QXmlStreamReader) DestroyQXmlStreamReader() {
	if ptr.Pointer() != nil {
		C.QXmlStreamReader_DestroyQXmlStreamReader(C.QtObjectPtr(ptr.Pointer()))
	}
}
