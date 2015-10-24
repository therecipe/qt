package core

//#include "qsocketnotifier.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QSocketNotifier struct {
	QObject
}

type QSocketNotifierITF interface {
	QObjectITF
	QSocketNotifierPTR() *QSocketNotifier
}

func PointerFromQSocketNotifier(ptr QSocketNotifierITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSocketNotifierPTR().Pointer()
	}
	return nil
}

func QSocketNotifierFromPointer(ptr unsafe.Pointer) *QSocketNotifier {
	var n = new(QSocketNotifier)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QSocketNotifier_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSocketNotifier) QSocketNotifierPTR() *QSocketNotifier {
	return ptr
}

//QSocketNotifier::Type
type QSocketNotifier__Type int

var (
	QSocketNotifier__Read      = QSocketNotifier__Type(0)
	QSocketNotifier__Write     = QSocketNotifier__Type(1)
	QSocketNotifier__Exception = QSocketNotifier__Type(2)
)

func (ptr *QSocketNotifier) ConnectActivated(f func(socket int)) {
	if ptr.Pointer() != nil {
		C.QSocketNotifier_ConnectActivated(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "activated", f)
	}
}

func (ptr *QSocketNotifier) DisconnectActivated() {
	if ptr.Pointer() != nil {
		C.QSocketNotifier_DisconnectActivated(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "activated")
	}
}

//export callbackQSocketNotifierActivated
func callbackQSocketNotifierActivated(ptrName *C.char, socket C.int) {
	qt.GetSignal(C.GoString(ptrName), "activated").(func(int))(int(socket))
}

func (ptr *QSocketNotifier) IsEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QSocketNotifier_IsEnabled(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSocketNotifier) SetEnabled(enable bool) {
	if ptr.Pointer() != nil {
		C.QSocketNotifier_SetEnabled(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QSocketNotifier) Type() QSocketNotifier__Type {
	if ptr.Pointer() != nil {
		return QSocketNotifier__Type(C.QSocketNotifier_Type(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSocketNotifier) DestroyQSocketNotifier() {
	if ptr.Pointer() != nil {
		C.QSocketNotifier_DestroyQSocketNotifier(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
