package gui

//#include "qimagewriter.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QImageWriter struct {
	ptr unsafe.Pointer
}

type QImageWriterITF interface {
	QImageWriterPTR() *QImageWriter
}

func (p *QImageWriter) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QImageWriter) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQImageWriter(ptr QImageWriterITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QImageWriterPTR().Pointer()
	}
	return nil
}

func QImageWriterFromPointer(ptr unsafe.Pointer) *QImageWriter {
	var n = new(QImageWriter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QImageWriter) QImageWriterPTR() *QImageWriter {
	return ptr
}

//QImageWriter::ImageWriterError
type QImageWriter__ImageWriterError int

var (
	QImageWriter__UnknownError           = QImageWriter__ImageWriterError(0)
	QImageWriter__DeviceError            = QImageWriter__ImageWriterError(1)
	QImageWriter__UnsupportedFormatError = QImageWriter__ImageWriterError(2)
)

func NewQImageWriter() *QImageWriter {
	return QImageWriterFromPointer(unsafe.Pointer(C.QImageWriter_NewQImageWriter()))
}

func NewQImageWriter2(device core.QIODeviceITF, format core.QByteArrayITF) *QImageWriter {
	return QImageWriterFromPointer(unsafe.Pointer(C.QImageWriter_NewQImageWriter2(C.QtObjectPtr(core.PointerFromQIODevice(device)), C.QtObjectPtr(core.PointerFromQByteArray(format)))))
}

func NewQImageWriter3(fileName string, format core.QByteArrayITF) *QImageWriter {
	return QImageWriterFromPointer(unsafe.Pointer(C.QImageWriter_NewQImageWriter3(C.CString(fileName), C.QtObjectPtr(core.PointerFromQByteArray(format)))))
}

func (ptr *QImageWriter) CanWrite() bool {
	if ptr.Pointer() != nil {
		return C.QImageWriter_CanWrite(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QImageWriter) Compression() int {
	if ptr.Pointer() != nil {
		return int(C.QImageWriter_Compression(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QImageWriter) Device() *core.QIODevice {
	if ptr.Pointer() != nil {
		return core.QIODeviceFromPointer(unsafe.Pointer(C.QImageWriter_Device(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QImageWriter) Error() QImageWriter__ImageWriterError {
	if ptr.Pointer() != nil {
		return QImageWriter__ImageWriterError(C.QImageWriter_Error(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QImageWriter) ErrorString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QImageWriter_ErrorString(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QImageWriter) FileName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QImageWriter_FileName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QImageWriter) OptimizedWrite() bool {
	if ptr.Pointer() != nil {
		return C.QImageWriter_OptimizedWrite(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QImageWriter) ProgressiveScanWrite() bool {
	if ptr.Pointer() != nil {
		return C.QImageWriter_ProgressiveScanWrite(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QImageWriter) Quality() int {
	if ptr.Pointer() != nil {
		return int(C.QImageWriter_Quality(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QImageWriter) SetCompression(compression int) {
	if ptr.Pointer() != nil {
		C.QImageWriter_SetCompression(C.QtObjectPtr(ptr.Pointer()), C.int(compression))
	}
}

func (ptr *QImageWriter) SetDevice(device core.QIODeviceITF) {
	if ptr.Pointer() != nil {
		C.QImageWriter_SetDevice(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQIODevice(device)))
	}
}

func (ptr *QImageWriter) SetFileName(fileName string) {
	if ptr.Pointer() != nil {
		C.QImageWriter_SetFileName(C.QtObjectPtr(ptr.Pointer()), C.CString(fileName))
	}
}

func (ptr *QImageWriter) SetFormat(format core.QByteArrayITF) {
	if ptr.Pointer() != nil {
		C.QImageWriter_SetFormat(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(format)))
	}
}

func (ptr *QImageWriter) SetOptimizedWrite(optimize bool) {
	if ptr.Pointer() != nil {
		C.QImageWriter_SetOptimizedWrite(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(optimize)))
	}
}

func (ptr *QImageWriter) SetProgressiveScanWrite(progressive bool) {
	if ptr.Pointer() != nil {
		C.QImageWriter_SetProgressiveScanWrite(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(progressive)))
	}
}

func (ptr *QImageWriter) SetQuality(quality int) {
	if ptr.Pointer() != nil {
		C.QImageWriter_SetQuality(C.QtObjectPtr(ptr.Pointer()), C.int(quality))
	}
}

func (ptr *QImageWriter) SetSubType(ty core.QByteArrayITF) {
	if ptr.Pointer() != nil {
		C.QImageWriter_SetSubType(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(ty)))
	}
}

func (ptr *QImageWriter) SetText(key string, text string) {
	if ptr.Pointer() != nil {
		C.QImageWriter_SetText(C.QtObjectPtr(ptr.Pointer()), C.CString(key), C.CString(text))
	}
}

func (ptr *QImageWriter) SetTransformation(transform QImageIOHandler__Transformation) {
	if ptr.Pointer() != nil {
		C.QImageWriter_SetTransformation(C.QtObjectPtr(ptr.Pointer()), C.int(transform))
	}
}

func (ptr *QImageWriter) SupportsOption(option QImageIOHandler__ImageOption) bool {
	if ptr.Pointer() != nil {
		return C.QImageWriter_SupportsOption(C.QtObjectPtr(ptr.Pointer()), C.int(option)) != 0
	}
	return false
}

func (ptr *QImageWriter) Transformation() QImageIOHandler__Transformation {
	if ptr.Pointer() != nil {
		return QImageIOHandler__Transformation(C.QImageWriter_Transformation(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QImageWriter) Write(image QImageITF) bool {
	if ptr.Pointer() != nil {
		return C.QImageWriter_Write(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQImage(image))) != 0
	}
	return false
}

func (ptr *QImageWriter) DestroyQImageWriter() {
	if ptr.Pointer() != nil {
		C.QImageWriter_DestroyQImageWriter(C.QtObjectPtr(ptr.Pointer()))
	}
}
