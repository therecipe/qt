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

type QImageWriter_ITF interface {
	QImageWriter_PTR() *QImageWriter
}

func (p *QImageWriter) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QImageWriter) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQImageWriter(ptr QImageWriter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QImageWriter_PTR().Pointer()
	}
	return nil
}

func NewQImageWriterFromPointer(ptr unsafe.Pointer) *QImageWriter {
	var n = new(QImageWriter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QImageWriter) QImageWriter_PTR() *QImageWriter {
	return ptr
}

//QImageWriter::ImageWriterError
type QImageWriter__ImageWriterError int64

const (
	QImageWriter__UnknownError           = QImageWriter__ImageWriterError(0)
	QImageWriter__DeviceError            = QImageWriter__ImageWriterError(1)
	QImageWriter__UnsupportedFormatError = QImageWriter__ImageWriterError(2)
)

func NewQImageWriter() *QImageWriter {
	return NewQImageWriterFromPointer(C.QImageWriter_NewQImageWriter())
}

func NewQImageWriter2(device core.QIODevice_ITF, format core.QByteArray_ITF) *QImageWriter {
	return NewQImageWriterFromPointer(C.QImageWriter_NewQImageWriter2(core.PointerFromQIODevice(device), core.PointerFromQByteArray(format)))
}

func NewQImageWriter3(fileName string, format core.QByteArray_ITF) *QImageWriter {
	return NewQImageWriterFromPointer(C.QImageWriter_NewQImageWriter3(C.CString(fileName), core.PointerFromQByteArray(format)))
}

func (ptr *QImageWriter) CanWrite() bool {
	if ptr.Pointer() != nil {
		return C.QImageWriter_CanWrite(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QImageWriter) Compression() int {
	if ptr.Pointer() != nil {
		return int(C.QImageWriter_Compression(ptr.Pointer()))
	}
	return 0
}

func (ptr *QImageWriter) Device() *core.QIODevice {
	if ptr.Pointer() != nil {
		return core.NewQIODeviceFromPointer(C.QImageWriter_Device(ptr.Pointer()))
	}
	return nil
}

func (ptr *QImageWriter) Error() QImageWriter__ImageWriterError {
	if ptr.Pointer() != nil {
		return QImageWriter__ImageWriterError(C.QImageWriter_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QImageWriter) ErrorString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QImageWriter_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QImageWriter) FileName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QImageWriter_FileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QImageWriter) Format() *core.QByteArray {
	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QImageWriter_Format(ptr.Pointer()))
	}
	return nil
}

func (ptr *QImageWriter) OptimizedWrite() bool {
	if ptr.Pointer() != nil {
		return C.QImageWriter_OptimizedWrite(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QImageWriter) ProgressiveScanWrite() bool {
	if ptr.Pointer() != nil {
		return C.QImageWriter_ProgressiveScanWrite(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QImageWriter) Quality() int {
	if ptr.Pointer() != nil {
		return int(C.QImageWriter_Quality(ptr.Pointer()))
	}
	return 0
}

func (ptr *QImageWriter) SetCompression(compression int) {
	if ptr.Pointer() != nil {
		C.QImageWriter_SetCompression(ptr.Pointer(), C.int(compression))
	}
}

func (ptr *QImageWriter) SetDevice(device core.QIODevice_ITF) {
	if ptr.Pointer() != nil {
		C.QImageWriter_SetDevice(ptr.Pointer(), core.PointerFromQIODevice(device))
	}
}

func (ptr *QImageWriter) SetFileName(fileName string) {
	if ptr.Pointer() != nil {
		C.QImageWriter_SetFileName(ptr.Pointer(), C.CString(fileName))
	}
}

func (ptr *QImageWriter) SetFormat(format core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QImageWriter_SetFormat(ptr.Pointer(), core.PointerFromQByteArray(format))
	}
}

func (ptr *QImageWriter) SetOptimizedWrite(optimize bool) {
	if ptr.Pointer() != nil {
		C.QImageWriter_SetOptimizedWrite(ptr.Pointer(), C.int(qt.GoBoolToInt(optimize)))
	}
}

func (ptr *QImageWriter) SetProgressiveScanWrite(progressive bool) {
	if ptr.Pointer() != nil {
		C.QImageWriter_SetProgressiveScanWrite(ptr.Pointer(), C.int(qt.GoBoolToInt(progressive)))
	}
}

func (ptr *QImageWriter) SetQuality(quality int) {
	if ptr.Pointer() != nil {
		C.QImageWriter_SetQuality(ptr.Pointer(), C.int(quality))
	}
}

func (ptr *QImageWriter) SetSubType(ty core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QImageWriter_SetSubType(ptr.Pointer(), core.PointerFromQByteArray(ty))
	}
}

func (ptr *QImageWriter) SetText(key string, text string) {
	if ptr.Pointer() != nil {
		C.QImageWriter_SetText(ptr.Pointer(), C.CString(key), C.CString(text))
	}
}

func (ptr *QImageWriter) SetTransformation(transform QImageIOHandler__Transformation) {
	if ptr.Pointer() != nil {
		C.QImageWriter_SetTransformation(ptr.Pointer(), C.int(transform))
	}
}

func (ptr *QImageWriter) SubType() *core.QByteArray {
	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QImageWriter_SubType(ptr.Pointer()))
	}
	return nil
}

func (ptr *QImageWriter) SupportsOption(option QImageIOHandler__ImageOption) bool {
	if ptr.Pointer() != nil {
		return C.QImageWriter_SupportsOption(ptr.Pointer(), C.int(option)) != 0
	}
	return false
}

func (ptr *QImageWriter) Transformation() QImageIOHandler__Transformation {
	if ptr.Pointer() != nil {
		return QImageIOHandler__Transformation(C.QImageWriter_Transformation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QImageWriter) Write(image QImage_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QImageWriter_Write(ptr.Pointer(), PointerFromQImage(image)) != 0
	}
	return false
}

func (ptr *QImageWriter) DestroyQImageWriter() {
	if ptr.Pointer() != nil {
		C.QImageWriter_DestroyQImageWriter(ptr.Pointer())
	}
}
