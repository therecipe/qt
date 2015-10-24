package nfc

//#include "qnearfieldtarget.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QNearFieldTarget struct {
	core.QObject
}

type QNearFieldTargetITF interface {
	core.QObjectITF
	QNearFieldTargetPTR() *QNearFieldTarget
}

func PointerFromQNearFieldTarget(ptr QNearFieldTargetITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNearFieldTargetPTR().Pointer()
	}
	return nil
}

func QNearFieldTargetFromPointer(ptr unsafe.Pointer) *QNearFieldTarget {
	var n = new(QNearFieldTarget)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QNearFieldTarget_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QNearFieldTarget) QNearFieldTargetPTR() *QNearFieldTarget {
	return ptr
}

//QNearFieldTarget::AccessMethod
type QNearFieldTarget__AccessMethod int

var (
	QNearFieldTarget__UnknownAccess         = QNearFieldTarget__AccessMethod(0x00)
	QNearFieldTarget__NdefAccess            = QNearFieldTarget__AccessMethod(0x01)
	QNearFieldTarget__TagTypeSpecificAccess = QNearFieldTarget__AccessMethod(0x02)
	QNearFieldTarget__LlcpAccess            = QNearFieldTarget__AccessMethod(0x04)
)

//QNearFieldTarget::Error
type QNearFieldTarget__Error int

var (
	QNearFieldTarget__NoError                = QNearFieldTarget__Error(0)
	QNearFieldTarget__UnknownError           = QNearFieldTarget__Error(1)
	QNearFieldTarget__UnsupportedError       = QNearFieldTarget__Error(2)
	QNearFieldTarget__TargetOutOfRangeError  = QNearFieldTarget__Error(3)
	QNearFieldTarget__NoResponseError        = QNearFieldTarget__Error(4)
	QNearFieldTarget__ChecksumMismatchError  = QNearFieldTarget__Error(5)
	QNearFieldTarget__InvalidParametersError = QNearFieldTarget__Error(6)
	QNearFieldTarget__NdefReadError          = QNearFieldTarget__Error(7)
	QNearFieldTarget__NdefWriteError         = QNearFieldTarget__Error(8)
)

//QNearFieldTarget::Type
type QNearFieldTarget__Type int

var (
	QNearFieldTarget__ProprietaryTag = QNearFieldTarget__Type(0)
	QNearFieldTarget__NfcTagType1    = QNearFieldTarget__Type(1)
	QNearFieldTarget__NfcTagType2    = QNearFieldTarget__Type(2)
	QNearFieldTarget__NfcTagType3    = QNearFieldTarget__Type(3)
	QNearFieldTarget__NfcTagType4    = QNearFieldTarget__Type(4)
	QNearFieldTarget__MifareTag      = QNearFieldTarget__Type(5)
)

func (ptr *QNearFieldTarget) AccessMethods() QNearFieldTarget__AccessMethod {
	if ptr.Pointer() != nil {
		return QNearFieldTarget__AccessMethod(C.QNearFieldTarget_AccessMethods(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNearFieldTarget) ConnectDisconnected(f func()) {
	if ptr.Pointer() != nil {
		C.QNearFieldTarget_ConnectDisconnected(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "disconnected", f)
	}
}

func (ptr *QNearFieldTarget) DisconnectDisconnected() {
	if ptr.Pointer() != nil {
		C.QNearFieldTarget_DisconnectDisconnected(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "disconnected")
	}
}

//export callbackQNearFieldTargetDisconnected
func callbackQNearFieldTargetDisconnected(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "disconnected").(func())()
}

func (ptr *QNearFieldTarget) HasNdefMessage() bool {
	if ptr.Pointer() != nil {
		return C.QNearFieldTarget_HasNdefMessage(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QNearFieldTarget) IsProcessingCommand() bool {
	if ptr.Pointer() != nil {
		return C.QNearFieldTarget_IsProcessingCommand(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QNearFieldTarget) ConnectNdefMessagesWritten(f func()) {
	if ptr.Pointer() != nil {
		C.QNearFieldTarget_ConnectNdefMessagesWritten(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "ndefMessagesWritten", f)
	}
}

func (ptr *QNearFieldTarget) DisconnectNdefMessagesWritten() {
	if ptr.Pointer() != nil {
		C.QNearFieldTarget_DisconnectNdefMessagesWritten(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "ndefMessagesWritten")
	}
}

//export callbackQNearFieldTargetNdefMessagesWritten
func callbackQNearFieldTargetNdefMessagesWritten(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "ndefMessagesWritten").(func())()
}

func (ptr *QNearFieldTarget) Type() QNearFieldTarget__Type {
	if ptr.Pointer() != nil {
		return QNearFieldTarget__Type(C.QNearFieldTarget_Type(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNearFieldTarget) Url() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QNearFieldTarget_Url(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QNearFieldTarget) DestroyQNearFieldTarget() {
	if ptr.Pointer() != nil {
		C.QNearFieldTarget_DestroyQNearFieldTarget(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
