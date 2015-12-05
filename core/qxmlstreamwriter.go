package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
	"unsafe"
)

type QXmlStreamWriter struct {
	ptr unsafe.Pointer
}

type QXmlStreamWriter_ITF interface {
	QXmlStreamWriter_PTR() *QXmlStreamWriter
}

func (p *QXmlStreamWriter) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlStreamWriter) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlStreamWriter(ptr QXmlStreamWriter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlStreamWriter_PTR().Pointer()
	}
	return nil
}

func NewQXmlStreamWriterFromPointer(ptr unsafe.Pointer) *QXmlStreamWriter {
	var n = new(QXmlStreamWriter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlStreamWriter) QXmlStreamWriter_PTR() *QXmlStreamWriter {
	return ptr
}

func (ptr *QXmlStreamWriter) AutoFormattingIndent() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlStreamWriter::autoFormattingIndent")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QXmlStreamWriter_AutoFormattingIndent(ptr.Pointer()))
	}
	return 0
}

func (ptr *QXmlStreamWriter) SetAutoFormattingIndent(spacesOrTabs int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlStreamWriter::setAutoFormattingIndent")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_SetAutoFormattingIndent(ptr.Pointer(), C.int(spacesOrTabs))
	}
}

func (ptr *QXmlStreamWriter) AutoFormatting() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlStreamWriter::autoFormatting")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QXmlStreamWriter_AutoFormatting(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlStreamWriter) Codec() *QTextCodec {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlStreamWriter::codec")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQTextCodecFromPointer(C.QXmlStreamWriter_Codec(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlStreamWriter) Device() *QIODevice {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlStreamWriter::device")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQIODeviceFromPointer(C.QXmlStreamWriter_Device(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlStreamWriter) HasError() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlStreamWriter::hasError")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QXmlStreamWriter_HasError(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlStreamWriter) SetAutoFormatting(enable bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlStreamWriter::setAutoFormatting")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_SetAutoFormatting(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QXmlStreamWriter) SetCodec(codec QTextCodec_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlStreamWriter::setCodec")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_SetCodec(ptr.Pointer(), PointerFromQTextCodec(codec))
	}
}

func (ptr *QXmlStreamWriter) SetCodec2(codecName string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlStreamWriter::setCodec")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_SetCodec2(ptr.Pointer(), C.CString(codecName))
	}
}

func (ptr *QXmlStreamWriter) SetDevice(device QIODevice_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlStreamWriter::setDevice")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_SetDevice(ptr.Pointer(), PointerFromQIODevice(device))
	}
}

func (ptr *QXmlStreamWriter) WriteAttribute(namespaceUri string, name string, value string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlStreamWriter::writeAttribute")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_WriteAttribute(ptr.Pointer(), C.CString(namespaceUri), C.CString(name), C.CString(value))
	}
}

func (ptr *QXmlStreamWriter) WriteAttribute2(qualifiedName string, value string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlStreamWriter::writeAttribute")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_WriteAttribute2(ptr.Pointer(), C.CString(qualifiedName), C.CString(value))
	}
}

func (ptr *QXmlStreamWriter) WriteAttribute3(attribute QXmlStreamAttribute_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlStreamWriter::writeAttribute")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_WriteAttribute3(ptr.Pointer(), PointerFromQXmlStreamAttribute(attribute))
	}
}

func (ptr *QXmlStreamWriter) WriteAttributes(attributes QXmlStreamAttributes_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlStreamWriter::writeAttributes")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_WriteAttributes(ptr.Pointer(), PointerFromQXmlStreamAttributes(attributes))
	}
}

func (ptr *QXmlStreamWriter) WriteCDATA(text string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlStreamWriter::writeCDATA")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_WriteCDATA(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QXmlStreamWriter) WriteCharacters(text string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlStreamWriter::writeCharacters")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_WriteCharacters(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QXmlStreamWriter) WriteComment(text string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlStreamWriter::writeComment")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_WriteComment(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QXmlStreamWriter) WriteCurrentToken(reader QXmlStreamReader_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlStreamWriter::writeCurrentToken")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_WriteCurrentToken(ptr.Pointer(), PointerFromQXmlStreamReader(reader))
	}
}

