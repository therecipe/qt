package qt

//#include "qstatusbar.h"
import "C"

type qstatusbar struct {
	qwidget
}

type QStatusBar interface {
	QWidget
	AddPermanentWidget_QWidget_Int(widget QWidget, stretch int)
	AddWidget_QWidget_Int(widget QWidget, stretch int)
	CurrentMessage() string
	InsertPermanentWidget_Int_QWidget_Int(index int, widget QWidget, stretch int) int
	InsertWidget_Int_QWidget_Int(index int, widget QWidget, stretch int) int
	IsSizeGripEnabled() bool
	RemoveWidget_QWidget(widget QWidget)
	SetSizeGripEnabled_Bool(enabled bool)
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

func NewQStatusBar_QWidget(parent QWidget) QStatusBar {
	var parentPtr C.QtObjectPtr = nil
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qstatusbar = new(qstatusbar)
	qstatusbar.SetPointer(C.QStatusBar_New_QWidget(parentPtr))
	qstatusbar.SetObjectName_String("QStatusBar_" + randomIdentifier())
	return qstatusbar
}

func (p *qstatusbar) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QStatusBar_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qstatusbar) AddPermanentWidget_QWidget_Int(widget QWidget, stretch int) {
	if p.Pointer() == nil {
	} else {
		var widgetPtr C.QtObjectPtr = nil
		if widget != nil {
			widgetPtr = widget.Pointer()
		}
		C.QStatusBar_AddPermanentWidget_QWidget_Int(p.Pointer(), widgetPtr, C.int(stretch))
	}
}

func (p *qstatusbar) AddWidget_QWidget_Int(widget QWidget, stretch int) {
	if p.Pointer() == nil {
	} else {
		var widgetPtr C.QtObjectPtr = nil
		if widget != nil {
			widgetPtr = widget.Pointer()
		}
		C.QStatusBar_AddWidget_QWidget_Int(p.Pointer(), widgetPtr, C.int(stretch))
	}
}

func (p *qstatusbar) CurrentMessage() string {
	if p.Pointer() == nil {
		return ""
	} else {
		return C.GoString(C.QStatusBar_CurrentMessage(p.Pointer()))
	}
}

func (p *qstatusbar) InsertPermanentWidget_Int_QWidget_Int(index int, widget QWidget, stretch int) int {
	if p.Pointer() == nil {
		return 0
	} else {
		var widgetPtr C.QtObjectPtr = nil
		if widget != nil {
			widgetPtr = widget.Pointer()
		}
		return int(C.QStatusBar_InsertPermanentWidget_Int_QWidget_Int(p.Pointer(), C.int(index), widgetPtr, C.int(stretch)))
	}
}

func (p *qstatusbar) InsertWidget_Int_QWidget_Int(index int, widget QWidget, stretch int) int {
	if p.Pointer() == nil {
		return 0
	} else {
		var widgetPtr C.QtObjectPtr = nil
		if widget != nil {
			widgetPtr = widget.Pointer()
		}
		return int(C.QStatusBar_InsertWidget_Int_QWidget_Int(p.Pointer(), C.int(index), widgetPtr, C.int(stretch)))
	}
}

func (p *qstatusbar) IsSizeGripEnabled() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QStatusBar_IsSizeGripEnabled(p.Pointer()) != 0
	}
}

func (p *qstatusbar) RemoveWidget_QWidget(widget QWidget) {
	if p.Pointer() == nil {
	} else {
		var widgetPtr C.QtObjectPtr = nil
		if widget != nil {
			widgetPtr = widget.Pointer()
		}
		C.QStatusBar_RemoveWidget_QWidget(p.Pointer(), widgetPtr)
	}
}

func (p *qstatusbar) SetSizeGripEnabled_Bool(enabled bool) {
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
