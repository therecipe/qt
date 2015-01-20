package qt

//#include "qstatusbar.h"
import "C"

type qstatusbar struct {
	qwidget
}

type QStatusBar interface {
	QWidget
	AddPermanentWidget(widget QWidget, stretch int)
	AddWidget(widget QWidget, stretch int)
	CurrentMessage() string
	InsertPermanentWidget(index int, widget QWidget, stretch int) int
	InsertWidget(index int, widget QWidget, stretch int) int
	IsSizeGripEnabled() bool
	RemoveWidget(widget QWidget)
	SetSizeGripEnabled(enabled bool)
	ConnectSlotClearMessage()
	DisconnectSlotClearMessage()
	SlotClearMessage()
	ConnectSignalMessageChanged(f func())
	DisconnectSignalMessageChanged()
	SignalMessageChanged() func()
}

func (p *qstatusbar) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qstatusbar) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQStatusBar(parent QWidget) QStatusBar {
	var parentPtr C.QtObjectPtr
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qstatusbar = new(qstatusbar)
	qstatusbar.SetPointer(C.QStatusBar_New_QWidget(parentPtr))
	qstatusbar.SetObjectName("QStatusBar_" + randomIdentifier())
	return qstatusbar
}

func (p *qstatusbar) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QStatusBar_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qstatusbar) AddPermanentWidget(widget QWidget, stretch int) {
	if p.Pointer() != nil {
		var widgetPtr C.QtObjectPtr
		if widget != nil {
			widgetPtr = widget.Pointer()
		}
		C.QStatusBar_AddPermanentWidget_QWidget_Int(p.Pointer(), widgetPtr, C.int(stretch))
	}
}

func (p *qstatusbar) AddWidget(widget QWidget, stretch int) {
	if p.Pointer() != nil {
		var widgetPtr C.QtObjectPtr
		if widget != nil {
			widgetPtr = widget.Pointer()
		}
		C.QStatusBar_AddWidget_QWidget_Int(p.Pointer(), widgetPtr, C.int(stretch))
	}
}

func (p *qstatusbar) CurrentMessage() string {
	if p.Pointer() == nil {
		return ""
	}
	return C.GoString(C.QStatusBar_CurrentMessage(p.Pointer()))
}

func (p *qstatusbar) InsertPermanentWidget(index int, widget QWidget, stretch int) int {
	if p.Pointer() == nil {
		return 0
	} else {
		var widgetPtr C.QtObjectPtr
		if widget != nil {
			widgetPtr = widget.Pointer()
		}
		return int(C.QStatusBar_InsertPermanentWidget_Int_QWidget_Int(p.Pointer(), C.int(index), widgetPtr, C.int(stretch)))
	}
}

func (p *qstatusbar) InsertWidget(index int, widget QWidget, stretch int) int {
	if p.Pointer() == nil {
		return 0
	} else {
		var widgetPtr C.QtObjectPtr
		if widget != nil {
			widgetPtr = widget.Pointer()
		}
		return int(C.QStatusBar_InsertWidget_Int_QWidget_Int(p.Pointer(), C.int(index), widgetPtr, C.int(stretch)))
	}
}

func (p *qstatusbar) IsSizeGripEnabled() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QStatusBar_IsSizeGripEnabled(p.Pointer()) != 0
}

func (p *qstatusbar) RemoveWidget(widget QWidget) {
	if p.Pointer() != nil {
		var widgetPtr C.QtObjectPtr
		if widget != nil {
			widgetPtr = widget.Pointer()
		}
		C.QStatusBar_RemoveWidget_QWidget(p.Pointer(), widgetPtr)
	}
}

func (p *qstatusbar) SetSizeGripEnabled(enabled bool) {
	if p.Pointer() != nil {
		C.QStatusBar_SetSizeGripEnabled_Bool(p.Pointer(), goBoolToCInt(enabled))
	}
}

func (p *qstatusbar) ConnectSlotClearMessage() {
	C.QStatusBar_ConnectSlotClearMessage(p.Pointer())
}

func (p *qstatusbar) DisconnectSlotClearMessage() {
	C.QStatusBar_DisconnectSlotClearMessage(p.Pointer())
}

func (p *qstatusbar) SlotClearMessage() {
	if p.Pointer() != nil {
		C.QStatusBar_ClearMessage(p.Pointer())
	}
}

func (p *qstatusbar) ConnectSignalMessageChanged(f func()) {
	C.QStatusBar_ConnectSignalMessageChanged(p.Pointer())
	connectSignal(p.ObjectName(), "messageChanged", f)
}

func (p *qstatusbar) DisconnectSignalMessageChanged() {
	C.QStatusBar_DisconnectSignalMessageChanged(p.Pointer())
	disconnectSignal(p.ObjectName(), "messageChanged")
}

func (p *qstatusbar) SignalMessageChanged() func() {
	return getSignal(p.ObjectName(), "messageChanged")
}

//export callbackQStatusBar
func callbackQStatusBar(ptr C.QtObjectPtr, msg *C.char) {
	var qstatusbar = new(qstatusbar)
	qstatusbar.SetPointer(ptr)
	getSignal(qstatusbar.ObjectName(), C.GoString(msg))()
}