func (ptr *QXmlStreamWriter) WriteDTD(dtd string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlStreamWriter::writeDTD")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_WriteDTD(ptr.Pointer(), C.CString(dtd))
	}
}

func (ptr *QXmlStreamWriter) WriteDefaultNamespace(namespaceUri string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlStreamWriter::writeDefaultNamespace")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_WriteDefaultNamespace(ptr.Pointer(), C.CString(namespaceUri))
	}
}

func (ptr *QXmlStreamWriter) WriteEmptyElement(namespaceUri string, name string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlStreamWriter::writeEmptyElement")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_WriteEmptyElement(ptr.Pointer(), C.CString(namespaceUri), C.CString(name))
	}
}

func (ptr *QXmlStreamWriter) WriteEmptyElement2(qualifiedName string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlStreamWriter::writeEmptyElement")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_WriteEmptyElement2(ptr.Pointer(), C.CString(qualifiedName))
	}
}

func (ptr *QXmlStreamWriter) WriteEndDocument() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlStreamWriter::writeEndDocument")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_WriteEndDocument(ptr.Pointer())
	}
}

func (ptr *QXmlStreamWriter) WriteEndElement() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlStreamWriter::writeEndElement")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_WriteEndElement(ptr.Pointer())
	}
}

func (ptr *QXmlStreamWriter) WriteEntityReference(name string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlStreamWriter::writeEntityReference")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_WriteEntityReference(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QXmlStreamWriter) WriteNamespace(namespaceUri string, prefix string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlStreamWriter::writeNamespace")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_WriteNamespace(ptr.Pointer(), C.CString(namespaceUri), C.CString(prefix))
	}
}

func (ptr *QXmlStreamWriter) WriteProcessingInstruction(target string, data string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlStreamWriter::writeProcessingInstruction")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_WriteProcessingInstruction(ptr.Pointer(), C.CString(target), C.CString(data))
	}
}

func (ptr *QXmlStreamWriter) WriteStartDocument3() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlStreamWriter::writeStartDocument")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_WriteStartDocument3(ptr.Pointer())
	}
}

func (ptr *QXmlStreamWriter) WriteStartDocument(version string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlStreamWriter::writeStartDocument")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_WriteStartDocument(ptr.Pointer(), C.CString(version))
	}
}

func (ptr *QXmlStreamWriter) WriteStartDocument2(version string, standalone bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlStreamWriter::writeStartDocument")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_WriteStartDocument2(ptr.Pointer(), C.CString(version), C.int(qt.GoBoolToInt(standalone)))
	}
}

func (ptr *QXmlStreamWriter) WriteStartElement(namespaceUri string, name string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlStreamWriter::writeStartElement")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_WriteStartElement(ptr.Pointer(), C.CString(namespaceUri), C.CString(name))
	}
}

func (ptr *QXmlStreamWriter) WriteStartElement2(qualifiedName string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlStreamWriter::writeStartElement")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_WriteStartElement2(ptr.Pointer(), C.CString(qualifiedName))
	}
}

func (ptr *QXmlStreamWriter) WriteTextElement(namespaceUri string, name string, text string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlStreamWriter::writeTextElement")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_WriteTextElement(ptr.Pointer(), C.CString(namespaceUri), C.CString(name), C.CString(text))
	}
}

func (ptr *QXmlStreamWriter) WriteTextElement2(qualifiedName string, text string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlStreamWriter::writeTextElement")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_WriteTextElement2(ptr.Pointer(), C.CString(qualifiedName), C.CString(text))
	}
}

func (ptr *QXmlStreamWriter) DestroyQXmlStreamWriter() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlStreamWriter::~QXmlStreamWriter")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlStreamWriter_DestroyQXmlStreamWriter(ptr.Pointer())
	}
}
