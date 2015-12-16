package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QXmlStreamReader struct {
	ptr unsafe.Pointer
}

type QXmlStreamReader_ITF interface {
	QXmlStreamReader_PTR() *QXmlStreamReader
}

func (p *QXmlStreamReader) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlStreamReader) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlStreamReader(ptr QXmlStreamReader_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlStreamReader_PTR().Pointer()
	}
	return nil
}

func NewQXmlStreamReaderFromPointer(ptr unsafe.Pointer) *QXmlStreamReader {
	var n = new(QXmlStreamReader)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlStreamReader) QXmlStreamReader_PTR() *QXmlStreamReader {
	return ptr
}

//QXmlStreamReader::Error
type QXmlStreamReader__Error int64

const (
	QXmlStreamReader__NoError                     = QXmlStreamReader__Error(0)
	QXmlStreamReader__UnexpectedElementError      = QXmlStreamReader__Error(1)
	QXmlStreamReader__CustomError                 = QXmlStreamReader__Error(2)
	QXmlStreamReader__NotWellFormedError          = QXmlStreamReader__Error(3)
	QXmlStreamReader__PrematureEndOfDocumentError = QXmlStreamReader__Error(4)
)

//QXmlStreamReader::ReadElementTextBehaviour
type QXmlStreamReader__ReadElementTextBehaviour int64

const (
	QXmlStreamReader__ErrorOnUnexpectedElement = QXmlStreamReader__ReadElementTextBehaviour(0)
	QXmlStreamReader__IncludeChildElements     = QXmlStreamReader__ReadElementTextBehaviour(1)
	QXmlStreamReader__SkipChildElements        = QXmlStreamReader__ReadElementTextBehaviour(2)
)

//QXmlStreamReader::TokenType
type QXmlStreamReader__TokenType int64

