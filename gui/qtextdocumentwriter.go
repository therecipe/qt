package gui

//#include "qtextdocumentwriter.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QTextDocumentWriter struct {
	ptr unsafe.Pointer
}

type QTextDocumentWriterITF interface {
	QTextDocumentWriterPTR() *QTextDocumentWriter
}

func (p *QTextDocumentWriter) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTextDocumentWriter) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTextDocumentWriter(ptr QTextDocumentWriterITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextDocumentWriterPTR().Pointer()
	}
	return nil
}

func QTextDocumentWriterFromPointer(ptr unsafe.Pointer) *QTextDocumentWriter {
	var n = new(QTextDocumentWriter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTextDocumentWriter) QTextDocumentWriterPTR() *QTextDocumentWriter {
	return ptr
}

func NewQTextDocumentWriter() *QTextDocumentWriter {
	return QTextDocumentWriterFromPointer(unsafe.Pointer(C.QTextDocumentWriter_NewQTextDocumentWriter()))
}

func NewQTextDocumentWriter2(device core.QIODeviceITF, format core.QByteArrayITF) *QTextDocumentWriter {
	return QTextDocumentWriterFromPointer(unsafe.Pointer(C.QTextDocumentWriter_NewQTextDocumentWriter2(C.QtObjectPtr(core.PointerFromQIODevice(device)), C.QtObjectPtr(core.PointerFromQByteArray(format)))))
}

func NewQTextDocumentWriter3(fileName string, format core.QByteArrayITF) *QTextDocumentWriter {
	return QTextDocumentWriterFromPointer(unsafe.Pointer(C.QTextDocumentWriter_NewQTextDocumentWriter3(C.CString(fileName), C.QtObjectPtr(core.PointerFromQByteArray(format)))))
}

func (ptr *QTextDocumentWriter) Codec() *core.QTextCodec {
	if ptr.Pointer() != nil {
		return core.QTextCodecFromPointer(unsafe.Pointer(C.QTextDocumentWriter_Codec(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QTextDocumentWriter) Device() *core.QIODevice {
	if ptr.Pointer() != nil {
		return core.QIODeviceFromPointer(unsafe.Pointer(C.QTextDocumentWriter_Device(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QTextDocumentWriter) FileName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTextDocumentWriter_FileName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QTextDocumentWriter) SetCodec(codec core.QTextCodecITF) {
	if ptr.Pointer() != nil {
		C.QTextDocumentWriter_SetCodec(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQTextCodec(codec)))
	}
}

func (ptr *QTextDocumentWriter) SetDevice(device core.QIODeviceITF) {
	if ptr.Pointer() != nil {
		C.QTextDocumentWriter_SetDevice(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQIODevice(device)))
	}
}

func (ptr *QTextDocumentWriter) SetFileName(fileName string) {
	if ptr.Pointer() != nil {
		C.QTextDocumentWriter_SetFileName(C.QtObjectPtr(ptr.Pointer()), C.CString(fileName))
	}
}

func (ptr *QTextDocumentWriter) SetFormat(format core.QByteArrayITF) {
	if ptr.Pointer() != nil {
		C.QTextDocumentWriter_SetFormat(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(format)))
	}
}

func (ptr *QTextDocumentWriter) Write(document QTextDocumentITF) bool {
	if ptr.Pointer() != nil {
		return C.QTextDocumentWriter_Write(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTextDocument(document))) != 0
	}
	return false
}

func (ptr *QTextDocumentWriter) Write2(fragment QTextDocumentFragmentITF) bool {
	if ptr.Pointer() != nil {
		return C.QTextDocumentWriter_Write2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTextDocumentFragment(fragment))) != 0
	}
	return false
}

func (ptr *QTextDocumentWriter) DestroyQTextDocumentWriter() {
	if ptr.Pointer() != nil {
		C.QTextDocumentWriter_DestroyQTextDocumentWriter(C.QtObjectPtr(ptr.Pointer()))
	}
}
