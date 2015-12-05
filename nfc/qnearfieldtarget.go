package nfc

//#include "nfc.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QNearFieldTarget struct {
	core.QObject
}

type QNearFieldTarget_ITF interface {
	core.QObject_ITF
	QNearFieldTarget_PTR() *QNearFieldTarget
}

func PointerFromQNearFieldTarget(ptr QNearFieldTarget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNearFieldTarget_PTR().Pointer()
	}
	return nil
}

func NewQNearFieldTargetFromPointer(ptr unsafe.Pointer) *QNearFieldTarget {
	var n = new(QNearFieldTarget)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QNearFieldTarget_") {
		n.SetObjectName("QNearFieldTarget_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QNearFieldTarget) QNearFieldTarget_PTR() *QNearFieldTarget {
	return ptr
}

//QNearFieldTarget::AccessMethod
type QNearFieldTarget__AccessMethod int64

const (
	QNearFieldTarget__UnknownAccess         = QNearFieldTarget__AccessMethod(0x00)
	QNearFieldTarget__NdefAccess            = QNearFieldTarget__AccessMethod(0x01)
	QNearFieldTarget__TagTypeSpecificAccess = QNearFieldTarget__AccessMethod(0x02)
	QNearFieldTarget__LlcpAccess            = QNearFieldTarget__AccessMethod(0x04)
)

//QNearFieldTarget::Error
type QNearFieldTarget__Error int64

const (
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
type QNearFieldTarget__Type int64

const (
	QNearFieldTarget__ProprietaryTag = QNearFieldTarget__Type(0)
	QNearFieldTarget__NfcTagType1    = QNearFieldTarget__Type(1)
	QNearFieldTarget__NfcTagType2    = QNearFieldTarget__Type(2)
	QNearFieldTarget__NfcTagType3    = QNearFieldTarget__Type(3)
	QNearFieldTarget__NfcTagType4    = QNearFieldTarget__Type(4)
	QNearFieldTarget__MifareTag      = QNearFieldTarget__Type(5)
)

func (ptr *QNearFieldTarget) AccessMethods() QNearFieldTarget__AccessMethod {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNearFieldTarget::accessMethods")
		}
	}()

	if ptr.Pointer() != nil {
		return QNearFieldTarget__AccessMethod(C.QNearFieldTarget_AccessMethods(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNearFieldTarget) ConnectDisconnected(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNearFieldTarget::disconnected")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNearFieldTarget_ConnectDisconnected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "disconnected", f)
	}
}

func (ptr *QNearFieldTarget) DisconnectDisconnected() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNearFieldTarget::disconnected")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNearFieldTarget_DisconnectDisconnected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "disconnected")
	}
}

//export callbackQNearFieldTargetDisconnected
func callbackQNearFieldTargetDisconnected(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNearFieldTarget::disconnected")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "disconnected").(func())()
}

func (ptr *QNearFieldTarget) HasNdefMessage() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNearFieldTarget::hasNdefMessage")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QNearFieldTarget_HasNdefMessage(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNearFieldTarget) IsProcessingCommand() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNearFieldTarget::isProcessingCommand")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QNearFieldTarget_IsProcessingCommand(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNearFieldTarget) ConnectNdefMessagesWritten(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNearFieldTarget::ndefMessagesWritten")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNearFieldTarget_ConnectNdefMessagesWritten(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "ndefMessagesWritten", f)
	}
}

func (ptr *QNearFieldTarget) DisconnectNdefMessagesWritten() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNearFieldTarget::ndefMessagesWritten")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNearFieldTarget_DisconnectNdefMessagesWritten(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "ndefMessagesWritten")
	}
}

//export callbackQNearFieldTargetNdefMessagesWritten
func callbackQNearFieldTargetNdefMessagesWritten(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNearFieldTarget::ndefMessagesWritten")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "ndefMessagesWritten").(func())()
}

func (ptr *QNearFieldTarget) Type() QNearFieldTarget__Type {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNearFieldTarget::type")
		}
	}()

	if ptr.Pointer() != nil {
		return QNearFieldTarget__Type(C.QNearFieldTarget_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNearFieldTarget) Uid() *core.QByteArray {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNearFieldTarget::uid")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QNearFieldTarget_Uid(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNearFieldTarget) DestroyQNearFieldTarget() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNearFieldTarget::~QNearFieldTarget")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNearFieldTarget_DestroyQNearFieldTarget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
