package qt

//#include "qstackedlayout.h"
import "C"

type qstackedlayout struct {
	qlayout
}

type QStackedLayout interface {
	QLayout
	CurrentIndex() int
	CurrentWidget() QWidget
	InsertWidget_Int_QWidget(index int, widget QWidget) int
	Widget_Int(index int) QWidget
	ConnectSlotSetCurrentIndex()
	DisconnectSlotSetCurrentIndex()
	SlotSetCurrentIndex_Int(index int)
	ConnectSignalCurrentChanged(f func())
	DisconnectSignalCurrentChanged()
	SignalCurrentChanged() func()
	ConnectSignalWidgetRemoved(f func())
	DisconnectSignalWidgetRemoved()
	SignalWidgetRemoved() func()
}

func (p *qstackedlayout) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qstackedlayout) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQStackedLayout() QStackedLayout {
	var qstackedlayout = new(qstackedlayout)
	qstackedlayout.SetPointer(C.QStackedLayout_New())
	qstackedlayout.SetObjectName_String("QStackedLayout_" + randomIdentifier())
	return qstackedlayout
}

func NewQStackedLayout_QWidget(parent QWidget) QStackedLayout {
	var parentPtr C.QtObjectPtr = nil
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qstackedlayout = new(qstackedlayout)
	qstackedlayout.SetPointer(C.QStackedLayout_New_QWidget(parentPtr))
	qstackedlayout.SetObjectName_String("QStackedLayout_" + randomIdentifier())
	return qstackedlayout
}

func NewQStackedLayout_QLayout(parentLayout QLayout) QStackedLayout {
	var parentLayoutPtr C.QtObjectPtr = nil
	if parentLayout != nil {
		parentLayoutPtr = parentLayout.Pointer()
	}
	var qstackedlayout = new(qstackedlayout)
	qstackedlayout.SetPointer(C.QStackedLayout_New_QLayout(parentLayoutPtr))
	qstackedlayout.SetObjectName_String("QStackedLayout_" + randomIdentifier())
	return qstackedlayout
}

func (p *qstackedlayout) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QStackedLayout_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qstackedlayout) CurrentIndex() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QStackedLayout_CurrentIndex(p.Pointer()))
	}
}

func (p *qstackedlayout) CurrentWidget() QWidget {
	if p.Pointer() == nil {
		return nil
	} else {
		var qwidget = new(qwidget)
		qwidget.SetPointer(C.QStackedLayout_CurrentWidget(p.Pointer()))
		if qwidget.ObjectName() == "" {
			qwidget.SetObjectName_String("QWidget_" + randomIdentifier())
		}
		return qwidget
	}
}

func (p *qstackedlayout) InsertWidget_Int_QWidget(index int, widget QWidget) int {
	if p.Pointer() == nil {
		return 0
	} else {
		var widgetPtr C.QtObjectPtr = nil
		if widget != nil {
			widgetPtr = widget.Pointer()
		}
		return int(C.QStackedLayout_InsertWidget_Int_QWidget(p.Pointer(), C.int(index), widgetPtr))
	}
}

func (p *qstackedlayout) Widget_Int(index int) QWidget {
	if p.Pointer() == nil {
		return nil
	} else {
		var qwidget = new(qwidget)
		qwidget.SetPointer(C.QStackedLayout_Widget_Int(p.Pointer(), C.int(index)))
		if qwidget.ObjectName() == "" {
			qwidget.SetObjectName_String("QWidget_" + randomIdentifier())
		}
		return qwidget
	}
}

func (p *qstackedlayout) ConnectSlotSetCurrentIndex() {
	C.QStackedLayout_ConnectSlotSetCurrentIndex(p.Pointer())
}

func (p *qstackedlayout) DisconnectSlotSetCurrentIndex() {
	C.QStackedLayout_DisconnectSlotSetCurrentIndex(p.Pointer())
}

func (p *qstackedlayout) SlotSetCurrentIndex_Int(index int) {
	if p.Pointer() != nil {
		C.QStackedLayout_SetCurrentIndex_Int(p.Pointer(), C.int(index))
	}
}

func (p *qstackedlayout) ConnectSignalCurrentChanged(f func()) {
	C.QStackedLayout_ConnectSignalCurrentChanged(p.Pointer())
	connectSignal(p.ObjectName(), "currentChanged", f)
}

func (p *qstackedlayout) DisconnectSignalCurrentChanged() {
	C.QStackedLayout_DisconnectSignalCurrentChanged(p.Pointer())
	disconnectSignal(p.ObjectName(), "currentChanged")
}

func (p *qstackedlayout) SignalCurrentChanged() func() {
	return getSignal(p.ObjectName(), "currentChanged")
}

func (p *qstackedlayout) ConnectSignalWidgetRemoved(f func()) {
	C.QStackedLayout_ConnectSignalWidgetRemoved(p.Pointer())
	connectSignal(p.ObjectName(), "widgetRemoved", f)
}

func (p *qstackedlayout) DisconnectSignalWidgetRemoved() {
	C.QStackedLayout_DisconnectSignalWidgetRemoved(p.Pointer())
	disconnectSignal(p.ObjectName(), "widgetRemoved")
}

func (p *qstackedlayout) SignalWidgetRemoved() func() {
	return getSignal(p.ObjectName(), "widgetRemoved")
}

//export callbackQStackedLayout
func callbackQStackedLayout(ptr C.QtObjectPtr, msg *C.char) {
	var qstackedlayout = new(qstackedlayout)
	qstackedlayout.SetPointer(ptr)
	getSignal(qstackedlayout.ObjectName(), C.GoString(msg))()
}
