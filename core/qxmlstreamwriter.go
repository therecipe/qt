package core

//#include "qxmlstreamwriter.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QXmlStreamWriter struct {
	ptr unsafe.Pointer
}

type QXmlStreamWriterITF interface {
	QXmlStreamWriterPTR() *QXmlStreamWriter
}

func (p *QXmlStreamWriter) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlStreamWriter) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlStreamWriter(ptr QXmlStreamWriterITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlStreamWriterPTR().Pointer()
	}
	return nil
}

func QXmlStreamWriterFromPointer(ptr unsafe.Pointer) *QXmlStreamWriter {
	var n = new(QXmlStreamWriter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlStreamWriter) QXmlStreamWriterPTR() *QXmlStreamWriter {
	return ptr
}

func (ptr *QXmlStreamWriter) AutoFormattingIndent() int {
	if ptr.Pointer() != nil {
		return int(C.QXmlStreamWriter_AutoFormattingIndent(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QXmlStreamWriter) SetAutoFormattingIndent(spacesOrTabs int) {
	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_SetAutoFormattingIndent(C.QtObjectPtr(ptr.Pointer()), C.int(spacesOrTabs))
	}
}

func (ptr *QXmlStreamWriter) AutoFormatting() bool {
	if ptr.Pointer() != nil {
		return C.QXmlStreamWriter_AutoFormatting(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QXmlStreamWriter) Codec() *QTextCodec {
	if ptr.Pointer() != nil {
		return QTextCodecFromPointer(unsafe.Pointer(C.QXmlStreamWriter_Codec(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QXmlStreamWriter) Device() *QIODevice {
	if ptr.Pointer() != nil {
		return QIODeviceFromPointer(unsafe.Pointer(C.QXmlStreamWriter_Device(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QXmlStreamWriter) HasError() bool {
	if ptr.Pointer() != nil {
		return C.QXmlStreamWriter_HasError(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QXmlStreamWriter) SetAutoFormatting(enable bool) {
	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_SetAutoFormatting(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QXmlStreamWriter) SetCodec(codec QTextCodecITF) {
	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_SetCodec(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTextCodec(codec)))
	}
}

func (ptr *QXmlStreamWriter) SetCodec2(codecName string) {
	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_SetCodec2(C.QtObjectPtr(ptr.Pointer()), C.CString(codecName))
	}
}

func (ptr *QXmlStreamWriter) SetDevice(device QIODeviceITF) {
	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_SetDevice(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQIODevice(device)))
	}
}

func (ptr *QXmlStreamWriter) WriteAttribute(namespaceUri string, name string, value string) {
	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_WriteAttribute(C.QtObjectPtr(ptr.Pointer()), C.CString(namespaceUri), C.CString(name), C.CString(value))
	}
}

func (ptr *QXmlStreamWriter) WriteAttribute2(qualifiedName string, value string) {
	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_WriteAttribute2(C.QtObjectPtr(ptr.Pointer()), C.CString(qualifiedName), C.CString(value))
	}
}

func (ptr *QXmlStreamWriter) WriteAttribute3(attribute QXmlStreamAttributeITF) {
	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_WriteAttribute3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlStreamAttribute(attribute)))
	}
}

func (ptr *QXmlStreamWriter) WriteAttributes(attributes QXmlStreamAttributesITF) {
	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_WriteAttributes(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlStreamAttributes(attributes)))
	}
}

func (ptr *QXmlStreamWriter) WriteCDATA(text string) {
	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_WriteCDATA(C.QtObjectPtr(ptr.Pointer()), C.CString(text))
	}
}

func (ptr *QXmlStreamWriter) WriteCharacters(text string) {
	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_WriteCharacters(C.QtObjectPtr(ptr.Pointer()), C.CString(text))
	}
}

func (ptr *QXmlStreamWriter) WriteComment(text string) {
	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_WriteComment(C.QtObjectPtr(ptr.Pointer()), C.CString(text))
	}
}

func (ptr *QXmlStreamWriter) WriteCurrentToken(reader QXmlStreamReaderITF) {
	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_WriteCurrentToken(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlStreamReader(reader)))
	}
}

func (ptr *QXmlStreamWriter) WriteDTD(dtd string) {
	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_WriteDTD(C.QtObjectPtr(ptr.Pointer()), C.CString(dtd))
	}
}

func (ptr *QXmlStreamWriter) WriteDefaultNamespace(namespaceUri string) {
	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_WriteDefaultNamespace(C.QtObjectPtr(ptr.Pointer()), C.CString(namespaceUri))
	}
}

func (ptr *QXmlStreamWriter) WriteEmptyElement(namespaceUri string, name string) {
	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_WriteEmptyElement(C.QtObjectPtr(ptr.Pointer()), C.CString(namespaceUri), C.CString(name))
	}
}

func (ptr *QXmlStreamWriter) WriteEmptyElement2(qualifiedName string) {
	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_WriteEmptyElement2(C.QtObjectPtr(ptr.Pointer()), C.CString(qualifiedName))
	}
}

func (ptr *QXmlStreamWriter) WriteEndDocument() {
	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_WriteEndDocument(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QXmlStreamWriter) WriteEndElement() {
	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_WriteEndElement(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QXmlStreamWriter) WriteEntityReference(name string) {
	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_WriteEntityReference(C.QtObjectPtr(ptr.Pointer()), C.CString(name))
	}
}

func (ptr *QXmlStreamWriter) WriteNamespace(namespaceUri string, prefix string) {
	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_WriteNamespace(C.QtObjectPtr(ptr.Pointer()), C.CString(namespaceUri), C.CString(prefix))
	}
}

func (ptr *QXmlStreamWriter) WriteProcessingInstruction(target string, data string) {
	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_WriteProcessingInstruction(C.QtObjectPtr(ptr.Pointer()), C.CString(target), C.CString(data))
	}
}

func (ptr *QXmlStreamWriter) WriteStartDocument3() {
	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_WriteStartDocument3(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QXmlStreamWriter) WriteStartDocument(version string) {
	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_WriteStartDocument(C.QtObjectPtr(ptr.Pointer()), C.CString(version))
	}
}

func (ptr *QXmlStreamWriter) WriteStartDocument2(version string, standalone bool) {
	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_WriteStartDocument2(C.QtObjectPtr(ptr.Pointer()), C.CString(version), C.int(qt.GoBoolToInt(standalone)))
	}
}

func (ptr *QXmlStreamWriter) WriteStartElement(namespaceUri string, name string) {
	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_WriteStartElement(C.QtObjectPtr(ptr.Pointer()), C.CString(namespaceUri), C.CString(name))
	}
}

func (ptr *QXmlStreamWriter) WriteStartElement2(qualifiedName string) {
	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_WriteStartElement2(C.QtObjectPtr(ptr.Pointer()), C.CString(qualifiedName))
	}
}

func (ptr *QXmlStreamWriter) WriteTextElement(namespaceUri string, name string, text string) {
	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_WriteTextElement(C.QtObjectPtr(ptr.Pointer()), C.CString(namespaceUri), C.CString(name), C.CString(text))
	}
}

func (ptr *QXmlStreamWriter) WriteTextElement2(qualifiedName string, text string) {
	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_WriteTextElement2(C.QtObjectPtr(ptr.Pointer()), C.CString(qualifiedName), C.CString(text))
	}
}

func (ptr *QXmlStreamWriter) DestroyQXmlStreamWriter() {
	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_DestroyQXmlStreamWriter(C.QtObjectPtr(ptr.Pointer()))
	}
}
