package multimedia

//#include "qmetadatawritercontrol.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QMetaDataWriterControl struct {
	QMediaControl
}

type QMetaDataWriterControl_ITF interface {
	QMediaControl_ITF
	QMetaDataWriterControl_PTR() *QMetaDataWriterControl
}

func PointerFromQMetaDataWriterControl(ptr QMetaDataWriterControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMetaDataWriterControl_PTR().Pointer()
	}
	return nil
}

func NewQMetaDataWriterControlFromPointer(ptr unsafe.Pointer) *QMetaDataWriterControl {
	var n = new(QMetaDataWriterControl)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QMetaDataWriterControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMetaDataWriterControl) QMetaDataWriterControl_PTR() *QMetaDataWriterControl {
	return ptr
}

func (ptr *QMetaDataWriterControl) AvailableMetaData() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QMetaDataWriterControl_AvailableMetaData(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QMetaDataWriterControl) IsMetaDataAvailable() bool {
	if ptr.Pointer() != nil {
		return C.QMetaDataWriterControl_IsMetaDataAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMetaDataWriterControl) IsWritable() bool {
	if ptr.Pointer() != nil {
		return C.QMetaDataWriterControl_IsWritable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMetaDataWriterControl) MetaData(key string) *core.QVariant {
	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QMetaDataWriterControl_MetaData(ptr.Pointer(), C.CString(key)))
	}
	return nil
}

func (ptr *QMetaDataWriterControl) ConnectMetaDataAvailableChanged(f func(available bool)) {
	if ptr.Pointer() != nil {
		C.QMetaDataWriterControl_ConnectMetaDataAvailableChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "metaDataAvailableChanged", f)
	}
}

func (ptr *QMetaDataWriterControl) DisconnectMetaDataAvailableChanged() {
	if ptr.Pointer() != nil {
		C.QMetaDataWriterControl_DisconnectMetaDataAvailableChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "metaDataAvailableChanged")
	}
}

//export callbackQMetaDataWriterControlMetaDataAvailableChanged
func callbackQMetaDataWriterControlMetaDataAvailableChanged(ptrName *C.char, available C.int) {
	qt.GetSignal(C.GoString(ptrName), "metaDataAvailableChanged").(func(bool))(int(available) != 0)
}

func (ptr *QMetaDataWriterControl) ConnectMetaDataChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QMetaDataWriterControl_ConnectMetaDataChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "metaDataChanged", f)
	}
}

func (ptr *QMetaDataWriterControl) DisconnectMetaDataChanged() {
	if ptr.Pointer() != nil {
		C.QMetaDataWriterControl_DisconnectMetaDataChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "metaDataChanged")
	}
}

//export callbackQMetaDataWriterControlMetaDataChanged
func callbackQMetaDataWriterControlMetaDataChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "metaDataChanged").(func())()
}

func (ptr *QMetaDataWriterControl) SetMetaData(key string, value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QMetaDataWriterControl_SetMetaData(ptr.Pointer(), C.CString(key), core.PointerFromQVariant(value))
	}
}

func (ptr *QMetaDataWriterControl) ConnectWritableChanged(f func(writable bool)) {
	if ptr.Pointer() != nil {
		C.QMetaDataWriterControl_ConnectWritableChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "writableChanged", f)
	}
}

func (ptr *QMetaDataWriterControl) DisconnectWritableChanged() {
	if ptr.Pointer() != nil {
		C.QMetaDataWriterControl_DisconnectWritableChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "writableChanged")
	}
}

//export callbackQMetaDataWriterControlWritableChanged
func callbackQMetaDataWriterControlWritableChanged(ptrName *C.char, writable C.int) {
	qt.GetSignal(C.GoString(ptrName), "writableChanged").(func(bool))(int(writable) != 0)
}

func (ptr *QMetaDataWriterControl) DestroyQMetaDataWriterControl() {
	if ptr.Pointer() != nil {
		C.QMetaDataWriterControl_DestroyQMetaDataWriterControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
