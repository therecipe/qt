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

type QMimeData_ITF interface {
	QObject_ITF
	QMimeData_PTR() *QMimeData
}

func PointerFromQMimeData(ptr QMimeData_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMimeData_PTR().Pointer()
	}
	return nil
}

func NewQMimeDataFromPointer(ptr unsafe.Pointer) *QMimeData {
	var n = new(QMimeData)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QMimeData_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMimeData) QMimeData_PTR() *QMimeData {
	return ptr
}

func NewQMimeData() *QMimeData {
	return NewQMimeDataFromPointer(C.QMimeData_NewQMimeData())
}

func (ptr *QMimeData) Clear() {
	if ptr.Pointer() != nil {
		C.QMimeData_Clear(ptr.Pointer())
	}
}

func (ptr *QMimeData) ColorData() *QVariant {
	if ptr.Pointer() != nil {
		return NewQVariantFromPointer(C.QMimeData_ColorData(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMimeData) Data(mimeType string) *QByteArray {
	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QMimeData_Data(ptr.Pointer(), C.CString(mimeType)))
	}
	return nil
}

func (ptr *QMimeData) Formats() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QMimeData_Formats(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QMimeData) HasColor() bool {
	if ptr.Pointer() != nil {
		return C.QMimeData_HasColor(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMimeData) HasFormat(mimeType string) bool {
	if ptr.Pointer() != nil {
		return C.QMimeData_HasFormat(ptr.Pointer(), C.CString(mimeType)) != 0
	}
	return false
}

func (ptr *QMimeData) HasHtml() bool {
	if ptr.Pointer() != nil {
		return C.QMimeData_HasHtml(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMimeData) HasImage() bool {
	if ptr.Pointer() != nil {
		return C.QMimeData_HasImage(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMimeData) HasText() bool {
	if ptr.Pointer() != nil {
		return C.QMimeData_HasText(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMimeData) HasUrls() bool {
	if ptr.Pointer() != nil {
		return C.QMimeData_HasUrls(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMimeData) Html() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMimeData_Html(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMimeData) ImageData() *QVariant {
	if ptr.Pointer() != nil {
		return NewQVariantFromPointer(C.QMimeData_ImageData(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMimeData) RemoveFormat(mimeType string) {
	if ptr.Pointer() != nil {
		C.QMimeData_RemoveFormat(ptr.Pointer(), C.CString(mimeType))
	}
}

func (ptr *QMimeData) SetColorData(color QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QMimeData_SetColorData(ptr.Pointer(), PointerFromQVariant(color))
	}
}

func (ptr *QMimeData) SetData(mimeType string, data QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QMimeData_SetData(ptr.Pointer(), C.CString(mimeType), PointerFromQByteArray(data))
	}
}

func (ptr *QMimeData) SetHtml(html string) {
	if ptr.Pointer() != nil {
		C.QMimeData_SetHtml(ptr.Pointer(), C.CString(html))
	}
}

func (ptr *QMimeData) SetImageData(image QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QMimeData_SetImageData(ptr.Pointer(), PointerFromQVariant(image))
	}
}

func (ptr *QMimeData) SetText(text string) {
	if ptr.Pointer() != nil {
		C.QMimeData_SetText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QMimeData) Text() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMimeData_Text(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMimeData) DestroyQMimeData() {
	if ptr.Pointer() != nil {
		C.QMimeData_DestroyQMimeData(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
