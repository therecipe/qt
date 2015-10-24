package webchannel

//#include "qwebchannel.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QWebChannel struct {
	core.QObject
}

type QWebChannelITF interface {
	core.QObjectITF
	QWebChannelPTR() *QWebChannel
}

func PointerFromQWebChannel(ptr QWebChannelITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebChannelPTR().Pointer()
	}
	return nil
}

func QWebChannelFromPointer(ptr unsafe.Pointer) *QWebChannel {
	var n = new(QWebChannel)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QWebChannel_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QWebChannel) QWebChannelPTR() *QWebChannel {
	return ptr
}

func (ptr *QWebChannel) BlockUpdates() bool {
	if ptr.Pointer() != nil {
		return C.QWebChannel_BlockUpdates(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWebChannel) SetBlockUpdates(block bool) {
	if ptr.Pointer() != nil {
		C.QWebChannel_SetBlockUpdates(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(block)))
	}
}

func NewQWebChannel(parent core.QObjectITF) *QWebChannel {
	return QWebChannelFromPointer(unsafe.Pointer(C.QWebChannel_NewQWebChannel(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QWebChannel) ConnectBlockUpdatesChanged(f func(block bool)) {
	if ptr.Pointer() != nil {
		C.QWebChannel_ConnectBlockUpdatesChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "blockUpdatesChanged", f)
	}
}

func (ptr *QWebChannel) DisconnectBlockUpdatesChanged() {
	if ptr.Pointer() != nil {
		C.QWebChannel_DisconnectBlockUpdatesChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "blockUpdatesChanged")
	}
}

//export callbackQWebChannelBlockUpdatesChanged
func callbackQWebChannelBlockUpdatesChanged(ptrName *C.char, block C.int) {
	qt.GetSignal(C.GoString(ptrName), "blockUpdatesChanged").(func(bool))(int(block) != 0)
}

func (ptr *QWebChannel) ConnectTo(transport QWebChannelAbstractTransportITF) {
	if ptr.Pointer() != nil {
		C.QWebChannel_ConnectTo(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWebChannelAbstractTransport(transport)))
	}
}

func (ptr *QWebChannel) DeregisterObject(object core.QObjectITF) {
	if ptr.Pointer() != nil {
		C.QWebChannel_DeregisterObject(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQObject(object)))
	}
}

func (ptr *QWebChannel) DisconnectFrom(transport QWebChannelAbstractTransportITF) {
	if ptr.Pointer() != nil {
		C.QWebChannel_DisconnectFrom(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWebChannelAbstractTransport(transport)))
	}
}

func (ptr *QWebChannel) RegisterObject(id string, object core.QObjectITF) {
	if ptr.Pointer() != nil {
		C.QWebChannel_RegisterObject(C.QtObjectPtr(ptr.Pointer()), C.CString(id), C.QtObjectPtr(core.PointerFromQObject(object)))
	}
}

func (ptr *QWebChannel) DestroyQWebChannel() {
	if ptr.Pointer() != nil {
		C.QWebChannel_DestroyQWebChannel(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
