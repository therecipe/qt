package multimedia

//#include "qmetadatareadercontrol.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QMetaDataReaderControl struct {
	QMediaControl
}

type QMetaDataReaderControl_ITF interface {
	QMediaControl_ITF
	QMetaDataReaderControl_PTR() *QMetaDataReaderControl
}

func PointerFromQMetaDataReaderControl(ptr QMetaDataReaderControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMetaDataReaderControl_PTR().Pointer()
	}
	return nil
}

func NewQMetaDataReaderControlFromPointer(ptr unsafe.Pointer) *QMetaDataReaderControl {
	var n = new(QMetaDataReaderControl)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QMetaDataReaderControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMetaDataReaderControl) QMetaDataReaderControl_PTR() *QMetaDataReaderControl {
	return ptr
}

func (ptr *QMetaDataReaderControl) AvailableMetaData() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QMetaDataReaderControl_AvailableMetaData(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QMetaDataReaderControl) IsMetaDataAvailable() bool {
	if ptr.Pointer() != nil {
		return C.QMetaDataReaderControl_IsMetaDataAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMetaDataReaderControl) MetaData(key string) *core.QVariant {
	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QMetaDataReaderControl_MetaData(ptr.Pointer(), C.CString(key)))
	}
	return nil
}

func (ptr *QMetaDataReaderControl) ConnectMetaDataAvailableChanged(f func(available bool)) {
	if ptr.Pointer() != nil {
		C.QMetaDataReaderControl_ConnectMetaDataAvailableChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "metaDataAvailableChanged", f)
	}
}

func (ptr *QMetaDataReaderControl) DisconnectMetaDataAvailableChanged() {
	if ptr.Pointer() != nil {
		C.QMetaDataReaderControl_DisconnectMetaDataAvailableChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "metaDataAvailableChanged")
	}
}

//export callbackQMetaDataReaderControlMetaDataAvailableChanged
func callbackQMetaDataReaderControlMetaDataAvailableChanged(ptrName *C.char, available C.int) {
	qt.GetSignal(C.GoString(ptrName), "metaDataAvailableChanged").(func(bool))(int(available) != 0)
}

func (ptr *QMetaDataReaderControl) ConnectMetaDataChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QMetaDataReaderControl_ConnectMetaDataChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "metaDataChanged", f)
	}
}

func (ptr *QMetaDataReaderControl) DisconnectMetaDataChanged() {
	if ptr.Pointer() != nil {
		C.QMetaDataReaderControl_DisconnectMetaDataChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "metaDataChanged")
	}
}

//export callbackQMetaDataReaderControlMetaDataChanged
func callbackQMetaDataReaderControlMetaDataChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "metaDataChanged").(func())()
}

func (ptr *QMetaDataReaderControl) DestroyQMetaDataReaderControl() {
	if ptr.Pointer() != nil {
		C.QMetaDataReaderControl_DestroyQMetaDataReaderControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
