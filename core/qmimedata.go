package core

//#include "qmimedata.h"
import "C"
import (
	"github.com/therecipe/qt"
	"strings"
	"unsafe"
)

type QMimeData struct {
	QObject
}

type QMimeDataITF interface {
	QObjectITF
	QMimeDataPTR() *QMimeData
}

func PointerFromQMimeData(ptr QMimeDataITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMimeDataPTR().Pointer()
	}
	return nil
}

func QMimeDataFromPointer(ptr unsafe.Pointer) *QMimeData {
	var n = new(QMimeData)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QMimeData_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMimeData) QMimeDataPTR() *QMimeData {
	return ptr
}

func NewQMimeData() *QMimeData {
	return QMimeDataFromPointer(unsafe.Pointer(C.QMimeData_NewQMimeData()))
}

func (ptr *QMimeData) Clear() {
	if ptr.Pointer() != nil {
		C.QMimeData_Clear(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QMimeData) ColorData() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMimeData_ColorData(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QMimeData) Formats() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QMimeData_Formats(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QMimeData) HasColor() bool {
	if ptr.Pointer() != nil {
		return C.QMimeData_HasColor(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMimeData) HasFormat(mimeType string) bool {
	if ptr.Pointer() != nil {
		return C.QMimeData_HasFormat(C.QtObjectPtr(ptr.Pointer()), C.CString(mimeType)) != 0
	}
	return false
}

func (ptr *QMimeData) HasHtml() bool {
	if ptr.Pointer() != nil {
		return C.QMimeData_HasHtml(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMimeData) HasImage() bool {
	if ptr.Pointer() != nil {
		return C.QMimeData_HasImage(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMimeData) HasText() bool {
	if ptr.Pointer() != nil {
		return C.QMimeData_HasText(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMimeData) HasUrls() bool {
	if ptr.Pointer() != nil {
		return C.QMimeData_HasUrls(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMimeData) Html() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMimeData_Html(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QMimeData) ImageData() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMimeData_ImageData(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QMimeData) RemoveFormat(mimeType string) {
	if ptr.Pointer() != nil {
		C.QMimeData_RemoveFormat(C.QtObjectPtr(ptr.Pointer()), C.CString(mimeType))
	}
}

func (ptr *QMimeData) SetColorData(color string) {
	if ptr.Pointer() != nil {
		C.QMimeData_SetColorData(C.QtObjectPtr(ptr.Pointer()), C.CString(color))
	}
}

func (ptr *QMimeData) SetData(mimeType string, data QByteArrayITF) {
	if ptr.Pointer() != nil {
		C.QMimeData_SetData(C.QtObjectPtr(ptr.Pointer()), C.CString(mimeType), C.QtObjectPtr(PointerFromQByteArray(data)))
	}
}

func (ptr *QMimeData) SetHtml(html string) {
	if ptr.Pointer() != nil {
		C.QMimeData_SetHtml(C.QtObjectPtr(ptr.Pointer()), C.CString(html))
	}
}

func (ptr *QMimeData) SetImageData(image string) {
	if ptr.Pointer() != nil {
		C.QMimeData_SetImageData(C.QtObjectPtr(ptr.Pointer()), C.CString(image))
	}
}

func (ptr *QMimeData) SetText(text string) {
	if ptr.Pointer() != nil {
		C.QMimeData_SetText(C.QtObjectPtr(ptr.Pointer()), C.CString(text))
	}
}

func (ptr *QMimeData) Text() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMimeData_Text(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QMimeData) DestroyQMimeData() {
	if ptr.Pointer() != nil {
		C.QMimeData_DestroyQMimeData(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
