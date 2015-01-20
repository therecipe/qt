package qt

//#include "qtoolbox.h"
import "C"

type qtoolbox struct {
	qframe
}

type QToolBox interface {
	QFrame
	AddItem(w QWidget, text string) int
	Count() int
	CurrentIndex() int
	CurrentWidget() QWidget
	IndexOf(widget QWidget) int
	InsertItem(index int, widget QWidget, text string) int
	IsItemEnabled(index int) bool
	ItemText(index int) string
	ItemToolTip(index int) string
	RemoveItem(index int)
	SetItemEnabled(index int, enabled bool)
	SetItemText(index int, text string)
	SetItemToolTip(index int, toolTip string)
	Widget(index int) QWidget
	ConnectSlotSetCurrentIndex()
	DisconnectSlotSetCurrentIndex()
	SlotSetCurrentIndex(index int)
	ConnectSignalCurrentChanged(f func())
	DisconnectSignalCurrentChanged()
	SignalCurrentChanged() func()
}

func (p *qtoolbox) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qtoolbox) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQToolBox(parent QWidget, f WindowType) QToolBox {
	var parentPtr C.QtObjectPtr
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qtoolbox = new(qtoolbox)
	qtoolbox.SetPointer(C.QToolBox_New_QWidget_WindowType(parentPtr, C.int(f)))
	qtoolbox.SetObjectName("QToolBox_" + randomIdentifier())
	return qtoolbox
}

func (p *qtoolbox) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QToolBox_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qtoolbox) AddItem(w QWidget, text string) int {
	if p.Pointer() == nil {
		return 0
	} else {
		var wPtr C.QtObjectPtr
		if w != nil {
			wPtr = w.Pointer()
		}
		return int(C.QToolBox_AddItem_QWidget_String(p.Pointer(), wPtr, C.CString(text)))
	}
}

func (p *qtoolbox) Count() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QToolBox_Count(p.Pointer()))
}

func (p *qtoolbox) CurrentIndex() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QToolBox_CurrentIndex(p.Pointer()))
}

func (p *qtoolbox) CurrentWidget() QWidget {
	if p.Pointer() == nil {
		return nil
	} else {
		var qwidget = new(qwidget)
		qwidget.SetPointer(C.QToolBox_CurrentWidget(p.Pointer()))
		if qwidget.ObjectName() == "" {
			qwidget.SetObjectName("QWidget_" + randomIdentifier())
		}
		return qwidget
	}
}

func (p *qtoolbox) IndexOf(widget QWidget) int {
	if p.Pointer() == nil {
		return 0
	} else {
		var widgetPtr C.QtObjectPtr
		if widget != nil {
			widgetPtr = widget.Pointer()
		}
		return int(C.QToolBox_IndexOf_QWidget(p.Pointer(), widgetPtr))
	}
}

func (p *qtoolbox) InsertItem(index int, widget QWidget, text string) int {
	if p.Pointer() == nil {
		return 0
	} else {
		var widgetPtr C.QtObjectPtr
		if widget != nil {
			widgetPtr = widget.Pointer()
		}
		return int(C.QToolBox_InsertItem_Int_QWidget_String(p.Pointer(), C.int(index), widgetPtr, C.CString(text)))
	}
}

func (p *qtoolbox) IsItemEnabled(index int) bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QToolBox_IsItemEnabled_Int(p.Pointer(), C.int(index)) != 0
}

func (p *qtoolbox) ItemText(index int) string {
	if p.Pointer() == nil {
		return ""
	}
	return C.GoString(C.QToolBox_ItemText_Int(p.Pointer(), C.int(index)))
}

func (p *qtoolbox) ItemToolTip(index int) string {
	if p.Pointer() == nil {
		return ""
	}
	return C.GoString(C.QToolBox_ItemToolTip_Int(p.Pointer(), C.int(index)))
}

func (p *qtoolbox) RemoveItem(index int) {
	if p.Pointer() != nil {
		C.QToolBox_RemoveItem_Int(p.Pointer(), C.int(index))
	}
}

func (p *qtoolbox) SetItemEnabled(index int, enabled bool) {
	if p.Pointer() != nil {
		C.QToolBox_SetItemEnabled_Int_Bool(p.Pointer(), C.int(index), goBoolToCInt(enabled))
	}
}

func (p *qtoolbox) SetItemText(index int, text string) {
	if p.Pointer() != nil {
		C.QToolBox_SetItemText_Int_String(p.Pointer(), C.int(index), C.CString(text))
	}
}

func (p *qtoolbox) SetItemToolTip(index int, toolTip string) {
	if p.Pointer() != nil {
		C.QToolBox_SetItemToolTip_Int_String(p.Pointer(), C.int(index), C.CString(toolTip))
	}
}

func (p *qtoolbox) Widget(index int) QWidget {
	if p.Pointer() == nil {
		return nil
	} else {
		var qwidget = new(qwidget)
		qwidget.SetPointer(C.QToolBox_Widget_Int(p.Pointer(), C.int(index)))
		if qwidget.ObjectName() == "" {
			qwidget.SetObjectName("QWidget_" + randomIdentifier())
		}
		return qwidget
	}
}

func (p *qtoolbox) ConnectSlotSetCurrentIndex() {
	C.QToolBox_ConnectSlotSetCurrentIndex(p.Pointer())
}

func (p *qtoolbox) DisconnectSlotSetCurrentIndex() {
	C.QToolBox_DisconnectSlotSetCurrentIndex(p.Pointer())
}

func (p *qtoolbox) SlotSetCurrentIndex(index int) {
	if p.Pointer() != nil {
		C.QToolBox_SetCurrentIndex_Int(p.Pointer(), C.int(index))
	}
}

func (p *qtoolbox) ConnectSignalCurrentChanged(f func()) {
	C.QToolBox_ConnectSignalCurrentChanged(p.Pointer())
	connectSignal(p.ObjectName(), "currentChanged", f)
}

func (p *qtoolbox) DisconnectSignalCurrentChanged() {
	C.QToolBox_DisconnectSignalCurrentChanged(p.Pointer())
	disconnectSignal(p.ObjectName(), "currentChanged")
}

func (p *qtoolbox) SignalCurrentChanged() func() {
	return getSignal(p.ObjectName(), "currentChanged")
}

//export callbackQToolBox
func callbackQToolBox(ptr C.QtObjectPtr, msg *C.char) {
	var qtoolbox = new(qtoolbox)
	qtoolbox.SetPointer(ptr)
	getSignal(qtoolbox.ObjectName(), C.GoString(msg))()
}
