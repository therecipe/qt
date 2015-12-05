package webchannel

//#include "webchannel.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QWebChannel struct {
	core.QObject
}

type QWebChannel_ITF interface {
	core.QObject_ITF
	QWebChannel_PTR() *QWebChannel
}

func PointerFromQWebChannel(ptr QWebChannel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebChannel_PTR().Pointer()
	}
	return nil
}

func NewQWebChannelFromPointer(ptr unsafe.Pointer) *QWebChannel {
	var n = new(QWebChannel)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QWebChannel_") {
		n.SetObjectName("QWebChannel_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QWebChannel) QWebChannel_PTR() *QWebChannel {
	return ptr
}

func (ptr *QWebChannel) BlockUpdates() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebChannel::blockUpdates")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QWebChannel_BlockUpdates(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebChannel) SetBlockUpdates(block bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebChannel::setBlockUpdates")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWebChannel_SetBlockUpdates(ptr.Pointer(), C.int(qt.GoBoolToInt(block)))
	}
}

func NewQWebChannel(parent core.QObject_ITF) *QWebChannel {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebChannel::QWebChannel")
		}
	}()

	return NewQWebChannelFromPointer(C.QWebChannel_NewQWebChannel(core.PointerFromQObject(parent)))
}

func (ptr *QWebChannel) ConnectBlockUpdatesChanged(f func(block bool)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebChannel::blockUpdatesChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWebChannel_ConnectBlockUpdatesChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "blockUpdatesChanged", f)
	}
}

func (ptr *QWebChannel) DisconnectBlockUpdatesChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebChannel::blockUpdatesChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWebChannel_DisconnectBlockUpdatesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "blockUpdatesChanged")
	}
}

//export callbackQWebChannelBlockUpdatesChanged
func callbackQWebChannelBlockUpdatesChanged(ptrName *C.char, block C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebChannel::blockUpdatesChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "blockUpdatesChanged").(func(bool))(int(block) != 0)
}

func (ptr *QWebChannel) ConnectTo(transport QWebChannelAbstractTransport_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebChannel::connectTo")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWebChannel_ConnectTo(ptr.Pointer(), PointerFromQWebChannelAbstractTransport(transport))
	}
}

func (ptr *QWebChannel) DeregisterObject(object core.QObject_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebChannel::deregisterObject")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWebChannel_DeregisterObject(ptr.Pointer(), core.PointerFromQObject(object))
	}
}

func (ptr *QWebChannel) DisconnectFrom(transport QWebChannelAbstractTransport_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebChannel::disconnectFrom")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWebChannel_DisconnectFrom(ptr.Pointer(), PointerFromQWebChannelAbstractTransport(transport))
	}
}

func (ptr *QWebChannel) RegisterObject(id string, object core.QObject_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebChannel::registerObject")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWebChannel_RegisterObject(ptr.Pointer(), C.CString(id), core.PointerFromQObject(object))
	}
}

func (ptr *QWebChannel) DestroyQWebChannel() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebChannel::~QWebChannel")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWebChannel_DestroyQWebChannel(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