const (
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
	defer qt.Recovering("QXmlStreamReader::namespaceProcessing")

	if ptr.Pointer() != nil {
		return C.QXmlStreamReader_NamespaceProcessing(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlStreamReader) SetNamespaceProcessing(v bool) {
	defer qt.Recovering("QXmlStreamReader::setNamespaceProcessing")

	if ptr.Pointer() != nil {
		C.QXmlStreamReader_SetNamespaceProcessing(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func NewQXmlStreamReader() *QXmlStreamReader {
	defer qt.Recovering("QXmlStreamReader::QXmlStreamReader")

	return NewQXmlStreamReaderFromPointer(C.QXmlStreamReader_NewQXmlStreamReader())
}

func NewQXmlStreamReader2(device QIODevice_ITF) *QXmlStreamReader {
	defer qt.Recovering("QXmlStreamReader::QXmlStreamReader")

	return NewQXmlStreamReaderFromPointer(C.QXmlStreamReader_NewQXmlStreamReader2(PointerFromQIODevice(device)))
}

func NewQXmlStreamReader3(data QByteArray_ITF) *QXmlStreamReader {
	defer qt.Recovering("QXmlStreamReader::QXmlStreamReader")

	return NewQXmlStreamReaderFromPointer(C.QXmlStreamReader_NewQXmlStreamReader3(PointerFromQByteArray(data)))
}

func NewQXmlStreamReader4(data string) *QXmlStreamReader {
	defer qt.Recovering("QXmlStreamReader::QXmlStreamReader")

	return NewQXmlStreamReaderFromPointer(C.QXmlStreamReader_NewQXmlStreamReader4(C.CString(data)))
}

func NewQXmlStreamReader5(data string) *QXmlStreamReader {
	defer qt.Recovering("QXmlStreamReader::QXmlStreamReader")

	return NewQXmlStreamReaderFromPointer(C.QXmlStreamReader_NewQXmlStreamReader5(C.CString(data)))
}

func (ptr *QXmlStreamReader) AddData(data QByteArray_ITF) {
	defer qt.Recovering("QXmlStreamReader::addData")

	if ptr.Pointer() != nil {
		C.QXmlStreamReader_AddData(ptr.Pointer(), PointerFromQByteArray(data))
	}
}

func (ptr *QXmlStreamReader) AddData2(data string) {
	defer qt.Recovering("QXmlStreamReader::addData")

	if ptr.Pointer() != nil {
		C.QXmlStreamReader_AddData2(ptr.Pointer(), C.CString(data))
	}
}

func (ptr *QXmlStreamReader) AddData3(data string) {
	defer qt.Recovering("QXmlStreamReader::addData")

	if ptr.Pointer() != nil {
		C.QXmlStreamReader_AddData3(ptr.Pointer(), C.CString(data))
	}
}

func (ptr *QXmlStreamReader) AddExtraNamespaceDeclaration(extraNamespaceDeclaration QXmlStreamNamespaceDeclaration_ITF) {
	defer qt.Recovering("QXmlStreamReader::addExtraNamespaceDeclaration")

	if ptr.Pointer() != nil {
		C.QXmlStreamReader_AddExtraNamespaceDeclaration(ptr.Pointer(), PointerFromQXmlStreamNamespaceDeclaration(extraNamespaceDeclaration))
	}
}

func (ptr *QXmlStreamReader) AtEnd() bool {
	defer qt.Recovering("QXmlStreamReader::atEnd")

	if ptr.Pointer() != nil {
		return C.QXmlStreamReader_AtEnd(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlStreamReader) CharacterOffset() int64 {
	defer qt.Recovering("QXmlStreamReader::characterOffset")

	if ptr.Pointer() != nil {
		return int64(C.QXmlStreamReader_CharacterOffset(ptr.Pointer()))
	}
	return 0
}

func (ptr *QXmlStreamReader) Clear() {
	defer qt.Recovering("QXmlStreamReader::clear")

	if ptr.Pointer() != nil {
		C.QXmlStreamReader_Clear(ptr.Pointer())
	}
}

func (ptr *QXmlStreamReader) ColumnNumber() int64 {
	defer qt.Recovering("QXmlStreamReader::columnNumber")

	if ptr.Pointer() != nil {
		return int64(C.QXmlStreamReader_ColumnNumber(ptr.Pointer()))
	}
	return 0
}

func (ptr *QXmlStreamReader) Device() *QIODevice {
	defer qt.Recovering("QXmlStreamReader::device")

	if ptr.Pointer() != nil {
		return NewQIODeviceFromPointer(C.QXmlStreamReader_Device(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlStreamReader) DocumentEncoding() *QStringRef {
	defer qt.Recovering("QXmlStreamReader::documentEncoding")

	if ptr.Pointer() != nil {
		return NewQStringRefFromPointer(C.QXmlStreamReader_DocumentEncoding(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlStreamReader) DocumentVersion() *QStringRef {
	defer qt.Recovering("QXmlStreamReader::documentVersion")

	if ptr.Pointer() != nil {
		return NewQStringRefFromPointer(C.QXmlStreamReader_DocumentVersion(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlStreamReader) DtdName() *QStringRef {
	defer qt.Recovering("QXmlStreamReader::dtdName")

	if ptr.Pointer() != nil {
		return NewQStringRefFromPointer(C.QXmlStreamReader_DtdName(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlStreamReader) DtdPublicId() *QStringRef {
	defer qt.Recovering("QXmlStreamReader::dtdPublicId")

	if ptr.Pointer() != nil {
		return NewQStringRefFromPointer(C.QXmlStreamReader_DtdPublicId(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlStreamReader) DtdSystemId() *QStringRef {
	defer qt.Recovering("QXmlStreamReader::dtdSystemId")

	if ptr.Pointer() != nil {
		return NewQStringRefFromPointer(C.QXmlStreamReader_DtdSystemId(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlStreamReader) EntityResolver() *QXmlStreamEntityResolver {
	defer qt.Recovering("QXmlStreamReader::entityResolver")

	if ptr.Pointer() != nil {
		return NewQXmlStreamEntityResolverFromPointer(C.QXmlStreamReader_EntityResolver(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlStreamReader) Error() QXmlStreamReader__Error {
	defer qt.Recovering("QXmlStreamReader::error")

	if ptr.Pointer() != nil {
		return QXmlStreamReader__Error(C.QXmlStreamReader_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QXmlStreamReader) ErrorString() string {
	defer qt.Recovering("QXmlStreamReader::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlStreamReader_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlStreamReader) HasError() bool {
	defer qt.Recovering("QXmlStreamReader::hasError")

	if ptr.Pointer() != nil {
		return C.QXmlStreamReader_HasError(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlStreamReader) IsCDATA() bool {
	defer qt.Recovering("QXmlStreamReader::isCDATA")

	if ptr.Pointer() != nil {
		return C.QXmlStreamReader_IsCDATA(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlStreamReader) IsCharacters() bool {
	defer qt.Recovering("QXmlStreamReader::isCharacters")

	if ptr.Pointer() != nil {
		return C.QXmlStreamReader_IsCharacters(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlStreamReader) IsComment() bool {
	defer qt.Recovering("QXmlStreamReader::isComment")

	if ptr.Pointer() != nil {
		return C.QXmlStreamReader_IsComment(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlStreamReader) IsDTD() bool {
	defer qt.Recovering("QXmlStreamReader::isDTD")

	if ptr.Pointer() != nil {
		return C.QXmlStreamReader_IsDTD(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlStreamReader) IsEndDocument() bool {
	defer qt.Recovering("QXmlStreamReader::isEndDocument")

	if ptr.Pointer() != nil {
		return C.QXmlStreamReader_IsEndDocument(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlStreamReader) IsEndElement() bool {
	defer qt.Recovering("QXmlStreamReader::isEndElement")

	if ptr.Pointer() != nil {
		return C.QXmlStreamReader_IsEndElement(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlStreamReader) IsEntityReference() bool {
	defer qt.Recovering("QXmlStreamReader::isEntityReference")

	if ptr.Pointer() != nil {
		return C.QXmlStreamReader_IsEntityReference(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlStreamReader) IsProcessingInstruction() bool {
	defer qt.Recovering("QXmlStreamReader::isProcessingInstruction")

	if ptr.Pointer() != nil {
		return C.QXmlStreamReader_IsProcessingInstruction(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlStreamReader) IsStandaloneDocument() bool {
	defer qt.Recovering("QXmlStreamReader::isStandaloneDocument")

	if ptr.Pointer() != nil {
		return C.QXmlStreamReader_IsStandaloneDocument(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlStreamReader) IsStartDocument() bool {
	defer qt.Recovering("QXmlStreamReader::isStartDocument")

	if ptr.Pointer() != nil {
		return C.QXmlStreamReader_IsStartDocument(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlStreamReader) IsStartElement() bool {
	defer qt.Recovering("QXmlStreamReader::isStartElement")

	if ptr.Pointer() != nil {
		return C.QXmlStreamReader_IsStartElement(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlStreamReader) IsWhitespace() bool {
	defer qt.Recovering("QXmlStreamReader::isWhitespace")

	if ptr.Pointer() != nil {
		return C.QXmlStreamReader_IsWhitespace(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlStreamReader) LineNumber() int64 {
	defer qt.Recovering("QXmlStreamReader::lineNumber")

	if ptr.Pointer() != nil {
		return int64(C.QXmlStreamReader_LineNumber(ptr.Pointer()))
	}
	return 0
}

func (ptr *QXmlStreamReader) Name() *QStringRef {
	defer qt.Recovering("QXmlStreamReader::name")

	if ptr.Pointer() != nil {
		return NewQStringRefFromPointer(C.QXmlStreamReader_Name(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlStreamReader) NamespaceUri() *QStringRef {
	defer qt.Recovering("QXmlStreamReader::namespaceUri")

	if ptr.Pointer() != nil {
		return NewQStringRefFromPointer(C.QXmlStreamReader_NamespaceUri(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlStreamReader) Prefix() *QStringRef {
	defer qt.Recovering("QXmlStreamReader::prefix")

	if ptr.Pointer() != nil {
		return NewQStringRefFromPointer(C.QXmlStreamReader_Prefix(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlStreamReader) ProcessingInstructionData() *QStringRef {
	defer qt.Recovering("QXmlStreamReader::processingInstructionData")

	if ptr.Pointer() != nil {
		return NewQStringRefFromPointer(C.QXmlStreamReader_ProcessingInstructionData(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlStreamReader) ProcessingInstructionTarget() *QStringRef {
	defer qt.Recovering("QXmlStreamReader::processingInstructionTarget")

	if ptr.Pointer() != nil {
		return NewQStringRefFromPointer(C.QXmlStreamReader_ProcessingInstructionTarget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlStreamReader) QualifiedName() *QStringRef {
	defer qt.Recovering("QXmlStreamReader::qualifiedName")

	if ptr.Pointer() != nil {
		return NewQStringRefFromPointer(C.QXmlStreamReader_QualifiedName(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlStreamReader) RaiseError(message string) {
	defer qt.Recovering("QXmlStreamReader::raiseError")

	if ptr.Pointer() != nil {
		C.QXmlStreamReader_RaiseError(ptr.Pointer(), C.CString(message))
	}
}

func (ptr *QXmlStreamReader) ReadElementText(behaviour QXmlStreamReader__ReadElementTextBehaviour) string {
	defer qt.Recovering("QXmlStreamReader::readElementText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlStreamReader_ReadElementText(ptr.Pointer(), C.int(behaviour)))
	}
	return ""
}

func (ptr *QXmlStreamReader) ReadNext() QXmlStreamReader__TokenType {
	defer qt.Recovering("QXmlStreamReader::readNext")

	if ptr.Pointer() != nil {
		return QXmlStreamReader__TokenType(C.QXmlStreamReader_ReadNext(ptr.Pointer()))
	}
	return 0
}

func (ptr *QXmlStreamReader) ReadNextStartElement() bool {
	defer qt.Recovering("QXmlStreamReader::readNextStartElement")

	if ptr.Pointer() != nil {
		return C.QXmlStreamReader_ReadNextStartElement(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlStreamReader) SetDevice(device QIODevice_ITF) {
	defer qt.Recovering("QXmlStreamReader::setDevice")

	if ptr.Pointer() != nil {
		C.QXmlStreamReader_SetDevice(ptr.Pointer(), PointerFromQIODevice(device))
	}
}

func (ptr *QXmlStreamReader) SetEntityResolver(resolver QXmlStreamEntityResolver_ITF) {
	defer qt.Recovering("QXmlStreamReader::setEntityResolver")

	if ptr.Pointer() != nil {
		C.QXmlStreamReader_SetEntityResolver(ptr.Pointer(), PointerFromQXmlStreamEntityResolver(resolver))
	}
}

func (ptr *QXmlStreamReader) SkipCurrentElement() {
	defer qt.Recovering("QXmlStreamReader::skipCurrentElement")

	if ptr.Pointer() != nil {
		C.QXmlStreamReader_SkipCurrentElement(ptr.Pointer())
	}
}

func (ptr *QXmlStreamReader) Text() *QStringRef {
	defer qt.Recovering("QXmlStreamReader::text")

	if ptr.Pointer() != nil {
		return NewQStringRefFromPointer(C.QXmlStreamReader_Text(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlStreamReader) TokenString() string {
	defer qt.Recovering("QXmlStreamReader::tokenString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlStreamReader_TokenString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlStreamReader) TokenType() QXmlStreamReader__TokenType {
	defer qt.Recovering("QXmlStreamReader::tokenType")

	if ptr.Pointer() != nil {
		return QXmlStreamReader__TokenType(C.QXmlStreamReader_TokenType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QXmlStreamReader) DestroyQXmlStreamReader() {
	defer qt.Recovering("QXmlStreamReader::~QXmlStreamReader")

	if ptr.Pointer() != nil {
		C.QXmlStreamReader_DestroyQXmlStreamReader(ptr.Pointer())
	}
}
